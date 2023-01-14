package com.beer.game.adapters.`in`.api

import com.beer.game.application.service.OrderService
import com.beer.game.application.events.OrderEvenListener
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.core.scheduler.Schedulers

@Component
class OrderApiAdapter(
    private val orderService: OrderService,
    private val orderEvenListener: OrderEvenListener
) {

    fun createOrder(
        boardId: String,
        receiverId: String
    ): Mono<OrderGraph> {
        return Mono.fromCallable {
            orderService.createOrder(boardId, receiverId)
        }.map {
            OrderGraph.fromOrder(it, boardId, receiverId)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun deliverOrder(
        orderId: String,
        boardId: String,
        amount: Int? = null
    ): Mono<Response> {
        return Mono.fromCallable {
            orderService.deliverOrder(orderId, boardId, amount)
        }.map {
            Response(message = "Order delivered", status = 200)
        }
    }

    fun newOrderSubscription(boardId: String, playerId: String): Flux<OrderGraph> {
        return orderEvenListener.subscribeNewOrder(boardId, playerId)
            .map { OrderGraph.fromOrder(it, boardId, it.receiver) }
    }

    fun orderDeliverySubscription(boardId: String, playerId: String): Flux<OrderGraph> {
        return orderEvenListener.subscribeUpdateDelivery(boardId, playerId)
            .map { OrderGraph.fromOrder(it, boardId, it.receiver) }
    }
}
