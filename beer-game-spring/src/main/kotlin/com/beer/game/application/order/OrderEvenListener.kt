package com.beer.game.application.order

import com.beer.game.application.events.InternalEventListener
import com.beer.game.domain.Order
import mu.KotlinLogging
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux

private val logger = KotlinLogging.logger {}

@Component
class OrderEvenListener(
    private val internalEventListener: InternalEventListener,
) {

    fun subscribeNewOrder(playerId: String): Flux<Order> {
        return internalEventListener
            .subscribe()
            .filter { it.isRelevantForNewOrder() }
            .map { it.document as Order }
            .filter { it.isPlayerInvolved(playerId) }
            .doOnError { logger.error("Something when wrong filtering the event", it) }
    }

    fun subscribeUpdateDelivery(playerId: String): Flux<Order> {
        return internalEventListener
            .subscribe()
            .filter { it.isRelevantForUpdateOrder() }
            .map { it.document as Order }
            .filter { it.isPlayerInvolved(playerId) }
            .doOnError { logger.error("Something when wrong filtering the event", it) }
    }
}
