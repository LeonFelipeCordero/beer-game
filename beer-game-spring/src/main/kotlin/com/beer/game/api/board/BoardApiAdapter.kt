package com.beer.game.api.board

import com.beer.game.api.order.OrderGraph
import com.beer.game.api.player.PlayerGraph
import com.beer.game.application.board.BoardEvenListener
import com.beer.game.application.board.BoardService
import com.beer.game.application.order.OrderService
import com.beer.game.application.player.PlayerService
import com.beer.game.common.Role
import org.springframework.stereotype.Component
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.core.scheduler.Schedulers

@Component
class BoardApiAdapter(
    private val boardService: BoardService,
    private val playerService: PlayerService,
    private val orderService: OrderService,
    private val boardEvenListener: BoardEvenListener,
) {

    fun createBoard(name: String): Mono<BoardGraph> {
        return Mono.fromCallable {
            boardService.createBoard(name)
        }.map {
            BoardGraph.fromBoard(it)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getBoard(id: String): Mono<BoardGraph> {
        return Mono.fromCallable {
            boardService.loadBoard(id)
        }.map {
            BoardGraph.fromBoard(it)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getBoardByName(name: String): Mono<BoardGraph> {
        return Mono.fromCallable {
            boardService.loadBoardByName(name)
        }.map {
            BoardGraph.fromBoard(it)
        }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getPlayerForBoard(boardGraph: BoardGraph): Flux<PlayerGraph> {
        return Flux.fromIterable(playerService.getPlayersInBoard(boardGraph.id!!))
            .map {
                PlayerGraph.fromPlayer(it, boardGraph.id)
            }.subscribeOn(Schedulers.boundedElastic())
    }

    fun getOrdersForBoard(boardGraph: BoardGraph): Flux<OrderGraph> {
        return Flux.fromIterable(orderService.getOrdersByBoard(boardGraph.id!!))
            .map {
                OrderGraph.fromOrder(it, boardGraph.id, it.receiver)
            }.subscribeOn(Schedulers.boundedElastic())
    }

    fun subscribeToBoard(boardId: String): Flux<BoardGraph> {
        return boardEvenListener.subscribe(boardId)
            .map { BoardGraph.fromBoard(it) }
    }

    fun getAvailableRoles(boardGraph: BoardGraph): Flux<Role> {
        return Flux.fromIterable(boardService.loadBoard(boardGraph.id!!).availableRoles())
    }
}
