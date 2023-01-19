package com.beer.game.adapters.out.mongo

import com.beer.game.domain.Board
import com.beer.game.common.BoardState
import com.beer.game.domain.exceptions.NotFoundException
import org.bson.types.ObjectId
import org.springframework.stereotype.Component

@Component
class BoardMongoAdapter(
    private val boardRepository: BoardRepository
) {

    fun loadBoard(boardId: String): Board {
        return loadBoardDocument(boardId).toBoard()
    }

    fun loadBoardByName(name: String): Board? {
        return boardRepository.findOneByName(name)?.toBoard()
    }

    fun loadBoardDocument(boardId: String): BoardDocument {
        return boardRepository.findOneById(ObjectId(boardId))
            ?: throw NotFoundException("Board with id $boardId doesn't exit", "Verify the id provided")
    }

    fun saveBoard(name: String): Board {
        val saveDocument = upsertBoard(BoardDocument(name = name))
        return saveDocument.toBoard()
    }

    fun upsertBoard(boardDocument: BoardDocument): BoardDocument {
        return boardRepository.save(boardDocument)
    }

    fun loadActiveBoards(): List<Board> {
        return loadActiveBoardDocuments()
            .map { it.toBoard() }
    }

    fun loadActiveBoardDocuments(): List<BoardDocument> {
        return boardRepository.findBoardDocumentByState(BoardState.RUNNING.toString())
    }

    fun loadBoardByPlayerId(playerId: String): Board {
        return boardRepository.findOneByPlayersId(playerId)?.toBoard()
            ?: throw NotFoundException("Board with player id $playerId doesn't exit", "Verify the id provided")
    }

    fun loadBoardByOrderId(orderId: String): Board {
        return boardRepository.findOneByOrdersId(orderId)?.toBoard()
            ?: throw NotFoundException("Board with order id $orderId doesn't exit", "Verify the id provided")
    }
}