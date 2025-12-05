defmodule TakeANumber do
  def child_function(state) do
    receive do
      {:report_state, parentPid} ->
        send(parentPid, state)
        child_function(state)

      {:take_a_number, parentPid} ->
        send(parentPid, state + 1)
        child_function(state + 1)

      :stop ->
        nil

      _ ->
        nil
        child_function(state)
    end
  end

  def start() do
    spawn(TakeANumber, :child_function, [0])
  end
end
