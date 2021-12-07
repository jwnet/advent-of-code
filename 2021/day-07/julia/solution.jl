using DelimitedFiles
using Statistics

function solve(data)
	x = 0.0
	pos1
	first = true
	for pos in data
		if first
			last = pos
			first = false
			continue
		end
		x = 
end

part1(data) = solve(data)
# part2(data) = 

sample = readdlm("sample", ',', Int)
# input = readdlm("input", ',', Int)

println("Part 1")
# println("  sample: ", part1(sample))
# println("  input : ", part1(input))
# println("Part 2")
# println("  sample: ", part2(sample))
# println("  input : ", part2(input))