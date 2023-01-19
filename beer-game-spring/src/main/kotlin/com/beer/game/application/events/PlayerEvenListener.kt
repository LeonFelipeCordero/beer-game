package com.beer.game.application.events

import com.beer.game.domain.Player
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux

@Component
class PlayerEvenListener(
    private val internalEventListener: InternalEventListener
) {

    companion object {
        val logger: Logger = LoggerFactory.getLogger(PlayerEvenListener::class.java)
    }

    fun subscribe(playerId: String): Flux<Player> {
        return internalEventListener
            .subscribe()
            .filter { it.isSamePlayer(playerId) }
            .filter { it.isRelevantForPlayer() }
            .map { it.document as Player }
            .doOnError { logger.error("Something when wrong filtering the event", it) }
    }
}