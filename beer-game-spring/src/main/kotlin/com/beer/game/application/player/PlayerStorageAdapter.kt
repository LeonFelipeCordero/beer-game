package com.beer.game.application.player

import com.beer.game.domain.Board
import com.beer.game.domain.Player

interface PlayerStorageAdapter {
    fun savePlayer(board: Board, player: Player): Player
    fun loadPlayer(playerId: String): Pair<Player, String>
}
