defmodule LibraryFees do
  def datetime_from_string(string) do
    NaiveDateTime.from_iso8601!(string)
  end

  def before_noon?(datetime) do
    datetime.hour < 12
  end

  def return_date(checkout_datetime) do
    days = if before_noon?(checkout_datetime), do: 28, else: 29

    checkout_datetime
    |> NaiveDateTime.to_date()
    |> Date.add(days)
  end

  def days_late(planned_return_date, actual_return_datetime) do
    diff =
      NaiveDateTime.to_date(actual_return_datetime)
      |> Date.diff(planned_return_date)

    if diff < 0, do: 0, else: diff
  end

  def monday?(datetime) do
    NaiveDateTime.to_date(datetime)
    |> Date.day_of_week()
    |> Kernel.===(1)
  end

  def calculate_late_fee(checkout, return, rate) do
    checkout_datetime = datetime_from_string(checkout)
    return_datetime = datetime_from_string(return)

    fee =
      return_date(checkout_datetime)
      |> days_late(return_datetime)
      |> Kernel.*(rate)

    if monday?(return_datetime), do: floor(fee / 2), else: fee
  end
end
