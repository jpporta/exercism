defmodule Newsletter do
  def read_emails(path) do
    File.read!(path)
    |> String.split("\n")
    |> Enum.filter(&(&1 !== ""))
  end

  def open_log(path) do
    File.open!(path, [:append])
  end

  def log_sent_email(pid, email) do
    IO.write(pid, "#{email}\n")
  end

  def close_log(pid) do
    File.close(pid)
  end

  def should_log(:ok, email, log_path) do
    pid = open_log(log_path)
    log_sent_email(pid, email)
    close_log(pid)
  end

  def should_log(_, _email, _log_path), do: nil

  def send_newsletter(emails_path, log_path, send_fun) do
    File.write!(log_path, "")

    read_emails(emails_path)
    |> Enum.map(fn email -> send_fun.(email) |> should_log(email, log_path) end)

    :ok
  end
end
