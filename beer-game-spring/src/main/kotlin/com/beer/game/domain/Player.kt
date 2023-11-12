package com.beer.game.domain

import com.beer.game.common.Role
import com.beer.game.events.DocumentType
import com.beer.game.events.Event
import com.beer.game.events.EventType
import com.beer.game.application.events.InternalEventListener
import java.util.UUID

data class Player(
    val id: String = UUID.randomUUID().toString(),
    val name: String,
    val role: Role,
    var stock: Int,
    var backlog: Int,
    var weeklyOrder: Int,
    var lastOrder: Int,
    val cpu: Boolean,
    val orders: MutableList<Order>
) {
    companion object {
        fun retailPlayer(): Player {
            return Player(
                name = Role.RETAILER.toString(),
                role = Role.RETAILER,
                stock = 80,
                backlog = 80,
                weeklyOrder = 40,
                lastOrder = 40,
                cpu = false,
                orders = mutableListOf()
            )
        }

        fun wholesalerPlayer(): Player {
            return Player(
                name = Role.WHOLESALER.toString(),
                role = Role.WHOLESALER,
                stock = 1200,
                backlog = 1200,
                weeklyOrder = 600,
                lastOrder = 600,
                cpu = false,
                orders = mutableListOf()
            )
        }

        fun factoryPlayer(): Player {
            return Player(
                name = Role.FACTORY.name,
                role = Role.FACTORY,
                stock = 12000,
                backlog = 12000,
                weeklyOrder = 6000,
                lastOrder = 6000,
                cpu = false,
                orders = mutableListOf()
            )
        }
    }

    fun addOrder(order: Order) {
        orders.add(order)
    }

    fun hasEnoughStock(amount: Int): Boolean {
        return stock >= amount
    }

    fun emitCreation(listener: InternalEventListener, board: Board) {
        listener.publish(
            Event(
                document = board,
                documentId = board.id,
                entityId = id,
                documentType = DocumentType.PLAYER,
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
                documentType = DocumentType.PLAYER,
                eventType = EventType.UPDATE
            )
        )
    }
}
