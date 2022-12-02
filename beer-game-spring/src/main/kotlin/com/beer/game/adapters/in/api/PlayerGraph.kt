package com.beer.game.adapters.`in`.api

import com.beer.game.common.Role
import com.beer.game.domain.Player

data class PlayerGraph(
    val id: String?,
    val name: String?,
    val role: Role?,
    var stock: Int?,
    var backlog: Int?,
    var weeklyOrder: Int?,
    var lastOrder: Int?,
    val cpu: Boolean?,
    val boardId: String?,
    val ordersId: List<String>?
) {
    companion object {
        fun fromPlayer(player: Player, boardId: String): PlayerGraph {
            return PlayerGraph(
                id = player.id,
                name = player.name,
                role = player.role,
                stock = player.stock,
                backlog = player.backlog,
                weeklyOrder = player.weeklyOrder,
                lastOrder = player.lastOrder,
                cpu = player.cpu,
                boardId = boardId,
                ordersId = player.orders.map { it.id }
            )
        }
    }
}
