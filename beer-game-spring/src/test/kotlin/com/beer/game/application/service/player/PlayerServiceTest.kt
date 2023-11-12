package com.beer.game.application.service.player

import com.beer.game.IntegrationTestBase
import com.beer.game.common.BoardState
import com.beer.game.common.Role
import com.beer.game.common.Role.RETAILER
import com.beer.game.domain.exceptions.ImpossibleActionException
import org.assertj.core.api.Assertions.assertThat
import org.assertj.core.api.Assertions.assertThatExceptionOfType
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test

internal class PlayerServiceTest : IntegrationTestBase() {

    @AfterEach
    fun setup() {
        boardRepository.deleteAll()
    }

    @Test
    fun `should add players to board and start game and validate impossible actions`() {
        val board = boardService.createBoard(BOARD_NAME)

        playerService.addPlayer(board.id, RETAILER)
        var loadedBoard = boardService.loadBoard(board.id)
        assertThat(loadedBoard.players).hasSize(1)
        assertThat(loadedBoard.players[0].role).isEqualTo(RETAILER)

        assertThatExceptionOfType(ImpossibleActionException::class.java)
            .isThrownBy { playerService.addPlayer(board.id, RETAILER) }
            .withMessage("Board ${board.id} already has role $RETAILER")

        playerService.addPlayer(board.id, Role.WHOLESALER)
        loadedBoard = boardService.loadBoard(board.id)
        assertThat(loadedBoard.players).hasSize(2)
        assertThat(loadedBoard.players[1].role).isEqualTo(Role.WHOLESALER)

        playerService.addPlayer(board.id, Role.FACTORY)
        loadedBoard = boardService.loadBoard(board.id)
        assertThat(loadedBoard.players).hasSize(3)
        assertThat(loadedBoard.players[2].role).isEqualTo(Role.FACTORY)

        assertThat(loadedBoard.full).isTrue
        assertThat(loadedBoard.state).isEqualTo(BoardState.RUNNING)
        assertThat(loadedBoard.finished).isFalse

        assertThatExceptionOfType(ImpossibleActionException::class.java)
            .isThrownBy { playerService.addPlayer(board.id, Role.FACTORY) }
            .withMessage("Board ${board.id} is full")
        assertThat(loadedBoard.players).hasSize(3)
    }

    @Test
    fun `should fetch player`() {
        val board = boardService.createBoard(BOARD_NAME)

        val player = playerService.addPlayer(board.id, RETAILER)
        val loadedPayer = playerService.getPlayer(player.id).first

        assertThat(player).isEqualTo(loadedPayer)
    }

    @Test
    fun `should fetch all players in board`() {
        val board = boardService.createBoard(BOARD_NAME)

        playerService.addPlayer(board.id, RETAILER)
        playerService.addPlayer(board.id, Role.WHOLESALER)
        playerService.addPlayer(board.id, Role.FACTORY)

        val players = playerService.getPlayersInBoard(board.id)

        assertThat(players).hasSize(3)
    }


    @Test
    fun `should update weekly order for customer`() {
        val board = boardService.createBoard(BOARD_NAME)

        val player = playerService.addPlayer(board.id, RETAILER)
        playerService.changeWeeklyOrder(player.id, 200)
        val loadedPayer = playerService.getPlayer(player.id).first

        assertThat(loadedPayer.weeklyOrder).isEqualTo(200)
    }
}