package com.beer.game.domain

import com.beer.game.common.OrderState
import com.beer.game.common.OrderType
import java.time.LocalDateTime
import java.util.UUID

data class Order(
    val id: String = UUID.randomUUID().toString(),
    var amount: Int,
    val originalAmount: Int,
    var state: OrderState = OrderState.PENDING,
    val type: OrderType,
    val sender: String,
    val receiver: String? = null,
    val createdAt: LocalDateTime
)
