package com.beer.game.repositories.player

import com.beer.game.application.player.PlayerStorageAdapter
import com.beer.game.domain.Board
import com.beer.game.domain.Player
import com.beer.game.repositories.board.BoardMongoAdapter
import org.springframework.stereotype.Component

@Component
class PlayerMongoAdapter(
    private val boardMongoAdapter: BoardMongoAdapter,
) : PlayerStorageAdapter {

    override fun savePlayer(board: Board, player: Player): Player {
        boardMongoAdapter.upsertBoard(board)
        return player
    }

    override fun loadPlayer(playerId: String): Pair<Player, String> {
        val board = boardMongoAdapter.loadBoardByPlayerId(playerId)
        val player = board.players.first { it.id == playerId }
        return Pair(player, board.id)
    }
}
