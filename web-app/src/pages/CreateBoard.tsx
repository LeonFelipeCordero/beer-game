import { useNavigate } from "solid-app-router";
import { createSignal, Show } from "solid-js";
import createLocalStorage from "../localStorage/useLocalStorage";
import { boardQueryType } from "../client/BoardQueries";
import boardClient from "../client/BoardClient";

function CreateBoard() {
  const navigate = useNavigate();
  const [boardName, setBoardName] = createSignal("")
  const [error, setError] = createSignal(undefined)
  const [_, setState] = createLocalStorage({ board: undefined, player: undefined });

  const createBoard = async () => {
    setError(undefined)
    boardClient.doMutation(boardQueryType.createBoard, { name: boardName() })
      .then(board => {
        setState({ board: board.id })
        navigate("/player/selection")
      })
      .catch(e => {
        setError(e)
        console.log(e)
      })
  }

  return (
    <div class="h-screen grid place-content-center">
      <div class="w-full max-w-xs">
        <form class="bg-slate-100 shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <div class="mb-4">
            <label for="name" class="block text-gray-700 text-xl font-bold mb-2">Board name.</label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              type="text"
              name="name"
              value={boardName()}
              placeholder="e.g. team-AB"
              autofocus
              autocomplete="off"
              onchange={(e) => setBoardName(e.currentTarget.value)}
            />
          </div>
          <div>
            <input
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
              value="Join Board"
              onclick={(e) => {
                e.preventDefault()
                createBoard()
              }}
            />
            <br />
            <Show when={error()} keyed>
              <span class="text-red-600">Something went wrong</span>
              <span class="text-red-600">{error()}</span>
            </Show>
            <br />
          </div>
        </form>
      </div>
    </div>
  )
}

export default CreateBoard
