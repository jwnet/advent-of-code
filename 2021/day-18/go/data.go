package main

var sample = []snailnum{
	{lp: &snailnum{lp: &snailnum{ll: 0, rp: &snailnum{ll: 5, rl: 8}}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 7}, rp: &snailnum{ll: 9, rl: 6}}}, rp: &snailnum{lp: &snailnum{ll: 4, rp: &snailnum{ll: 1, rl: 2}}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 4}, rl: 2}}},
	{lp: &snailnum{lp: &snailnum{ll: 5, rp: &snailnum{ll: 2, rl: 8}}, rl: 4}, rp: &snailnum{ll: 5, rp: &snailnum{lp: &snailnum{ll: 9, rl: 9}, rl: 0}}},
	{ll: 6, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 2}, rp: &snailnum{ll: 5, rl: 6}}, rp: &snailnum{lp: &snailnum{ll: 7, rl: 6}, rp: &snailnum{ll: 4, rl: 7}}}},
	{lp: &snailnum{lp: &snailnum{ll: 6, rp: &snailnum{ll: 0, rl: 7}}, rp: &snailnum{ll: 0, rl: 9}}, rp: &snailnum{ll: 4, rp: &snailnum{ll: 9, rp: &snailnum{ll: 9, rl: 0}}}},
	{lp: &snailnum{lp: &snailnum{ll: 7, rp: &snailnum{ll: 6, rl: 4}}, rp: &snailnum{ll: 3, rp: &snailnum{ll: 1, rl: 3}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 5}, rl: 1}, rl: 9}},
	{lp: &snailnum{ll: 6, rp: &snailnum{lp: &snailnum{ll: 7, rl: 3}, rp: &snailnum{ll: 3, rl: 2}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 8}, rp: &snailnum{ll: 5, rl: 7}}, rl: 4}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 4}, rp: &snailnum{ll: 7, rl: 7}}, rl: 8}, rp: &snailnum{lp: &snailnum{ll: 8, rl: 3}, rl: 8}},
	{lp: &snailnum{ll: 9, rl: 3}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 9}, rp: &snailnum{ll: 6, rp: &snailnum{ll: 4, rl: 9}}}},
	{lp: &snailnum{ll: 2, rp: &snailnum{lp: &snailnum{ll: 7, rl: 7}, rl: 7}}, rp: &snailnum{lp: &snailnum{ll: 5, rl: 8}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 3}, rp: &snailnum{ll: 0, rl: 2}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 2}, rl: 5}, rp: &snailnum{ll: 8, rp: &snailnum{ll: 3, rl: 7}}}, rp: &snailnum{lp: &snailnum{ll: 5, rp: &snailnum{ll: 7, rl: 5}}, rp: &snailnum{ll: 4, rl: 4}}},
}

