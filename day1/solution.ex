defmodule Solution do
  def part1() do
    File.stream!("input.txt")
    |> Enum.map(&String.trim/1)
    |> Enum.map(&String.to_integer/1)
    |> Enum.reduce(0, fn freq, accumulator -> accumulator + freq end)
    |> IO.inspect()
  end

  def part2() do
    File.stream!("input.txt")
    |> Stream.cycle()
    |> Enum.reduce_while({Map.new(), 0}, fn freq, {map, accumulator} ->
      {value, _} = Integer.parse(freq)

      if Map.has_key?(map, accumulator) do
        {:halt, accumulator}
      else
        updated_set = Map.put(map, accumulator, 1)
        accumulator = accumulator + value
        {:cont, {updated_set, accumulator}}
      end
    end)
    |> IO.inspect()
  end
end
