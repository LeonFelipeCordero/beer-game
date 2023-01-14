package com.beer.game.api

object GraphQlDocuments {

    private const val boardData = """
        id
        name
        state
        full
        finished
        createdAt
    """

    private const val playerData = """
        id
        name
        role
        stock
        backlog
        weeklyOrder
        lastOrder
        cpu
    """

    private const val orderData = """
        id
        amount
        originalAmount
        state
        type
        createdAt
    """

    val documentCreateBoard = """
            mutation createBoard(${'$'}name: String) {
                createBoard(name: ${'$'}name) {
                    $boardData                
                }
            }
        """.trimIndent()

    val documentGetBoard = """
             query getBoard(${'$'}id: String) {
                getBoard(id: ${'$'}id) {
                   $boardData 
                }
            }
        """.trimIndent()

    val documentAddPlayer = """
        mutation addPlayer(${'$'}boardId: String, ${'$'}role: Role) {
            addPlayer(boardId: ${'$'}boardId, role: ${'$'}role) {
                $playerData
                board {
                    $boardData
                }
            }
        }
    """.trimIndent()

    val documentAddPlayerMinimal = """
        mutation addPlayer(${'$'}boardId: String, ${'$'}role: Role) {
            addPlayer(boardId: ${'$'}boardId, role: ${'$'}role) {
                $playerData
            }
        }
    """.trimIndent()

    val documentCreateOrder = """
       mutation createOrder(${'$'}boardId: String, ${'$'}receiverId: String)  {
           createOrder(boardId: ${'$'}boardId, receiverId: ${'$'}receiverId) {
               $orderData
               sender {
                   $playerData
               }
               receiver {
                   $playerData
               }
               board {
                   $boardData
               }
           }
       }
    """.trimIndent()

}