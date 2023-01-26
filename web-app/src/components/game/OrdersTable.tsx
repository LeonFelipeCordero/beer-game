import {Orders} from "../../types/order/Orders";
import {Board, Order, OrderState, Player} from "../../gql/graphql";
import {For, Show} from "solid-js";

function OrdersTable(props: { orders: Orders, player: Player, board: Board, deliver: Function }) {

    const deliverOrder = (orderId: string, amount: number) => {
        props.deliver(orderId, amount)
    }

    return (
        <div class="flex grid-rows-2">
            <div class="bg-slate-100 shadow-md rounded p-5 mr-5 w-full">
                <strong class="text-xl">
                    Incoming.
                </strong>
                <div class="overflow-x-auto">
                    <table class="table w-full">
                        <thead>
                        <tr>
                            <th>Order Number</th>
                            <th>Quantity</th>
                            <th>Status</th>
                        </tr>
                        </thead>
                        <tbody>
                        <For each={props.orders.value.filter(o => o.receiver?.id == props.player.id)}>
                            {(order: Order) => (
                                <tr class="hover">
                                    <td>{order.id}</td>
                                    <td>{order.originalAmount}</td>
                                    <td>{order.state}</td>
                                </tr>
                            )}
                        </For>
                        </tbody>
                    </table>
                </div>
            </div>
            <div class="bg-slate-100 shadow-md rounded p-5 w-full">
                <strong class="text-xl">
                    Outgoing
                </strong>
                <div class="overflow-x-auto">
                    <table class="table w-full">
                        <thead>
                        <tr>
                            <th>Order Number</th>
                            <th>Quantity</th>
                            <th>Status</th>
                            <th>Action</th>
                        </tr>
                        </thead>
                        <tbody>
                        <For each={props.orders.value.filter(o => o.sender?.id == props.player.id)}>
                            {(order: Order) => (
                                <tr class="hover">
                                    <td>
                                        {order.id}
                                    </td>
                                    <td>
                                        {order.originalAmount}
                                    </td>
                                    <td>
                                        {order.state}
                                    </td>
                                    <Show
                                        when={order.state === OrderState.Pending}
                                        keyed
                                        fallback={
                                            <td>N/A</td>
                                        }>
                                        <td>
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
                                                            deliverOrder(order.id, order.amount)
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