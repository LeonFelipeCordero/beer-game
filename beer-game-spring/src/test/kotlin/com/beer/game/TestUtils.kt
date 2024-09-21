package com.beer.game

import com.fasterxml.jackson.databind.DeserializationFeature
import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.databind.SerializationFeature
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import java.io.File
import java.nio.file.Files
import java.text.SimpleDateFormat

object TestUtils {

    private fun createObjectMapper(): ObjectMapper = jacksonObjectMapper().apply {
        configure(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS, true)
        configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)
        registerModule(JavaTimeModule())
        dateFormat = SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ssZ")
    }

    fun objectToString(obj: Any): String {
        val objectMapper = createObjectMapper()
        return objectMapper.writeValueAsString(obj)
    }

    fun mockData(path: String): String = Files.readString(getFile(path).toPath())

    private fun getFile(path: String): File = File(this.javaClass.getResource(path).toURI())
}
