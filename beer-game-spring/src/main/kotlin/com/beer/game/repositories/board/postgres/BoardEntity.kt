package com.beer.game.repositories.board.postgres

import com.beer.game.common.BoardState
import com.beer.game.domain.Board
import com.beer.game.repositories.player.postgres.PlayerEntity
import org.springframework.data.annotation.CreatedDate
import org.springframework.data.annotation.Id
import org.springframework.data.annotation.LastModifiedDate
import org.springframework.data.annotation.PersistenceConstructor
import org.springframework.data.annotation.PersistenceCreator
import org.springframework.data.annotation.Transient
import org.springframework.data.relational.core.mapping.Table
import java.time.LocalDateTime
import java.util.UUID

@Table(name = "board")
class BoardEntity(
    @Id
    val boardId: UUID? = null,
    val name: String,
    val state: String,
    val isFull: Boolean,
    val isFinished: Boolean,
    @CreatedDate
    val createdAt: LocalDateTime = LocalDateTime.now(),
    @LastModifiedDate
    val updatedAt: LocalDateTime = LocalDateTime.now(),
    @Transient
    var players: MutableList<PlayerEntity>? = mutableListOf()
) {

    @PersistenceCreator
    constructor(
        boardId: UUID?,
        name: String,
        state: String,
        isFull: Boolean,
        isFinished: Boolean,
        createdAt: LocalDateTime,
        updatedAt: LocalDateTime
    ) : this(
        boardId = boardId,
        name = name,
        state = state,
        isFull = isFull,
        isFinished = isFinished,
        createdAt = createdAt,
        updatedAt = updatedAt,
        players = mutableListOf()
    )

    constructor(name: String) : this(
        name = name,
        state = BoardState.CREATED.toString(),
        isFull = false,
        isFinished = false
    )

    companion object {
        fun fromBoard(board: Board): BoardEntity {
            return BoardEntity(
                boardId = UUID.fromString(board.id),
                name = board.name,
                state = board.state.toString(),
                isFull = board.isFull,
                isFinished = board.isFinished,
                createdAt = board.createdAt,
//                players = board.players.map { PlayerEntity.fromPlayer(it) }.toMutableList(),
            )
        }
    }

    fun toBoard(): Board {
        return Board(
            id = boardId!!.toString(),
            name = name,
            state = BoardState.valueOf(state),
            isFull = isFull,
            isFinished = isFinished,
            createdAt = createdAt,
            players = players?.map { it.toPlayer(this.toBoard()) }?.toMutableList() ?: mutableListOf(),
        )
    }
}
