defmodule BoutiqueInventory do
  def sort_by_price(inventory) do
    inventory
    |> Enum.sort(fn item_a, item_b -> item_a[:price] <= item_b[:price] end)
  end

  def with_missing_price(inventory) do
    inventory
    |> Enum.filter(fn item -> item[:price] === nil end)
  end

  def increase_quantity(item, count) do
    %{
      item
      | quantity_by_size:
          Enum.map(item[:quantity_by_size], fn {key, value} -> {key, value + count} end)
          |> Map.new(fn {key, value} -> {key, value} end)
    }
  end

  def total_quantity(item) do
    item[:quantity_by_size]
    |> Enum.reduce(0, fn {key, value}, acc -> acc + value end)
  end
end
