package com.beer.game.application.service.order

import com.beer.game.IntegrationTestBase
import com.beer.game.common.OrderState
import com.beer.game.common.Role
import com.beer.game.domain.exceptions.ImpossibleActionException
import org.assertj.core.api.Assertions.assertThat
import org.assertj.core.api.Assertions.assertThatExceptionOfType
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test

internal class OrderServiceTest : IntegrationTestBase() {

    @AfterEach
    fun setup() {
        boardRepository.deleteAll()
    }

    @Test
    fun `should create cpu orders for all players in a board`() {
        var board = boardService.createBoard(boardName)
        playerService.addPlayer(board.id, Role.RETAILER)
        playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)

        orderService.createCpuOrders()

        board = boardService.loadBoard(board.id)

        assertThat(board.players[0].orders).hasSize(1)
        assertThat(board.players[1].orders).hasSize(1)
        assertThat(board.players[2].orders).hasSize(1)

        assertThat(board.players[0].orders[0].amount).isEqualTo(10)
        assertThat(board.players[1].orders[0].amount).isEqualTo(150)
        assertThat(board.players[2].orders[0].amount).isEqualTo(1500)

        orderService.createCpuOrders()

        board = boardService.loadBoard(board.id)

        assertThat(board.players[0].orders).hasSize(2)
        assertThat(board.players[1].orders).hasSize(2)
        assertThat(board.players[2].orders).hasSize(2)
    }

    @Test
    fun `should deliver factory batch`() {
        var board = boardService.createBoard(boardName)
        playerService.addPlayer(board.id, Role.RETAILER)
        playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)

        orderService.deliverFactoryBatches()

        board = boardService.loadBoard(board.id)

        assertThat(board.players[2].stock).isEqualTo(18000)
    }

    @Test
    fun `should deliver cpu orders`() {
        var board = boardService.createBoard(boardName)
        playerService.addPlayer(board.id, Role.RETAILER)
        playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)

        orderService.createCpuOrders()

        board = boardService.loadBoard(board.id)
        board.players
            .flatMap { it.orders }
            .forEach { order ->
                orderService.deliverOrder(order.id, order.amount)
            }

        board = boardService.loadBoard(board.id)
        assertThat(board.players[0].stock).isEqualTo(70)
        assertThat(board.players[1].stock).isEqualTo(1050)
        assertThat(board.players[2].stock).isEqualTo(10500)
        assertThat(board.players[0].orders[0].state).isEqualTo(OrderState.DELIVERED)
        assertThat(board.players[1].orders[0].state).isEqualTo(OrderState.DELIVERED)
        assertThat(board.players[2].orders[0].state).isEqualTo(OrderState.DELIVERED)
    }

    @Test
    fun `should fail if order to deliver amount is bigger than original amount`() {
        var board = boardService.createBoard(boardName)
        playerService.addPlayer(board.id, Role.RETAILER)
        playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)

        orderService.createCpuOrders()

        board = boardService.loadBoard(board.id)
        val orderToDeliver = board.players
            .flatMap { it.orders }
            .first()

        assertThatExceptionOfType(ImpossibleActionException::class.java)
            .isThrownBy { orderService.deliverOrder(orderToDeliver.id, orderToDeliver.amount + 1) }
            .withMessage("deliver amount can't be bigger than original amount")
    }


    @Test
    fun `should create order with sender and receiver`() {
        var board = boardService.createBoard(boardName)
        val retailer = playerService.addPlayer(board.id, Role.RETAILER)
        val wholesaler = playerService.addPlayer(board.id, Role.WHOLESALER)
        val factory = playerService.addPlayer(board.id, Role.FACTORY)

        orderService.createOrder(retailer.id)
        orderService.createOrder(wholesaler.id)

        board = boardService.loadBoard(board.id)
        assertThat(board.players[0].orders).hasSize(1)
        assertThat(board.players[1].orders).hasSize(2)
        assertThat(board.players[2].orders).hasSize(1)
    }

    @Test
    fun `should fail if order amount bigger than sender stock`() {
        var board = boardService.createBoard(boardName)
        playerService.addPlayer(board.id, Role.RETAILER)
        playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)


        for (i in 1.rangeTo(8)) {
            orderService.createCpuOrders()
        }

        board = boardService.loadBoard(board.id)
        var retailer = board.players.first()
        retailer.orders
            .filter { it.sender == retailer.id }
            .forEach {
                orderService.deliverOrder(it.id, it.amount)
            }

        orderService.createCpuOrders()
        board = boardService.loadBoard(board.id)
        retailer = board.players.first()
        retailer.orders
            .filter { it.sender == retailer.id && it.state == OrderState.PENDING }
            .forEach {
                assertThatExceptionOfType(ImpossibleActionException::class.java)
                    .isThrownBy { orderService.deliverOrder(it.id, it.amount) }
                    .withMessage("Sender doesn't have enough stock to deliver the order")
            }
    }

    @Test
    fun `should fetch all orders from a board`() {
        val board = boardService.createBoard(boardName)
        playerService.addPlayer(board.id, Role.RETAILER)
        playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)


        for (i in 1.rangeTo(8)) {
            orderService.createCpuOrders()
        }

        val orders = orderService.getOrdersByBoard(board.id)

        assertThat(orders).hasSize(24)
    }

    @Test
    fun `should get order with id`() {
        val board = boardService.createBoard(boardName)
        val retailer = playerService.addPlayer(board.id, Role.RETAILER)
        val wholesaler = playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)

        val order = orderService.createOrder(retailer.id)

        val loadOrder = orderService.getOrder(order.first.id)
        assertThat(order.first.id).isEqualTo(loadOrder.first.id)
    }
}