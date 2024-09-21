package com.beer.game.domain

import com.beer.game.application.events.InternalEventListener
import com.beer.game.common.BoardState
import com.beer.game.common.Role
import com.beer.game.events.DocumentType
import com.beer.game.events.Event
import com.beer.game.events.EventType
import java.lang.RuntimeException
import java.time.LocalDateTime

class Board(
    val id: String,
    val name: String,
    var state: BoardState,
    var full: Boolean,
    var finished: Boolean,
    val createdAt: LocalDateTime,
    val players: MutableList<Player>,
) {

    companion object {
        fun fromName(name: String): Board {
            return Board(
                id = "",
                name = name,
                state = BoardState.CREATED,
                full = false,
                finished = false,
                createdAt = LocalDateTime.now(),
                players = mutableListOf(),
            )
        }
    }

    fun addPlayer(player: Player) {
        players.add(player)
    }

    fun findPlayer(playerId: String): Player {
        return players.first { it.id == playerId }
    }

    fun findOrder(orderId: String): Order {
        return players
            .flatMap { it.orders }
            .first { it.id == orderId }
    }

    private fun findPlayerByRole(role: Role): Player {
        return players.first { it.role == role }
    }

    fun findContraPart(receiver: Player): Player {
        return when (receiver.role) {
            Role.RETAILER -> findPlayerByRole(Role.WHOLESALER)
            Role.WHOLESALER -> findPlayerByRole(Role.FACTORY)
            else -> throw RuntimeException("This role ${receiver.role} can't create orders")
        }
    }

    fun availableRoles(): List<Role> {
        return Role.values()
            .asList()
            .filter { role ->
                !this.players.map { player -> player.role }.contains(role)
            }
    }

    fun emitUpdate(listener: InternalEventListener) {
        listener.publish(
            Event(
                document = this,
                documentId = id,
                documentType = DocumentType.PLAYER,
                eventType = EventType.UPDATE,
            ),
        )
    }
}
