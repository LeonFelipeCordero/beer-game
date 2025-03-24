import { gql, TypedDocumentNode } from "@urql/core"

export const playerQueryType = {
  addPlayer: "addPlayer",
  getPlayer: "getPlayer",
  player: "player",
  updateWeeklyOrder: "updateWeeklyOrder"
}

function PlayerQueries(query: string): TypedDocumentNode {
  const queries = new Map<string, TypedDocumentNode>([
    [playerQueryType.addPlayer, createPlayerMutation],
    [playerQueryType.getPlayer, getPlayerQuery],
    [playerQueryType.player, playerSubscription],
    [playerQueryType.updateWeeklyOrder, updateWeeklyOrderMutation]
  ])
  if (queries.has(query)) {
    return queries.get(query)!!
  } else {
    throw Error("Query doesn't exist")
  }
}

export default PlayerQueries

const createPlayerMutation = gql`
mutation addPlayer($boardId: String, $role: Role) {
  addPlayer(boardId: $boardId, role: $role) {
    id
    role
  }
}
`;

const getPlayerQuery = gql`
query getPlayer($playerId: String) {
  getPlayer(playerId: $playerId) {
    id
    role
    stock
    backlog
    weeklyOrder
    lastOrder
    orders {
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
}
`;

const updateWeeklyOrderMutation = gql`
mutation updateWeeklyOrder($playerId: String, $amount: Int) {
  updateWeeklyOrder(playerId: $playerId, amount: $amount) {
    message
    status 
  }
}
`

const playerSubscription = gql`
subscription player($playerId: String) {
  player(playerId: $playerId) {
    id
    backlog
    stock
    lastOrder
    weeklyOrder
  }
}
`;
