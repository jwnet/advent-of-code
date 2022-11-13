package main

// following this video: https://www.youtube.com/watch?v=hmq6veCFo0Y

import "fmt"

func main() {
	prog := compute()
	opt(prog)
	force(prog)
	opt(prog)
	dump(prog)
	fmt.Println(maxModel(prog))
	fmt.Println(minModel(prog))
}

func maxModel(prog []*val) string {
	m := make([]byte, 14)
	for i := range m {
		m[i] = '?'
	}
	for _, v := range prog {
		var a, b, i, j int
		switch {
		case bin(bin(bin(inp(&i), "+", num(&a)), "+", num(&b)), "force", inp(&j))(v),
			bin(bin(inp(&i), "+", num(&b)), "force", inp(&j))(v):
			a += b
			fmt.Println(i, a, j)
			if a > 0 {
				m[i] = byte('9' - a)
				m[j] = '9'
			} else {
				m[i] = '9'
				m[j] = byte('9' + a)
			}
		}
	}
	return string(m)
}

func minModel(prog []*val) string {
	m := make([]byte, 14)
	for i := range m {
		m[i] = '?'
	}
	for _, v := range prog {
		var a, b, i, j int
		switch {
		case bin(bin(bin(inp(&i), "+", num(&a)), "+", num(&b)), "force", inp(&j))(v),
			bin(bin(inp(&i), "+", num(&b)), "force", inp(&j))(v):
			a += b
			fmt.Println(i, a, j)
			if a > 0 {
				m[i] = '1'
				m[j] = byte('1' + a)
			} else {
				m[j] = '1'
				m[i] = byte('1' - a)
			}
		}
	}
	return string(m)
}


func opt(prog[]*val) {
	for _, v := range prog {
		v.fwd = v

		if v.l != nil && v.l.fwd != nil {
			v.l = v.l.fwd
		}
		if v.r != nil && v.r.fwd != nil {
			v.r = v.r.fwd
		}

		var a, b int
		var x, y *val
		switch {
		case bin(num(&a), "+", num(&b))(v):
			setnum(v, a+b)
		case bin(num(&a), "*", num(&b))(v):
			setnum(v, a*b)
		case bin(num(&a), "/", num(&b))(v):
			setnum(v, a/b)
		case bin(num(&a), "%", num(&b))(v):
			setnum(v, a%b)
		case bin(any(&x), "*", con(0))(v),
			bin(con(0), "*", any(&x))(v):
			setnum(v, 0)
		case bin(any(&x), "+", con(0))(v),
				bin(con(0), "+", any(&x))(v),
				bin(any(&x), "/", con(1))(v),
				bin(any(&x), "*", con(1))(v),
				bin(con(1), "*", any(&x))(v):
			v.fwd = x
		case bin(any(&x), "==", any(&y))(v) && (x.max < y.min || y.max < x.min):
			setnum(v, 0)
		case bin(num(&a), "==", num(&b))(v) && a == b:
			setnum(v, 1)
		case bin(bin(bin(any(&y), "*", con(26)), "+", any(&x)), "%", con(26))(v) && x.max < 26:
			v.fwd = x
		case bin(bin(bin(any(&y), "*", con(26)), "+", any(&x)), "/", con(26))(v) && x.max < 26:
			v.fwd = y
		case bin(any(&x), "%", con(26))(v) && x.max < 26:
			v.fwd = x
		}

		switch v.op {
		case "num":
			v.min = v.n
			v.max = v.n	
		case "inp":
			v.min = 1
			v.max = 9
		case "+":
			v.min = v.l.min + v.r.min
			v.max = v.l.max + v.r.max
		case "*":
			v.min = v.l.min * v.r.min
			v.max = v.l.max * v.r.max
		case "/":
			v.min = v.l.min / v.r.n
			v.max = v.l.max / v.r.n
		case "%":
			v.min = 0
			v.max = v.r.n - 1
		case "==":
			v.min = 0
			v.max = 1
		case "force":
			v.min = 0
			v.max = 0
		}

	}
}

