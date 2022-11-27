package com.beer.game.domain

import com.beer.game.common.Role
import java.util.UUID

data class Player(
    val id: String = UUID.randomUUID().toString(),
    val name: String,
    val role: Role,
    var stock: Int,
    var backlog: Int,
    var weeklyOrder: Int,
    var lastOrder: Int,
    val cpu: Boolean,
    val orders: MutableList<Order>
) {
    fun addOrder(order: Order) {
        orders.add(order)
    }
}
