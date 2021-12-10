package main

type signal struct {
	patterns, output []string
}

var sample = []signal{
	{[]string{"be", "abcdefg", "bcdefg", "acdefg", "bceg", "cdefg", "abdefg", "bcdef", "abcdf", "bde"}, []string{"abcdefg", "bcdef", "bcdefg", "bceg"}},
	{[]string{"abdefg", "bcdeg", "bcg", "cg", "abcdefg", "bdefg", "abcdfg", "abcde", "bcdefg", "cefg"}, []string{"bcdefg", "bcg", "abcdefg", "cg"}},
	{[]string{"abdefg", "cg", "abcde", "abdfg", "abcdfg", "bcdefg", "abcdg", "acfg", "bcg", "abcdefg"}, []string{"cg", "cg", "abcdfg", "bcg"}},
	{[]string{"bcdefg", "bcd", "abcdef", "abdeg", "abcf", "bc", "acdef", "abcde", "acdefg", "abcdefg"}, []string{"abcdef", "abcde", "acdefg", "bc"}},
	{[]string{"abcdefg", "bfg", "fg", "abefg", "abdef", "cefg", "abceg", "abcefg", "abcdeg", "abcdfg"}, []string{"cefg", "abcdefg", "bfg", "abefg"}},
	{[]string{"abefg", "ac", "abcefg", "abcdefg", "acdefg", "bcdfg", "abce", "abdefg", "abcfg", "acf"}, []string{"abcdefg", "abce", "ac", "abcdefg"}},
	{[]string{"bcdfg", "dfg", "abcdefg", "cefg", "abdefg", "abcdef", "bcdef", "abcdg", "bcdefg", "fg"}, []string{"cefg", "bcdef", "cefg", "abcdefg"}},
	{[]string{"bcdefg", "abcefg", "bcefg", "acdefg", "abcdg", "de", "bdef", "cde", "abcdefg", "bcdeg"}, []string{"de", "abcefg", "abcdg", "bcefg"}},
	{[]string{"abdefg", "bcdefg", "cdeg", "abcef", "bcg", "abcdefg", "cg", "abcdfg", "bdefg", "bcefg"}, []string{"abcdefg", "bcg", "cg", "bcg"}},
	{[]string{"abcfg", "cfg", "abcdefg", "abceg", "fg", "abcdeg", "aefg", "abcefg", "abcdf", "bcdefg"}, []string{"aefg", "abcfg", "fg", "abceg"}},
}

