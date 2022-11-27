package com.beer.game.domain.exceptions

class ControlledErrorResponse(status: Int, message: String, hint: String) : GenericErrorResponse(status, message)
