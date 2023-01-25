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
import org.springframework.graphql.test.tester.GraphQlTester
import org.springframework.test.context.ActiveProfiles
import reactor.core.publisher.Mono
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

    protected fun getPort() =
        port ?: throw IllegalStateException("Spring context is not initialized correctly for the test.")

    protected fun createBoard(): Mono<BoardGraph> {
        return boardController.createBoard(boardName)
    }

    protected fun createBoardAndPlayers(): BoardGraph? {
        val board = createBoard().block()
        playerController.addPlayer(board?.id.toString(), Role.RETAILER).block()
        playerController.addPlayer(board?.id.toString(), Role.WHOLESALER).block()
        playerController.addPlayer(board?.id.toString(), Role.FACTORY).block()
        return boardController.getBoard(board?.id.toString()).block()
    }
}