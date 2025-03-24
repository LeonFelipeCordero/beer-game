package com.beer.game.repositories.player.postgres

import org.springframework.data.r2dbc.repository.R2dbcRepository

interface PlayerRepository : R2dbcRepository<PlayerEntity, String>
