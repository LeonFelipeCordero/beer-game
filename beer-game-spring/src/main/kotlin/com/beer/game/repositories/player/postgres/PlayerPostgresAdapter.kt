package com.beer.game.repositories.player.postgres

import com.beer.game.application.player.PlayerStorageAdapter
import com.beer.game.domain.Board
import com.beer.game.domain.Player
import com.beer.game.repositories.board.postgres.BoardRepository
import org.springframework.stereotype.Component
import reactor.core.publisher.Mono

@Component
class PlayerPostgresAdapter(
    private val playerRepository: PlayerRepository,
    private val boardRepository: BoardRepository
) : PlayerStorageAdapter {

    override fun savePlayer(player: Player): Mono<Player> {
        return playerRepository.save(PlayerEntity.fromPlayer(player))
            .zipWith(boardRepository.findById(player.board.getUUID()))
            .map { tuple -> tuple.t1.toPlayer(tuple.t2.toBoard()) }
    }

    override fun loadPlayer(playerId: String): Mono<Player> {
//        val player = playerRepository.findById(playerId)
//        return Pair(player.toPlayer(player.board), player.board?.boardId!!.toString())
        return Mono.empty<Player>()
    }

    private fun mergeWithBoard(player: Player) {

    }
}
