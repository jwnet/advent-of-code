package main

import "fmt"

type game struct {
	p1, p2 *player
	d      *die
}

type die struct {
	val    int
	nrolls int
}

func (d *die) roll() int {
	d.nrolls += 1
	d.val = (d.val % 100) + 1
	return d.val
}

func (g *game) TakeTurns() (won bool) {
	if g.p1.turn(g.d) >= 1000 {
		return true
	}
	if g.p2.turn(g.d) >= 1000 {
		return true
	}
	return false
}

func (g *game) LosingScore() int {
	if g.p1.score <= g.p2.score {
		return g.p1.score
	}
	return g.p2.score
}

func (g *game) NRolls() int {
	return g.d.nrolls
}

type player struct {
	pos   int
	score int
}

func (p *player) turn(d *die) (score int) {
	moves := d.roll()
	moves += d.roll()
	moves += d.roll()
	p.pos = ((moves + p.pos - 1) % 10) + 1
	p.score += p.pos
	return p.score
}

func part1(g game) int {
	for !g.TakeTurns() {
	}
	return g.LosingScore() * g.NRolls()
}

// -------------------------------------------------------------------

type qGame struct {
	p1, p2 player
}

func (p *player) Split(roll int) player {
	np := player{}
	np.pos = ((p.pos + roll - 1) % 10) + 1
	np.score = np.pos + p.score
	return np
}

var rolls = [...]int{3, 4, 4, 4, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8, 8, 8, 9}

func playQ(g qGame) (p1Wins, p2Wins int) {
	var qGames = map[qGame]int{}
	qGames[g] = 1
	for len(qGames) > 0 {
		newGames := make(map[qGame]int, len(qGames))
		for ga, n := range qGames {
			p1 := ga.p1
			p2 := ga.p2

		ROLL1:
			for _, roll1 := range rolls {
				np1 := p1.Split(roll1)
				if np1.score >= 21 {
					p1Wins += n
					continue ROLL1
				}
			ROLL2:
				for _, roll2 := range rolls {
					np2 := p2.Split(roll2)
					if np2.score >= 21 {
						p2Wins += n
						continue ROLL2
					}
					newGames[qGame{p1: np1, p2: np2}] += n
				}
			}
		}
		qGames = newGames
	}
	return
}

func part2(g qGame) int {
	p1Wins, p2Wins := playQ(g)
	if p1Wins > p2Wins {
		return p1Wins
	}
	return p2Wins
}

func main() {
	sample := game{p1: &player{pos: 4}, p2: &player{pos: 8}, d: &die{}}
	input := game{p1: &player{pos: 3}, p2: &player{pos: 5}, d: &die{}}

	qsample := qGame{p1: player{pos: 4}, p2: player{pos: 8}}
	qinput := qGame{p1: player{pos: 3}, p2: player{pos: 5}}

	fmt.Println("Part One")
	fmt.Printf("\tsample: %d\n", part1(sample))
	fmt.Printf("\t input: %d\n", part1(input))
	fmt.Println("Part Two")
	fmt.Printf("\tsample: %d\n", part2(qsample))
	fmt.Printf("\t input: %d\n", part2(qinput))

	// Part One
	// 	sample: 739785
	// 	 input: 720750
	// Part Two
	// 	sample: 444356092776315
	// 	 input: 275067741811212
}
