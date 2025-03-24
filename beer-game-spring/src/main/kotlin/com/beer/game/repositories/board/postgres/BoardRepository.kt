package com.beer.game.repositories.board.postgres

import org.springframework.data.r2dbc.repository.Query
import org.springframework.data.r2dbc.repository.R2dbcRepository
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import java.util.UUID

interface BoardRepository : R2dbcRepository<BoardEntity, UUID> {
    fun findBoardEntityByState(state: String): Flux<BoardEntity>
    fun findByName(name: String): Mono<BoardEntity>

    @Query(
        """
        SELECT b
        FROM boards b
        left join players p on b.id = p.board.id
        left join orders os on p.id = os.sender.id
        left join orders or on p.id = or.receiver.id
        where os.id = :oroderId
           or or.receiver=:orderId
    """)
    fun findOneByPlayersId(playerId: String): Mono<BoardEntity>

    @Query(
        """
        SELECT b
        FROM boards b
        left join players p on b.id = p.board.id
        left join orders os on p.id = os.sender.id
        left join orders or on p.id = or.receiver.id
        where os.id = :oroderId
           or or.receiver=:orderId
    """
    )
    fun findByOrderId(orderId: String): Mono<BoardEntity>
}
