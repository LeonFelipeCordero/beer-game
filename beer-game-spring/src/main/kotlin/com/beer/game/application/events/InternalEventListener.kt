package com.beer.game.application.events

import com.beer.game.events.Event
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux
import reactor.core.publisher.Sinks
import java.lang.RuntimeException

@Component
class InternalEventListener : EventEmitter<Event> {

    companion object {
        val logger: Logger = LoggerFactory.getLogger(InternalEventListener::class.java)
    }

    private val sink = Sinks.many().multicast().onBackpressureBuffer<Event>()

    override fun publish(event: Event) {
        logger.info("Receive event type ${event.eventType.name} for document type ${event.documentType.name}")
        sink.tryEmitNext(event)
    }

    override fun subscribe(): Flux<Event> {
        return sink.asFlux()
            .cache()
            .doOnComplete { logger.info("Stream completed") }
            .doOnError { logger.error("Something when wrong with the stream", it) }
    }
}
