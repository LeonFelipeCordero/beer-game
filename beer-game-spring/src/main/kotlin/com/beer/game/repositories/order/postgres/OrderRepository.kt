package com.beer.game.repositories.order.postgres

import org.springframework.data.r2dbc.repository.R2dbcRepository

interface OrderRepository : R2dbcRepository<OrderEntity, String>

