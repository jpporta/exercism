defmodule BasketballWebsite do
  def extract_from_path(data, ""), do: data
  def extract_from_path(nil, _), do: nil

  def extract_from_path(data, path) do
    [head | tail] = String.split(path, ".")
    extract_from_path(data[head], Enum.join(tail, "."))
  end

  def get_in_path(data, path) do
    Kernel.get_in(data, String.split(path, "."))
  end
end
