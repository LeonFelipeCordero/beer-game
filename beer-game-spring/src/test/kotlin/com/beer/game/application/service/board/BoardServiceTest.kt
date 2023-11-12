package com.beer.game.application.service.board

import com.beer.game.IntegrationTestBase
import com.beer.game.common.BoardState
import com.beer.game.common.Role
import com.beer.game.domain.exceptions.ImpossibleActionException
import com.beer.game.domain.exceptions.NotFoundException
import org.assertj.core.api.Assertions.assertThat
import org.assertj.core.api.Assertions.assertThatExceptionOfType
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test
import java.time.LocalDate

internal class BoardServiceTest : IntegrationTestBase() {

    @AfterEach
    fun setup() {
        boardRepository.deleteAll()
    }

    @Test
    fun `should fail if loading board that doesn't exist`() {
        assertThatExceptionOfType(NotFoundException::class.java)
            .isThrownBy { boardService.loadBoard("507f1f77bcf86cd799439011") }
            .withMessage("Board with id 507f1f77bcf86cd799439011 doesn't exit")
    }

    @Test
    fun `should save board document in repository and return a board`() {
        val board = boardService.createBoard(BOARD_NAME)

        assertThat(board.id).isNotNull.isNotEmpty
        assertThat(board.name).isEqualTo(BOARD_NAME)
        assertThat(board.state).isEqualTo(BoardState.CREATED)
        assertThat(board.full).isFalse
        assertThat(board.finished).isFalse
        assertThat(board.createdAt.toLocalDate()).isEqualTo(LocalDate.now())
        assertThat(board.players).hasSize(0)

        assertThatExceptionOfType(ImpossibleActionException::class.java)
            .isThrownBy { boardService.createBoard(BOARD_NAME) }
            .withMessage("Name is already used by another board")
    }

    @Test
    fun `should create board and only returned as active when starting the game`() {
        val board = boardService.createBoard(BOARD_NAME)

        playerService.addPlayer(board.id, Role.RETAILER)
        var loadedBoards = boardService.loadActiveBoards()
        assertThat(loadedBoards).hasSize(0)

        playerService.addPlayer(board.id, Role.WHOLESALER)
        loadedBoards = boardService.loadActiveBoards()
        assertThat(loadedBoards).hasSize(0)

        playerService.addPlayer(board.id, Role.FACTORY)
        loadedBoards = boardService.loadActiveBoards()
        assertThat(loadedBoards).hasSize(1)
    }
}