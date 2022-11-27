package com.beer.game.adapters.out.mongo

import org.bson.types.ObjectId
import org.springframework.data.mongodb.repository.MongoRepository

interface BoardRepository : MongoRepository<BoardDocument, String> {
    fun findOneById(id: ObjectId): BoardDocument?
    fun findBoardDocumentByState(state: String): List<BoardDocument>
}