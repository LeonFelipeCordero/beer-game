package com.beer.game.events


data class Event(
    val document: Any,
    val documentId: String,
    val entityId: String? = null,
    val documentType: DocumentType,
    val eventType: EventType,
)

enum class EventType {
    NEW,
    UPDATE
}

enum class DocumentType {
    BOARD,
    PLAYER,
    ORDER
}
