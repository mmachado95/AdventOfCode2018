defmodule Solution do
  defp chars_ocurrence(id) do
    id
    |> String.codepoints()
    |> Enum.reduce(Map.new(), fn codepoint, map_accumulator ->
      Map.update(map_accumulator, codepoint, 1, & &1 + 1)
    end)
  end

  defp count_ocurrences(ids_ocurrences, value) do
    # Erlangs better approach
    # :maps.filter fn _, ocurrence -> ocurrence == value end, %{foo: "bar", biz: nil, baz: 4}
    Enum.map(ids_ocurrences, fn id ->
      Enum.filter(id, fn {_, ocurrence} -> value == ocurrence end)
    end)
    |> Enum.reject(&(&1 == []))
  end

  def part1() do
    char_id_ocurrence =
    File.stream!("input.txt")
    |> Enum.map(&String.trim/1)
    |> Enum.map(&chars_ocurrence/1)

    two = count_ocurrences(char_id_ocurrence, 2) |> Enum.count()
    three = count_ocurrences(char_id_ocurrence, 3) |> Enum.count()

    (two * three) |> IO.inspect()
  end

  def is_similar?(id1, id2) do
    difference = String.myers_difference(id1, id2)
    byte_size(difference[:del]) <= 1 or byte_size(difference[:ins]) <= 1
  end

  def similar_ids(id, [head|tail]) do
    case is_similar?(id, head) do
      true -> # como
      false -> similar_ids(id, tail)
    end

  end

  def part2() do

  end
end
