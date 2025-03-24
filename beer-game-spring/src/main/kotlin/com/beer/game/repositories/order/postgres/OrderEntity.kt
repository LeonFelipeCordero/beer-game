package com.beer.game.repositories.order.postgres

import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import com.beer.game.domain.Order
import com.beer.game.repositories.player.postgres.PlayerEntity
import org.springframework.data.annotation.CreatedDate
import org.springframework.data.annotation.Id
import org.springframework.data.annotation.LastModifiedDate
import org.springframework.data.relational.core.mapping.Table
import java.time.LocalDateTime
import java.util.UUID

@Table(name = "orders")
class OrderEntity(
    @Id
    val orderId: UUID? = null,
    var amount: Int,
    val originalAmount: Int,
    var state: String = OrderState.PENDING.toString(),
    val type: String,
    @CreatedDate
    val createAt: LocalDateTime = LocalDateTime.now(),
    @LastModifiedDate
    val updatedAt: LocalDateTime = LocalDateTime.now(),
    @Transient
    val sender: PlayerEntity? = null,
    @Transient
    val receiver: PlayerEntity? = null,
) {

    companion object {
        fun fromOrder(order: Order): OrderEntity {
            return OrderEntity(
                orderId = UUID.fromString(order.id),
                amount = order.amount,
                originalAmount = order.originalAmount,
                state = order.state.toString(),
                type = order.type.toString(),
//                sender = order.sender,
//                receiver = order.receiver,
                createAt = order.createdAt,
            )
        }
    }

    fun toOrder(): Order {
        return Order(
            id = orderId!!.toString(),
            amount = amount,
            originalAmount = originalAmount,
            state = OrderState.valueOf(state),
            type = OrderType.valueOf(type),
            sender = "",
            receiver = "",
            createdAt = createAt,
        )
    }
}
