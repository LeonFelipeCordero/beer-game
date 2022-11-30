package com.beer.game.api

import com.beer.game.adapters.`in`.api.BoardApiAdapter
import com.beer.game.adapters.`in`.api.BoardGraph
import com.beer.game.adapters.`in`.api.OrderGraph
import com.beer.game.adapters.`in`.api.PlayerGraph
import org.springframework.graphql.data.method.annotation.Argument
import org.springframework.graphql.data.method.annotation.MutationMapping
import org.springframework.graphql.data.method.annotation.QueryMapping
import org.springframework.graphql.data.method.annotation.SchemaMapping
import org.springframework.graphql.data.method.annotation.SubscriptionMapping
import org.springframework.stereotype.Controller
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import java.time.Duration
import java.time.LocalTime

@Controller
class BoardController(
    private val boardApiAdapter: BoardApiAdapter
) {

    @MutationMapping
    fun createBoard(@Argument name: String): Mono<BoardGraph> {
        return boardApiAdapter.createBoard(name)
    }

    @QueryMapping
    fun getBoard(@Argument id: String): Mono<BoardGraph> {
        return boardApiAdapter.getBoard(id)
    }

    @SubscriptionMapping
    fun board(@Argument boardId: String): Flux<BoardGraph> {
        return boardApiAdapter.subscribeToBoard(boardId)
    }


    @SubscriptionMapping
    fun test(): Flux<String> {
        return Flux.interval(Duration.ofSeconds(1))
            .map { "Greetings" + LocalTime.now() }
    }


    @SchemaMapping(typeName = "Board", field = "players")
    fun players(boardGraph: BoardGraph): Flux<PlayerGraph> {
        return boardApiAdapter.getPlayerForBoard(boardGraph)
    }

    @SchemaMapping(typeName = "Board", field = "orders")
    fun orders(boardGraph: BoardGraph): Flux<OrderGraph> {
        return boardApiAdapter.getOrdersForBoard(boardGraph)
    }
}