var input = []signal{
	{[]string{"ce", "abcef", "abdef", "abdefg", "abcfg", "abcdef", "cdef", "abcdefg", "bce", "abcdeg"}, []string{"bce", "bce", "abcef", "ce"}},
	{[]string{"acdfg", "ae", "abce", "abcdefg", "bcefg", "aef", "bcdefg", "abdefg", "acefg", "abcefg"}, []string{"bcdefg", "abce", "acdfg", "abcdefg"}},
	{[]string{"defg", "acdef", "acdefg", "eg", "acdeg", "abcdg", "abcdefg", "abcdef", "ceg", "abcefg"}, []string{"acdeg", "defg", "abcdg", "acdef"}},
	{[]string{"ae", "acdefg", "abef", "ace", "abcdefg", "abcfg", "bcdeg", "abcdfg", "abceg", "abcefg"}, []string{"abcdefg", "abcdefg", "abceg", "bcdeg"}},
	{[]string{"bcdefg", "bdf", "abcdefg", "abcefg", "abcdf", "abcef", "bd", "abcdef", "abde", "acdfg"}, []string{"bdf", "abcef", "abcdf", "abcef"}},
	{[]string{"abcf", "afg", "abefg", "bcdefg", "abcdefg", "abcefg", "acdefg", "af", "bcefg", "abdeg"}, []string{"abdeg", "abcdefg", "af", "abcdefg"}},
	{[]string{"adefg", "abcfg", "acde", "abdefg", "acdfg", "acdefg", "cd", "bcdefg", "cdf", "abcdefg"}, []string{"abdefg", "acdfg", "abcdefg", "acde"}},
	{[]string{"bcdeg", "cd", "bcd", "abcdefg", "abcdeg", "abcdfg", "acde", "bcefg", "abdefg", "abdeg"}, []string{"bcd", "abcdeg", "abcdeg", "bcdeg"}},
	{[]string{"ae", "bcdeg", "bcdefg", "ace", "abcefg", "abcdefg", "adeg", "abcde", "abcdf", "abcdeg"}, []string{"ace", "abcdeg", "ace", "adeg"}},
	{[]string{"abcfg", "bcdef", "ae", "abcdefg", "abcefg", "abdefg", "abcef", "abe", "aceg", "abcdfg"}, []string{"ae", "abcdfg", "abcdfg", "bcdef"}},
	{[]string{"abcd", "cd", "bcefg", "abcdefg", "bcdef", "abcdef", "abdefg", "cdf", "abdef", "acdefg"}, []string{"cd", "cdf", "acdefg", "cd"}},
	{[]string{"abcdeg", "bcf", "abceg", "cf", "abcef", "abdef", "abcdfg", "cefg", "abcefg", "abcdefg"}, []string{"abcdefg", "cf", "cf", "abcdeg"}},
	{[]string{"bcefg", "abcdeg", "abcefg", "bcdefg", "abcdefg", "bdef", "acdfg", "deg", "de", "cdefg"}, []string{"bcefg", "abcdeg", "abcdefg", "bdef"}},
	{[]string{"adefg", "cef", "abcde", "cdfg", "cf", "acdef", "abdefg", "acdefg", "abcdefg", "abcefg"}, []string{"abdefg", "adefg", "abdefg", "abcefg"}},
	{[]string{"acfg", "abcdfg", "cf", "cdf", "bcdfg", "abcdeg", "bdefg", "abcdefg", "abcdef", "abcdg"}, []string{"abcdg", "cdf", "bdefg", "cdf"}},
	{[]string{"befg", "abcdeg", "abcdefg", "abcdef", "aeg", "abcefg", "acefg", "eg", "acdfg", "abcef"}, []string{"abcdefg", "abcefg", "abcefg", "befg"}},
	{[]string{"abcdefg", "abcefg", "abef", "bfg", "acefg", "bcefg", "bf", "acdefg", "bcdeg", "abcdfg"}, []string{"abcdefg", "abef", "abcdefg", "acefg"}},
	{[]string{"cdf", "df", "adef", "abcef", "abcdefg", "bcdefg", "abcdf", "abcdg", "abcdef", "abcefg"}, []string{"abcef", "cdf", "cdf", "abcefg"}},
	{[]string{"abdg", "abceg", "abcdefg", "abcef", "abcdeg", "acdefg", "ag", "bcdeg", "aeg", "bcdefg"}, []string{"abdg", "abdg", "ag", "aeg"}},
	{[]string{"aceg", "abcdefg", "abcdef", "bcefg", "bce", "abcdfg", "abcefg", "ce", "abcfg", "bdefg"}, []string{"ce", "abcdfg", "abcfg", "abcdefg"}},
	{[]string{"abcefg", "bdeg", "bcg", "abcdg", "acdefg", "abcdefg", "abcdeg", "abcdf", "acdeg", "bg"}, []string{"bg", "abcdefg", "bcg", "bcg"}},
	{[]string{"bcefg", "abcdefg", "abceg", "abcefg", "bcdfg", "bef", "ef", "abcdeg", "abcdef", "aefg"}, []string{"aefg", "abcdefg", "bcdfg", "bcefg"}},
	{[]string{"abcefg", "abcdg", "abdefg", "ce", "abefg", "bcef", "acdefg", "abcdefg", "abceg", "ace"}, []string{"abcefg", "abcefg", "abdefg", "abceg"}},
	{[]string{"acdefg", "acfg", "acdeg", "abcdef", "abcdefg", "cdg", "cg", "bcdefg", "abdeg", "acdef"}, []string{"acdef", "bcdefg", "acfg", "abcdefg"}},
	{[]string{"bcdefg", "abcdeg", "bcdeg", "abcefg", "bcf", "bcdef", "abcdefg", "bf", "bdfg", "acdef"}, []string{"bcf", "abcdefg", "bf", "bcdef"}},
	{[]string{"adefg", "abcdef", "bd", "bcdg", "abd", "abdeg", "abcdefg", "abceg", "abcdeg", "abcefg"}, []string{"adefg", "bd", "abcdeg", "abcdef"}},
	{[]string{"abcdg", "bdfg", "abcdef", "bg", "bcg", "abcdfg", "abcdefg", "abcefg", "abcdf", "acdeg"}, []string{"acdeg", "bcg", "abcdefg", "abcdg"}},
	{[]string{"bcdef", "abdefg", "abcdef", "bd", "abcdefg", "acdef", "bdf", "abcd", "bcefg", "acdefg"}, []string{"abdefg", "acdef", "bd", "bdf"}},
	{[]string{"abdefg", "abcdefg", "acdeg", "bcdeg", "abcefg", "acdefg", "adefg", "acdf", "ac", "ace"}, []string{"adefg", "ace", "adefg", "acdeg"}},
	{[]string{"abdef", "def", "abcde", "bcdefg", "aefg", "abdefg", "ef", "abdfg", "abcdfg", "abcdefg"}, []string{"abcdfg", "def", "aefg", "abcde"}},
	{[]string{"bcde", "abcdefg", "abdefg", "acdef", "cd", "acd", "abcdef", "abcdfg", "acefg", "abdef"}, []string{"acefg", "bcde", "acd", "cd"}},
	{[]string{"bcefg", "abg", "abcde", "abcdef", "acdg", "abcdefg", "abcdeg", "ag", "abceg", "abdefg"}, []string{"ag", "bcefg", "ag", "abcdef"}},
	{[]string{"bcdf", "cdefg", "bcdeg", "abceg", "bcdefg", "bde", "abcdefg", "acdefg", "abdefg", "bd"}, []string{"bd", "bde", "abcdefg", "bde"}},
	{[]string{"adeg", "def", "abdef", "abdefg", "de", "abefg", "abcefg", "abcdefg", "bcdefg", "abcdf"}, []string{"de", "abcefg", "adeg", "abcdefg"}},
	{[]string{"abcef", "abcdef", "efg", "bcfg", "acdeg", "abcefg", "acefg", "fg", "abcdefg", "abdefg"}, []string{"efg", "abcef", "acefg", "acdeg"}},
	{[]string{"acdef", "abcdef", "acefg", "cg", "acdefg", "cdfg", "abcdeg", "acg", "abcdefg", "abefg"}, []string{"acg", "acefg", "cdfg", "acg"}},
	{[]string{"ace", "abceg", "bcef", "abcefg", "abefg", "ce", "abcdg", "abdefg", "acdefg", "abcdefg"}, []string{"abcdefg", "ace", "ace", "ace"}},
	{[]string{"abefg", "abf", "af", "abceg", "abcefg", "abcdfg", "acef", "abcdeg", "abcdefg", "bdefg"}, []string{"abcdefg", "abcefg", "abcefg", "abf"}},
	{[]string{"bdg", "cdefg", "abcfg", "abdefg", "bd", "bcdfg", "bcde", "abcdefg", "acdefg", "bcdefg"}, []string{"bd", "abcdefg", "bcde", "bd"}},
	{[]string{"abcef", "abcdeg", "abg", "ag", "abceg", "acdg", "abdefg", "bcdefg", "abcdefg", "bcdeg"}, []string{"abcdefg", "abcef", "acdg", "abg"}},
	{[]string{"acde", "acdfg", "ac", "bcdefg", "acdefg", "abcefg", "acf", "abcdefg", "abdfg", "cdefg"}, []string{"abcdefg", "abcdefg", "acf", "abcdefg"}},
	{[]string{"abcdefg", "abdefg", "abcde", "bdefg", "abfg", "acdefg", "bcdefg", "aef", "abdef", "af"}, []string{"aef", "abcde", "abdefg", "abdefg"}},
	{[]string{"bcd", "abcdeg", "abdeg", "cdefg", "bcdeg", "abcdefg", "abcdfg", "abdefg", "abce", "bc"}, []string{"abdeg", "bc", "abce", "abcdeg"}},
	{[]string{"abcefg", "abcdf", "abcde", "abcfg", "abcdefg", "abcdfg", "adf", "cdfg", "abdefg", "df"}, []string{"abcdefg", "df", "df", "abcefg"}},
	{[]string{"bcdeg", "ae", "abcdefg", "aeg", "abcefg", "abdeg", "abcdfg", "adef", "abdfg", "abdefg"}, []string{"abdeg", "aeg", "ae", "aeg"}},
	{[]string{"abcdef", "abde", "acdef", "ace", "abcdefg", "ae", "acdfg", "bcdef", "abcefg", "bcdefg"}, []string{"ae", "abcefg", "acdef", "abde"}},
	{[]string{"bcdefg", "acdefg", "bdefg", "abdefg", "adf", "abdef", "abcde", "af", "abfg", "abcdefg"}, []string{"abdef", "abcdefg", "af", "bdefg"}},
	{[]string{"bcdfg", "abcdefg", "ef", "efg", "abcdfg", "acdefg", "bcdefg", "bdef", "abceg", "bcefg"}, []string{"ef", "abcdefg", "bcefg", "abcdefg"}},
	{[]string{"acdeg", "acdf", "abcdeg", "cf", "cef", "cdefg", "abcdefg", "acdefg", "bdefg", "abcefg"}, []string{"cef", "acdefg", "acdf", "abcdefg"}},
	{[]string{"abdfg", "bceg", "abcdef", "be", "acefg", "abe", "acdefg", "abefg", "abcdefg", "abcefg"}, []string{"abcdefg", "abefg", "abe", "abdfg"}},
	{[]string{"abcdef", "abcfg", "abdefg", "bcdefg", "eg", "efg", "bcefg", "cdeg", "abcdefg", "bcdef"}, []string{"abcdefg", "eg", "cdeg", "abcfg"}},
	{[]string{"ace", "bcef", "ce", "acdfg", "abdefg", "acefg", "abcdefg", "abcdeg", "abcefg", "abefg"}, []string{"acefg", "ce", "acdfg", "abcdefg"}},
	{[]string{"abcef", "bcdefg", "abcdefg", "acdg", "cfg", "abcdfg", "abdfg", "abdefg", "cg", "abcfg"}, []string{"cfg", "abcdefg", "abcdefg", "acdg"}},
	{[]string{"abcdfg", "dfg", "abcdf", "fg", "acdefg", "abdfg", "abcdefg", "abcdef", "bcfg", "abdeg"}, []string{"abcdefg", "abdfg", "abcdefg", "abcdef"}},
	{[]string{"bcefg", "abefg", "abcdefg", "cef", "ce", "bcdfg", "abcefg", "abdefg", "aceg", "abcdef"}, []string{"abefg", "abcefg", "bcdfg", "abefg"}},
	{[]string{"bcf", "bcdeg", "bcdefg", "bf", "bdfg", "bcefg", "acefg", "abcdeg", "abcdefg", "abcdef"}, []string{"abcdefg", "abcdefg", "bcf", "bf"}},
	{[]string{"dg", "bdg", "abcdefg", "adeg", "abcdg", "abceg", "abcdeg", "abcefg", "bcdefg", "abcdf"}, []string{"abcdefg", "adeg", "abcefg", "dg"}},
	{[]string{"abdef", "abdfg", "adg", "abcdeg", "bcdfg", "bcdefg", "acfg", "abcdefg", "abcdfg", "ag"}, []string{"acfg", "adg", "abcdefg", "adg"}},
	{[]string{"abdeg", "abcdfg", "acdfg", "acdefg", "abf", "bf", "abcdefg", "abcdef", "bcfg", "abdfg"}, []string{"abf", "abcdefg", "bf", "abdeg"}},
	{[]string{"bdef", "abdefg", "abcefg", "adefg", "abdfg", "bf", "abcdefg", "abf", "abcdg", "acdefg"}, []string{"abf", "abcdefg", "bf", "bdef"}},
	{[]string{"acdefg", "abcdfg", "ab", "abcfg", "acdfg", "abcdefg", "abcdeg", "bcefg", "abdf", "abc"}, []string{"acdefg", "abcdfg", "abcfg", "acdfg"}},
	{[]string{"abcdef", "abcg", "abdefg", "bc", "abcdefg", "abefg", "cdefg", "bcf", "abcefg", "bcefg"}, []string{"abcdef", "bcefg", "bcefg", "cdefg"}},
	{[]string{"abcdfg", "abcdef", "cefg", "bcdef", "bfg", "bcdefg", "fg", "bdefg", "abdeg", "abcdefg"}, []string{"abcdfg", "cefg", "fg", "abcdefg"}},
	{[]string{"acdg", "bcdefg", "abcdeg", "abcdefg", "cd", "abcefg", "abcde", "abceg", "bcd", "abdef"}, []string{"abcdeg", "bcd", "bcd", "abcdefg"}},
	{[]string{"abcdefg", "adg", "acdefg", "dg", "acefg", "cdeg", "abcefg", "abdef", "abcdfg", "adefg"}, []string{"abcefg", "cdeg", "adg", "cdeg"}},
	{[]string{"bcdefg", "abdefg", "adefg", "acdefg", "acdef", "acdg", "cd", "cdf", "abcef", "abcdefg"}, []string{"adefg", "acdefg", "abcdefg", "cd"}},
	{[]string{"bcfg", "abcdfg", "acdfg", "abdeg", "acdefg", "bf", "abdfg", "abcdef", "abf", "abcdefg"}, []string{"abf", "bf", "bcfg", "abcdfg"}},
	{[]string{"abcefg", "abdefg", "abc", "acef", "abceg", "abefg", "abcdfg", "ac", "bcdeg", "abcdefg"}, []string{"ac", "ac", "abc", "abc"}},
	{[]string{"abfg", "acdef", "abdefg", "abcdefg", "ag", "aeg", "bdefg", "abcdeg", "adefg", "bcdefg"}, []string{"abcdeg", "aeg", "abfg", "bcdefg"}},
	{[]string{"abcfg", "ef", "abdeg", "abefg", "acef", "bef", "bcdefg", "abcdfg", "abcefg", "abcdefg"}, []string{"abcdefg", "abcfg", "abcefg", "ef"}},
	{[]string{"adfg", "abceg", "cdg", "abcdef", "abcdfg", "dg", "abcdefg", "abcdg", "bcdefg", "abcdf"}, []string{"adfg", "dg", "cdg", "cdg"}},
	{[]string{"bcdefg", "abcdefg", "cdf", "acdef", "abdefg", "abcdef", "acdeg", "abdef", "cf", "abcf"}, []string{"acdeg", "abcdefg", "cf", "cdf"}},
	{[]string{"ace", "abcdfg", "abef", "abceg", "abcdefg", "acdefg", "abcfg", "ae", "bcdeg", "abcefg"}, []string{"acdefg", "ace", "abcdefg", "abceg"}},
	{[]string{"eg", "abcdef", "abcefg", "abcef", "acdefg", "befg", "abceg", "abcdg", "ceg", "abcdefg"}, []string{"abcef", "ceg", "abcdefg", "abcdefg"}},
	{[]string{"cd", "acefg", "abdefg", "cdf", "abdef", "abcd", "acdef", "abcdefg", "bcdefg", "abcdef"}, []string{"cdf", "acefg", "abcdefg", "abdef"}},
	{[]string{"bcdef", "bef", "abcdef", "abcefg", "abcde", "cdefg", "abcdefg", "abdf", "bf", "abcdeg"}, []string{"cdefg", "bf", "bef", "bf"}},
	{[]string{"abcef", "abcefg", "bf", "abdefg", "abcdeg", "abcdefg", "abf", "abceg", "acdef", "bcfg"}, []string{"abcefg", "abcdeg", "bcfg", "abcefg"}},
	{[]string{"abcdeg", "abcdef", "abcdefg", "acefg", "cde", "bcdf", "acdef", "cd", "abdef", "abdefg"}, []string{"abcdeg", "abdefg", "abcdefg", "abcdefg"}},
	{[]string{"ae", "bcdeg", "abcdefg", "abcdeg", "abcdef", "abcfg", "abceg", "ace", "bcdefg", "adeg"}, []string{"ace", "abceg", "abcdeg", "ae"}},
	{[]string{"abcdefg", "acdefg", "abcefg", "cde", "bcdfg", "cdefg", "adef", "acefg", "de", "abcdeg"}, []string{"adef", "adef", "cde", "cdefg"}},
	{[]string{"acdefg", "abcdefg", "bcdfg", "abdfg", "afg", "abdeg", "abef", "af", "abcdeg", "abdefg"}, []string{"abcdefg", "abef", "af", "abcdefg"}},
	{[]string{"bc", "abce", "abcdfg", "bdefg", "acdefg", "abcefg", "abcdefg", "bcg", "acefg", "bcefg"}, []string{"bc", "bc", "acefg", "bc"}},
	{[]string{"abcdeg", "abcefg", "cdefg", "abcdfg", "abcdg", "ae", "aeg", "abcdefg", "acdeg", "abde"}, []string{"abcdefg", "abcdefg", "aeg", "ae"}},
	{[]string{"abdeg", "abdef", "adfg", "abcefg", "abcdeg", "abf", "af", "abcdefg", "abdefg", "bcdef"}, []string{"abdefg", "adfg", "abcdeg", "af"}},
	{[]string{"bcg", "abcdefg", "bcefg", "acdefg", "abcefg", "bcdef", "acefg", "bg", "abeg", "abcdfg"}, []string{"bg", "abeg", "bcefg", "bcdef"}},
	{[]string{"abg", "abcef", "acdeg", "abcefg", "abcdfg", "abcdefg", "abcdef", "abceg", "befg", "bg"}, []string{"abcdfg", "bg", "bg", "abg"}},
	{[]string{"abcdefg", "acdefg", "defg", "acdeg", "abcdeg", "acefg", "abcdfg", "abcef", "fg", "cfg"}, []string{"defg", "acdeg", "cfg", "acdefg"}},
	{[]string{"abcdefg", "abdeg", "acdeg", "abcd", "abcdeg", "acdefg", "abg", "ab", "bdefg", "abcefg"}, []string{"abcdeg", "abcd", "abcdefg", "abcd"}},
	{[]string{"bcdfg", "bcdefg", "bcf", "bc", "bcde", "cdefg", "acdefg", "abcdefg", "abdfg", "abcefg"}, []string{"bcf", "bcf", "bc", "abcdefg"}},
	{[]string{"bcefg", "abdeg", "af", "abcdef", "bcdefg", "abcefg", "abcdefg", "abefg", "aef", "acfg"}, []string{"bcdefg", "abcdefg", "af", "acfg"}},
	{[]string{"abcef", "eg", "abcefg", "abcdeg", "abcdefg", "abdfg", "cefg", "abcdef", "beg", "abefg"}, []string{"abcdef", "abcdefg", "abcef", "cefg"}},
	{[]string{"acdg", "abdfg", "abcdefg", "abcdfg", "bcdef", "cg", "bcdfg", "cfg", "abcefg", "abdefg"}, []string{"abcdefg", "abcdefg", "abcdefg", "bcdef"}},
	{[]string{"bdeg", "abcdefg", "bcdefg", "abcdfg", "cdefg", "bcdfg", "def", "de", "abcdef", "acefg"}, []string{"bdeg", "bcdfg", "bdeg", "bdeg"}},
	{[]string{"acdfg", "acefg", "dfg", "df", "abcdfg", "abcdeg", "abcdg", "bcdefg", "abdf", "abcdefg"}, []string{"abcdefg", "bcdefg", "abcdefg", "abdf"}},
	{[]string{"abcdef", "abcdefg", "ab", "acdeg", "abdefg", "abd", "bdefg", "abfg", "bcdefg", "abdeg"}, []string{"abfg", "abcdef", "abfg", "abd"}},
	{[]string{"bcf", "cf", "abcdeg", "abcdfg", "bdefg", "abcdg", "abcefg", "bcdfg", "abcdefg", "acdf"}, []string{"bdefg", "cf", "bcf", "abcdfg"}},
	{[]string{"bcdefg", "abdeg", "ef", "acef", "acdefg", "def", "adefg", "abcdfg", "abcdefg", "acdfg"}, []string{"def", "acdefg", "acdefg", "ef"}},
	{[]string{"abdef", "abcefg", "ae", "acde", "abcdfg", "abcdf", "abe", "bdefg", "abcdefg", "abcdef"}, []string{"abe", "abcefg", "abcdef", "ae"}},
	{[]string{"abcef", "abefg", "abdefg", "adefg", "bg", "bdeg", "acdefg", "abcdfg", "abg", "abcdefg"}, []string{"abg", "bdeg", "bdeg", "abcdefg"}},
	{[]string{"adefg", "ab", "abefg", "abcf", "abcefg", "bcefg", "abcdeg", "abcdefg", "abg", "bcdefg"}, []string{"ab", "abefg", "abcdefg", "abcdefg"}},
	{[]string{"abdg", "bcdeg", "beg", "abcdef", "abcefg", "abcdeg", "abcdefg", "cdefg", "abcde", "bg"}, []string{"abdg", "bg", "beg", "abcdeg"}},
	{[]string{"abcdeg", "cdeg", "abcdefg", "abdeg", "abdefg", "abcdf", "cg", "abcdg", "abcefg", "bcg"}, []string{"abcefg", "abdeg", "cdeg", "abcdeg"}},
	{[]string{"abcdefg", "abdefg", "abdef", "defg", "abefg", "abcefg", "bdf", "df", "abcdfg", "abcde"}, []string{"abcdefg", "abdefg", "abdefg", "df"}},
	{[]string{"abcdef", "abdfg", "abcef", "bcde", "cdf", "abcdefg", "abcefg", "cd", "acdefg", "abcdf"}, []string{"acdefg", "bcde", "cd", "abcdef"}},
	{[]string{"aceg", "abcdef", "adefg", "acdefg", "efg", "abdfg", "acdef", "abcdefg", "eg", "bcdefg"}, []string{"adefg", "acdef", "acdefg", "abcdefg"}},
	{[]string{"abcdefg", "abcefg", "abcdef", "abe", "aefg", "ae", "bcdeg", "abcfg", "abcdfg", "abceg"}, []string{"bcdeg", "aefg", "ae", "aefg"}},
	{[]string{"bcdefg", "adf", "abdefg", "abdfg", "abcdg", "af", "aefg", "abcdef", "bdefg", "abcdefg"}, []string{"bdefg", "aefg", "abcdefg", "abdefg"}},
	{[]string{"abcde", "abceg", "acdf", "abdefg", "bcdef", "abcdefg", "bcdefg", "abd", "ad", "abcdef"}, []string{"ad", "acdf", "abcdef", "abd"}},
	{[]string{"abcdfg", "abcdefg", "acdeg", "abefg", "bdg", "abdefg", "bd", "abdeg", "abcefg", "bdef"}, []string{"bdg", "bd", "bd", "bd"}},
	{[]string{"bdg", "abcdefg", "abcdef", "abcdfg", "bdefg", "dg", "adeg", "abdef", "abdefg", "bcefg"}, []string{"abcdefg", "abdefg", "adeg", "adeg"}},
	{[]string{"bcdefg", "abdefg", "abcdefg", "ad", "bdefg", "abde", "acefg", "adefg", "adf", "abcdfg"}, []string{"ad", "abde", "adefg", "ad"}},
	{[]string{"abdfg", "abcfg", "adefg", "bcdf", "abcdeg", "abd", "abcdfg", "bd", "abcefg", "abcdefg"}, []string{"abd", "abcdefg", "adefg", "abcdefg"}},
	{[]string{"abceg", "abcdef", "abcfg", "eg", "abcdefg", "adeg", "abcde", "beg", "bcdefg", "abcdeg"}, []string{"abcdefg", "beg", "eg", "beg"}},
	{[]string{"abcdefg", "bdfg", "cdg", "dg", "abceg", "bcdefg", "abcdef", "bcdef", "acdefg", "bcdeg"}, []string{"cdg", "cdg", "cdg", "abcdef"}},
	{[]string{"abdefg", "ab", "acdeg", "abcdefg", "abg", "bcdefg", "abef", "abcdfg", "abdeg", "bdefg"}, []string{"acdeg", "abef", "abef", "abef"}},
	{[]string{"abcdfg", "cd", "abcde", "abcdefg", "bcd", "abdef", "abdefg", "abcdef", "abceg", "cdef"}, []string{"abcde", "abdefg", "abcdef", "cd"}},
	{[]string{"abcdeg", "abcdfg", "eg", "bcdefg", "abcdefg", "aceg", "abdef", "abcdg", "deg", "abdeg"}, []string{"abdef", "bcdefg", "eg", "abdeg"}},
	{[]string{"abcefg", "ceg", "acdeg", "abcde", "bdeg", "abcdeg", "abcdefg", "abcdef", "acdfg", "eg"}, []string{"ceg", "abcdefg", "acdfg", "abcdeg"}},
	{[]string{"dg", "abcdfg", "abdefg", "bdeg", "abcdef", "abcdefg", "acefg", "adefg", "abdef", "dfg"}, []string{"dg", "acefg", "adefg", "abcdef"}},
	{[]string{"fg", "cdefg", "befg", "bcdefg", "cfg", "abcdefg", "abcdfg", "abcdef", "bcdef", "acdeg"}, []string{"befg", "abcdefg", "cfg", "bcdefg"}},
	{[]string{"abcfg", "abdefg", "abcdefg", "af", "abf", "abcdeg", "acef", "bcdfg", "abcefg", "abceg"}, []string{"af", "acef", "abcefg", "abcfg"}},
	{[]string{"bf", "bcef", "abcdefg", "acdefg", "abdefg", "bcdfg", "cdefg", "bcdefg", "bdf", "abcdg"}, []string{"bcef", "bf", "bcef", "bf"}},
	{[]string{"bcefg", "bcdefg", "abcdfg", "bdefg", "abcdefg", "bdf", "abdeg", "df", "abcefg", "cdef"}, []string{"df", "abcdefg", "abcefg", "df"}},
	{[]string{"acdef", "abcdf", "bcdfg", "abcdeg", "abcdef", "acdefg", "abef", "abd", "ab", "abcdefg"}, []string{"ab", "acdef", "abef", "abcdefg"}},
	{[]string{"bcdeg", "abcdef", "abceg", "acdefg", "ab", "abe", "acefg", "abcefg", "abcdefg", "abfg"}, []string{"abceg", "abcefg", "acdefg", "abcdefg"}},
	{[]string{"adefg", "be", "abef", "abdefg", "beg", "bdefg", "abcdeg", "bcdfg", "abcdefg", "acdefg"}, []string{"abcdefg", "adefg", "beg", "abdefg"}},
	{[]string{"abcdg", "abcefg", "acg", "abcdefg", "abdfg", "cg", "cdfg", "abcdfg", "abcde", "abdefg"}, []string{"abcdfg", "abcde", "abdfg", "abcdefg"}},
	{[]string{"abcde", "bdefg", "ag", "abcdeg", "acdg", "abcdefg", "abdeg", "aeg", "abcefg", "abcdef"}, []string{"abcde", "abcdefg", "abcde", "abcdefg"}},
	{[]string{"abcdefg", "abcdeg", "abcef", "abcfg", "bcdef", "abe", "abcefg", "aefg", "ae", "abcdfg"}, []string{"aefg", "aefg", "abcdefg", "aefg"}},
	{[]string{"abceg", "abdeg", "bcdefg", "abdefg", "abcfg", "abcdeg", "acde", "ceg", "abcdefg", "ce"}, []string{"ceg", "acde", "abcdefg", "acde"}},
	{[]string{"adeg", "bcdfg", "abcefg", "eg", "acdefg", "acdef", "abcdef", "cdefg", "abcdefg", "ceg"}, []string{"cdefg", "abcdef", "abcdef", "abcdefg"}},
	{[]string{"bcdeg", "abde", "abcdeg", "abcdefg", "abcefg", "abceg", "de", "bcdfg", "cde", "acdefg"}, []string{"cde", "cde", "abcdeg", "abcefg"}},
	{[]string{"abcdefg", "abcfg", "bcdefg", "abdefg", "abcef", "ag", "acdg", "abcdfg", "abg", "bcdfg"}, []string{"bcdfg", "ag", "abdefg", "abdefg"}},
	{[]string{"acfg", "abcdfg", "abg", "bcdeg", "abcdf", "abdefg", "abcdefg", "abcdef", "ag", "abcdg"}, []string{"abcdefg", "abcdef", "abg", "abcdefg"}},
	{[]string{"acdefg", "cdef", "abcdeg", "abcdfg", "ace", "acefg", "abcdefg", "abefg", "ce", "acdfg"}, []string{"abcdfg", "acefg", "abcdefg", "abefg"}},
	{[]string{"abdefg", "abcefg", "abcdfg", "abcde", "abcdefg", "cdfg", "df", "adf", "abcdf", "abcfg"}, []string{"adf", "abcfg", "abcdfg", "df"}},
	{[]string{"abcdeg", "abcdg", "abeg", "ag", "bcdfg", "abcde", "acdefg", "abcdef", "abcdefg", "adg"}, []string{"ag", "abcdg", "ag", "abcdg"}},
	{[]string{"abcdefg", "acdeg", "bcdf", "abdefg", "bd", "bde", "abcde", "abcef", "abcefg", "abcdef"}, []string{"abdefg", "abcef", "bcdf", "bcdf"}},
	{[]string{"deg", "abcdef", "abceg", "bcdef", "abdefg", "bcdeg", "dg", "bcdefg", "cdfg", "abcdefg"}, []string{"cdfg", "cdfg", "cdfg", "deg"}},
	{[]string{"abcdf", "bcde", "abcdefg", "abcdef", "abcefg", "abdfg", "acdef", "acdefg", "bc", "abc"}, []string{"acdef", "abcdef", "abcefg", "acdef"}},
	{[]string{"abdefg", "abcg", "bcefg", "acdefg", "bg", "bfg", "bcdef", "acefg", "abcefg", "abcdefg"}, []string{"bg", "abdefg", "abcg", "acefg"}},
	{[]string{"abcdefg", "adeg", "acdefg", "acefg", "dg", "abcdfg", "cdefg", "abcefg", "cdg", "bcdef"}, []string{"dg", "dg", "acdefg", "cdg"}},
	{[]string{"cdefg", "abdeg", "abcdefg", "adefg", "acdefg", "adf", "abcdfg", "bcdefg", "acef", "af"}, []string{"abcdefg", "abcdefg", "adefg", "acdefg"}},
	{[]string{"abcdefg", "bcdefg", "acdefg", "abcdeg", "abde", "bcdeg", "adg", "abcfg", "abcdg", "ad"}, []string{"abcdg", "abcdefg", "abcdefg", "abde"}},
	{[]string{"abd", "bcdefg", "abcdeg", "ad", "abcdfg", "acdf", "abefg", "abcdefg", "abdfg", "bcdfg"}, []string{"abefg", "ad", "bcdefg", "abcdefg"}},
	{[]string{"ae", "acefg", "abcefg", "abcdefg", "abef", "bcefg", "abcdeg", "aeg", "bcdefg", "acdfg"}, []string{"abef", "acdfg", "ae", "abcdefg"}},
	{[]string{"abce", "acdeg", "be", "bde", "abdfg", "bcdefg", "abcdeg", "acdefg", "abdeg", "abcdefg"}, []string{"acdefg", "abcdefg", "abcdefg", "abdfg"}},
	{[]string{"de", "bdeg", "abcdfg", "bcdfg", "abcdefg", "bcdefg", "cde", "acdefg", "abcef", "bcdef"}, []string{"bcdefg", "abcdefg", "bcdefg", "acdefg"}},
	{[]string{"cdeg", "abcdefg", "abcdfg", "acdefg", "acdef", "acefg", "acd", "abcefg", "cd", "abdef"}, []string{"abcdefg", "acdefg", "cd", "abcdefg"}},
	{[]string{"abcdfg", "abcf", "abcdg", "cf", "cfg", "bdefg", "abcdeg", "acdefg", "abcdefg", "bcdfg"}, []string{"abcdeg", "cf", "abcdeg", "cf"}},
	{[]string{"be", "cdefg", "abcdfg", "abcdef", "abcdeg", "bcdeg", "bce", "abeg", "abcdg", "abcdefg"}, []string{"abcdeg", "abcdefg", "be", "abcdg"}},
	{[]string{"abcdef", "bdefg", "abcdefg", "ceg", "abcdeg", "abcde", "bcdeg", "abcefg", "cg", "acdg"}, []string{"abcdef", "abcefg", "ceg", "abcdefg"}},
	{[]string{"abcdfg", "acefg", "abcde", "bdef", "abcdefg", "bcf", "abcef", "abcdeg", "bf", "abcdef"}, []string{"bf", "bcf", "bdef", "abcdeg"}},
	{[]string{"ce", "cdeg", "abcdefg", "abcdf", "abcdeg", "ace", "abcde", "abcefg", "abdefg", "abdeg"}, []string{"abdefg", "abcde", "abcdefg", "ace"}},
	{[]string{"abfg", "beg", "cdefg", "abdef", "abdefg", "abcdeg", "bdefg", "abcdefg", "bg", "abcdef"}, []string{"abfg", "abfg", "bg", "beg"}},
	{[]string{"bc", "abcf", "bcdef", "cdefg", "abcdef", "abcdeg", "bce", "abcdefg", "abdef", "abdefg"}, []string{"abcdef", "abcf", "bce", "bcdef"}},
	{[]string{"bdef", "abefg", "abcdfg", "abceg", "abf", "bf", "adefg", "abcdefg", "acdefg", "abdefg"}, []string{"abcdefg", "bf", "bf", "abf"}},
	{[]string{"abcdefg", "abcefg", "abdefg", "abcdfg", "cefg", "acf", "abefg", "abcde", "abcef", "cf"}, []string{"cefg", "abcdefg", "abcdfg", "abdefg"}},
	{[]string{"abdfg", "adefg", "abcdefg", "efg", "acdef", "abcefg", "eg", "abcdef", "acdefg", "cdeg"}, []string{"eg", "abcdefg", "cdeg", "efg"}},
	{[]string{"abdefg", "abcefg", "ceg", "abefg", "bcef", "acefg", "acdfg", "abcdefg", "ce", "abcdeg"}, []string{"abefg", "abcdefg", "abcdeg", "ceg"}},
	{[]string{"abfg", "abdeg", "bcdef", "abdefg", "aef", "abdef", "abcdefg", "acdefg", "af", "abcdeg"}, []string{"acdefg", "aef", "af", "aef"}},
	{[]string{"abcef", "cf", "abcdfg", "abdef", "bcf", "abceg", "abdefg", "abcdefg", "cdef", "abcdef"}, []string{"bcf", "abcdfg", "bcf", "abcef"}},
	{[]string{"abcdef", "abcefg", "bcdefg", "abdfg", "aceg", "cg", "abcef", "abcfg", "bcg", "abcdefg"}, []string{"bcdefg", "bcg", "cg", "abcdefg"}},
	{[]string{"acefg", "abcd", "abcdef", "bcdef", "abdefg", "ade", "acdef", "bcdefg", "abcdefg", "ad"}, []string{"bcdefg", "abcdefg", "bcdef", "abcdef"}},
	{[]string{"abcdfg", "abcef", "abcfg", "abcefg", "befg", "ef", "abcde", "cef", "acdefg", "abcdefg"}, []string{"cef", "abcfg", "ef", "abcdefg"}},
	{[]string{"abcdefg", "abeg", "bcdef", "acdefg", "abcdg", "ae", "abcdfg", "abcdeg", "abcde", "ace"}, []string{"ae", "abcdeg", "abcdefg", "ae"}},
	{[]string{"abcdefg", "bcdeg", "acd", "abdfg", "abcdg", "abcdef", "bcdefg", "ac", "abcdeg", "aceg"}, []string{"aceg", "aceg", "acd", "ac"}},
	{[]string{"acdefg", "abcdfg", "abdefg", "abcef", "abg", "abcdefg", "bdeg", "bg", "adefg", "abefg"}, []string{"bg", "abefg", "abdefg", "bdeg"}},
	{[]string{"abdeg", "acdfg", "abcdefg", "ce", "bcdefg", "abcdeg", "cde", "acdeg", "abce", "abdefg"}, []string{"ce", "ce", "abce", "abdeg"}},
	{[]string{"bcef", "abdeg", "bc", "abcdefg", "bcg", "abcdfg", "acefg", "abceg", "abcefg", "acdefg"}, []string{"bcef", "abcdfg", "bcg", "bcef"}},
	{[]string{"abcdef", "bcfg", "abefg", "acdefg", "fg", "abdeg", "abcdefg", "abcef", "abcefg", "afg"}, []string{"bcfg", "abcef", "afg", "abdeg"}},
	{[]string{"abdef", "aefg", "adg", "bcdfg", "abdfg", "abdefg", "ag", "abcdefg", "abcdeg", "abcdef"}, []string{"abcdefg", "aefg", "aefg", "aefg"}},
	{[]string{"abcdef", "cdg", "bcdef", "abcdfg", "abcdefg", "bdeg", "acefg", "bcdefg", "dg", "cdefg"}, []string{"abcdefg", "cdg", "bdeg", "cdg"}},
	{[]string{"abcefg", "bcg", "abdefg", "cdfg", "abcde", "bcdeg", "cg", "bcdefg", "abcdefg", "bdefg"}, []string{"abdefg", "cg", "cg", "abcdefg"}},
	{[]string{"cdefg", "be", "acdefg", "abcdf", "abcdefg", "abcefg", "bcdefg", "bcdef", "bef", "bdeg"}, []string{"bef", "cdefg", "bdeg", "bef"}},
	{[]string{"befg", "abcdfg", "abdeg", "acdef", "abdefg", "dfg", "adefg", "abcdeg", "abcdefg", "fg"}, []string{"dfg", "abcdefg", "dfg", "abdefg"}},
	{[]string{"abcefg", "abcde", "abcdfg", "abf", "acefg", "abcdefg", "abcef", "befg", "bf", "acdefg"}, []string{"befg", "abf", "abcde", "befg"}},
	{[]string{"acdefg", "abcdefg", "adefg", "cg", "abcefg", "abdefg", "acg", "acdeg", "abcde", "cdfg"}, []string{"cg", "acg", "abcefg", "adefg"}},
	{[]string{"abdfg", "bdefg", "bde", "be", "cdefg", "abcdefg", "abcdef", "bceg", "bcdefg", "acdefg"}, []string{"be", "be", "abcdefg", "abcdefg"}},
	{[]string{"abcdfg", "acef", "ce", "abcde", "abcdef", "bce", "abdeg", "bcdefg", "abcdf", "abcdefg"}, []string{"ce", "bce", "acef", "bce"}},
	{[]string{"acdefg", "acdeg", "abcefg", "cdefg", "ac", "abdeg", "acg", "abcdefg", "bcdefg", "acdf"}, []string{"acdf", "acdefg", "cdefg", "acg"}},
	{[]string{"acefg", "bcdf", "abdefg", "abcdfg", "bf", "bfg", "abcdeg", "abcdefg", "abcdg", "abcfg"}, []string{"bfg", "abcfg", "bf", "acefg"}},
	{[]string{"abdefg", "acef", "abcdefg", "adefg", "acdefg", "ceg", "cdefg", "bcdfg", "ce", "abcdeg"}, []string{"acef", "acef", "acdefg", "acef"}},
	{[]string{"acdefg", "abcdefg", "acde", "abdefg", "ac", "adefg", "acefg", "acg", "bcefg", "abcdfg"}, []string{"adefg", "abdefg", "ac", "acde"}},
	{[]string{"acdeg", "acf", "cefg", "abcdeg", "abcdef", "acdefg", "abdfg", "acdfg", "cf", "abcdefg"}, []string{"acdefg", "cefg", "acdfg", "acf"}},
	{[]string{"abcdefg", "ceg", "bcefg", "abcdeg", "defg", "eg", "bcdef", "abcfg", "abcdef", "bcdefg"}, []string{"ceg", "abcdefg", "defg", "eg"}},
	{[]string{"abcefg", "abcdeg", "abdefg", "bdefg", "abefg", "abcdefg", "bcdef", "dg", "adfg", "deg"}, []string{"abcdefg", "abcefg", "bcdef", "abdefg"}},
	{[]string{"acdefg", "cg", "abcdefg", "bcdfg", "abcdef", "abcdfg", "cdg", "bdefg", "abcdf", "abcg"}, []string{"cg", "acdefg", "acdefg", "abcdefg"}},
	{[]string{"bceg", "abefg", "acefg", "acdef", "acg", "cg", "abcefg", "abdefg", "abcdfg", "abcdefg"}, []string{"abefg", "abcdfg", "abefg", "abdefg"}},
	{[]string{"abcef", "abcdef", "eg", "beg", "abcefg", "aceg", "abcdefg", "bcdfg", "bcefg", "abdefg"}, []string{"abdefg", "aceg", "eg", "bcefg"}},
	{[]string{"abdfg", "bcdefg", "bdefg", "abcefg", "de", "cdef", "deg", "abcdeg", "abcdefg", "bcefg"}, []string{"cdef", "abcdeg", "de", "de"}},
	{[]string{"abcdef", "bcdefg", "bcdef", "ab", "abe", "abcf", "acdeg", "abdefg", "abcde", "abcdefg"}, []string{"abcf", "abcdefg", "abcdef", "abcdefg"}},
	{[]string{"abdef", "be", "abcdefg", "acdefg", "abdefg", "bcdefg", "abeg", "abcdf", "bef", "adefg"}, []string{"abcdefg", "bcdefg", "be", "abeg"}},
	{[]string{"abcdef", "abdefg", "abcdefg", "acdfg", "cefg", "acdefg", "abcdg", "acdef", "dfg", "fg"}, []string{"abcdg", "cefg", "cefg", "acdefg"}},
	{[]string{"bdefg", "bcdeg", "adef", "abcdefg", "abefg", "abcdfg", "dfg", "df", "abcefg", "abdefg"}, []string{"dfg", "abcdefg", "abcdefg", "df"}},
	{[]string{"abcdfg", "abdefg", "ac", "abcdf", "bcdef", "abcdefg", "abdfg", "acfg", "acd", "abcdeg"}, []string{"abcdeg", "abcdefg", "acd", "abdfg"}},
	{[]string{"cdeg", "abcdefg", "abcdef", "acdfg", "acefg", "abdfg", "abcefg", "cdf", "cd", "acdefg"}, []string{"abcdefg", "cdf", "acdfg", "abcdefg"}},
	{[]string{"ade", "acdefg", "abcef", "abcefg", "bcdeg", "abcde", "ad", "abcdef", "abdf", "abcdefg"}, []string{"abcdefg", "abcefg", "abdf", "ad"}},
	{[]string{"bcefg", "abcdef", "acfg", "abcefg", "abcdeg", "abcdefg", "cg", "ceg", "abcef", "bdefg"}, []string{"abcdefg", "abcdeg", "cg", "abcdeg"}},
	{[]string{"abcefg", "acdefg", "abef", "ab", "abcdefg", "acefg", "bcdeg", "abc", "abcdfg", "abceg"}, []string{"abef", "abcdfg", "abceg", "abceg"}},
}
