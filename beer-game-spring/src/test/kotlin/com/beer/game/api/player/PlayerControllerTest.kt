package com.beer.game.api.player

import com.beer.game.IntegrationTestBase
import com.beer.game.common.Role
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test
import reactor.test.StepVerifier

class PlayerControllerTest : IntegrationTestBase() {

    @AfterEach
    fun afterEach() {
        deleteAll()
    }

    @Test
    fun `should add player to board`() {
        val board = createBoard().block()
        StepVerifier.create(
            playerController.addPlayer(board?.id.toString(), Role.RETAILER)
        ).assertNext {
            assertThat(it.name).isEqualTo(Role.RETAILER.toString())
            assertThat(it.role).isEqualTo(Role.RETAILER)
            assertThat(it.stock).isEqualTo(80)
            assertThat(it.backlog).isEqualTo(80)
            assertThat(it.weeklyOrder).isEqualTo(40)
            assertThat(it.lastOrder).isEqualTo(40)
            assertThat(it.cpu).isFalse
            assertThat(it.boardId).isEqualTo(board?.id)
            assertThat(it.ordersId).isEmpty()
        }.verifyComplete()
    }

    @Test
    fun `should update weekly order of a player`() {
        val board = createBoardAndPlayers()
        playerController.updateWeeklyOrder(
            playerId = board?.playersId?.first()!!,
            amount = 1,
        ).block()
        StepVerifier.create(
            playerController.getPlayer(
                board.playersId?.first()!!
            )
        ).assertNext {
            assertThat(it.weeklyOrder).isEqualTo(1)
        }.verifyComplete()
    }

    @Test
    fun `should get all players in a board`() {
        val board = createBoardAndPlayers()
        StepVerifier.create(
            playerController.getPlayersByBoard(board?.id.toString())
        ).assertNext {
            assertThat(it.role).isEqualTo(Role.RETAILER)
        }.assertNext {
            assertThat(it.role).isEqualTo(Role.WHOLESALER)
        }.assertNext {
            assertThat(it.role).isEqualTo(Role.FACTORY)
        }.verifyComplete()
    }

    @Test
    fun `should get the board of player graph`() {
        val board = createBoardAndPlayers()
        val player = playerController.getPlayer(
            board?.playersId?.first()!!
        ).block()
        StepVerifier.create(
            playerController.board(player!!)
        ).assertNext {
            assertThat(it.id).isEqualTo(board.id)
        }.verifyComplete()
    }

    @Test
    fun `should get the in and out orders of player`() {
    }
}