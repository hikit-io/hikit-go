package types

import "errors"

type Set struct {
	data   map[Any]Void
	TypeFc func(elem Any) bool
}

var (
	TypeErr = errors.New("type limit err")
)

func NewSet(typeFc TypeFc) *Set {
	return &Set{
		TypeFc: typeFc,
		data:   map[Any]Void{},
	}
}

func (s *Set) Add(elem Any) (ok bool) {
	if s.TypeFc(elem) {
		s.data[elem] = _Empty
		return true
	}
	panic(TypeErr)
}

func (s *Set) Del(elem Any) (ok bool) {
	if s.TypeFc(elem) {
		delete(s.data, elem)
		return true
	}
	panic(TypeErr)
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) Exist(elem Any) bool {
	if s.TypeFc(elem) {
		_, ok := s.data[elem]
		return ok
	}
	panic(TypeErr)
}

func (s *Set) Strings() []string {
	var res []string
	for any := range s.data {
		res = append(res, any.(string))
	}
	return res
}
