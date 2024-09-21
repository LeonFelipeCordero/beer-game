package com.beer.game.application.player

import com.beer.game.application.board.BoardStorageAdapter
import com.beer.game.application.events.InternalEventListener
import com.beer.game.common.BoardState
import com.beer.game.common.Role
import com.beer.game.domain.Player
import com.beer.game.domain.exceptions.ImpossibleActionException
import org.springframework.stereotype.Service

@Service
class PlayerService(
    private val playerStorageAdapter: PlayerStorageAdapter,
    private val boardStorageAdapter: BoardStorageAdapter,
    private val internalEventListener: InternalEventListener,
) {

    fun addPlayer(boardId: String, role: Role): Player {
        val board = boardStorageAdapter.loadBoard(boardId)
        if (board.full) {
            throw ImpossibleActionException("Board $boardId is full", "Verify the id provided")
        }

        board.players.find { it.role == role }?.let {
            throw ImpossibleActionException(
                "Board $boardId already has role $role",
                "Verify the id provided",
            )
        }

        val player = createPlayer(role)
        board.addPlayer(player)

        if (board.players.size == 3) {
            board.full = true
            board.state = BoardState.RUNNING
            board.emitUpdate(internalEventListener)
        }
        val savedPlayer = playerStorageAdapter.savePlayer(board, player)
        player.emitCreation(internalEventListener, board)

        return savedPlayer
    }

    fun getPlayer(playerId: String): Pair<Player, String> {
        return playerStorageAdapter.loadPlayer(playerId)
    }

    fun getPlayersInBoard(boardId: String): MutableList<Player> {
        return boardStorageAdapter.loadBoard(boardId).players
    }

    fun changeWeeklyOrder(playerId: String, amount: Int) {
        val board = boardStorageAdapter.loadBoardByPlayerId(playerId)
        val player = board.players.first { it.id == playerId }
        player.weeklyOrder = amount
        playerStorageAdapter.savePlayer(board, player)
        player.emitUpdate(internalEventListener, board.id)
    }

    private fun createPlayer(role: Role) =
        when (role) {
            Role.RETAILER -> Player.retailPlayer()
            Role.WHOLESALER -> Player.wholesalerPlayer()
            else -> Player.factoryPlayer()
        }
}
