import { gql, TypedDocumentNode } from "@urql/core"

export const orderQueryType = {
  createOrder: "createOrder",
  newOrder: "newOrder",
  deliverOrder: "deliverOrder",
  orderDelivery: "orderDelivery"
}

function orderQueries(query: string): TypedDocumentNode {
  const queries = new Map<string, TypedDocumentNode>([
    [orderQueryType.createOrder, createOrderMutation],
    [orderQueryType.deliverOrder, deliverOrderMutation],
    [orderQueryType.newOrder, newOrderSubscription],
    [orderQueryType.orderDelivery, orderDeliverySubscription]
  ])
  if (queries.has(query)) {
    return queries.get(query)!!
  } else {
    throw Error("Query doesn't exist")
  }
}

export default orderQueries

const createOrderMutation = gql`
mutation createOrder($receiverId: String) {
  createOrder(receiverId: $receiverId) {
    id
    originalAmount
    state
    type
    createdAt
  }
}
`;

const deliverOrderMutation = gql`
mutation deliverOrder($orderId: String, $amount: Int) {
  deliverOrder(orderId: $orderId, amount: $amount) {
    message
    status
  }
}
`;

const newOrderSubscription = gql`
subscription newOrder($playerId: String) {
  newOrder(playerId: $playerId) {
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

const orderDeliverySubscription = gql`
subscription orderDelivery($playerId: String) {
  orderDelivery(playerId: $playerId) {
    id
    amount
    originalAmount
    state
  }
}
`;
