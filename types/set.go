package types

type Set struct {
	data      map[Any]Empty
	validType func(elem Any) bool
}

func NewSet(typeFc TypeFc) *Set {
	return &Set{
		validType: typeFc,
		data:      map[Any]Empty{},
	}
}

func (s *Set) Add(elem Any) (exist bool, err error) {
	if s.validType(elem) {
		s.data[elem] = _Empty
	}
	return false, nil
}

func (s *Set) Del(elem Any) (exist bool, err error) {
	if s.validType(elem) {
		delete(s.data, elem)
	}
	return false, nil
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Exist(elem Any) bool {
	if s.validType(elem) {
		_, ok := s.data[elem]
		return ok
	}
	return false
}

func (s *Set) Strings() []string {
	var res []string
	for any := range s.data {
		res = append(res, any.(string))
	}
	return res
}
