package com.beer.game.api

import com.beer.game.IntegrationTestBase
import com.beer.game.TestUtils
import com.beer.game.api.board.BoardGraph
import com.beer.game.api.GraphQlDocuments.documentAddPlayer
import com.beer.game.api.GraphQlDocuments.documentCreateBoard
import com.beer.game.common.Role
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test
import org.skyscreamer.jsonassert.Customization
import org.skyscreamer.jsonassert.JSONAssert
import org.skyscreamer.jsonassert.JSONCompareMode
import org.skyscreamer.jsonassert.comparator.CustomComparator

class PlayerControllerTest : IntegrationTestBase() {

    @AfterEach
    fun setup() {
        boardRepository.deleteAll()
    }

    @Test
    fun `should add all players to game and start it`() {
        val board = graphQlTester
            .document(documentCreateBoard)
            .variable("name", boardName)
            .execute()
            .path("createBoard")
            .entity(BoardGraph::class.java)
            .get()

        val retailer = addPlayer(board.id!!, Role.RETAILER)
        val wholesaler = addPlayer(board.id!!, Role.WHOLESALER)
        val factory = addPlayer(board.id!!, Role.FACTORY)

        comparePlayer(retailer, TestUtils.mockData("/responses/retailer.json"))
        comparePlayer(wholesaler, TestUtils.mockData("/responses/wholesaler.json"))
        comparePlayer(factory, TestUtils.mockData("/responses/factory.json"))
    }

    private fun addPlayer(boardId: String, role: Role): String {
        val player = graphQlTester
            .document(documentAddPlayer)
            .variable("boardId", boardId)
            .variable("role", role)
            .execute()
            .path("addPlayer")
            .entity(Any::class.java)
            .get()

        return TestUtils.objectToString(player)
    }

    private fun comparePlayer(left: String, right: String) {
        JSONAssert.assertEquals(
            left,
            right,
            CustomComparator(
                JSONCompareMode.LENIENT,
                Customization("createdAt") { _, _ -> true },
                Customization("id") { _, _ -> true },
                Customization("board.id") { _, _ -> true },
                Customization("board.createdAt") { _, _ -> true },
            )
        )
    }
}