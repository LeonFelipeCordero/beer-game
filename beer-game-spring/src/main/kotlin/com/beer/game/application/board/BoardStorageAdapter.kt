package com.beer.game.application.board

import com.beer.game.domain.Board

interface BoardStorageAdapter {
    fun loadBoard(boardId: String): Board
    fun loadBoardByName(name: String): Board?
    fun saveBoard(name: String): Board
    fun upsertBoard(board: Board): Board
    fun loadActiveBoards(): List<Board>
    fun loadBoardByPlayerId(playerId: String): Board
    fun loadBoardByOrderId(orderId: String): Board
}