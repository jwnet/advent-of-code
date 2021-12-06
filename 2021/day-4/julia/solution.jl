using DelimitedFiles

function splitCards(m)
    cardSize = size(m, 2)
    ncards = Int.(size(m, 1) / cardSize)
    cards = Matrix{Int}[]
    for i = 1:ncards
        last = i * cardSize
        first = last - cardSize + 1
        push!(cards, m[first:last, :])
    end
    return cards
end

winningRun(run) = findfirst(x -> x >= 0, run) === nothing
winningSpot(card, idx) = winningRun(card[idx[1], :]) || winningRun(card[:, idx[2]])
score(card, winningNumber) = sum(filter(x -> x > 0, card)) * winningNumber

function playBingo(cards, calls)
    ncards = size(cards, 1)
    winningCards = fill(false, 1, ncards)
    firstScore, lastScore, nwins = 0, 0, 0
    for num in calls
        for cardIdx in eachindex(cards)
            if winningCards[cardIdx]
                continue # skip cards already won
            end

            card = cards[cardIdx]

            numIdx = findfirst(x -> x == num, card)
            if numIdx === nothing
                continue # no luck, move on to next card
            end

            # mark spot on card with -1
            setindex!(card, -1, numIdx)

            if winningSpot(card, numIdx)
                lastScore = score(card, num)
                if nwins == 0
                    firstScore = lastScore
                end

                setindex!(winningCards, true, cardIdx)
                if (nwins += 1) == ncards
                    @goto gameOver
                end
            end
        end
    end
    @label gameOver
    println("part 1 // first winning score: ", firstScore)
    println("part 2 // last winning score: ", lastScore)
end

sampleCards = splitCards(readdlm("sample-cards", Int))
sampleNums = readdlm("sample-nums", ',', Int)
println("sample:")
playBingo(sampleCards, sampleNums)

cards = splitCards(readdlm("input-cards", Int))
nums = readdlm("input-nums", ',', Int)
println("input:")
playBingo(cards, nums)

# Result:
#   sample:
#   part 1 // first winning score: 4512
#   part 2 // last winning score: 1924
#   input:
#   part 1 // first winning score: 16674
#   part 2 // last winning score: 7075
