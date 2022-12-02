package com.beer.game.events

import com.beer.game.domain.Order
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux

@Component
class OrderEvenListener(
    private val internalEventListener: InternalEventListener
) {

    companion object {
        val logger: Logger = LoggerFactory.getLogger(OrderEvenListener::class.java)
    }

    fun subscribeNewOrder(boardId: String, playerId: String): Flux<Order> {
        return internalEventListener
            .subscribe()
            .filter { it.isSameBoard(boardId) }
            .filter { it.isRelevantForNewOrder() }
            .map { it.document as Order }
            .filter { it.isPlayerInvolved(playerId) }
            .doOnError { logger.error("Something when wrong filtering the event", it) }
    }

    fun subscribeUpdateDelivery(boardId: String, playerId: String): Flux<Order> {
        return internalEventListener
            .subscribe()
            .filter { it.isSameBoard(boardId) }
            .filter { it.isRelevantForUpdateOrder() }
            .map { it.document as Order }
            .filter { it.isPlayerInvolved(playerId) }
            .doOnError { logger.error("Something when wrong filtering the event", it) }
    }
}