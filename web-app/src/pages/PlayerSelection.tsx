import {createSignal, For, Show} from "solid-js";
import createLocalStorage from "../localStorage/useLocalStorage";
import {Board, Role} from "../gql/graphql";
import boardClient from "../client/BoardClient";
import playerClient from "../client/PLayerClient";
import {useNavigate} from "solid-app-router";
import {playerQueryType} from "../client/PlayerQueries";
import {boardQueryType} from "../client/BoardQueries";

function PlayerSelection() {
    const navigate = useNavigate();
    const [state, setState] = createLocalStorage({board: undefined});
    const [loading, setLoading] = createSignal(false)
    const [board, setBoard] = createSignal<Board | undefined>(undefined)
    const [error, setError] = createSignal(undefined)

    boardClient.doQuery(boardQueryType.getBoard, {id: state.board})
        .then(result => {
            setError(undefined)
            setBoard(result)
        })

    boardClient.doSubscription(boardQueryType.board, {boardId: state.board}, setBoard)

    const selectPlayer = (role: String) => {
        setLoading(true)
        playerClient.doMutation(playerQueryType.addPlayer, {boardId: board()?.id, role: role})
            .then(player => {
                console.log(player)
                setState({player: player.id})
                setLoading(false)
                navigate("/game")
            })
            .catch(e => {
                setError(e)
                console.log(e)
                setLoading(false)
            })
    }

    return (
        <div>
            <Show
                when={board()}
                fallback={<button class="btn btn-primary w-full loading"/>}
                keyed>
                <div class="h-screen grid place-content-center">
                    <div class="w-full max-w-xs">
                        <div class="bg-slate-100 shadow-md rounded px-4 pt-3 mb-2">
                            <div class="mb-4">
                                <strong class="text-4xl ">
                                    Board {board()!!.name}
                                </strong>
                                <p class="text-xl font-bold text-cyan-600">Available roles:</p>
                                <Show when={board()!!.full}>
                                    <p class="font-bold">The board is full</p>
                                </Show>
                                <Show when={!loading()} keyed fallback={<h1>Loading...</h1>}>
                                    <For each={board()?.availableRoles}>
                                        {(item: Role) => (
                                            <p
                                                class="font-bold hover:cursor-pointer hover:text-blue-500"
                                                onclick={(e) => selectPlayer(e.currentTarget.innerText)}
                                            >{item}</p>
                                        )
                                        }
                                    </For>
                                </Show>
                            </div>
                        </div>
                    </div>
                </div>
            </Show>
            <Show when={error()} keyed>
                <span class="text-red-600">Something went wrong</span>
                <span class="text-red-600">{error()}</span>
            </Show>
        </div>
    )
}

export default PlayerSelection