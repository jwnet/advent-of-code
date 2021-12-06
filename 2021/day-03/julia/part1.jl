using DelimitedFiles

data = readdlm("input", Int)
count = sum(data, dims=1)
half = size(data, 1) / 2 # half the number of rows

gammaStr = join(map((x) -> x > half ? "1" : "0", count))
gamma = parse(Int, gammaStr, base=2)

epsilonStr = join(map((x) -> x < half ? "1" : "0", count))
epsilon = parse(Int, epsilonStr, base=2)

println("power consumption rate: ", gamma * epsilon) 
