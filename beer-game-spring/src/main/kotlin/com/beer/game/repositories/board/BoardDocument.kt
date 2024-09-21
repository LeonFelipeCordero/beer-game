package com.beer.game.repositories.board

import com.beer.game.common.BoardState
import com.beer.game.domain.Board
import com.beer.game.repositories.order.OrderDocument
import com.beer.game.repositories.player.PlayerDocument
import org.bson.types.ObjectId
import org.springframework.data.annotation.Id
import org.springframework.data.mongodb.core.mapping.Document
import java.time.LocalDateTime

@Document("board")
data class BoardDocument(
    @Id
    val id: ObjectId = ObjectId(),
    val name: String,
    var state: String = BoardState.CREATED.toString(),
    var full: Boolean = false,
    var finished: Boolean = false,
    val createdAt: LocalDateTime = LocalDateTime.now(),
    val updatedAt: LocalDateTime = LocalDateTime.now(),
    val players: MutableList<PlayerDocument> = mutableListOf(),
    val orders: MutableSet<OrderDocument> = mutableSetOf(),
) {

    companion object {
        fun fromBoard(board: Board): BoardDocument {
            return BoardDocument(
                id = ObjectId(board.id),
                name = board.name,
                state = board.state.toString(),
                full = board.full,
                finished = board.finished,
                createdAt = board.createdAt,
                players = board.players.map { PlayerDocument.fromPlayer(it) }.toMutableList(),
                orders = board.players.flatMap { it.orders }.map { OrderDocument.fromOrder(it) }.toMutableSet(),
            )
        }
    }

    fun toBoard(): Board {
        return Board(
            id = id.toString(),
            name = name,
            state = BoardState.valueOf(state),
            full = full,
            finished = finished,
            createdAt = createdAt,
            players = players.map { playerDoc ->
                val playerOrders =
                    orders.filter { orderDoc ->
                        orderDoc.sender == playerDoc.id || orderDoc.receiver == playerDoc.id
                    }
                playerDoc.toPlayer(playerOrders)
            }.toMutableList(),
        )
    }
}
