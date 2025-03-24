import BoardQueries from "./BoardQueries";
import { runMutation, runQuery, runSubscription } from "./BeerGameClient";
import { Board } from "../gql/graphql";

const boardClient = {
  doQuery: doQuery,
  doMutation: doMutation,
  doSubscription: doSubscription
}

export default boardClient

async function doQuery(
  query: string,
  data: {} = {}
): Promise<Board> {
  return runQuery<Board>(BoardQueries(query), query, data)
    .catch(e => {
      console.log(e)
      throw Error("Something went wrong calling mutation")
    })
}

async function doMutation(
  query: string,
  data: {} = {},
): Promise<Board> {
  return runMutation<Board>(BoardQueries(query), query, data)
    .catch(e => {
      console.log(e)
      throw Error("Something went wrong calling mutation")
    })
}

function doSubscription(
  query: string,
  data: {} = {},
  updater: any
) {
  runSubscription<Board>(BoardQueries(query), query, data, updater)
}
