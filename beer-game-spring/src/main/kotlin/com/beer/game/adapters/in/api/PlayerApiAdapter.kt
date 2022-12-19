package com.beer.game.adapters.`in`.api

import com.beer.game.application.service.BoardService
import com.beer.game.application.service.OrderService
import com.beer.game.application.service.PlayerService
import com.beer.game.common.Role
import com.beer.game.application.events.PlayerEvenListener
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.core.scheduler.Schedulers

@Component
class PlayerApiAdapter(
    private val playerService: PlayerService,
    private val boardService: BoardService,
    private val orderService: OrderService,
    private val playerEvenListener: PlayerEvenListener
) {

    fun addPlayer(boardId: String, role: Role): Mono<PlayerGraph> {
        return Mono.fromCallable {
            playerService.addPlayer(boardId, role)
        }.map {
            PlayerGraph.fromPlayer(it, boardId)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun updateWeeklyOrder(boardId: String, playerId: String, amount: Int): Mono<Response> {
        return Mono.fromCallable {
            playerService.changeWeeklyOrder(boardId, playerId, amount)
        }.map {
            Response(
                message = "weekly order updated",
                status = 200
            )
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getPlayer(boardId: String, playerId: String): Mono<PlayerGraph> {
        return Mono.fromCallable {
            playerService.getPlayer(boardId, playerId)
        }.map {
            PlayerGraph.fromPlayer(it, boardId)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getPlayersByOrder(boardId: String): Flux<PlayerGraph> {
        return Flux.fromIterable(
            playerService.getPlayersInBoard(boardId)
        ).map {
            PlayerGraph.fromPlayer(it, boardId)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun board(playerGraph: PlayerGraph): Mono<BoardGraph> {
        return Mono.fromCallable {
            val board = boardService.loadBoard(playerGraph.boardId!!)
            BoardGraph.fromBoard(board)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun orders(playerGraph: PlayerGraph): Flux<OrderGraph> {
        return Flux.fromIterable(orderService.getOrdersByPlayer(playerGraph.id!!, playerGraph.boardId!!))
            .map {
                OrderGraph.fromOrder(it, playerGraph.boardId, it.sender, it.receiver)
            }.subscribeOn(Schedulers.boundedElastic())
    }

    fun subscribeToPlayer(boardId: String, playerId: String): Flux<PlayerGraph> {
        return playerEvenListener.subscribe(boardId, playerId)
            .map { PlayerGraph.fromPlayer(it, boardId) }
    }
}