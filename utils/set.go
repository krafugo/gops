package utils

// Set is a list of elements non-repeated
type Set struct {
	Items map[interface{}]int
}

func (s *Set) append(item interface{}) {
	if _, ok := s.Items[item]; !ok {
		s.Items[item] = 1
	} else {
		s.Items[item]++
	}
}

func (s *Set) remove(item interface{}) {
	delete(s.Items, item)
}
