package main

// source: https://leetcode.com/problems/design-browser-history/

type BrowserHistory struct {
	cur int
	val []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		cur: 0,
		val: []string{homepage},
	}
}

func (h *BrowserHistory) Visit(url string) {
	if len(h.val) > h.cur+1 {
		h.val = h.val[:h.cur+1]
	}
	h.val = append(h.val, url)
	h.cur++
}

func (h *BrowserHistory) Back(steps int) string {
	if h.cur-steps > 0 {
		h.cur -= steps
	} else {
		h.cur = 0
	}
	return h.val[h.cur]
}

func (h *BrowserHistory) Forward(steps int) string {
	if h.cur+steps < len(h.val) {
		h.cur += steps
	} else {
		h.cur = len(h.val) - 1
	}
	return h.val[h.cur]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
