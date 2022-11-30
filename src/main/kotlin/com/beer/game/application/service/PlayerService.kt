package com.beer.game.application.service

import com.beer.game.adapters.out.mongo.*
import com.beer.game.common.BoardState
import com.beer.game.domain.Player
import com.beer.game.common.Role
import com.beer.game.domain.exceptions.ImpossibleActionException
import com.beer.game.events.InternalEventListener
import org.springframework.stereotype.Service

@Service
class PlayerService(
    private val playerMongoAdapter: PlayerMongoAdapter,
    private val boardMongoAdapter: BoardMongoAdapter,
    private val internalEventListener: InternalEventListener
) {

    fun addPlayer(boardId: String, role: Role): Player {
        val board = boardMongoAdapter.loadBoard(boardId)
        if (board.full) {
            throw ImpossibleActionException("Board $boardId is full", "Verify the id provided")
        }

        board.players.find { it.role == role }?.let {
            throw ImpossibleActionException(
                "Board $boardId already has role $role",
                "Verify the id provided"
            )
        }

        val player = createPlayer(role)
        board.addPlayer(player)

        val savedPlayer = playerMongoAdapter.savePlayer(board, player)
        player.emitCreation(internalEventListener, board)

        if (board.players.size == 3) {
            board.full = true
            board.state = BoardState.RUNNING
            board.emitUpdate(internalEventListener)
        }
        return savedPlayer
    }

    fun getPlayer(boardId: String, playerId: String): Player {
        return playerMongoAdapter.loadPlayer(boardId, playerId)
    }

    fun getPlayersInBoard(boardId: String): MutableList<Player> {
        return boardMongoAdapter.loadBoard(boardId)
            .players
    }

    fun changeWeeklyOrder(boardId: String, playerId: String, amount: Int) {
        val board = boardMongoAdapter.loadBoard(boardId)
        val player = board.players.first { it.id == playerId }
        player.weeklyOrder = amount
        playerMongoAdapter.savePlayer(board, player)
        player.emitUpdate(internalEventListener, board)
    }

    private fun createPlayer(role: Role) =
        when (role) {
            Role.RETAILER -> Player(
                name = role.toString(),
                role = role,
                stock = 80,
                backlog = 80,
                weeklyOrder = 40,
                lastOrder = 40,
                cpu = false,
                orders = mutableListOf()
            )

            Role.WHOLESALER -> Player(
                name = role.toString(),
                role = role,
                stock = 1200,
                backlog = 1200,
                weeklyOrder = 600,
                lastOrder = 600,
                cpu = false,
                orders = mutableListOf()
            )

            else -> Player(
                name = role.toString(),
                role = role,
                stock = 12000,
                backlog = 12000,
                weeklyOrder = 6000,
                lastOrder = 6000,
                cpu = false,
                orders = mutableListOf()
            )
        }
}