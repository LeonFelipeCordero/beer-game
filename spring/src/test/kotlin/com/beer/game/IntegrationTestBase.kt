package com.beer.game

import com.beer.game.adapters.out.mongo.BoardRepository
import com.beer.game.api.BoardController
import com.beer.game.api.BoardControllerTest
import com.beer.game.application.service.BoardService
import com.beer.game.application.service.OrderService
import com.beer.game.application.service.PlayerService
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.boot.test.web.server.LocalServerPort
import org.springframework.graphql.test.tester.AbstractDelegatingGraphQlTester
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