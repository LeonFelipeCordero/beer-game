import boardClient from "../client/BoardClient";
import {Board, Order, OrderState, Player} from "../gql/graphql";
import {boardQueryType} from "../client/BoardQueries";
import createLocalStorage from "../localStorage/useLocalStorage";
import {createSignal, For, Show} from "solid-js";
import Loading from "../components/common/Loading";
import WaitingForPlayersMessage from "../components/board/WaitingForPlayersMessage";
import playerClient from "../client/PLayerClient";
import {playerQueryType} from "../client/PlayerQueries";
import GameHeader from "../components/game/GameHeader";
import GameStatus from "../components/game/GameStatus";
import orderClient from "../client/OrderClient";
import orderQueries, {orderQueryType} from "../client/OrderQueries";
import {Orders} from "../types/order/Orders";
import OrdersTable from "../components/game/OrdersTable";

function Game() {
    const [state, _] = createLocalStorage({board: undefined, player: undefined});
    const [board, setBoard] = createSignal<Board | undefined>(undefined)
    const [player, setPlayer] = createSignal<Player | undefined>(undefined)
    const [orders, setOrders] = createSignal<Orders | undefined>(undefined)
    const [error, setError] = createSignal(undefined)

    const addOrder = (newOrder: Order) => {
        let onGoingOrders = orders()?.value!!
        onGoingOrders.push(newOrder)
        setOrders({value: onGoingOrders} as Orders)
    }

    boardClient.doQuery(boardQueryType.getBoard, {id: state.board})
        .then(result => {
            setBoard(result)
        })

    playerClient.doQuery(playerQueryType.getPlayer, {boardId: state.board, playerId: state.player})
        .then(result => {
            setOrders({value: result.orders} as Orders)
            result.orders = null
            setPlayer(result)
        })

    boardClient.doSubscription(boardQueryType.board, {boardId: state.board}, setBoard)

    orderClient.doSubscription(orderQueryType.newOrder, {boardId: state.board, playerId: state.player}, addOrder)

    const createOrder = (boardId: string, senderId: string, receiverId: string) => {
        orderClient.doQuery(orderQueryType.createOrder, {boardId: boardId, receiverId: receiverId})
    }

    return (
        <div>
            <Show when={board()} fallback={<Loading></Loading>} keyed>
                <Show when={board()?.full}
                      fallback={
                          <WaitingForPlayersMessage name={board()!!.name}></WaitingForPlayersMessage>
                      }>
                    <div class="h-screen">
                        <div class="flex justify-between p-5">
                            <Show when={player()} fallback={<Loading></Loading>} keyed>
                                <div class="bg-slate-100 shadow-md rounded p-2 mr-5 w-full">
                                    <div class="mb-4">
                                        <GameHeader boardName={board()!!.name}
                                                    playerRole={player()!!.role}></GameHeader>
                                        <GameStatus player={player()!!}></GameStatus>
                                    </div>
                                </div>
                                <OrdersTable orders={orders()!!} player={player()!!}></OrdersTable>
                            </Show>
                        </div>
                    </div>
                </Show>
            </Show>
            <Show when={error()} keyed>
                <span class="text-red-600">Something went wrong</span>
                <span class="text-red-600">{error()}</span>
            </Show>
        </div>
    )
}

export default Game
