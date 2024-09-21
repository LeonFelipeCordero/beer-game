package com.beer.game.application.board

import com.beer.game.domain.Board
import com.beer.game.domain.exceptions.ImpossibleActionException
import org.springframework.stereotype.Service

@Service
class BoardService(
    private val boardStorageAdapter: BoardStorageAdapter,
) {
    fun loadBoard(boardId: String): Board {
        return boardStorageAdapter.loadBoard(boardId)
    }

    fun loadBoardByName(name: String): Board {
        return boardStorageAdapter.loadBoardByName(name)!!
    }

    fun createBoard(name: String): Board {
        return boardStorageAdapter.loadBoardByName(name)
            ?.let {
                throw ImpossibleActionException(
                    "Name is already used by another board",
                    "find another name for your game",
                )
            }
            ?: boardStorageAdapter.saveBoard(name)
    }

    fun loadActiveBoards(): List<Board> {
        return boardStorageAdapter.loadActiveBoards()
    }

    fun loadBoardByPlayerId(playerId: String): Board {
        return boardStorageAdapter.loadBoardByPlayerId(playerId)
    }

    fun loadBoardByOrderId(orderId: String): Board {
        return boardStorageAdapter.loadBoardByOrderId(orderId)
    }
}
