package com.beer.game.adapters.out.mongo

import com.beer.game.domain.Board
import org.bson.types.ObjectId
import org.springframework.data.mongodb.repository.MongoRepository

interface BoardRepository : MongoRepository<BoardDocument, ObjectId> {
    fun findOneById(id: ObjectId): BoardDocument?
    fun findBoardDocumentByState(state: String): List<BoardDocument>

    fun findOneByName(name: String): BoardDocument?
}