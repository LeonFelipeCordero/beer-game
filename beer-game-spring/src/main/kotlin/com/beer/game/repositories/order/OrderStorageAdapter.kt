package com.beer.game.repositories.order

import com.beer.game.application.order.OrderStorageAdapter
import com.beer.game.domain.Board
import com.beer.game.domain.Order
import com.beer.game.domain.Player
import com.beer.game.repositories.board.BoardMongoAdapter
import org.springframework.stereotype.Component

@Component
class OrderStorageAdapter(
    private val boardMongoAdapter: BoardMongoAdapter,
) : OrderStorageAdapter {

    override fun createOrder(board: Board, order: Order): Order {
        boardMongoAdapter.upsertBoard(board)
        return order
    }

    override fun deliverOrder(order: Order, board: Board) {
        val sender = board.players.first { it.id == order.sender }
        sender.orders.removeIf { it.id == order.id }
        sender.addOrder(order)

        val receiver: Player? = board.players.find { it.id == order.receiver }
        receiver?.let { receiver.orders.removeIf { it.id == order.id } }
        receiver?.let { receiver.addOrder(order) }

        boardMongoAdapter.upsertBoard(board)
    }

    override fun deliverFactoryBatch(board: Board) {
        boardMongoAdapter.upsertBoard(board)
    }
}
