package com.beer.game.events

import reactor.core.publisher.Flux

interface EventEmitter<T> {
    fun publish(event: T)
    fun subscribe(): Flux<T>
}