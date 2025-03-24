import { createClient, defaultExchanges, subscriptionExchange, TypedDocumentNode } from "@urql/core";
import { createClient as createWSClient } from 'graphql-ws';
import { pipe, subscribe } from "wonka";

const wsClient = createWSClient({
  url: "ws://localhost:8080/graphql"
})

const apiClient = createClient({
  url: "http://localhost:8080/graphql",
  exchanges: [
    ...defaultExchanges,
    subscriptionExchange({
      forwardSubscription: (operation) => ({
        subscribe: (sink) => ({
          unsubscribe: wsClient.subscribe(operation, sink),
        }),
      }),
    }),
  ]
})

export async function runQuery<T>(
  query: TypedDocumentNode,
  queryName: string,
  data: {} = {}
): Promise<T> {
  return apiClient.query(query, data)
    .toPromise()
    .then(result => {
      return result.data!![queryName]
    })
    .then(result => result as T)
    .catch((e) => {
      console.log(e)
      throw new Error("Something went wrong")
    })
}

export async function runMutation<T>(
  query: TypedDocumentNode,
  queryName: string,
  data: {} = {},
): Promise<T> {
  return apiClient.mutation(query, data)
    .toPromise()
    .then(result => result.data!![queryName])
    .then(result => result as T)
    .catch((e) => {
      console.log(e)
      throw new Error("Something went wrong")
    })
}

export function runSubscription<T>(
  query: TypedDocumentNode,
  queryName: string,
  data: {} = {},
  updater: any
) {
  pipe(
    apiClient.subscription(query, data),
    subscribe((result) => {
      const entity = result.data!![queryName] as T
      updater(entity)
    })
  )
}
