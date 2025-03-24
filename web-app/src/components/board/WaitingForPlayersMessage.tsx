function WaitingForPlayersMessage(props: { name: string }) {
  return (
    <div class="h-screen grid place-content-center">
      <div class="w-full max-w-xs">
        <div class="bg-slate-100 shadow-md rounded px-4 pt-3 mb-2">
          <div class="mb-4">
            <strong class="text-4xl ">
              Board {props.name}
            </strong>
            <p class="text-xl font-bold text-cyan-600">Waiting for other players</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default WaitingForPlayersMessage
