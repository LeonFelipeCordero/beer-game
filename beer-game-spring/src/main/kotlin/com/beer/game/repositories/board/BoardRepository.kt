package com.beer.game.repositories.board

import org.bson.types.ObjectId
import org.springframework.data.mongodb.repository.MongoRepository
import org.springframework.data.mongodb.repository.Query

interface BoardRepository : MongoRepository<BoardDocument, ObjectId> {
    fun findOneById(id: ObjectId): BoardDocument?
    fun findBoardDocumentByState(state: String): List<BoardDocument>
    fun findOneByName(name: String): BoardDocument?

    @Query(value = "{'players._id': ?0}")
    fun findOneByPlayersId(playerId: String): BoardDocument?

    @Query(value = "{'orders._id': ?0}")
    fun findOneByOrdersId(orderId: String): BoardDocument?
}