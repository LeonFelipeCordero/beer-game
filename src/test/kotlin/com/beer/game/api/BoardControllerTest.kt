package com.beer.game.api

import com.beer.game.IntegrationTestBase
import com.beer.game.TestUtils
import com.beer.game.adapters.`in`.api.BoardGraph
import com.beer.game.api.GraphQlDocuments.documentCreateBoard
import com.beer.game.api.GraphQlDocuments.documentGetBoard
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Test
import org.skyscreamer.jsonassert.Customization
import org.skyscreamer.jsonassert.JSONAssert
import org.skyscreamer.jsonassert.JSONCompareMode
import org.skyscreamer.jsonassert.comparator.CustomComparator
import java.time.LocalDateTime


internal class BoardControllerTest : IntegrationTestBase() {

    @AfterEach
    fun setup() {
        boardRepository.deleteAll()
    }

    @Test
    fun `should create and fetch board from graphql interface`() {
        val board = graphQlTester
            .document(documentCreateBoard)
            .variable("name", boardName)
            .execute()
            .path("createBoard")
            .entity(BoardGraph::class.java)
            .get()

        val loadedBoard = graphQlTester
            .document(documentGetBoard)
            .variable("id", board.id)
            .execute()
            .path("getBoard")
            .entity(BoardGraph::class.java)
            .get()

        JSONAssert.assertEquals(
            TestUtils.objectToString(board),
            TestUtils.objectToString(loadedBoard),
            CustomComparator(
                JSONCompareMode.LENIENT,
                Customization("createdAt") { o1, o2 ->
                    val date1 = LocalDateTime.parse(o1.toString())
                    val date2 = LocalDateTime.parse(o2.toString())
                    date1.toLocalDate() == date2.toLocalDate()
                }
            )
        )
    }
}