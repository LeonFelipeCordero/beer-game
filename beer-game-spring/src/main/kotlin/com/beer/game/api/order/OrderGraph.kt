package com.beer.game.api.order

import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import com.beer.game.domain.Order
import java.time.LocalDateTime

data class OrderGraph(
    val id: String?,
    var amount: Int?,
    val originalAmount: Int?,
    var state: OrderState?,
    val type: OrderType?,
    val senderId: String?,
    val receiverId: String?,
    val boardId: String?,
    val createdAt: LocalDateTime?,
) {
    companion object {
        fun fromOrder(order: Order, boardId: String, receiverId: String? = null): OrderGraph {
            return OrderGraph(
                id = order.id,
                amount = order.amount,
                originalAmount = order.originalAmount,
                state = order.state,
                type = order.type,
                senderId = order.sender,
                receiverId = receiverId,
                boardId = boardId,
                createdAt = order.createdAt,
            )
        }
    }
}
