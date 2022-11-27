package com.beer.game.api

import com.beer.game.adapters.`in`.api.*
import com.beer.game.common.Role
import org.springframework.graphql.data.method.annotation.Argument
import org.springframework.graphql.data.method.annotation.MutationMapping
import org.springframework.graphql.data.method.annotation.QueryMapping
import org.springframework.graphql.data.method.annotation.SchemaMapping
import org.springframework.stereotype.Controller
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono

@Controller
class PlayerController(
    private val playerApiAdapter: PlayerApiAdapter,
) {

    @MutationMapping
    fun addPlayer(@Argument boardId: String, @Argument role: Role): Mono<PlayerGraph> {
        return playerApiAdapter.addPlayer(boardId, role)
    }

    @MutationMapping
    fun updateWeeklyOrder(
        @Argument boardId: String,
        @Argument playerId: String,
        @Argument amount: Int
    ): Mono<Response> {
        return playerApiAdapter.updateWeeklyOrder(boardId, playerId, amount)
    }

    @QueryMapping
    fun getPlayer(@Argument boardId: String, @Argument playerId: String): Mono<PlayerGraph> {
        return playerApiAdapter.getPlayer(boardId, playerId)
    }

    @QueryMapping
    fun getPlayersByOrder(@Argument boardId: String): Flux<PlayerGraph> {
        return playerApiAdapter.getPlayersByOrder(boardId)
    }

    @SchemaMapping(typeName = "Player", field = "board")
    fun board(playerGraph: PlayerGraph): Mono<BoardGraph> {
        return playerApiAdapter.board(playerGraph)
    }

    @SchemaMapping(typeName = "Player", field = "orders")
    fun orders(playerGraph: PlayerGraph): Flux<OrderGraph> {
        return playerApiAdapter.orders(playerGraph)
    }
}