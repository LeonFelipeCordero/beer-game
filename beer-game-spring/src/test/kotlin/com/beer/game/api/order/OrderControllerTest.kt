package com.beer.game.api.order

import com.beer.game.IntegrationTestBase
import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test
import reactor.test.StepVerifier
import java.time.LocalDate

class OrderControllerTest : IntegrationTestBase() {

    @AfterEach
    fun afterEach() {
        deleteAll()
    }

    @Test
    fun `should create a new order`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val senderId = board.playersId!![1]
        StepVerifier.create(
            orderController.createOrder(receiverId),
        ).assertNext {
            assertThat(it.amount).isEqualTo(40)
            assertThat(it.originalAmount).isEqualTo(40)
            assertThat(it.state).isEqualTo(OrderState.PENDING)
            assertThat(it.type).isEqualTo(OrderType.PLAYER_ORDER)
            assertThat(it.senderId).isEqualTo(senderId)
            assertThat(it.receiverId).isEqualTo(receiverId)
            assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
        }.verifyComplete()
    }

    @Test
    fun `should deliver an order`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val senderId = board.playersId!![1]
        val order = orderController.createOrder(receiverId).block()
        orderController.deliverOrder(order?.id.toString(), order?.amount).block()
        StepVerifier.create(
            boardController.orders(board)
                .filter { it.id == order?.id },
        ).assertNext {
            assertThat(it.amount).isEqualTo(40)
            assertThat(it.originalAmount).isEqualTo(40)
            assertThat(it.state).isEqualTo(OrderState.DELIVERED)
            assertThat(it.type).isEqualTo(OrderType.PLAYER_ORDER)
            assertThat(it.senderId).isEqualTo(senderId)
            assertThat(it.receiverId).isEqualTo(receiverId)
            assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
        }.assertNext {
            assertThat(it.amount).isEqualTo(40)
            assertThat(it.originalAmount).isEqualTo(40)
            assertThat(it.state).isEqualTo(OrderState.DELIVERED)
            assertThat(it.type).isEqualTo(OrderType.PLAYER_ORDER)
            assertThat(it.senderId).isEqualTo(senderId)
            assertThat(it.receiverId).isEqualTo(receiverId)
            assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
        }.verifyComplete()
    }

    @Test
    fun `should deliver an order and update amount`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val senderId = board.playersId!![1]
        val order = orderController.createOrder(receiverId).block()
        orderController.deliverOrder(order?.id.toString(), 1).block()
        StepVerifier.create(
            boardController.orders(board)
                .filter { it.id == order?.id },
        ).assertNext {
            assertThat(it.amount).isEqualTo(1)
            assertThat(it.originalAmount).isEqualTo(40)
            assertThat(it.state).isEqualTo(OrderState.DELIVERED)
            assertThat(it.type).isEqualTo(OrderType.PLAYER_ORDER)
            assertThat(it.senderId).isEqualTo(senderId)
            assertThat(it.receiverId).isEqualTo(receiverId)
            assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
        }.assertNext {
            assertThat(it.amount).isEqualTo(1)
            assertThat(it.originalAmount).isEqualTo(40)
            assertThat(it.state).isEqualTo(OrderState.DELIVERED)
            assertThat(it.type).isEqualTo(OrderType.PLAYER_ORDER)
            assertThat(it.senderId).isEqualTo(senderId)
            assertThat(it.receiverId).isEqualTo(receiverId)
            assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
        }.verifyComplete()
    }

    @Test
    fun `should deliver an order and update players`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val order = orderController.createOrder(receiverId).block()
        orderController.deliverOrder(order?.id.toString(), 10).block()
        StepVerifier.create(
            boardController.players(board),
        ).assertNext {
            assertThat(it.stock).isEqualTo(90)
            assertThat(it.lastOrder).isEqualTo(10)
        }.assertNext {
            assertThat(it.stock).isEqualTo(1190)
        }.assertNext {
            assertThat(it.stock).isEqualTo(12000)
        }.verifyComplete()
    }

    @Test
    fun `should get a sender player from an order`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val senderId = board.playersId!![1]
        val order = orderController.createOrder(receiverId).block()
        StepVerifier.create(
            orderController.sender(order!!),
        ).assertNext {
            assertThat(it.id).isEqualTo(senderId)
        }.verifyComplete()
    }

    @Test
    fun `should get a receiver player from an order`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val order = orderController.createOrder(receiverId).block()
        StepVerifier.create(
            orderController.receiver(order!!),
        ).assertNext {
            assertThat(it.id).isEqualTo(receiverId)
        }.verifyComplete()
    }

    @Test
    fun `should get the board from an order`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val order = orderController.createOrder(receiverId).block()
        StepVerifier.create(
            orderController.board(order!!),
        ).assertNext {
            assertThat(it.id).isEqualTo(board.id)
        }.verifyComplete()
    }
}
