import {gql, TypedDocumentNode} from "@urql/core"

export const playerQueryType = {
    addPlayer: "addPlayer",
    getPlayer: "getPlayer",
    player: "player"
}

function PlayerQueries(query: string): TypedDocumentNode {
    const queries = new Map<string, TypedDocumentNode>([
        [playerQueryType.addPlayer, createPlayerMutation],
        [playerQueryType.getPlayer, getPlayerQuery],
        [playerQueryType.player, playerSubscription]
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
query getPlayer($boardId: String, $playerId: String) {
  getPlayer(boardId: $boardId, playerId: $playerId) {
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

const playerSubscription = gql`
subscription player($boardId: String, $playerId: String) {
  player(boardId: $boardId, playerId: $playerId) {
    id
    backlog
    stock
    lastOrder
    weeklyOrder
  }
}
`;