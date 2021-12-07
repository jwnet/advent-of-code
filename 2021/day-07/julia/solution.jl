using DelimitedFiles
using Statistics


function part1(data)
	median = median(sort(data[1:,]))
	cost = 0.0
	for pos in data
		cost += abs(pos-median)
	end
	return Int.(cost)
end

# part2(data) = 

sample = readdlm("sample", ',', Int)
input = readdlm("input", ',', Int)

println("Part 1")
println("  sample: ", part1(sample))
println("  input : ", part1(input))
# println("Part 2")
# println("  sample: ", part2(sample))
# println("  input : ", part2(input))