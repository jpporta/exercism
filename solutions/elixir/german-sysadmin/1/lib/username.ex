defmodule Username do
  def sanitize(username) do
    username
    |> Enum.map(fn char ->
      case char do
        ?ä -> 'ae'
        ?ö -> 'oe'
        ?ü -> 'ue'
        ?ß -> 'ss'
        _ -> char
      end
    end)
    |> List.flatten()
    |> Enum.filter(fn char ->
      case char do
        char when char >= ?a and char <= ?z -> true
        ?_ -> true
        _ -> false
      end
    end)
  end
end
