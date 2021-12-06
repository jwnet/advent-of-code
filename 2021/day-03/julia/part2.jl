using DelimitedFiles

function oxyRating(nums, col)
    len = size(nums, 1)
    if len == 1
        return parse(Int, join(string.(nums[1, :])), base = 2)
    end
    keepBit = sum(nums[:, col]) < (len / 2) ? 0 : 1
    oxyRating(nums[nums[:, col].==keepBit, :], col + 1)
end

function co2Rating(nums, col)
    len = size(nums, 1)
    if len == 1
        return parse(Int, join(string.(nums[1, :])), base = 2)
    end
    keepBit = sum(nums[:, col]) < (len / 2) ? 1 : 0
    co2Rating(nums[nums[:, col].==keepBit, :], col + 1)
end

data = readdlm("input", Int)
println("life support rating: ", oxyRating(data, 1) * co2Rating(data, 1))
