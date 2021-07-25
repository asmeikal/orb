package wkt

import (
	"errors"
	"github.com/paulmach/orb"
	"strconv"
	"strings"
)

var (
	ErrEmptyGeometry              = errors.New("empty geometry")
	ErrUnMarshalPoint             = errors.New("unmarshal point error")
	ErrUnMarshalMultiPoint        = errors.New("unmarshal multipoint error")
	ErrUnMarshaLineString         = errors.New("unmarshal linestring error")
	ErrUnMarshaMultiLineString    = errors.New("unmarshal multilinestring error")
	ErrUnMarshaPolygon            = errors.New("unmarshal polygon error")
	ErrUnMarshaMultiPolygon       = errors.New("unmarshal multipolygon error")
	ErrUnMarshaGeometryCollection = errors.New("unmarshal collection error")

	ErrConvertToPoint              = errors.New("convert to point error")
	ErrConvertToMultiPoint         = errors.New("convert to multi point error")
	ErrConvertToLineString         = errors.New("convert to line string error")
	ErrConvertToMultiLineString    = errors.New("convert to multi line string error")
	ErrConvertToPolygon            = errors.New("convert to polygon error")
	ErrConvertToMultiPolygon       = errors.New("convert to multi polygon error")
	ErrConvertToGeometryCollection = errors.New("convert to geometry collection error")
)

// ErrWrap errWarp
func ErrWrap(err error, es ...error) error {
	s := make([]string, 0)
	if err != nil {
		s = append(s, err.Error())
	}

	for _, e := range es {
		if e != nil {
			s = append(s, e.Error())
		}
	}

	return errors.New(strings.Join(s, "\n"))
}

