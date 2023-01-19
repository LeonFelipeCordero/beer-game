package com.beer.game.events


data class Event(
    val document: Any,
    val documentId: String,
    val entityId: String? = null,
    val documentType: DocumentType,
    val eventType: EventType,
) {
    fun isSameBoard(id: String) = documentId == id

    fun isSamePlayer(playerId: String): Boolean {
        return entityId == playerId
    }

    fun isRelevantForBoard(): Boolean {
        return (documentType == DocumentType.BOARD && eventType == EventType.UPDATE) ||
                (documentType == DocumentType.PLAYER && eventType == EventType.NEW)
    }

    fun isRelevantForPlayer(): Boolean {
        return documentType == DocumentType.PLAYER && eventType == EventType.UPDATE
    }

    fun isRelevantForNewOrder(): Boolean {
        return documentType == DocumentType.ORDER && eventType == EventType.NEW
    }

    fun isRelevantForUpdateOrder(): Boolean {
        return documentType == DocumentType.ORDER && eventType == EventType.UPDATE
    }
}

enum class EventType {
    NEW,
    UPDATE
}

enum class DocumentType {
    BOARD,
    PLAYER,
    ORDER
}
