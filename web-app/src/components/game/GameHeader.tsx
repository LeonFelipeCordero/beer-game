function GameHeader(props: { boardName: string, playerRole: string }) {
  return (
    <div>
      <strong class="text-6xl ">
        Board {props.boardName}
      </strong>
      <p>
        Role {props.playerRole}
      </p>
      <button
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        Quit Game
      </button>
    </div>
  )
}

export default GameHeader
