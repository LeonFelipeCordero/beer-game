package com.beer.game.application.service

import com.beer.game.adapters.out.mongo.BoardMongoAdapter
import com.beer.game.domain.Board
import org.springframework.stereotype.Service

@Service
class BoardService(
    private val boardMongoAdapter: BoardMongoAdapter
) {
    fun loadBoard(boardId: String): Board {
        return boardMongoAdapter.loadBoard(boardId)
    }

    fun createBoard(name: String): Board {
        return boardMongoAdapter.saveBoard(name)
    }

    fun loadActiveBoards(): List<Board> {
        return boardMongoAdapter.loadActiveBoards()
    }
}