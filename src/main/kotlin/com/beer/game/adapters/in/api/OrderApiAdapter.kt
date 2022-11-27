package com.beer.game.adapters.`in`.api

import com.beer.game.application.service.OrderService
import com.beer.game.application.service.PlayerService
import org.springframework.stereotype.Component
import reactor.core.publisher.Mono
import reactor.core.scheduler.Schedulers

@Component
class OrderApiAdapter(
    private val orderService: OrderService,
    private val playerService: PlayerService
) {

    fun createOrder(
        boardId: String,
        senderId: String,
        receiverId: String
    ): Mono<OrderGraph> {
        return Mono.fromCallable {
            orderService.createOrder(boardId, senderId, receiverId)
        }.map {
            OrderGraph.fromOrder(it, boardId, senderId, receiverId)
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
}
