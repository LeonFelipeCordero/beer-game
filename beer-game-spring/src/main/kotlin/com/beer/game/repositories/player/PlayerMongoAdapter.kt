package com.beer.game.repositories.player

import com.beer.game.domain.Board
import com.beer.game.domain.Player
import com.beer.game.repositories.board.BoardDocument
import com.beer.game.repositories.board.BoardMongoAdapter
import org.springframework.stereotype.Component

@Component
class PlayerMongoAdapter(
    private val boardMongoAdapter: BoardMongoAdapter,
) {

    fun savePlayer(board: Board, player: Player): Player {
        val boardDocument = BoardDocument.fromBoard(board)
        boardMongoAdapter.upsertBoard(boardDocument)
        return player
    }

    fun loadPlayer(playerId: String): Pair<Player, String> {
        val board = boardMongoAdapter.loadBoardByPlayerId(playerId)
        val player = board.players.first { it.id == playerId }
        return Pair(player, board.id)
    }
}