/*
Unmarshal return a geometry by parse wkt string
order:
	GEOMETRYCOLLECTION
	MULTIPOINT
	POINT
	MULTILINESTRING
	LINESTRING
	MULTIPOLYGON
	POLYGON
*/
func Unmarshal(s string) (geom orb.Geometry, err error) {
	s = strings.ToUpper(strings.Trim(s, " "))
	switch {
	case strings.Contains(s, "GEOMETRYCOLLECTION"):
		if s == "GEOMETRYCOLLECTION " {
			return orb.Collection{}, nil
		}
		s = strings.Replace(s, "GEOMETRYCOLLECTION", "", -1)
		c := orb.Collection{}
		ms := splitGeometryCollection(s)
		if len(ms) == 0 {
			return nil, ErrUnMarshaGeometryCollection
		}
		for _, v := range ms {
			if len(v) == 0 {
				continue
			}
			g, err := Unmarshal(v)
			if err != nil {
				return nil, ErrWrap(ErrUnMarshaGeometryCollection, err)
			}
			c = append(c, g)
		}
		geom = c

	case strings.Contains(s, "MULTIPOINT"):
		if s == "MULTIPOINT EMPTY" {
			return orb.MultiPoint{}, nil
		}
		s = strings.Replace(s, "MULTIPOINT", "", -1)
		s = trimSpaceBrackets(s)
		ps := strings.Split(s, ",")
		mp := orb.MultiPoint{}
		for _, p := range ps {
			tp, err := parsePoint(trimSpaceBrackets(p))
			if err != nil {
				return nil, ErrWrap(ErrUnMarshalPoint, err)
			}
			mp = append(mp, tp)
		}
		geom = mp

	case strings.Contains(s, "POINT"):
		s = strings.Replace(s, "POINT", "", -1)
		tp, err := parsePoint(trimSpaceBrackets(s))
		if err != nil {
			return nil, ErrWrap(ErrUnMarshalPoint, err)
		}
		geom = tp

	case strings.Contains(s, "MULTILINESTRING"):
		if s == "MULTILINESTRING EMPTY" {
			return orb.MultiLineString{}, nil
		}
		s = strings.Replace(s, "MULTILINESTRING", "", -1)
		ml := orb.MultiLineString{}
		for _, l := range strings.Split(trimSpaceBrackets(s), "),(") {
			tl := orb.LineString{}
			for _, p := range strings.Split(trimSpaceBrackets(l), ",") {
				tp, err := parsePoint(trimSpaceBrackets(p))
				if err != nil {
					return nil, ErrWrap(ErrUnMarshaMultiLineString, err)
				}
				tl = append(tl, tp)
			}
			ml = append(ml, tl)
		}
		geom = ml

	case strings.Contains(s, "LINESTRING"):
		if s == "LINESTRING EMPTY" {
			return orb.LineString{}, nil
		}
		s = strings.Replace(s, "LINESTRING", "", -1)
		s = trimSpaceBrackets(s)
		ps := strings.Split(s, ",")
		ls := orb.LineString{}
		for _, p := range ps {
			tp, err := parsePoint(trimSpaceBrackets(p))
			if err != nil {
				return nil, ErrWrap(ErrUnMarshaLineString, err)
			}
			ls = append(ls, tp)
		}
		geom = ls

	case strings.Contains(s, "MULTIPOLYGON"):
		if s == "MULTIPOLYGON EMPTY" {
			return orb.MultiPolygon{}, nil
		}
		s = strings.Replace(s, "MULTIPOLYGON", "", -1)
		mpol := orb.MultiPolygon{}
		for _, ps := range strings.Split(trimSpaceBrackets(s), ")),((") {
			pol := orb.Polygon{}
			for _, ls := range strings.Split(trimSpaceBrackets(ps), "),(") {
				ring := orb.Ring{}
				for _, p := range strings.Split(ls, ",") {
					tp, err := parsePoint(trimSpaceBrackets(p))
					if err != nil {
						return nil, ErrWrap(ErrUnMarshaMultiPolygon, err)
					}
					ring = append(ring, tp)
				}
				pol = append(pol, ring)
			}
			mpol = append(mpol, pol)
		}
		geom = mpol

	case strings.Contains(s, "POLYGON"):
		if s == "POLYGON EMPTY" {
			return orb.Polygon{}, nil
		}
		s = strings.Replace(s, "POLYGON", "", -1)
		s = trimSpaceBrackets(s)
		rs := strings.Split(s, "),(")
		if len(rs) == 1 {
			// ring
			ps := strings.Split(trimSpaceBrackets(s), ",")
			ring := orb.Ring{}
			for _, p := range ps {
				tp, err := parsePoint(trimSpaceBrackets(p))
				if err != nil {
					return nil, ErrWrap(ErrUnMarshaLineString, err)
				}
				ring = append(ring, tp)
			}
			geom = orb.Polygon{ring}
		} else {
			// more ring
			pol := orb.Polygon{}
			for _, r := range rs {
				ps := strings.Split(trimSpaceBrackets(r), ",")
				ring := orb.Ring{}
				for _, p := range ps {
					tp, err := parsePoint(trimSpaceBrackets(p))
					if err != nil {
						return nil, ErrWrap(ErrUnMarshaLineString, err)
					}
					ring = append(ring, tp)
				}
				pol = append(pol, ring)
			}
			geom = pol
		}
	default:
		return nil, errors.New("wkt: unsupported geometry")
	}

	return
}

// trimSpaceBrackets trim space and brackets
func trimSpaceBrackets(s string) string {
	s = strings.Trim(s, " ")
	if s[0] == '(' {
		s = s[1:]
	}
	if s[len(s)-1] == ')' {
		s = s[:len(s)-1]
	}
	s = strings.Trim(s, " ")
	return s
}

// parsePoint pase point by (x y)
func parsePoint(s string) (p orb.Point, err error) {
	ps := strings.Split(s, " ")
	if len(ps) != 2 {
		return orb.Point{}, errors.New("can't get x,y")
	}
	x, err := strconv.ParseFloat(ps[0], 64)
	if err != nil {
		return orb.Point{}, err
	}
	y, err := strconv.ParseFloat(ps[1], 64)
	if err != nil {
		return orb.Point{}, err
	}
	p = orb.Point{x, y}
	return p, nil
}

// splitGeometryCollection split GEOMETRYCOLLECTION to more geometry
func splitGeometryCollection(s string) (r []string) {
	r = make([]string, 0)
	stack := make([]rune, 0)
	l := len(s)
	for i, v := range s {
		if !strings.Contains(string(stack), "(") {
			stack = append(stack, v)
			continue
		}
		if v >= 'A' && v < 'Z' {
			t := string(stack)
			r = append(r, t[:len(t)-1])
			stack = make([]rune, 0)
			stack = append(stack, v)
			continue
		}
		if i == l-1 {
			r = append(r, string(stack))
			continue
		}
		stack = append(stack, v)
	}
	return
}
