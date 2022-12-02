package com.beer.game.application.service

import com.beer.game.adapters.out.mongo.OrderMongoAdapter
import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import com.beer.game.common.Role
import com.beer.game.domain.*
import com.beer.game.domain.exceptions.ImpossibleActionException
import com.beer.game.events.InternalEventListener
import org.springframework.stereotype.Component
import java.time.LocalDateTime

@Component
class OrderService(
    private val orderMongoAdapter: OrderMongoAdapter,
    private val boardService: BoardService,
    private val playerService: PlayerService,
    private val internalEventListener: InternalEventListener
) {

    fun createCpuOrders() {
        boardService.loadActiveBoards()
            .forEach { createCpuOrderForBoard(it) }
    }

    fun createOrder(boardId: String, senderId: String, receiverId: String): Order {
        val board = boardService.loadBoard(boardId)
        val sender = board.findPlayer(senderId)
        val receiver = board.findPlayer(receiverId)
        val order = Order(
            amount = receiver.weeklyOrder,
            originalAmount = receiver.weeklyOrder,
            type = OrderType.PLAYER_ORDER,
            sender = sender.id,
            receiver = receiver.id,
            createdAt = LocalDateTime.now()
        )
        sender.addOrder(order)
        receiver.addOrder(order)
        order.emitCreation(internalEventListener, board.id)
        return orderMongoAdapter.createOrder(board, order)
    }

    fun deliverOrder(orderId: String, boardId: String, amount: Int? = null) {
        val board = boardService.loadBoard(boardId)
        val order = board.findOrder(orderId)
        amount?.let { order.amount = it }
        val sender = board.players.first { order.sender == it.id }
        val receiver = order.receiver?.let { board.players.first { order.receiver == it.id } }

        if (!order.validOderAmount()) {
            throw ImpossibleActionException(
                "deliver amount can't be bigger than original amount",
                "verify order amount"
            )
        }
        if (!sender.hasEnoughStock(order.amount)) {
            sender.emitUpdate(internalEventListener, board.id)
            throw ImpossibleActionException(
                "Sender doesn't have enough stock to deliver the order",
                "wait until the sender has enough stock, you can order more"
            )
        }

        sender.stock -= order.amount
        receiver?.let {
            it.stock += order.amount
            it.lastOrder = order.amount
        }
        order.state = OrderState.DELIVERED

        order.emitUpdate(internalEventListener, board.id)
        sender.emitUpdate(internalEventListener, board.id)
        receiver?.apply { emitUpdate(internalEventListener, board.id) }

        orderMongoAdapter.deliverOrder(order, board)
    }

    fun deliverFactoryBatches() {
        boardService.loadActiveBoards().forEach { deliverFactoryBatch(it) }
    }

    fun getOrder(orderId: String, boardId: String): Order {
        return boardService.loadBoard(boardId).findOrder(orderId)
    }

    fun getOrdersByBoard(boardId: String): List<Order> {
        return boardService.loadBoard(boardId)
            .players
            .flatMap { it.orders }
    }

    fun getOrdersByPlayer(playerId: String, boardId: String): List<Order> {
        return playerService.getPlayer(boardId, playerId).orders
    }

    private fun createCpuOrderForBoard(board: Board) {
        board.players
            .forEach { player ->
                val order = Order(
                    amount = player.weeklyOrder / 4,
                    originalAmount = player.weeklyOrder / 4,
                    type = OrderType.CPU_ORDER,
                    sender = player.id,
                    createdAt = LocalDateTime.now()
                )
                player.addOrder(order)
                order.emitCreation(internalEventListener, board.id)
                orderMongoAdapter.createOrder(board, order)
            }
    }

    private fun deliverFactoryBatch(board: Board) {
        board
            .players
            .first { it.role == Role.FACTORY }
            .let { factory ->
                factory.stock += factory.weeklyOrder
                factory.backlog = factory.weeklyOrder
                factory.lastOrder = factory.weeklyOrder
            }
        orderMongoAdapter.deliverFactoryBatch(board)
    }
}