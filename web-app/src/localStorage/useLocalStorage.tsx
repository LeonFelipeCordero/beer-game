import { createEffect } from "solid-js";
import { createStore, SetStoreFunction } from "solid-js/store";

export type State = {
  board: string | undefined
  player: string | undefined
}

function createLocalStorage(initState: {}): [State, SetStoreFunction<{}>] {
  const [state, setState] = createStore(initState);
  if (localStorage.localState) setState(JSON.parse(localStorage.localState));
  createEffect(() => (localStorage.localState = JSON.stringify(state)));
  return [state as State, setState as SetStoreFunction<{}>];
}

export default createLocalStorage
