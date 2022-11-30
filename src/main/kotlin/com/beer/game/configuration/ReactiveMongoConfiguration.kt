package com.beer.game.configuration

import com.mongodb.client.MongoClient
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.data.mongodb.core.ReactiveMongoTemplate
import org.springframework.data.mongodb.repository.config.EnableReactiveMongoRepositories

@Configuration
//@EnableReactiveMongoRepositories
class ReactiveMongoConfiguration {

//    @Bean
//    fun reactiveMongoTemplate(mongoClient: MongoClient): ReactiveMongoTemplate {
//        return ReactiveMongoTemplate(mongoClient, "beer_game")
//    }
}