func force(prog []*val) {
	max := make(map[*val]int)
	max[prog[len(prog)-1]] = 0

	updateMax := func(v *val, m int) {
		if old, ok := max[v]; ok && old < m {
			return
		}
		max[v] = m
	}

	for i := len(prog)-1; i >= 0; i-- {
		v := prog[i]
		m, ok := max[v]
		if !ok {
			continue
		}

		if v.max < m {
			continue
		}

		var x, y *val
		var a  int
		switch {
		default:
			//panic("force "+v.op)
		case num(&a)(v):
			if a > m {
				panic("force impossible")
			}
		case bin(any(&x), "+", any(&y))(v):
			updateMax(x, m-y.min)
			updateMax(y, m-x.min)
		case bin(any(&x), "*", any(&y))(v):
			if y.min > 0 {
				updateMax(x, m/y.min)
			}
			if x.min > 0 {
				updateMax(y, m/x.min)
			}
		case bin(any(&x), "/", num(&a))(v):
			updateMax(x, m*a+a-1)
		case bin(bin(any(&x), "==", any(&y)), "==", con(0))(v):
			v.op = "force"
			v.l = x
			v.r = y
			v.min = 0
			v.max = 0
		}
	}
}

func any(p **val) matcher {
	return func(v *val) bool {
		*p = v
		return true
	}
}

func setnum(v *val, n int) {
	*v = val{op: "num", n: n}
}

type matcher func(v *val) bool

func con(n int) matcher {
	return func(v *val) bool {
		return v.op == "num" && v.n == n ||
			v.min == n && v.max == n
	}
}

func num(n *int) matcher {
	return func(v *val) bool {
		if v.op == "num" {
			*n = v.n
			return true
		}
		if v.min == v.max {
			*n = v.min
			return true
		}
		return false
	}
}

func inp(n *int) matcher {
	return func(v *val) bool {
		if v.op == "inp" {
			*n = v.n
			return true
		}
		return false
	}
}

func bin(l matcher, op string, r matcher) matcher {
	return func(v *val) bool {
		return v.op == op && l(v.l) && r(v.r)
	}
}

func dump(prog []*val) {
	count := make(map[*val]int)
	for i := len(prog) -1; i >= 0; i -- {
		v := prog[i]
		if count[v] == 0 && i != len(prog)-1 {
			continue
		}
		count[v.l]++
		count[v.r]++
	}
	str := make(map[*val]string)
	for _, v := range prog {
		var x string
		switch v.op {
		case "inp", "num":
			x = v.Init()
		default:
			x = fmt.Sprintf("(%v %v %v)", str[v.l], v.op, str[v.r])
			if v.op == "force" {
				x = fmt.Sprintf("force %v+%d == %v", str[v.l], v.n,str[v.r])
			}
			if count[v] >= 2 || v.op =="force" {
				fmt.Println(v.Name(), "=", x)
				x = v.Name()
			}
		}
		str[v] = x
	}
	fmt.Println(str[prog[len(prog)-1]])
}

type val struct {
	t int
	op string
	n int
	l, r *val
	min, max int
	fwd *val
}

func (v *val) Name() string {
	return fmt.Sprint("t", v.t)
}

func (v *val) Init() string {
	switch v.op {
	case "num":
		return fmt.Sprint(v.n)
	case "inp":
		return fmt.Sprint("m", v.n)
	case "force":
		return fmt.Sprintf("force %v + %d == %v", v.l.Name(), v.n, v.r.Name())
	default:
		return fmt.Sprintf("(%v %v %v)", v.l.Name(), v.op, v.r.Name())
	}
}

func (v *val) String() string {
	return fmt.Sprintf("%v = %v", v.Name(), v.Init())
}

func compute() []*val {
	var prog []*val

	t := 0
	emit := func(v *val) *val {
		t++
		v.t = t
		prog = append(prog, v)
		return v
	}

	i := 0
	inp := func() *val {
		v := emit(&val{op: "inp", n: i})
		i++
		return v
	}

	bin := func(l *val, op string, r *val) *val {
		return emit(&val{l: l, op: op, r: r})
	}

	add := func(l, r *val) *val { return bin(l, "+", r) }
	mul := func(l, r *val) *val { return bin(l, "*", r) }
	mod := func(l, r *val) *val { return bin(l, "%", r) }
	div := func(l, r *val) *val { return bin(l, "/", r) }
	eql := func(l, r *val) *val { return bin(l, "==", r) }
	num := func(n int) *val {
		return emit(&val{op:"num", n: n})
	}

	var w, x, y, z = num(0), num(0), num(0), num(0)

	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(13))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(10))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(11))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(16))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(11))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(0))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(10))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(13))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-14))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(7))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-4))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(11))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(11))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(11))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-3))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(10))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(16))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(8))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(13))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(15))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(2))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-15))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(5))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(10))
	y = mul(y, x)
	z = add(z, y)

	return prog
}