var input = []snailnum{
	{lp: &snailnum{lp: &snailnum{ll: 7, rl: 1}, rl: 2}, rl: 3},
	{lp: &snailnum{ll: 1, rl: 7}, rl: 7},
	{lp: &snailnum{ll: 6, rl: 8}, rp: &snailnum{lp: &snailnum{ll: 6, rp: &snailnum{ll: 3, rl: 6}}, rp: &snailnum{ll: 0, rl: 5}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 1}, rl: 8}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 4}, rl: 8}}, rp: &snailnum{lp: &snailnum{ll: 6, rl: 5}, rl: 4}},
	{lp: &snailnum{ll: 1, rp: &snailnum{lp: &snailnum{ll: 3, rl: 8}, rp: &snailnum{ll: 9, rl: 1}}}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 1}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 7}, rl: 0}}},
	{lp: &snailnum{lp: &snailnum{ll: 7, rl: 4}, rp: &snailnum{ll: 8, rp: &snailnum{ll: 7, rl: 6}}}, rp: &snailnum{ll: 9, rp: &snailnum{lp: &snailnum{ll: 6, rl: 3}, rp: &snailnum{ll: 7, rl: 8}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 0}, rl: 1}, rl: 4}, rp: &snailnum{lp: &snailnum{ll: 5, rp: &snailnum{ll: 6, rl: 9}}, rp: &snailnum{lp: &snailnum{ll: 4, rl: 3}, rl: 2}}},
	{lp: &snailnum{lp: &snailnum{ll: 3, rl: 8}, rl: 8}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 2}, rl: 8}, rp: &snailnum{ll: 9, rp: &snailnum{ll: 0, rl: 5}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 8}, rp: &snailnum{ll: 3, rl: 9}}, rp: &snailnum{ll: 7, rp: &snailnum{ll: 1, rl: 4}}}, rp: &snailnum{ll: 6, rl: 1}},
	{ll: 3, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 3}, rl: 9}, rp: &snailnum{ll: 0, rl: 7}}},
	{lp: &snailnum{lp: &snailnum{ll: 6, rl: 9}, rl: 1}, rp: &snailnum{lp: &snailnum{ll: 0, rp: &snailnum{ll: 8, rl: 4}}, rp: &snailnum{lp: &snailnum{ll: 2, rl: 2}, rl: 9}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 2}, rl: 3}, rp: &snailnum{ll: 0, rl: 4}}, rl: 3},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 8}, rl: 7}, rp: &snailnum{lp: &snailnum{ll: 7, rl: 4}, rl: 0}}, rp: &snailnum{ll: 2, rp: &snailnum{ll: 5, rp: &snailnum{ll: 2, rl: 8}}}},
	{lp: &snailnum{ll: 4, rp: &snailnum{ll: 9, rp: &snailnum{ll: 8, rl: 0}}}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 5}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 3}, rl: 8}}},
	{lp: &snailnum{lp: &snailnum{ll: 8, rl: 5}, rp: &snailnum{ll: 3, rp: &snailnum{ll: 1, rl: 4}}}, rp: &snailnum{lp: &snailnum{ll: 6, rp: &snailnum{ll: 8, rl: 0}}, rp: &snailnum{lp: &snailnum{ll: 2, rl: 7}, rp: &snailnum{ll: 2, rl: 6}}}},
	{ll: 4, rl: 7},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 3}, rl: 0}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 9}, rp: &snailnum{ll: 4, rl: 1}}}, rp: &snailnum{lp: &snailnum{ll: 1, rp: &snailnum{ll: 4, rl: 2}}, rl: 3}},
	{lp: &snailnum{lp: &snailnum{ll: 8, rp: &snailnum{ll: 5, rl: 3}}, rp: &snailnum{lp: &snailnum{ll: 5, rl: 7}, rl: 7}}, rp: &snailnum{lp: &snailnum{ll: 5, rl: 6}, rp: &snailnum{ll: 6, rl: 4}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 4}, rl: 1}, rp: &snailnum{ll: 8, rl: 6}}, rp: &snailnum{lp: &snailnum{ll: 6, rl: 5}, rp: &snailnum{ll: 0, rp: &snailnum{ll: 9, rl: 1}}}},
	{lp: &snailnum{lp: &snailnum{ll: 1, rp: &snailnum{ll: 5, rl: 7}}, rl: 8}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 9, rl: 1}, rl: 9}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 2}, rl: 4}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 5}, rp: &snailnum{ll: 4, rl: 0}}, rp: &snailnum{ll: 4, rp: &snailnum{ll: 9, rl: 6}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 1}, rl: 1}, rl: 7}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 1, rl: 9}, rp: &snailnum{ll: 9, rl: 5}}, rp: &snailnum{lp: &snailnum{ll: 5, rl: 0}, rp: &snailnum{ll: 3, rl: 1}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 7}, rp: &snailnum{ll: 8, rl: 8}}, rp: &snailnum{lp: &snailnum{ll: 7, rl: 3}, rl: 0}}},
	{lp: &snailnum{ll: 6, rp: &snailnum{lp: &snailnum{ll: 6, rl: 7}, rp: &snailnum{ll: 9, rl: 0}}}, rp: &snailnum{lp: &snailnum{ll: 7, rl: 7}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 3}, rl: 0}}},
	{lp: &snailnum{ll: 0, rl: 6}, rp: &snailnum{ll: 5, rl: 2}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 8}, rl: 3}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 0}, rl: 8}}, rp: &snailnum{ll: 7, rl: 4}},
	{lp: &snailnum{ll: 0, rp: &snailnum{lp: &snailnum{ll: 9, rl: 9}, rp: &snailnum{ll: 9, rl: 4}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 1, rl: 1}, rl: 2}, rp: &snailnum{ll: 1, rp: &snailnum{ll: 6, rl: 7}}}},
	{ll: 0, rp: &snailnum{lp: &snailnum{ll: 5, rl: 7}, rl: 2}},
	{lp: &snailnum{ll: 2, rp: &snailnum{lp: &snailnum{ll: 5, rl: 4}, rl: 6}}, rp: &snailnum{ll: 1, rp: &snailnum{ll: 8, rp: &snailnum{ll: 7, rl: 6}}}},
	{lp: &snailnum{lp: &snailnum{ll: 1, rl: 7}, rp: &snailnum{ll: 8, rp: &snailnum{ll: 5, rl: 8}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 1}, rp: &snailnum{ll: 9, rl: 1}}, rp: &snailnum{lp: &snailnum{ll: 5, rl: 6}, rl: 9}}},
	{lp: &snailnum{ll: 1, rl: 8}, rp: &snailnum{ll: 9, rp: &snailnum{ll: 4, rl: 3}}},
	{ll: 5, rp: &snailnum{ll: 2, rp: &snailnum{lp: &snailnum{ll: 5, rl: 5}, rl: 9}}},
	{ll: 3, rp: &snailnum{ll: 8, rp: &snailnum{lp: &snailnum{ll: 2, rl: 8}, rp: &snailnum{ll: 4, rl: 8}}}},
	{lp: &snailnum{lp: &snailnum{ll: 4, rl: 9}, rp: &snailnum{lp: &snailnum{ll: 5, rl: 5}, rl: 0}}, rp: &snailnum{ll: 9, rp: &snailnum{ll: 8, rp: &snailnum{ll: 3, rl: 0}}}},
	{lp: &snailnum{lp: &snailnum{ll: 2, rp: &snailnum{ll: 6, rl: 4}}, rp: &snailnum{ll: 8, rp: &snailnum{ll: 9, rl: 9}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 0, rl: 4}, rl: 8}, rp: &snailnum{ll: 3, rp: &snailnum{ll: 9, rl: 7}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 8, rl: 1}, rp: &snailnum{ll: 2, rl: 4}}, rl: 3}, rp: &snailnum{ll: 1, rp: &snailnum{lp: &snailnum{ll: 3, rl: 3}, rp: &snailnum{ll: 6, rl: 3}}}},
	{lp: &snailnum{lp: &snailnum{ll: 8, rp: &snailnum{ll: 7, rl: 3}}, rp: &snailnum{ll: 1, rl: 8}}, rl: 2},
	{lp: &snailnum{ll: 8, rp: &snailnum{ll: 8, rl: 4}}, rp: &snailnum{lp: &snailnum{ll: 6, rp: &snailnum{ll: 4, rl: 7}}, rp: &snailnum{ll: 3, rl: 0}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 4, rl: 6}, rp: &snailnum{ll: 8, rl: 3}}, rl: 9}, rp: &snailnum{ll: 9, rp: &snailnum{lp: &snailnum{ll: 8, rl: 9}, rp: &snailnum{ll: 0, rl: 9}}}},
	{lp: &snailnum{ll: 3, rp: &snailnum{lp: &snailnum{ll: 2, rl: 7}, rp: &snailnum{ll: 4, rl: 4}}}, rl: 2},
	{ll: 8, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 8, rl: 6}, rl: 2}, rp: &snailnum{lp: &snailnum{ll: 8, rl: 9}, rl: 6}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 7}, rp: &snailnum{ll: 2, rl: 0}}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 2}, rp: &snailnum{ll: 5, rl: 5}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 8, rl: 5}, rl: 5}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 3}, rp: &snailnum{ll: 2, rl: 3}}}},
	{lp: &snailnum{ll: 1, rl: 6}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 8}, rp: &snailnum{ll: 9, rp: &snailnum{ll: 4, rl: 9}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 1, rl: 4}, rl: 5}, rl: 9}, rp: &snailnum{ll: 4, rp: &snailnum{ll: 6, rl: 8}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 4}, rp: &snailnum{ll: 9, rl: 0}}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 4}, rp: &snailnum{ll: 6, rl: 6}}}, rp: &snailnum{lp: &snailnum{ll: 9, rp: &snailnum{ll: 2, rl: 8}}, rl: 2}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 9}, rl: 2}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 0}, rl: 5}}, rp: &snailnum{ll: 2, rl: 1}},
	{ll: 6, rp: &snailnum{lp: &snailnum{ll: 3, rl: 2}, rp: &snailnum{lp: &snailnum{ll: 3, rl: 0}, rl: 0}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 7, rl: 4}, rl: 1}, rp: &snailnum{lp: &snailnum{ll: 4, rl: 1}, rl: 1}}, rp: &snailnum{lp: &snailnum{ll: 3, rl: 4}, rl: 4}},
	{ll: 3, rp: &snailnum{ll: 9, rp: &snailnum{ll: 9, rl: 7}}},
	{lp: &snailnum{lp: &snailnum{ll: 3, rp: &snailnum{ll: 3, rl: 3}}, rp: &snailnum{ll: 0, rl: 3}}, rp: &snailnum{ll: 1, rp: &snailnum{ll: 1, rl: 8}}},
	{lp: &snailnum{ll: 8, rp: &snailnum{ll: 8, rl: 7}}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 2}, rl: 5}},
	{lp: &snailnum{lp: &snailnum{ll: 1, rp: &snailnum{ll: 3, rl: 9}}, rp: &snailnum{ll: 5, rl: 9}}, rp: &snailnum{ll: 1, rl: 5}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 7, rl: 8}, rp: &snailnum{ll: 9, rl: 7}}, rl: 9}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 9, rl: 2}, rp: &snailnum{ll: 2, rl: 2}}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 6}, rl: 8}}},
	{ll: 4, rp: &snailnum{lp: &snailnum{ll: 3, rl: 5}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 3}, rp: &snailnum{ll: 5, rl: 5}}}},
	{ll: 7, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 0, rl: 1}, rl: 2}, rp: &snailnum{lp: &snailnum{ll: 3, rl: 6}, rl: 5}}},
	{ll: 0, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 4}, rp: &snailnum{ll: 3, rl: 4}}, rp: &snailnum{ll: 8, rl: 9}}},
	{lp: &snailnum{ll: 1, rp: &snailnum{lp: &snailnum{ll: 6, rl: 8}, rl: 1}}, rp: &snailnum{ll: 8, rl: 0}},
	{ll: 1, rl: 1},
	{ll: 7, rl: 0},
	{lp: &snailnum{ll: 1, rl: 2}, rp: &snailnum{lp: &snailnum{ll: 0, rp: &snailnum{ll: 8, rl: 3}}, rp: &snailnum{lp: &snailnum{ll: 4, rl: 5}, rp: &snailnum{ll: 9, rl: 7}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 3}, rp: &snailnum{ll: 5, rl: 9}}, rp: &snailnum{ll: 7, rp: &snailnum{ll: 1, rl: 9}}}, rl: 2},
	{lp: &snailnum{ll: 3, rl: 5}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 7}, rl: 9}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 9}, rp: &snailnum{ll: 4, rl: 8}}, rl: 6}, rl: 0},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 4}, rp: &snailnum{ll: 3, rl: 9}}, rp: &snailnum{ll: 2, rp: &snailnum{ll: 9, rl: 4}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 8, rl: 9}, rp: &snailnum{ll: 3, rl: 1}}, rl: 7}},
	{lp: &snailnum{ll: 5, rp: &snailnum{lp: &snailnum{ll: 0, rl: 2}, rl: 4}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 9, rl: 9}, rp: &snailnum{ll: 7, rl: 4}}, rp: &snailnum{ll: 1, rl: 5}}},
	{ll: 3, rp: &snailnum{ll: 6, rp: &snailnum{lp: &snailnum{ll: 5, rl: 4}, rl: 1}}},
	{lp: &snailnum{lp: &snailnum{ll: 2, rp: &snailnum{ll: 2, rl: 7}}, rl: 2}, rp: &snailnum{lp: &snailnum{ll: 4, rp: &snailnum{ll: 7, rl: 3}}, rl: 5}},
	{ll: 7, rp: &snailnum{lp: &snailnum{ll: 0, rp: &snailnum{ll: 2, rl: 0}}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 4}, rl: 6}}},
	{lp: &snailnum{ll: 4, rp: &snailnum{ll: 3, rp: &snailnum{ll: 6, rl: 2}}}, rl: 9},
	{lp: &snailnum{lp: &snailnum{ll: 0, rp: &snailnum{ll: 5, rl: 6}}, rp: &snailnum{ll: 8, rl: 3}}, rp: &snailnum{lp: &snailnum{ll: 7, rl: 9}, rp: &snailnum{ll: 0, rp: &snailnum{ll: 9, rl: 6}}}},
	{ll: 8, rp: &snailnum{lp: &snailnum{ll: 6, rl: 4}, rp: &snailnum{ll: 4, rl: 8}}},
	{lp: &snailnum{lp: &snailnum{ll: 8, rp: &snailnum{ll: 6, rl: 8}}, rp: &snailnum{ll: 5, rp: &snailnum{ll: 7, rl: 3}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 7, rl: 8}, rl: 5}, rl: 2}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 5}, rp: &snailnum{ll: 4, rl: 7}}, rl: 5}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 0}, rp: &snailnum{ll: 9, rp: &snailnum{ll: 1, rl: 9}}}},
	{lp: &snailnum{ll: 7, rp: &snailnum{lp: &snailnum{ll: 1, rl: 5}, rl: 9}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 4}, rp: &snailnum{ll: 1, rl: 7}}, rp: &snailnum{ll: 1, rp: &snailnum{ll: 7, rl: 9}}}},
	{lp: &snailnum{ll: 0, rp: &snailnum{ll: 3, rp: &snailnum{ll: 4, rl: 1}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 9}, rl: 3}, rp: &snailnum{ll: 4, rp: &snailnum{ll: 0, rl: 8}}}},
	{lp: &snailnum{lp: &snailnum{ll: 8, rp: &snailnum{ll: 1, rl: 6}}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 1}, rl: 7}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 1, rl: 1}, rp: &snailnum{ll: 0, rl: 2}}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 4}, rp: &snailnum{ll: 9, rl: 6}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 7}, rl: 0}, rp: &snailnum{lp: &snailnum{ll: 6, rl: 8}, rl: 9}}, rp: &snailnum{lp: &snailnum{ll: 1, rp: &snailnum{ll: 6, rl: 6}}, rp: &snailnum{lp: &snailnum{ll: 2, rl: 9}, rp: &snailnum{ll: 4, rl: 7}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 0}, rp: &snailnum{ll: 1, rl: 2}}, rp: &snailnum{ll: 1, rp: &snailnum{ll: 5, rl: 1}}}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 4}, rl: 1}},
	{lp: &snailnum{ll: 9, rl: 1}, rl: 6},
	{lp: &snailnum{ll: 7, rl: 2}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 5}, rp: &snailnum{ll: 4, rl: 3}}, rl: 6}},
	{lp: &snailnum{ll: 9, rp: &snailnum{lp: &snailnum{ll: 0, rl: 6}, rl: 9}}, rp: &snailnum{lp: &snailnum{ll: 7, rl: 9}, rp: &snailnum{ll: 7, rl: 1}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 7, rl: 3}, rp: &snailnum{ll: 6, rl: 4}}, rp: &snailnum{lp: &snailnum{ll: 2, rl: 5}, rp: &snailnum{ll: 7, rl: 2}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 4, rl: 4}, rl: 0}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 5}, rp: &snailnum{ll: 8, rl: 5}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 8, rl: 8}, rp: &snailnum{ll: 6, rl: 4}}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 2}, rp: &snailnum{ll: 9, rl: 5}}}, rl: 2},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 0}, rl: 7}, rp: &snailnum{ll: 9, rl: 2}}, rp: &snailnum{lp: &snailnum{ll: 0, rp: &snailnum{ll: 8, rl: 6}}, rp: &snailnum{lp: &snailnum{ll: 7, rl: 2}, rp: &snailnum{ll: 8, rl: 5}}}},
	{lp: &snailnum{ll: 0, rl: 6}, rp: &snailnum{ll: 1, rp: &snailnum{ll: 9, rp: &snailnum{ll: 4, rl: 3}}}},
	{lp: &snailnum{ll: 0, rl: 8}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 0}, rl: 6}, rp: &snailnum{ll: 5, rp: &snailnum{ll: 2, rl: 0}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 7, rl: 1}, rp: &snailnum{ll: 0, rl: 3}}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 9}, rp: &snailnum{ll: 3, rl: 5}}}, rp: &snailnum{ll: 4, rp: &snailnum{ll: 8, rl: 4}}},
	{ll: 7, rp: &snailnum{lp: &snailnum{ll: 1, rp: &snailnum{ll: 3, rl: 7}}, rp: &snailnum{lp: &snailnum{ll: 3, rl: 4}, rp: &snailnum{ll: 2, rl: 3}}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 2, rl: 2}, rp: &snailnum{ll: 4, rl: 8}}, rp: &snailnum{lp: &snailnum{ll: 3, rl: 4}, rl: 0}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 1, rl: 5}, rp: &snailnum{ll: 2, rl: 8}}, rl: 5}},
	{ll: 6, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 9, rl: 1}, rl: 5}, rp: &snailnum{ll: 9, rl: 9}}},
	{lp: &snailnum{lp: &snailnum{ll: 2, rp: &snailnum{ll: 8, rl: 6}}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 9}, rp: &snailnum{ll: 6, rl: 3}}}, rl: 4},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 2}, rp: &snailnum{ll: 9, rl: 3}}, rl: 8}, rl: 9},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 9}, rl: 0}, rp: &snailnum{lp: &snailnum{ll: 0, rl: 6}, rp: &snailnum{ll: 1, rl: 3}}}, rp: &snailnum{lp: &snailnum{ll: 5, rp: &snailnum{ll: 9, rl: 8}}, rp: &snailnum{lp: &snailnum{ll: 1, rl: 5}, rp: &snailnum{ll: 3, rl: 7}}}},
	{lp: &snailnum{ll: 2, rp: &snailnum{ll: 4, rp: &snailnum{ll: 2, rl: 3}}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 6, rl: 0}, rp: &snailnum{ll: 7, rl: 2}}, rl: 3}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 8, rl: 3}, rl: 4}, rp: &snailnum{ll: 6, rp: &snailnum{ll: 8, rl: 8}}}, rl: 4},
	{lp: &snailnum{lp: &snailnum{ll: 9, rl: 8}, rl: 5}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 4, rl: 4}, rp: &snailnum{ll: 6, rl: 3}}, rp: &snailnum{ll: 8, rl: 6}}},
	{ll: 9, rl: 2},
	{lp: &snailnum{lp: &snailnum{ll: 3, rl: 4}, rp: &snailnum{ll: 4, rp: &snailnum{ll: 7, rl: 0}}}, rp: &snailnum{ll: 0, rp: &snailnum{ll: 4, rp: &snailnum{ll: 6, rl: 9}}}},
	{lp: &snailnum{lp: &snailnum{ll: 0, rl: 8}, rp: &snailnum{ll: 3, rl: 9}}, rp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 3, rl: 8}, rl: 6}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 3}, rl: 6}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 5, rl: 6}, rp: &snailnum{ll: 0, rl: 3}}, rl: 1}, rp: &snailnum{ll: 8, rp: &snailnum{ll: 2, rl: 9}}},
	{lp: &snailnum{lp: &snailnum{lp: &snailnum{ll: 4, rl: 2}, rl: 8}, rp: &snailnum{lp: &snailnum{ll: 9, rl: 3}, rl: 7}}, rl: 0},
}
