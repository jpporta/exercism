defmodule LogLevel do
  def filter_log(nil, _) do
    {:unknown, false}
  end

  def filter_log({log, index}, level) do
    if index === level, do: log, else: {:unknown, false}
  end

  def get_log(level) do
    [
      {:trace, false},
      {:debug, true},
      {:info, true},
      {:warning, true},
      {:error, true},
      {:fatal, false}
    ]
    |> Enum.with_index()
    |> Enum.at(level)
    |> filter_log(level)
  end

  def to_label(level, legacy?) do
    cond do
      level === 0 -> if legacy?, do: :unknow, else: :trace
      level === 1 -> if legacy?, do: :debug
      level === 2 -> if legacy?, do: :info
      level === 3 -> if legacy?, do: :warning
      level === 4 -> if legacy?, do: :error
      level === 5 -> if legacy?, do: :fatal, else: :trace
      true -> :unknown
    end
  end

  def alert_recipient(level, legacy?) do
    label = to_label(level, legacy?)

    cond do
      label === :error or label === :fatal -> :ops
      label === :unknown and legacy? -> :dev1
      label === :unknown and not legacy? -> :dev2
      true -> false
    end
  end
end
