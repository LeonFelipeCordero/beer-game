package com.beer.game.application.order

import com.beer.game.application.board.BoardService
import com.beer.game.repositories.order.OrderStorageAdapter
import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import com.beer.game.common.Role
import com.beer.game.domain.*
import com.beer.game.domain.exceptions.ImpossibleActionException
import com.beer.game.application.events.InternalEventListener
import com.beer.game.application.player.PlayerService
import org.springframework.stereotype.Component
import java.time.LocalDateTime

@Component
class OrderService(
    private val orderStorageAdapter: OrderStorageAdapter,
    private val boardService: BoardService,
    private val playerService: PlayerService,
    private val internalEventListener: InternalEventListener
) {

    fun createCpuOrders() {
        boardService.loadActiveBoards()
            .forEach { createCpuOrderForBoard(it) }
    }

    fun createOrder(receiverId: String): Pair<Order, String> {
        val board = boardService.loadBoardByPlayerId(receiverId)
        val receiver = board.findPlayer(receiverId)
        val sender = board.findContraPart(receiver)
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
        return Pair(orderStorageAdapter.createOrder(board, order), board.id)
    }

    fun deliverOrder(orderId: String, amount: Int? = null) {
        val board = boardService.loadBoardByOrderId(orderId)
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

        // todo last order?
        sender.stock -= order.amount
        receiver?.let {
            it.stock += order.amount
            it.lastOrder = order.amount
        }
        order.state = OrderState.DELIVERED
        amount?.let { order.amount = amount }

        orderStorageAdapter.deliverOrder(order, board)
        order.emitUpdate(internalEventListener, board.id)
        sender.emitUpdate(internalEventListener, board.id)
        receiver?.apply { emitUpdate(internalEventListener, board.id) }
    }

    fun deliverFactoryBatches() {
        boardService.loadActiveBoards().forEach { deliverFactoryBatch(it) }
    }

    fun getOrder(orderId: String): Pair<Order, String> {
        val board = boardService.loadBoardByOrderId(orderId)
        val order = board.findOrder(orderId)
        return Pair(order, board.id)
    }

    fun getOrdersByBoard(boardId: String): List<Order> {
        return boardService.loadBoard(boardId)
            .players
            .flatMap { it.orders }
    }

    fun getOrdersByPlayer(playerId: String): List<Order> {
        val (player, _) = playerService.getPlayer(playerId)
        return player.orders
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
                orderStorageAdapter.createOrder(board, order)
                order.emitCreation(internalEventListener, board.id)
            }
    }

    private fun deliverFactoryBatch(board: Board) {
        val factory = board
            .players
            .first { it.role == Role.FACTORY }
        factory.stock += factory.weeklyOrder
        factory.backlog += factory.weeklyOrder
        factory.lastOrder = factory.weeklyOrder
        orderStorageAdapter.deliverFactoryBatch(board)
        factory.emitUpdate(internalEventListener, board.id)
    }
}