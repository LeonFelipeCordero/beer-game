package com.beer.game

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.scheduling.annotation.EnableScheduling

@EnableScheduling
@SpringBootApplication
class BeerGameSpringApplication

fun main(args: Array<String>) {
    runApplication<BeerGameSpringApplication>(*args)
}
