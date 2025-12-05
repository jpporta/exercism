defmodule GuessingGame do
  def compare(secret_number \\ nil, guess \\ :no_guess)

  def compare(secret_number, guess) when guess === :no_guess or is_nil(secret_number),
    do: "Make a guess"

  def compare(secret_number, guess) when secret_number === guess, do: "Correct"

  def compare(secret_number, guess)
      when secret_number === guess + 1 or secret_number === guess - 1,
      do: "So close"

  def compare(secret_number, guess) when secret_number > guess, do: "Too low"
  def compare(secret_number, guess) when secret_number < guess, do: "Too high"
end
