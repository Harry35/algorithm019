package Week_02

type Queue struct {
    list []int
}

func (q *Queue) Pop(val int) {
if len(q.list) != 0  && q.list[0] == val {
        q.list = q.list[1:]
    }
}

func (q *Queue) Push(val int) {
	for len(q.list) != 0 && q.list[len(q.list) - 1] < val {
        q.list = q.list[:len(q.list) - 1]
    }
    q.list = append(q.list, val)
}

func (q *Queue) Max() int {
	return q.list[0]
}

func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	var win Queue
	for idx, value := range nums {
			win.Push(value)
		if idx >= k-1 {
				result = append(result, win.Max())
				win.Pop(nums[idx-k+1])
			}
		}
	return result
}
