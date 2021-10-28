package hfmongo

import (
	"fmt"
	"strings"
)

type Doc interface {
	DocName() string
}

func (o updateOp) GenPlaceAll(prefix, suffix string) string {
	return strings.Join([]string{prefix, o.PlaceholderAll, suffix}, ".")
}
func (o updateOp) GenPlaceSome(prefix, suffix string) string {
	return strings.Join([]string{prefix, o.PlaceholderSome, suffix}, ".")
}

type PlaceIndex struct {
	base      string
	indexBase string
}

func (o updateOp) GenPlaceIndex(prefix, index string) *PlaceIndex {
	return &PlaceIndex{
		base:      strings.Join([]string{prefix, fmt.Sprintf(o.PlaceholderIndex, index)}, "."),
		indexBase: index,
	}
}

func (p *PlaceIndex) Base() string {
	return p.base
}

func (p *PlaceIndex) EBase() string {
	return p.indexBase
}

func (p *PlaceIndex) Index(i string) *PlaceIndex {
	return &PlaceIndex{base: strings.Join([]string{p.base, i}, "."), indexBase: p.indexBase}
}

func (p *PlaceIndex) EIndex(i string) *PlaceIndex {
	return &PlaceIndex{base: p.base, indexBase: strings.Join([]string{p.indexBase, i}, ".")}
}
