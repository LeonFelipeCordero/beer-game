package com.beer.game.repositories.player.mongo

import com.beer.game.common.Role
import com.beer.game.domain.Player
import com.beer.game.repositories.order.mongo.OrderDocument
import java.util.UUID

/**
 * sock: available beers bottles to sell
 * backlog: pending beers to be delivered to the player
 * weeklyOrder: amount of beers to be ordered in the next round
 * lastOrder: amount of beers deliver in the last order
 */
data class PlayerDocument(
    val id: String = UUID.randomUUID().toString(),
    val name: String,
    val role: String,
    var stock: Int,
    var backlog: Int,
    var weeklyOrder: Int,
    var lastOrder: Int,
    val cpu: Boolean,
) {

//    companion object {
//        fun fromPlayer(player: Player): PlayerDocument {
//            return PlayerDocument(
//                id = player.id,
//                name = player.name,
//                role = player.role.toString(),
//                stock = player.stock,
//                backlog = player.backlog,
//                weeklyOrder = player.weeklyOrder,
//                lastOrder = player.lastOrder,
//                cpu = player.cpu,
//            )
//        }
//    }
//
//    fun toPlayer(orders: List<OrderDocument>): Player {
//        return Player(
//            id = id,
//            name = name,
//            role = Role.valueOf(role),
//            stock = stock,
//            backlog = backlog,
//            weeklyOrder = weeklyOrder,
//            lastOrder = lastOrder,
//            cpu = cpu,
//            orders = orders.map { it.toOrder() }.toMutableList(),
//        )
//    }
}
