using DelimitedFiles
using DataStructures

function fishCountAfterNDays(data, n)
    fish = counter(data)
    for i = 1:n
        newCycle = counter(Int)
        for (timer, nfish) in fish
            if timer == 0
                newCycle[6] += nfish
                newCycle[8] = nfish
            else
                newCycle[timer-1] += nfish
            end
        end
        fish = newCycle
    end
    return sum(fish)
end

part1(data) = fishCountAfterNDays(data, 80)
part2(data) = fishCountAfterNDays(data, 256)

sample = readdlm("sample", ',', Int)
input = readdlm("input", ',', Int)

println("Part 1")
println("  sample: ", part1(sample))
println("  input : ", part1(input))
println("Part 2")
println("  sample: ", part2(sample))
println("  input : ", part2(input))

# Part 1
#   sample: 5934
#   input : 377263
# Part 2
#   sample: 26984457539
#   input : 1695929023803
