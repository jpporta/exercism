# Use the Plot struct as it is provided
defmodule Plot do
  @enforce_keys [:plot_id, :registered_to]
  defstruct [:plot_id, :registered_to]
end

defmodule CommunityGarden do
  use Agent

  def start(opts \\ []) do
    Agent.start_link(fn -> %{next: 1, list: []} end, opts)
  end

  def list_registrations(pid) do
    Agent.get(pid, fn state -> state.list end)
  end

  defp add_one_to_state(old_state, register_to) do
    %{
      next: old_state.next + 1,
      list: [%Plot{plot_id: old_state.next, registered_to: register_to} | old_state.list]
    }
  end

  def register(pid, register_to) do
    Agent.get_and_update(pid, fn state ->
      {%Plot{plot_id: state.next, registered_to: register_to},
       add_one_to_state(state, register_to)}
    end)
  end

  def release(pid, plot_id) do
    Agent.update(pid, fn state ->
      %{
        next: state.next,
        list: Enum.filter(state.list, fn e -> e.plot_id != plot_id end)
      }
    end)
  end

  def get_registration(pid, plot_id) do
    Agent.get(pid, fn state ->
      Enum.find(
        state.list,
        {:not_found, "plot is unregistered"},
        fn e -> e.plot_id == plot_id end
      )
    end)
  end
end
