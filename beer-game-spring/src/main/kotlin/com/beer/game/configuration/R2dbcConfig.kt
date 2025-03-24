package com.beer.game.configuration

import io.r2dbc.spi.ConnectionFactories
import io.r2dbc.spi.ConnectionFactory
import io.r2dbc.spi.ConnectionFactoryOptions
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.data.r2dbc.core.R2dbcEntityTemplate
import org.springframework.r2dbc.connection.R2dbcTransactionManager

@Configuration
class R2dbcConfig {

//    @Bean
//    fun connectionFactory(): ConnectionFactory {
//        return ConnectionFactories.get(
//            ConnectionFactoryOptions.builder()
//                .option(ConnectionFactoryOptions.DRIVER, "pool")
//                .option(ConnectionFactoryOptions.PROTOCOL, "postgresql")
//                .option(ConnectionFactoryOptions.HOST, "localhost")
//                .option(ConnectionFactoryOptions.PORT, 32770)
//                .option(ConnectionFactoryOptions.DATABASE, "beer_game")
//                .option(ConnectionFactoryOptions.USER, "beer_game")
//                .option(ConnectionFactoryOptions.PASSWORD, "beer_game")
//                .build()
//        )
//    }
//    @Bean
//    fun transactionManager(connectionFactory: ConnectionFactory): R2dbcTransactionManager {
//        return R2dbcTransactionManager(connectionFactory)
//    }
//
//    @Bean
//    fun r2dbcEntityTemplate(connectionFactory: ConnectionFactory): R2dbcEntityTemplate {
//        return R2dbcEntityTemplate(connectionFactory)
//    }
//
//    @Bean
//    fun connectionFactory(): ConnectionFactory {
//        ConnectionFactories.find(
//            ConnectionFactoryOptions.builder()
//                .option(ConnectionFactoryOptions.HOST)
//        )
//    }
}
