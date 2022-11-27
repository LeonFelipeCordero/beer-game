package com.beer.game.api

import com.beer.game.adapters.`in`.api.BoardApiAdapter
import com.beer.game.adapters.`in`.api.BoardGraph
import com.beer.game.adapters.`in`.api.OrderGraph
import com.beer.game.adapters.`in`.api.PlayerGraph
import org.springframework.graphql.data.method.annotation.Argument
import org.springframework.graphql.data.method.annotation.MutationMapping
import org.springframework.graphql.data.method.annotation.QueryMapping
import org.springframework.graphql.data.method.annotation.SchemaMapping
import org.springframework.stereotype.Controller
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono

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

    @SchemaMapping(typeName = "Board", field = "players")
    fun players(boardGraph: BoardGraph): Flux<PlayerGraph> {
        return boardApiAdapter.getPlayerForBoard(boardGraph)
    }

    @SchemaMapping(typeName = "Board", field = "orders")
    fun orders(boardGraph: BoardGraph): Flux<OrderGraph> {
        return boardApiAdapter.getOrdersForBoard(boardGraph)
    }
}