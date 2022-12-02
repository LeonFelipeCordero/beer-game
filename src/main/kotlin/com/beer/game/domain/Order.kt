package com.beer.game.domain

import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import com.beer.game.events.DocumentType
import com.beer.game.events.Event
import com.beer.game.events.EventType
import com.beer.game.events.InternalEventListener
import java.time.LocalDateTime
import java.util.UUID

data class Order(
    val id: String = UUID.randomUUID().toString(),
    var amount: Int,
    val originalAmount: Int,
    var state: OrderState = OrderState.PENDING,
    val type: OrderType,
    val sender: String,
    val receiver: String? = null,
    val createdAt: LocalDateTime
) {
    fun isPlayerInvolved(playerId: String): Boolean {
        return sender == playerId || receiver == playerId
    }

    fun validOderAmount(): Boolean {
        return amount <= originalAmount
    }

    fun emitCreation(listener: InternalEventListener, boardId: String) {
        listener.publish(
            Event(
                document = this,
                documentId = boardId,
                entityId = id,
                documentType = DocumentType.ORDER,
                eventType = EventType.NEW
            )
        )
    }

    fun emitUpdate(listener: InternalEventListener, boardId: String) {
        listener.publish(
            Event(
                document = this,
                documentId = boardId,
                entityId = id,
                documentType = DocumentType.ORDER,
                eventType = EventType.UPDATE
            )
        )
    }
}
