package _25_ImplementStackUsingQueues

type MyStack struct {
	queue1,queue2 []int
}

func Constrcutor() (s MyStack) {
	return
}

func (s *MyStack) Push(x int)  {
	s.queue2 = append(s.queue2,x)
	for len(s.queue1) > 0{
		s.queue2 = append(s.queue2,s.queue1[0])
		s.queue1 = s.queue1[1:]
	}
	s.queue1,s.queue2 = s.queue2,s.queue1
}

func (s *MyStack) Pop() int {
	v := s.queue1[0]
	s.queue1 = s.queue1[1:]
	return v
}

func (s *MyStack) Top() int {
	return s.queue1[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0
}