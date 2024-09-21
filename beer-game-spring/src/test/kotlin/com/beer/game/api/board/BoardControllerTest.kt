package com.beer.game.api.board

import com.beer.game.IntegrationTestBase
import com.beer.game.common.BoardState
import com.beer.game.common.Role
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test
import reactor.test.StepVerifier
import java.time.LocalDate

class BoardControllerTest : IntegrationTestBase() {

    @AfterEach
    fun afterEach() {
        deleteAll()
    }

    @Test
    fun `should create board`() {
        StepVerifier.create(createBoard())
            .assertNext {
                assertThat(it.name).isEqualTo(BOARD_NAME)
                assertThat(it.state).isEqualTo(BoardState.CREATED)
                assertThat(it.full).isEqualTo(false)
                assertThat(it.finished).isEqualTo(false)
                assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
                assertThat(it.playersId).isEmpty()
                assertThat(it.ordersId).isEmpty()
            }.verifyComplete()
    }

    @Test
    fun `should get board by id`() {
        val boardGraph = createBoard().block()
        StepVerifier.create(boardController.getBoard(boardGraph?.id.toString()))
            .assertNext {
                assertThat(it.name).isEqualTo(BOARD_NAME)
                assertThat(it.state).isEqualTo(BoardState.CREATED)
                assertThat(it.full).isEqualTo(false)
                assertThat(it.finished).isEqualTo(false)
                assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
                assertThat(it.playersId).isEmpty()
                assertThat(it.ordersId).isEmpty()
            }.verifyComplete()
    }

    @Test
    fun `should get board by name`() {
        createBoard().block()
        StepVerifier.create(boardController.getBoardByName(BOARD_NAME))
            .assertNext {
                assertThat(it.name).isEqualTo(BOARD_NAME)
                assertThat(it.state).isEqualTo(BoardState.CREATED)
                assertThat(it.full).isEqualTo(false)
                assertThat(it.finished).isEqualTo(false)
                assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
                assertThat(it.playersId).isEmpty()
                assertThat(it.ordersId).isEmpty()
            }.verifyComplete()
    }

    @Test
    fun `should update board if gets full`() {
        createBoardAndPlayers()
        StepVerifier.create(boardController.getBoardByName(BOARD_NAME))
            .assertNext {
                assertThat(it.name).isEqualTo(BOARD_NAME)
                assertThat(it.state).isEqualTo(BoardState.RUNNING)
                assertThat(it.full).isEqualTo(true)
                assertThat(it.finished).isEqualTo(false)
                assertThat(it.createdAt?.toLocalDate()).isEqualTo(LocalDate.now())
                assertThat(it.playersId).hasSize(3)
                assertThat(it.ordersId).isEmpty()
            }.verifyComplete()
    }

    @Test
    fun `should get corresponding players from board graph`() {
        val board = createBoardAndPlayers()
        StepVerifier.create(
            boardController.players(board!!),
        ).assertNext {
            assertThat(it.role).isEqualTo(Role.RETAILER)
        }.assertNext {
            assertThat(it.role).isEqualTo(Role.WHOLESALER)
        }.assertNext {
            assertThat(it.role).isEqualTo(Role.FACTORY)
        }.verifyComplete()
    }

    @Test
    fun `should get corresponding orders from board graph`() {
        val board = createBoardAndPlayers()
        val receiverId = board?.playersId?.first()!!
        val order = orderController.createOrder(receiverId).block()
        StepVerifier.create(
            boardController.orders(board),
        ).assertNext {
            assertThat(it.id).isEqualTo(order?.id)
        }.assertNext {
            assertThat(it.id).isEqualTo(order?.id)
        }.verifyComplete()
    }

    @Test
    fun `should get available roles from a board graph`() {
        val board = createBoard().block()
        StepVerifier.create(
            boardController.availableRoles(board!!),
        ).assertNext {
            assertThat(it).isEqualTo(Role.RETAILER)
        }.assertNext {
            assertThat(it).isEqualTo(Role.WHOLESALER)
        }.assertNext {
            assertThat(it).isEqualTo(Role.FACTORY)
        }.verifyComplete()
    }
}
