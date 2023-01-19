package com.beer.game.api

import com.beer.game.adapters.`in`.api.*
import org.springframework.graphql.data.method.annotation.Argument
import org.springframework.graphql.data.method.annotation.MutationMapping
import org.springframework.graphql.data.method.annotation.SchemaMapping
import org.springframework.graphql.data.method.annotation.SubscriptionMapping
import org.springframework.stereotype.Controller
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono

@Controller
class OrderController(
    private val orderApiAdapter: OrderApiAdapter,
    private val playerApiAdapter: PlayerApiAdapter,
    private val boardApiAdapter: BoardApiAdapter
) {

    @MutationMapping
    fun createOrder(
        @Argument receiverId: String
    ): Mono<OrderGraph> {
        return orderApiAdapter.createOrder(receiverId)
    }

    @MutationMapping
    fun deliverOrder(
        @Argument orderId: String,
        @Argument amount: Int? = null
    ): Mono<Response> {
        return orderApiAdapter.deliverOrder(orderId, amount)
    }

    @SubscriptionMapping
    fun newOrder(@Argument playerId: String): Flux<OrderGraph> {
        return orderApiAdapter.newOrderSubscription(playerId)
    }

    @SubscriptionMapping
    fun orderDelivery(@Argument playerId: String): Flux<OrderGraph> {
        return orderApiAdapter.orderDeliverySubscription(playerId)
    }

    @SchemaMapping(typeName = "Order", field = "sender")
    fun sender(order: OrderGraph): Mono<PlayerGraph> {
        return playerApiAdapter.getPlayer(order.senderId!!)
    }

    @SchemaMapping(typeName = "Order", field = "receiver")
    fun receiver(order: OrderGraph): Mono<PlayerGraph> {
        return order.receiverId?.let { playerApiAdapter.getPlayer(order.receiverId) } ?: Mono.empty()
    }

    @SchemaMapping(typeName = "Order", field = "board")
    fun board(order: OrderGraph): Mono<BoardGraph> {
        return boardApiAdapter.getBoard(order.boardId!!)
    }
}

