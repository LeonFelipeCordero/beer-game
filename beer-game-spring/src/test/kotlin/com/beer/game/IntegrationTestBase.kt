package com.beer.game

import com.beer.game.api.board.BoardController
import com.beer.game.api.board.BoardGraph
import com.beer.game.api.order.OrderController
import com.beer.game.api.player.PlayerController
import com.beer.game.repositories.board.BoardRepository
import com.beer.game.application.board.BoardService
import com.beer.game.application.order.OrderService
import com.beer.game.application.player.PlayerService
import com.beer.game.common.Role
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.boot.test.web.server.LocalServerPort
import org.springframework.boot.testcontainers.service.connection.ServiceConnection
import org.springframework.graphql.test.tester.GraphQlTester
import org.springframework.test.context.ActiveProfiles
import org.testcontainers.containers.MongoDBContainer
import org.testcontainers.junit.jupiter.Container
import org.testcontainers.junit.jupiter.Testcontainers
import reactor.core.publisher.Mono
import java.lang.IllegalStateException

@Testcontainers
@SpringBootTest(
    webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT
)
@ActiveProfiles("test")
class IntegrationTestBase {

    companion object {
        const val BOARD_NAME = "test"

        @Container
        @ServiceConnection
        val container: MongoDBContainer = MongoDBContainer("mongo:latest")
            .apply {
                portBindings.add("27017:27017")
                withReuse(true)
                start()
            }
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
    protected lateinit var boardController: BoardController

    @Autowired
    protected lateinit var playerController: PlayerController

    @Autowired
    protected lateinit var orderController: OrderController

    @Autowired
    protected lateinit var graphQlTester: GraphQlTester

    protected fun deleteAll() {
        boardRepository.deleteAll()
    }

    protected fun createBoard(): Mono<BoardGraph> {
        return boardController.createBoard(BOARD_NAME)
    }

    protected fun createBoardAndPlayers(): BoardGraph? {
        val board = createBoard().block()
        playerController.addPlayer(board?.id.toString(), Role.RETAILER).block()
        playerController.addPlayer(board?.id.toString(), Role.WHOLESALER).block()
        playerController.addPlayer(board?.id.toString(), Role.FACTORY).block()
        return boardController.getBoard(board?.id.toString()).block()
    }
}