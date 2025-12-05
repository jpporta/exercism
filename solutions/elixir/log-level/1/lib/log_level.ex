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
    {label, suport} = get_log(level)

    if legacy? and not suport, do: :unknown, else: label
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
