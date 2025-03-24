import { gql, TypedDocumentNode } from "@urql/core"

export const boardQueryType = {
  createBoard: "createBoard",
  getBoardByName: "getBoardByName",
  getBoard: "getBoard",
  board: "board"
}

function BoardQueries(query: string): TypedDocumentNode {
  const queries = new Map<string, TypedDocumentNode>([
    [boardQueryType.createBoard, createBoardMutation],
    [boardQueryType.getBoardByName, getBoardByNameQuery],
    [boardQueryType.getBoard, getBoardQuery],
    [boardQueryType.board, boardSubscription]
  ])
  if (queries.has(query)) {
    return queries.get(query)!!
  } else {
    throw Error("Query doesn't exist")
  }
}

export default BoardQueries

const createBoardMutation = gql`
mutation createBoard($name: String) {
  createBoard(name: $name) {
    id
  }
}
`

const getBoardByNameQuery = gql`
query getBoardByName($name: String) {
  getBoardByName(name: $name) {
    id
    full
  }
}
`

const getBoardQuery = gql`
query getBoard($id: String) {
  getBoard(id: $id) {
    id
    name
    state
    full
    players {
      id
      role
    }
    availableRoles
  }
}
`;

const boardSubscription = gql`
subscription board($boardId: String) {
  board(boardId: $boardId) {
    id
    state
    full
    players {
        id
        role
    }
    availableRoles
  }
}
`;
