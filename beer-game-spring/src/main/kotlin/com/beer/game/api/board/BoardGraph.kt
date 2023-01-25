package com.beer.game.api.board

import com.beer.game.common.BoardState
import com.beer.game.domain.Board
import java.time.LocalDateTime

class BoardGraph(
    val id: String?,
    val name: String?,
    var state: BoardState?,
    var full: Boolean?,
    var finished: Boolean?,
    val createdAt: LocalDateTime?,
    val playersId: List<String>?,
    val ordersId: List<String>?
) {
    companion object {
        fun fromBoard(board: Board): BoardGraph {
            return BoardGraph(
                id = board.id,
                name = board.name,
                state = board.state,
                full = board.full,
                finished = board.finished,
                createdAt = board.createdAt,
                playersId = board.players.map { it.id },
                ordersId = board.players.flatMap { it.orders }.map { it.id }
            )
        }
    }
}
