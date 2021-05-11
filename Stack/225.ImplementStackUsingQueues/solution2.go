package _25_ImplementStackUsingQueues

type Mystack struct {
	queue []int
}

func Constructor() (s MyStack) {
	return
}

func (s *Mystack) Push(x int)  {
	n := len(s.queue)
	s.queue = append(s.queue,x)
	for ; n>0;n-- {
		s.queue = append(s.queue,s.queue[0])
		s.queue = s.queue[1:]
	}
}

func (s *Mystack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

func (s *Mystack) Top() int {
	return s.queue[0]
}

func (s *Mystack) Empty() bool {
	return len(s.queue) == 0
}