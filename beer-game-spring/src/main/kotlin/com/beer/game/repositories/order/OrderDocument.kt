package com.beer.game.repositories.order

import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import com.beer.game.domain.Order
import java.time.LocalDateTime
import java.util.UUID

data class OrderDocument(
    val id: String = UUID.randomUUID().toString(),
    var amount: Int,
    val originalAmount: Int,
    var state: String = OrderState.PENDING.toString(),
    val type: String,
    val sender: String,
    val receiver: String? = null,
    val createAt: LocalDateTime,
) {

    companion object {
        fun fromOrder(order: Order): OrderDocument {
            return OrderDocument(
                id = order.id,
                amount = order.amount,
                originalAmount = order.originalAmount,
                state = order.state.toString(),
                type = order.type.toString(),
                sender = order.sender,
                receiver = order.receiver,
                createAt = order.createdAt,
            )
        }
    }

    fun toOrder(): Order {
        return Order(
            id = id,
            amount = amount,
            originalAmount = originalAmount,
            state = OrderState.valueOf(state),
            type = OrderType.valueOf(type),
            sender = sender,
            receiver = receiver,
            createdAt = createAt,
        )
    }
}
