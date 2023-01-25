package com.beer.game.application.board

import com.beer.game.application.events.InternalEventListener
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
            .filter { it.isSameBoard(boardId) }
            .filter { it.isRelevantForBoard() }
            .map { it.document as Board }
            .doOnError { logger.error("Something when wrong filtering the event", it) }
    }

}