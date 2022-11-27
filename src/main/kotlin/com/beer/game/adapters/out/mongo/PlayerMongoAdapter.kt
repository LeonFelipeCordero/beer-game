package com.beer.game.adapters.out.mongo

import com.beer.game.domain.Board
import com.beer.game.domain.Player
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

    fun loadPlayer(board: String, playerId: String): Player {
        return boardMongoAdapter.loadBoard(board)
            .players
            .first { it.id == playerId }
    }
}