package com.beer.game.api.player

import com.beer.game.api.board.BoardGraph
import com.beer.game.api.domain.Response
import com.beer.game.api.order.OrderGraph
import com.beer.game.common.Role
import org.springframework.graphql.data.method.annotation.Argument
import org.springframework.graphql.data.method.annotation.MutationMapping
import org.springframework.graphql.data.method.annotation.QueryMapping
import org.springframework.graphql.data.method.annotation.SchemaMapping
import org.springframework.graphql.data.method.annotation.SubscriptionMapping
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
        @Argument playerId: String,
        @Argument amount: Int,
    ): Mono<Response> {
        return playerApiAdapter.updateWeeklyOrder(playerId, amount)
    }

    @QueryMapping
    fun getPlayer(@Argument playerId: String): Mono<PlayerGraph> {
        return playerApiAdapter.getPlayer(playerId)
    }

    @QueryMapping
    fun getPlayersByBoard(@Argument boardId: String): Flux<PlayerGraph> {
        return playerApiAdapter.getPlayersByBoard(boardId)
    }

    @SubscriptionMapping
    fun player(@Argument playerId: String): Flux<PlayerGraph> {
        return playerApiAdapter.subscribeToPlayer(playerId)
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
