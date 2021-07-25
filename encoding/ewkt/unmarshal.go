package ewkt

import (
	"errors"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/internal/wkt"
	wkt2 "github.com/paulmach/orb/encoding/wkt"
	"regexp"
	"strconv"
	"strings"
)

var errInvalidSRID = errors.New("invalid SRID")

// UnmarshalPoint return point and srid by parsing ewkt point string
func UnmarshalPoint(s string) (p orb.Point, srid int, err error) {
	s, srid, err = splitEWKT(s)
	if err != nil {
		return orb.Point{}, 0, wkt.ErrWrap(err, errInvalidSRID)
	}
	p, err = wkt2.UnmarshalPoint(s)
	return p, srid, err
}

// UnmarshalMultiPoint return multipoint and srid by parsing ewkt multipoint string
func UnmarshalMultiPoint(s string) (p orb.MultiPoint, srid int, err error) {
	s, srid, err = splitEWKT(s)
	if err != nil {
		return orb.MultiPoint{}, 0, wkt.ErrWrap(err, errInvalidSRID)
	}
	p, err = wkt2.UnmarshalMultiPoint(s)
	return p, srid, err
}

// UnmarshalLineString return linestring and srid by parsing ewkt linestring string
func UnmarshalLineString(s string) (p orb.LineString, srid int, err error) {
	s, srid, err = splitEWKT(s)
	if err != nil {
		return orb.LineString{}, 0, wkt.ErrWrap(err, errInvalidSRID)
	}
	p, err = wkt2.UnmarshalLineString(s)
	return p, srid, err
}

// UnmarshalMultiLineString return multilinestring and srid by parsing ewkt multilinestring string
func UnmarshalMultiLineString(s string) (p orb.MultiLineString, srid int, err error) {
	s, srid, err = splitEWKT(s)
	if err != nil {
		return orb.MultiLineString{}, 0, wkt.ErrWrap(err, errInvalidSRID)
	}
	p, err = wkt2.UnmarshalMultiLineString(s)
	return p, srid, err
}

// UnmarshalPolygon return polygon and srid by parsing ewkt polygon string
func UnmarshalPolygon(s string) (p orb.Polygon, srid int, err error) {
	s, srid, err = splitEWKT(s)
	if err != nil {
		return orb.Polygon{}, 0, wkt.ErrWrap(err, errInvalidSRID)
	}
	p, err = wkt2.UnmarshalPolygon(s)
	return p, srid, err
}

// UnmarshalMultiPolygon return multipolygon and srid by parsing ewkt multipolygon string
func UnmarshalMultiPolygon(s string) (p orb.MultiPolygon, srid int, err error) {
	s, srid, err = splitEWKT(s)
	if err != nil {
		return orb.MultiPolygon{}, 0, wkt.ErrWrap(err, errInvalidSRID)
	}
	p, err = wkt2.UnmarshalMultiPolygon(s)
	return p, srid, err
}

// UnmarshalCollection return collection and srid by parsing ewkt collection string
func UnmarshalCollection(s string) (p orb.Collection, srid int, err error) {
	s, srid, err = splitEWKT(s)
	if err != nil {
		return orb.Collection{}, 0, wkt.ErrWrap(err, errInvalidSRID)
	}
	p, err = wkt2.UnmarshalCollection(s)
	return p, srid, err
}

// splitEWKT splits an EWKT string into the wkt part and the SRID
func splitEWKT(s string) (wkt string, srid int, err error) {
	wkt = strings.ToUpper(strings.Trim(s, " "))
	rgxp := regexp.MustCompile("^SRID=(\\d+);")
	match := rgxp.FindStringSubmatch(s)
	if match != nil {
		wkt = strings.Replace(wkt, match[0], "", -1)
		srid, err = strconv.Atoi(match[1])
	}
	return wkt, srid, err
}
