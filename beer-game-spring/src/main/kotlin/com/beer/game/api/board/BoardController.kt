package com.beer.game.api.board

import com.beer.game.api.order.OrderGraph
import com.beer.game.api.player.PlayerGraph
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

    @QueryMapping
    fun getBoardByName(@Argument name: String): Mono<BoardGraph> {
        return boardApiAdapter.getBoardByName(name)
    }

    @SubscriptionMapping
    fun board(@Argument boardId: String): Flux<BoardGraph> {
        return boardApiAdapter.subscribeToBoard(boardId)
    }


    @SchemaMapping(typeName = "Board", field = "players")
    fun players(boardGraph: BoardGraph): Flux<PlayerGraph> {
        return boardApiAdapter.getPlayerForBoard(boardGraph)
    }

    @SchemaMapping(typeName = "Board", field = "orders")
    fun orders(boardGraph: BoardGraph): Flux<OrderGraph> {
        return boardApiAdapter.getOrdersForBoard(boardGraph)
    }

    @SchemaMapping(typeName = "Board", field = "availableRoles")
    fun availableRoles(boardGraph: BoardGraph): Flux<Role> {
        return boardApiAdapter.getAvailableRoles(boardGraph)
    }
}