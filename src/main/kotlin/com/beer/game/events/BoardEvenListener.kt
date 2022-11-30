package com.beer.game.events

import com.beer.game.domain.Board
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux

@Component
class BoardEvenListener(
    private val internalEventListener: InternalEventListener
) {

    companion object {
        val logger: Logger = LoggerFactory.getLogger(BoardEvenListener::class.java)
    }

    fun subscribe(boardId: String): Flux<Board> {
        return internalEventListener
            .subscribe()
            .filter { it.documentId == boardId }
            .filter {
                (it.documentType == DocumentType.BOARD && it.eventType == EventType.UPDATE) ||
                        (it.documentType == DocumentType.PLAYER && it.eventType == EventType.NEW)
            }
            .map { it.document as Board }
            .doOnError {
                logger.error("Something when wrong filtering the event", it)
            }
    }
}