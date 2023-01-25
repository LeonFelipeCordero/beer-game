package com.beer.game

import com.beer.game.repositories.board.BoardRepository
import com.beer.game.application.board.BoardService
import com.beer.game.application.order.OrderService
import com.beer.game.application.player.PlayerService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.boot.test.web.server.LocalServerPort
import org.springframework.graphql.test.tester.GraphQlTester
import org.springframework.test.context.ActiveProfiles
import java.lang.IllegalStateException

@SpringBootTest(
    webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT
)
@ActiveProfiles("test")
class IntegrationTestBase {

    companion object {
        const val boardName = "test"
    }

    @LocalServerPort
    private var port: Int? = null

    @Autowired
    protected lateinit var boardService: BoardService

    @Autowired
    protected lateinit var playerService: PlayerService

    @Autowired
    protected lateinit var orderService: OrderService

    @Autowired
    protected lateinit var boardRepository: BoardRepository

    @Autowired
    protected lateinit var graphQlTester: GraphQlTester

    protected fun getPort() =
        port ?: throw IllegalStateException("Spring context is not initialized correctly for the test.")
}