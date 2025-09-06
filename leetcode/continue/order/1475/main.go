// type stack strcut {
//     top int
//     st []int{}
// }

// func NewStack() *stack {
//     return &stack {
//         top = -1
//     }
// }

// func (s *stack) empty() bool {
//     return s.top == -1
// }

// func (s *stack) top() (int, bool) {
//     if s.empty() {
//         return 0, false
//     }
//     return s.st[s.top], true
// }

// func (s *stack) pop() (int, bool) {
//     if (s.empty()) {
//         return 0, false
//     }
//     res := s.st[s.top]
//     top -= 1
//     return res, true
// }

// func (s *stack) push(x int) {
//     for !s.empty() && s.top() > x {
//         s.pop()
//     }
//     if s.top + 1 == len(s.st) {
//         s.top += 1
//         s.st = append(s.st, x)
//     } else {
//         s.top += 1
//         s[top] = x
//     }
// }

// func finalPrices(prices []int) []int {
//     ans := make([]int, len(prices))
//     for i, p := range prices {
//         j := i + 1
//         ans[i] = p
//         for j < len(prices) {
//             if prices[j] <= p {
//                 ans[i] = p - prices[j]
//                 break
//             }
//             j += 1
//         }
//     }

//     return ans
// }

func finalPrices(prices []int) []int {
	n := len(prices)
	ans := make([]int, len(prices))
	st := []int{0}
	for i := n - 1; i >= 0; i-- {
		p := prices[i]
		for len(st) > 1 && st[len(st)-1] > p {
			st = st[:len(st)-1]
		}
		ans[i] = p - st[len(st)-1]
		st = append(st, p)
	}
	return ans
}
