package com.beer.game.application.service

import com.beer.game.adapters.out.mongo.*
import com.beer.game.domain.Board
import com.beer.game.domain.exceptions.ImpossibleActionException
import org.springframework.stereotype.Service

@Service
class BoardService(
    private val boardMongoAdapter: BoardMongoAdapter,
) {
    fun loadBoard(boardId: String): Board {
        return boardMongoAdapter.loadBoard(boardId)
    }

    fun createBoard(name: String): Board {
        return boardMongoAdapter.loadBoardByName(name)
            ?.let {
                throw ImpossibleActionException(
                    "Name is already used by another board",
                    "find another name for your game"
                )
            }
            ?: boardMongoAdapter.saveBoard(name)
    }

    fun loadActiveBoards(): List<Board> {
        return boardMongoAdapter.loadActiveBoards()
    }
}