package hftypes

import "errors"

type Set struct {
	data   map[Any]Void
	TypeFc func(elem Any) B
	init   B
}

var (
	TypeErr = errors.New("type limit err")
)

func NewSet(typeFc TypeFc) *Set {
	s := &Set{
		TypeFc: typeFc,
		data:   map[Any]Void{},
	}
	s.Init()
	return s
}

func (s *Set) Init() {
	if !s.init {
		if s.TypeFc == nil {
			s.TypeFc = IsAny
		}
		if s.data == nil {
			s.data = map[Any]Void{}
		}
		s.init = true
	}
}

func (s *Set) Add(elem Any) (ok B) {
	s.Init()
	if s.TypeFc(elem) {
		s.data[elem] = Empty
		return true
	}
	panic(TypeErr)
}

func (s *Set) Del(elem Any) (ok B) {
	s.Init()
	if s.TypeFc(elem) {
		delete(s.data, elem)
		return true
	}
	panic(TypeErr)
}

func (s *Set) Size() I {
	return len(s.data)
}

func (s *Set) Exist(elem Any) B {
	if s.TypeFc(elem) {
		_, ok := s.data[elem]
		return ok
	}
	panic(TypeErr)
}

func (s *Set) Strings() []Str {
	var res []Str
	for any := range s.data {
		res = append(res, any.(Str))
	}
	return res
}
