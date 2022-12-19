import {gql, TypedDocumentNode} from "@urql/core"

export const orderQueryType = {
    createOrder: "createOrder",
    newOrder: "newOrder"
}

function orderQueries(query: string): TypedDocumentNode {
    const queries = new Map<string, TypedDocumentNode>([
        [orderQueryType.createOrder, createOrderMutation],
        [orderQueryType.newOrder, newOrderSubscription],
    ])
    if (queries.has(query)) {
        return queries.get(query)!!
    } else {
        throw Error("Query doesn't exist")
    }
}

export default orderQueries

const createOrderMutation = gql`
mutation createOrder($boardId: String, $receiverId: String) {
  createOrder(boardId: $boardId, receiverId: $receiverId) {
    id
    originalAmount
    state
    type
    createdAt
  }
}
`;

const newOrderSubscription = gql`
subscription newOrder($boardId: String, $playerId: String) {
  newOrder(boardId: $boardId, playerId: $playerId) {
    id
    originalAmount
    state
    createdAt
    receiver {
      id
    }
    sender {
      id
    }
  }
}
`;