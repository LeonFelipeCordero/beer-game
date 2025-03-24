package com.beer.game.repositories.player.postgres

import com.beer.game.common.Role
import com.beer.game.domain.Board
import com.beer.game.domain.Player
import com.beer.game.repositories.board.postgres.BoardEntity
import com.beer.game.repositories.order.postgres.OrderEntity
import org.springframework.data.annotation.CreatedDate
import org.springframework.data.annotation.Id
import org.springframework.data.annotation.LastModifiedDate
import org.springframework.data.annotation.PersistenceCreator
import org.springframework.data.annotation.Transient
import org.springframework.data.relational.core.mapping.Table
import java.time.LocalDateTime
import java.util.UUID

@Table(name = "player")
class PlayerEntity(
    @Id
    val playerId: UUID? = null,
    val name: String,
    val role: String,
    var stock: Int,
    var backlog: Int,
    var weeklyOrder: Int,
    var lastOrder: Int,
    val cpu: Boolean,
    @CreatedDate
    val createdAt: LocalDateTime = LocalDateTime.now(),
    @LastModifiedDate
    val updatedAt: LocalDateTime = LocalDateTime.now(),
    val boardId: UUID,
    @Transient
    val sendingOrders: MutableList<OrderEntity>? = mutableListOf(),
    @Transient
    val receivingOrders: MutableList<OrderEntity>? = mutableListOf(),
    @Transient
    val board: BoardEntity? = null
) {

    @PersistenceCreator
    constructor(
        playerId: UUID,
        name: String,
        role: String,
        stock: Int,
        backlog: Int,
        weeklyOrder: Int,
        lastOrder: Int,
        cpu: Boolean,
        boardId: UUID
    ) : this(
        playerId = playerId,
        name = name,
        role = role,
        stock = stock,
        backlog = backlog,
        weeklyOrder = weeklyOrder,
        lastOrder = lastOrder,
        cpu = cpu,
        boardId = boardId,
        sendingOrders = mutableListOf(),
        receivingOrders = mutableListOf(),
        board = null
    )

    companion object {
        fun fromPlayer(player: Player): PlayerEntity {
            return PlayerEntity(
                playerId = player.id?.let { UUID.fromString(player.id) },
                name = player.name,
                role = player.role.toString(),
                stock = player.stock,
                backlog = player.backlog,
                weeklyOrder = player.weeklyOrder,
                lastOrder = player.lastOrder,
                cpu = player.cpu,
                boardId = UUID.fromString(player.board.id)
            )
        }
    }

    fun toPlayer(board: Board): Player {
        return Player(
            id = playerId!!.toString(),
            name = name,
            role = Role.valueOf(role),
            stock = stock,
            backlog = backlog,
            weeklyOrder = weeklyOrder,
            lastOrder = lastOrder,
            cpu = cpu,
            board = board,
            orders = mutableListOf()// sendingOrders?.map { it.toOrder() }?.plus(receivingOrders?.map { it.toOrder() }).toMutableList(),
        )
    }
}
