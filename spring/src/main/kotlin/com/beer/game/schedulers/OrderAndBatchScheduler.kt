package com.beer.game.schedulers

import com.beer.game.application.service.OrderService
import org.slf4j.LoggerFactory
import org.springframework.context.annotation.Profile
import org.springframework.scheduling.annotation.Scheduled
import org.springframework.stereotype.Component


@Component
@Profile("!test")
class OrderAndBatchScheduler(
    private var orderService: OrderService
) {

    companion object {
        private val logger = LoggerFactory.getLogger(OrderAndBatchScheduler::class.java)
    }

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