package wkt

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/internal/wkt"
)

// UnmarshalPoint return point by parse wkt point string
func UnmarshalPoint(s string) (p orb.Point, err error) {
	geom, err := wkt.Unmarshal(s)
	if err != nil {
		return orb.Point{}, wkt.ErrWrap(err, wkt.ErrEmptyGeometry)
	}
	g, ok := geom.(orb.Point)
	if !ok {
		return orb.Point{}, wkt.ErrWrap(err, wkt.ErrConvertToPoint)
	}
	return g, nil
}

// UnmarshalMultiPoint return multipoint by parse wkt multipoint string
func UnmarshalMultiPoint(s string) (p orb.MultiPoint, err error) {
	geom, err := wkt.Unmarshal(s)
	if err != nil {
		return orb.MultiPoint{}, wkt.ErrWrap(err, wkt.ErrEmptyGeometry)
	}
	g, ok := geom.(orb.MultiPoint)
	if !ok {
		return orb.MultiPoint{}, wkt.ErrWrap(err, wkt.ErrConvertToMultiPoint)
	}
	return g, nil
}

// UnmarshalLineString return linestring by parse wkt linestring string
func UnmarshalLineString(s string) (p orb.LineString, err error) {
	geom, err := wkt.Unmarshal(s)
	if err != nil {
		return orb.LineString{}, wkt.ErrWrap(err, wkt.ErrEmptyGeometry)
	}
	g, ok := geom.(orb.LineString)
	if !ok {
		return orb.LineString{}, wkt.ErrWrap(err, wkt.ErrConvertToLineString)
	}
	return g, nil
}

// UnmarshalMultiLineString return linestring by parse wkt multilinestring string
func UnmarshalMultiLineString(s string) (p orb.MultiLineString, err error) {
	geom, err := wkt.Unmarshal(s)
	if err != nil {
		return orb.MultiLineString{}, wkt.ErrWrap(err, wkt.ErrEmptyGeometry)
	}
	g, ok := geom.(orb.MultiLineString)
	if !ok {
		return orb.MultiLineString{}, wkt.ErrWrap(err, wkt.ErrConvertToMultiLineString)
	}
	return g, nil
}

// UnmarshalPolygon return linestring by parse wkt polygon string
func UnmarshalPolygon(s string) (p orb.Polygon, err error) {
	geom, err := wkt.Unmarshal(s)
	if err != nil {
		return orb.Polygon{}, wkt.ErrWrap(err, wkt.ErrEmptyGeometry)
	}
	g, ok := geom.(orb.Polygon)
	if !ok {
		return orb.Polygon{}, wkt.ErrWrap(err, wkt.ErrConvertToPolygon)
	}
	return g, nil
}

// UnmarshalMultiPolygon return linestring by parse wkt multipolygon string
func UnmarshalMultiPolygon(s string) (p orb.MultiPolygon, err error) {
	geom, err := wkt.Unmarshal(s)
	if err != nil {
		return orb.MultiPolygon{}, wkt.ErrWrap(err, wkt.ErrEmptyGeometry)
	}
	g, ok := geom.(orb.MultiPolygon)
	if !ok {
		return orb.MultiPolygon{}, wkt.ErrWrap(err, wkt.ErrConvertToMultiPolygon)
	}
	return g, nil
}

// UnmarshalCollection return linestring by parse wkt collection string
func UnmarshalCollection(s string) (p orb.Collection, err error) {
	geom, err := wkt.Unmarshal(s)
	if err != nil {
		return orb.Collection{}, wkt.ErrWrap(err, wkt.ErrEmptyGeometry)
	}
	g, ok := geom.(orb.Collection)
	if !ok {
		return orb.Collection{}, wkt.ErrWrap(err, wkt.ErrConvertToGeometryCollection)
	}
	return g, nil
}
