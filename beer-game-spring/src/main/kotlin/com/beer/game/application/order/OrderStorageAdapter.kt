package com.beer.game.application.order

import com.beer.game.domain.Board
import com.beer.game.domain.Order

interface OrderStorageAdapter {
    fun createOrder(board: Board, order: Order): Order
    fun deliverOrder(order: Order, board: Board)
    fun deliverFactoryBatch(board: Board)
}
