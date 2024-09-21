package com.beer.game.api.player

import com.beer.game.api.board.BoardGraph
import com.beer.game.api.domain.Response
import com.beer.game.api.order.OrderGraph
import com.beer.game.application.board.BoardService
import com.beer.game.application.order.OrderService
import com.beer.game.application.player.PlayerEvenListener
import com.beer.game.application.player.PlayerService
import com.beer.game.common.Role
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.core.scheduler.Schedulers

@Component
class PlayerApiAdapter(
    private val playerService: PlayerService,
    private val boardService: BoardService,
    private val orderService: OrderService,
    private val playerEvenListener: PlayerEvenListener,
) {

    fun addPlayer(boardId: String, role: Role): Mono<PlayerGraph> {
        return Mono.fromCallable {
            playerService.addPlayer(boardId, role)
        }.map {
            PlayerGraph.fromPlayer(it, boardId)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun updateWeeklyOrder(playerId: String, amount: Int): Mono<Response> {
        return Mono.fromCallable {
            playerService.changeWeeklyOrder(playerId, amount)
        }.map {
            Response(
                message = "weekly order updated",
                status = 200,
            )
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getPlayer(playerId: String): Mono<PlayerGraph> {
        return Mono.fromCallable {
            playerService.getPlayer(playerId)
        }.map {
            PlayerGraph.fromPlayer(it.first, it.second)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getPlayersByBoard(boardId: String): Flux<PlayerGraph> {
        return Flux.fromIterable(
            playerService.getPlayersInBoard(boardId),
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
        return Flux.fromIterable(orderService.getOrdersByPlayer(playerGraph.id!!))
            .map {
                OrderGraph.fromOrder(it, playerGraph.boardId!!, it.receiver)
            }.subscribeOn(Schedulers.boundedElastic())
    }

    fun subscribeToPlayer(playerId: String): Flux<PlayerGraph> {
        val board = boardService.loadBoardByPlayerId(playerId)
        return playerEvenListener.subscribe(playerId)
            .map { PlayerGraph.fromPlayer(it, board.id) }
    }
}
