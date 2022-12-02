package com.beer.game.api

import com.beer.game.IntegrationTestBase
import com.beer.game.TestUtils
import com.beer.game.adapters.`in`.api.BoardGraph
import com.beer.game.adapters.`in`.api.PlayerGraph
import com.beer.game.api.GraphQlDocuments.documentAddPlayerMinimal
import com.beer.game.api.GraphQlDocuments.documentCreateBoard
import com.beer.game.api.GraphQlDocuments.documentCreateOrder
import com.beer.game.common.Role
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test
import org.skyscreamer.jsonassert.Customization
import org.skyscreamer.jsonassert.JSONAssert
import org.skyscreamer.jsonassert.JSONCompareMode
import org.skyscreamer.jsonassert.comparator.CustomComparator

class OrderControllerTest : IntegrationTestBase() {

    @AfterEach
    fun setup() {
        boardRepository.deleteAll()
    }

    @Test
    fun `should create board with players and receive orders`() {
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

        val order = graphQlTester
            .document(documentCreateOrder)
            .variable("boardId", board.id)
            .variable("senderId", wholesaler.id)
            .variable("receiverId", retailer.id)
            .execute()
            .path("createOrder")
            .entity(Any::class.java)
            .get()

        JSONAssert.assertEquals(
            TestUtils.objectToString(order),
            TestUtils.mockData("/responses/order.json"),
            CustomComparator(
                JSONCompareMode.LENIENT,
                Customization("createdAt") { _, _ -> true },
                Customization("id") { _, _ -> true },
                Customization("board.id") { _, _ -> true },
                Customization("board.createdAt") { _, _ -> true },
                Customization("sender.id") { _, _ -> true },
                Customization("sender.createdAt") { _, _ -> true },
                Customization("receiver.id") { _, _ -> true },
                Customization("receiver.createdAt") { _, _ -> true }
            )
        )
    }

    private fun addPlayer(boardId: String, role: Role): PlayerGraph {
        return graphQlTester
            .document(documentAddPlayerMinimal)
            .variable("boardId", boardId)
            .variable("role", role)
            .execute()
            .path("addPlayer")
            .entity(PlayerGraph::class.java)
            .get()
    }
}