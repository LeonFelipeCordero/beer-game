import { runMutation, runQuery, runSubscription } from "./BeerGameClient";
import PlayerQueries from "./PlayerQueries";
import { Player } from "../gql/graphql";

const playerClient = {
  doQuery: doQuery,
  doMutation: doMutation,
  doSubscription: doSubscription
}

export default playerClient

async function doQuery(
  query: string,
  data: {} = {}
): Promise<Player> {
  return runQuery<Player>(PlayerQueries(query), query, data)
}

async function doMutation(
  query: string,
  data: {} = {},
): Promise<Player> {
  return runMutation<Player>(PlayerQueries(query), query, data)
}

function doSubscription(
  query: string,
  data: {} = {},
  updater: any
) {
  runSubscription<Player>(PlayerQueries(query), query, data, updater)
}
