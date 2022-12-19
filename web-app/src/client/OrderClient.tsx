import {runMutation, runQuery, runSubscription} from "./BeerGameClient";
import orderQueries from "./OrderQueries";
import {Order} from "../gql/graphql";

const orderClient = {
    doQuery: doQuery,
    doMutation: doMutation,
    doSubscription: doSubscription
}

export default orderClient

async function doQuery(
    query: string,
    data: {} = {}
): Promise<Order> {
    return runQuery<Order>(orderQueries(query), query, data)
}

async function doMutation(
    query: string,
    data: {} = {},
): Promise<Order> {
    return runMutation<Order>(orderQueries(query), query, data)
}

function doSubscription(
    query: string,
    data: {} = {},
    updater: any
) {
    runSubscription<Order>(orderQueries(query), query, data, updater)
}
