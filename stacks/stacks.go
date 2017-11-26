package stacks

type Stack []string

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() string {
	return (*s)[len(*s) - 1]
}

func (s *Stack) Push(ss string) {
	*s = append(*s, ss)
}

func (s *Stack) Pop() string {
	ss := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return ss
}

func (s *Stack) Size() int {
	return len(*s)
}

type IntStack []int

func (s *IntStack) Empty() bool {
	return len(*s) == 0
}

func (s *IntStack) Peek() int {
	return (*s)[len(*s) - 1]
}

func (s *IntStack) Push(i int) {
	*s = append(*s, i)
}

func (s *IntStack) Pop() int {
	i := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return i
}

func (s *IntStack) Size() int {
	return len(*s)
}
