package com.beer.game

import com.beer.game.repositories.board.BoardRepository
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest

@SpringBootTest
class BeerGameSpringApplicationTests {

    @Autowired
    private lateinit var boardRepository: BoardRepository

    @Test
    fun contextLoads() {
        assertThat(boardRepository).isNotNull
    }
}
