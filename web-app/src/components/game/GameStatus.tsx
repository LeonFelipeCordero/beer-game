import {Player} from "../../gql/graphql";
import {Show} from "solid-js";

function GameStatus(props: { player: Player }) {
    return (
        <div class="mt-5">
            board
            <strong>Current status</strong>
            <table class="border-separate border border-slate-500 table-fixed w-full">
                <thead>
                <tr>
                    <th class="border border-slate-600">Stock</th>
                    <th class="border border-slate-600">Last week</th>
                </tr>
                </thead>
                <tbody>
                <tr>
                    <Show when={props.player.stock > 10} keyed
                          fallback={
                              <td class="border border-slate-700 text-red-600 font-bold">
                                  {props.player.stock}
                              </td>
                          }>
                        <td class="border border-slate-700">
                            {props.player.stock}
                        </td>
                    </Show>
                    <td class="border border-slate-700">
                        {props.player.lastOrder}
                    </td>
                </tr>
                </tbody>
            </table>
            <form phx-submit="create_order" class="flex flex-col mt-5">
                <Show when={props.player.role == "FACTORY"} keyed
                      fallback={
                          <label for="amount" class="text-gray-700 text-m font-bold mb-2">
                              Weekly order
                          </label>
                      }
                >
                    <label for="amount" class="text-gray-700 text-m font-bold mb-2">
                        Weekly production
                    </label>
                </Show>
                <input class="border rounded p-2 text-gray-700"
                       type="number"
                       name="amount"
                       value={props.player.weeklyOrder}/>
                <Show when={props.player.role != "FACTORY"} keyed>
                    <input
                        class="bg-blue-500 hover:bg-blue-700 text-white font-bold p-2 rounded mt-2"
                        type="submit"
                        value="order"
                        onClick={(e) => {
                            e.preventDefault()

                        }}
                    />
                </Show>
            </form>
        </div>
    )
}

export default GameStatus
