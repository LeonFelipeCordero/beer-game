import {Orders} from "../../types/order/Orders";
import {Board, Order, OrderState, Player} from "../../gql/graphql";
import {For, Show} from "solid-js";

function OrdersTable(props: { orders: Orders, player: Player, board: Board, deliver: Function }) {

    const deliverOrder = (orderId: string, boardId: string, amount: number) => {
        props.deliver(orderId, boardId, amount)
    }

    return (
        <div class="flex grid-rows-2">
            <div class="bg-slate-100 shadow-md rounded p-2 mr-5 w-full">
                <strong class="text-xl">
                    Incoming.
                </strong>
                <table class="border-separate border border-slate-500 table-fixed w-full">
                    <thead>
                    <tr>
                        <th class="border border-slate-600">#</th>
                        <th class="border border-slate-600">Quantity</th>
                        <th class="border border-slate-600">Status</th>
                    </tr>
                    </thead>
                    <tbody>
                    <For each={props.orders.value.filter(o => o.receiver?.id == props.player.id)}>
                        {(order: Order) => (
                            <tr class="p-2">
                                <td class="p-2 border border-slate-700">{order.id}</td>
                                <td class="p-2 border border-slate-700">{order.originalAmount}</td>
                                <td class="p-2 border border-slate-700">{order.state}</td>
                            </tr>
                        )}
                    </For>
                    </tbody>
                </table>
            </div>
            <div class="bg-slate-100 shadow-md rounded p-2 w-full">
                <strong class="text-xl">
                    Outgoing
                </strong>
                <div class="divide-y-4">
                    <table class="border-separate border border-slate-500 table-fixed w-full">
                        <thead>
                        <tr>
                            <th class="border border-slate-600">#</th>
                            <th class="border border-slate-600">Quantity</th>
                            <th class="border border-slate-600">Status</th>
                            <th class="border border-slate-600">Action</th>
                        </tr>
                        </thead>
                        <tbody>
                        <For each={props.orders.value.filter(o => o.sender?.id == props.player.id)}>
                            {(order: Order) => (
                                <tr>
                                    <td class="p-2 border border-slate-700">
                                        {order.id}
                                    </td>
                                    <td class="p-2 border border-slate-700">
                                        {order.originalAmount}
                                    </td>
                                    <td class="p-2 border border-slate-700">
                                        {order.state}
                                    </td>
                                    <Show
                                        when={order.state === OrderState.Pending}
                                        keyed
                                        fallback={
                                            <td class="p-2 border border-slate-600">N/A</td>
                                        }>
                                        <td class="p-2 border border-slate-600">
                                            <Show when={props.player.stock >= order.originalAmount} keyed
                                                  fallback={
                                                      <button disabled class="bg-gray-600 text-white py-1 px-2
                          rounded focus:outline-none focus:shadow-outline w-full">
                                                          No Stock
                                                      </button>
                                                  }>
                                                <button class="bg-blue-500 hover:bg-blue-700 text-white py-1 px-2
                          rounded focus:outline-none focus:shadow-outline w-full"
                                                        onclick={e => {
                                                            e.preventDefault()
                                                            deliverOrder(order.id, props.board.id, order.amount)
                                                        }}>
                                                    deliver
                                                </button>
                                            </Show>
                                        </td>
                                    </Show>
                                </tr>
                            )}
                        </For>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    )
}

export default OrdersTable