package main

type formula struct {
	template []rune
	rules    map[pattern]rune
}

type pattern struct {
	l, r rune
}

var sample = formula{
	template: []rune{'N', 'N', 'C', 'B'},
	rules: map[pattern]rune{
		{'C', 'H'}: 'B',
		{'H', 'H'}: 'N',
		{'C', 'B'}: 'H',
		{'N', 'H'}: 'C',
		{'H', 'B'}: 'C',
		{'H', 'C'}: 'B',
		{'H', 'N'}: 'C',
		{'N', 'N'}: 'C',
		{'B', 'H'}: 'H',
		{'N', 'C'}: 'B',
		{'N', 'B'}: 'B',
		{'B', 'N'}: 'B',
		{'B', 'B'}: 'N',
		{'B', 'C'}: 'B',
		{'C', 'C'}: 'N',
		{'C', 'N'}: 'C',
	},
}

var input = formula{
	template: []rune{'K', 'B', 'K', 'P', 'H', 'K', 'H', 'H', 'N', 'B', 'C', 'V', 'C', 'H', 'P', 'S', 'P', 'N', 'H', 'F'},
	rules: map[pattern]rune{
		{'O', 'P'}: 'H',
		{'C', 'F'}: 'C',
		{'B', 'B'}: 'V',
		{'K', 'H'}: 'O',
		{'C', 'V'}: 'S',
		{'F', 'V'}: 'O',
		{'F', 'S'}: 'K',
		{'K', 'O'}: 'C',
		{'P', 'P'}: 'S',
		{'S', 'H'}: 'K',
		{'F', 'H'}: 'O',
		{'N', 'F'}: 'H',
		{'P', 'N'}: 'P',
		{'B', 'O'}: 'H',
		{'O', 'K'}: 'K',
		{'P', 'O'}: 'P',
		{'S', 'F'}: 'K',
		{'B', 'F'}: 'P',
		{'H', 'H'}: 'S',
		{'K', 'P'}: 'H',
		{'H', 'B'}: 'N',
		{'N', 'P'}: 'V',
		{'K', 'K'}: 'P',
		{'P', 'F'}: 'P',
		{'B', 'K'}: 'V',
		{'O', 'F'}: 'H',
		{'F', 'O'}: 'S',
		{'V', 'C'}: 'P',
		{'F', 'K'}: 'B',
		{'N', 'K'}: 'S',
		{'C', 'B'}: 'B',
		{'P', 'V'}: 'C',
		{'C', 'O'}: 'N',
		{'B', 'N'}: 'C',
		{'H', 'V'}: 'H',
		{'O', 'C'}: 'N',
		{'N', 'B'}: 'O',
		{'C', 'S'}: 'S',
		{'H', 'K'}: 'C',
		{'V', 'S'}: 'F',
		{'B', 'H'}: 'C',
		{'P', 'C'}: 'S',
		{'K', 'C'}: 'O',
		{'V', 'O'}: 'P',
		{'F', 'B'}: 'K',
		{'B', 'V'}: 'V',
		{'V', 'N'}: 'N',
		{'O', 'N'}: 'F',
		{'V', 'H'}: 'H',
		{'C', 'N'}: 'O',
		{'H', 'O'}: 'O',
		{'S', 'V'}: 'O',
		{'S', 'S'}: 'H',
		{'K', 'F'}: 'N',
		{'S', 'P'}: 'C',
		{'N', 'S'}: 'V',
		{'S', 'O'}: 'F',
		{'B', 'C'}: 'P',
		{'H', 'C'}: 'C',
		{'F', 'P'}: 'H',
		{'O', 'H'}: 'S',
		{'O', 'B'}: 'S',
		{'H', 'F'}: 'V',
		{'S', 'C'}: 'B',
		{'S', 'N'}: 'N',
		{'V', 'K'}: 'C',
		{'N', 'C'}: 'V',
		{'V', 'V'}: 'S',
		{'S', 'K'}: 'K',
		{'P', 'K'}: 'K',
		{'P', 'S'}: 'N',
		{'K', 'B'}: 'S',
		{'K', 'S'}: 'C',
		{'N', 'N'}: 'C',
		{'O', 'O'}: 'C',
		{'B', 'S'}: 'B',
		{'N', 'V'}: 'H',
		{'F', 'F'}: 'P',
		{'F', 'C'}: 'N',
		{'O', 'S'}: 'H',
		{'K', 'N'}: 'N',
		{'V', 'P'}: 'B',
		{'P', 'H'}: 'N',
		{'N', 'H'}: 'S',
		{'O', 'V'}: 'O',
		{'F', 'N'}: 'V',
		{'C', 'P'}: 'B',
		{'N', 'O'}: 'V',
		{'C', 'K'}: 'C',
		{'V', 'F'}: 'B',
		{'H', 'S'}: 'B',
		{'K', 'V'}: 'K',
		{'V', 'B'}: 'H',
		{'S', 'B'}: 'S',
		{'B', 'P'}: 'S',
		{'C', 'C'}: 'F',
		{'H', 'P'}: 'B',
		{'P', 'B'}: 'P',
		{'H', 'N'}: 'P',
		{'C', 'H'}: 'O',
	},
}