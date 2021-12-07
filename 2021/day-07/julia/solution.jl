using DelimitedFiles
using Statistics

cost1(data, pos) = Int.(sum(map(x -> abs(x-pos), data)))
part1(data) = cost1(data, median(data))

cost2(data, pos) = Int.(sum(map(x -> abs(x-pos)*(abs(x-pos)+1)/2, data)))
function part2(data)
	l, r = first(data), last(data)
	cost = cost2(data, l)
	for p in (l+1):r
		c = cost2(data, p)
		c > cost && break
		cost = c
	end
	return cost
end

sample = sort(readdlm("sample", ',', Int)[1,:])
input = sort(readdlm("input", ',', Int)[1,:])

println("Part 1")
println("  sample: ", part1(sample))
println("  input : ", part1(input))
println("Part 2")
println("  sample: ", part2(sample))
println("  input : ", part2(input))
