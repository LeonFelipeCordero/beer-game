package com.beer.game.repositories.board

import com.beer.game.application.board.BoardStorageAdapter
import com.beer.game.domain.Board
import com.beer.game.common.BoardState
import com.beer.game.domain.exceptions.NotFoundException
import org.bson.types.ObjectId
import org.springframework.stereotype.Component

@Component
class BoardMongoAdapter(
    private val boardRepository: BoardRepository
) : BoardStorageAdapter {

    override fun loadBoard(boardId: String): Board {
        return loadBoardDocument(boardId).toBoard()
    }

    override fun loadBoardByName(name: String): Board? {
        return boardRepository.findOneByName(name)?.toBoard()
    }

    override fun saveBoard(name: String): Board {
        val board = BoardDocument(name = name).toBoard()
        return upsertBoard(board)
    }

    override fun upsertBoard(board: Board): Board {
        return boardRepository.save(BoardDocument.fromBoard(board)).toBoard()
    }

    override fun loadActiveBoards(): List<Board> {
        return loadActiveBoardDocuments()
            .map { it.toBoard() }
    }

    override fun loadBoardByPlayerId(playerId: String): Board {
        return boardRepository.findOneByPlayersId(playerId)?.toBoard()
            ?: throw NotFoundException("Board with player id $playerId doesn't exit", "Verify the id provided")
    }

    override fun loadBoardByOrderId(orderId: String): Board {
        return boardRepository.findOneByOrdersId(orderId)?.toBoard()
            ?: throw NotFoundException("Board with order id $orderId doesn't exit", "Verify the id provided")
    }

    private fun loadBoardDocument(boardId: String): BoardDocument {
        return boardRepository.findOneById(ObjectId(boardId))
            ?: throw NotFoundException("Board with id $boardId doesn't exit", "Verify the id provided")
    }

    private fun loadActiveBoardDocuments(): List<BoardDocument> {
        return boardRepository.findBoardDocumentByState(BoardState.RUNNING.toString())
    }
}