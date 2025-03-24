package com.beer.game.repositories.board.postgres

import com.beer.game.application.board.BoardStorageAdapter
import com.beer.game.common.BoardState
import com.beer.game.domain.Board
import com.beer.game.domain.exceptions.NotFoundException
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import java.util.UUID

@Component
class BoardPostgresAdapter(
    private val boardRepository: BoardRepository,
) : BoardStorageAdapter {

    override fun save(board: Board): Mono<Board> {
        return boardRepository.save(BoardEntity.fromBoard(board)).map { it.toBoard() }
    }

    override fun loadBoard(boardId: String): Mono<Board> =
        loadBoardEntity(UUID.fromString(boardId)).map { it.toBoard() }

    override fun loadBoardByName(name: String): Mono<Board> {
        return boardRepository.findByName(name).map { it.toBoard() }
    }

    override fun createBoard(name: String): Mono<Board> {
        val boardEntity = BoardEntity(name = name)
        return boardRepository.save(boardEntity).map { it.toBoard() }
    }

    override fun upsertBoard(board: Board): Mono<Board> {
        return boardRepository.save(BoardEntity.fromBoard(board)).map { it.toBoard() }
    }

    override fun loadActiveBoards(): Flux<Board> {
        return loadActiveBoardEntities().map { it.toBoard() }
    }

    override fun loadBoardByPlayerId(playerId: String): Mono<Board> {
        return boardRepository.findOneByPlayersId(playerId)
            .map { it.toBoard() }
            .switchIfEmpty(Mono.error {
                throw NotFoundException(
                    "Board with player id $playerId doesn't exit",
                    "Verify the id provided"
                )
            })
    }

    override fun loadBoardByOrderId(orderId: String): Mono<Board> {
        return boardRepository.findByOrderId(orderId).map { it.toBoard() }
            .switchIfEmpty(Mono.error {
                throw NotFoundException(
                    "Board with order id $orderId doesn't exit", "Verify the id provided"
                )
            })
    }

    private fun loadBoardEntity(boardId: UUID): Mono<BoardEntity> {
        return boardRepository.findById(boardId)
            .switchIfEmpty(Mono.error {
                throw NotFoundException(
                    "Board with id $boardId doesn't exit", "Verify the id provided"
                )
            })
    }

    private fun loadActiveBoardEntities(): Flux<BoardEntity> {
        return boardRepository.findBoardEntityByState(BoardState.RUNNING.toString())
    }
}
