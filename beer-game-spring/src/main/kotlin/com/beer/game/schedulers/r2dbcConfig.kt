package com.beer.game.schedulers

import com.beer.game.application.order.OrderService
import mu.KotlinLogging
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty
import org.springframework.scheduling.annotation.Scheduled
import org.springframework.stereotype.Component

private val logger = KotlinLogging.logger {}

@Component
@ConditionalOnProperty(name = ["beer-game.schedulers.orders.enabled"], havingValue = "true", matchIfMissing = true)
class OrderAndBatchScheduler(
    private var orderService: OrderService,
) {

    @Scheduled(fixedDelay = 60000)
    fun createOrders() {
        logger.info("running from order scheduler")
        orderService.createCpuOrders()
    }

    @Scheduled(fixedDelay = 60000)
    fun deliverFactoryBatch() {
        logger.info("running from batch schedulers")
        orderService.deliverFactoryBatches()
    }
}
