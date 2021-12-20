package main

func make2d[T any](m, n int) [][]T {
	if m < 0 {
		m = 0
	}
	if n < 0 {
		n = 0
	}
	s := make([][]T, m)
	for i := range s {
		s[i] = make([]T, n)
	}
	return s
}

func fill2d[T any](m, n int, val T) [][]T {
	s := make2d[T](m, n)
	for r, row := range s {
		for c := range row {
			s[r][c] = val
		}
	}
	return s
}

func size2d[T any](s [][]T) (m, n int) {
	if len(s) == 0 {
		return
	}
	m = len(s)
	n = len(s[0])
	return
}

//---------------------------------------------------------

type point struct {
	r, c int
}

type riskPoint struct {
	p     point
	r, rr int // running risk
}

type elem struct {
	rp         *riskPoint
	prev, next *elem
}

type rpqueue struct {
	pointers    map[point]*elem
	front, back *elem
	len         int
}

func rpqueueFromOrdered(rps []*riskPoint) *rpqueue {
	q := new(rpqueue)
	q.pointers = make(map[point]*elem, len(rps))
	var prev *elem
	for i, rp := range rps {
		e := &elem{rp: rp, prev: prev}
		q.pointers[rp.p] = e
		q.len += 1
		if i == 0 {
			q.front = e
		} else {
			prev.next = e
		}
		if i == len(rps)-1 {
			q.back = e
		}
		prev = e
	}
	return q
}

func (q *rpqueue) pop() (rp *riskPoint, ok bool) {
	if q.len == 0 {
		return
	}
	q.len--

	elem := q.front
	rp = elem.rp
	ok = true
	delete(q.pointers, rp.p)

	q.front = elem.next
	if q.front == nil {
		q.back = nil
	} else {
		q.front.prev = nil
	}

	return
}

func (q *rpqueue) get(p point) (rp *riskPoint, ok bool) {
	elem, ok := q.pointers[p]
	if ok {
		rp = elem.rp
	}
	return
}

func (q *rpqueue) update(p point, rr int) {
	elem, ok := q.pointers[p]
	if !ok {
		return
	}

	elem.rp.rr = rr
	if q.len == 0 || q.len == 1 {
		return
	}

	// remove elem
	prev, next := elem.prev, elem.next
	if prev != nil {
		prev.next = next
	} else {
		q.front = next
	}
	if next != nil {
		next.prev = prev
	} else {
		q.back = prev
	}
	elem.prev, elem.next = nil, nil
	q.len--

	// re-insert elem in order based on running risk (rr)
	for cur := q.front; cur != nil; cur = cur.next {
		if elem.rp.rr < cur.rp.rr {
			elem.next = cur
			elem.prev = cur.prev
			if cur.prev == nil {
				q.front = elem
			} else {
				elem.prev.next = elem
			}
			cur.prev = elem
			break
		}
	}

	// if not inserted above, append to back
	if elem.next == nil {
		q.back.next = elem
		elem.prev = q.back
		q.back = elem
	}
	q.len++
}
