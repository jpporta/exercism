defmodule HighScore do
  def new(), do: %{}

  def add_player(scores, name, score \\ 0), do: Map.put(scores, name, score)

  def remove_player(scores, name), do: Map.delete(scores, name)

  def reset_score(scores, name),
    do: if(name in scores, do: Map.replace(scores, name, 0), else: add_player(scores, name))

  def update_score(scores, name, score),
    do:
      Map.get_and_update(scores, name, fn current ->
        if is_nil(current), do: {current, score}, else: {current, current + score}
      end)
      |> (fn {_, scores} -> scores end).()

  def get_players(scores), do: Map.keys(scores)
end
