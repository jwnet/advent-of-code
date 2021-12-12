package main

type connectionData struct {
	a, b string
}

var sample1 = []connectionData {
	{"start","A"},
	{"start","b"},
	{"A","c"},
	{"A","b"},
	{"b","d"},
	{"A","end"},
	{"b","end"},
}
var sample2 = []connectionData {
	{"dc","end"},
	{"HN","start"},
	{"start","kj"},
	{"dc","start"},
	{"dc","HN"},
	{"LN","dc"},
	{"HN","end"},
	{"kj","sa"},
	{"kj","HN"},
	{"kj","dc"},
}
var sample3 = []connectionData {
	{"fs","end"},
	{"he","DX"},
	{"fs","he"},
	{"start","DX"},
	{"pj","DX"},
	{"end","zg"},
	{"zg","sl"},
	{"zg","pj"},
	{"pj","he"},
	{"RW","he"},
	{"fs","DX"},
	{"pj","RW"},
	{"zg","RW"},
	{"start","pj"},
	{"he","WI"},
	{"zg","he"},
	{"pj","fs"},
	{"start","RW"},
}
var input = []connectionData {
	{"QR","da"},
	{"QR","end"},
	{"QR","al"},
	{"start","op"},
	{"zh","iw"},
	{"zh","start"},
	{"da","PF"},
	{"op","bj"},
	{"iw","QR"},
	{"end","HR"},
	{"bj","PF"},
	{"da","LY"},
	{"op","PF"},
	{"bj","iw"},
	{"end","da"},
	{"bj","zh"},
	{"HR","iw"},
	{"zh","op"},
	{"zh","PF"},
	{"HR","bj"},
	{"start","PF"},
	{"HR","da"},
	{"QR","bj"},
}
