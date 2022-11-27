package com.beer.game.adapters.out.mongo

import com.beer.game.domain.*
import org.springframework.stereotype.Component

@Component
class OrderMongoAdapter(
    private val boardMongoAdapter: BoardMongoAdapter
) {

    fun loadOrder(orderId: String, board: Board): Order {
        return BoardDocument.fromBoard(board)
            .orders
            .first { it.id == orderId }
            .toOrder()
    }

    fun createOrder(board: Board, order: Order): Order {
        val boardDocument = BoardDocument.fromBoard(board)
        boardMongoAdapter.upsertBoard(boardDocument)
        return order
    }

    fun deliverOrder(order: Order, board: Board) {
        val boardDocument = BoardDocument.fromBoard(board)
        boardMongoAdapter.upsertBoard(boardDocument)
    }

    fun deliverFactoryBatch(board: Board) {
        val boardDocument = BoardDocument.fromBoard(board)
        boardMongoAdapter.upsertBoard(boardDocument)
    }
}