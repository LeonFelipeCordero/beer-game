package com.beer.game.application.player

import com.beer.game.application.events.InternalEventListener
import com.beer.game.domain.Player
import mu.KotlinLogging
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux

private val logger = KotlinLogging.logger {}

@Component
class PlayerEvenListener(
    private val internalEventListener: InternalEventListener,
) {

    fun subscribe(playerId: String): Flux<Player> {
        return internalEventListener
            .subscribe()
            .filter { it.isSamePlayer(playerId) }
            .filter { it.isRelevantForPlayer() }
            .map { it.document as Player }
            .doOnError { logger.error("Something when wrong filtering the event", it) }
    }
}
