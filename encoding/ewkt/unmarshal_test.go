package ewkt

import (
	"github.com/paulmach/orb"
	"testing"
)

func TestSplitEWKT(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected string
	}{
		{
			s:        "SRID=4326;POINT EMPTY",
			srid:     4326,
			expected: "POINT EMPTY",
		},
		{
			s:        "POINT EMPTY",
			srid:     0,
			expected: "POINT EMPTY",
		},
	}

	for _, tc := range cases {
		s, srid, err := splitEWKT(tc.s)
		if err != nil {
			t.Fatal(err)
		}
		if s != tc.expected {
			t.Log(s)
			t.Log(tc.expected)
			t.Errorf("incorrent wkt returned")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Errorf("incorrect srid returned")
		}
	}
}

func TestUnmarshalPoint(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected orb.Point
	}{
		// int
		{
			s:        "POINT(1 2)",
			srid:     0,
			expected: orb.Point{1, 2},
		},
		{
			s:        "SRID=4326;POINT(1 2)",
			srid:     4326,
			expected: orb.Point{1, 2},
		},
		// float64
		{
			s:        "POINT(1.34 2.35)",
			srid:     0,
			expected: orb.Point{1.34, 2.35},
		},
		{
			s:        "SRID=4326;POINT(1.34 2.35)",
			srid:     4326,
			expected: orb.Point{1.34, 2.35},
		},
	}

	for _, tc := range cases {
		geom, srid, err := UnmarshalPoint(tc.s)
		if err != nil {
			t.Fatal(err)
		}
		if !geom.Equal(tc.expected) {
			t.Log(geom)
			t.Log(tc.expected)
			t.Errorf("incorrect ewkt unmarshalling")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Error("incorrect srid")
		}
	}
}

func TestUnmarshalMultiPoint(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected orb.MultiPoint
	}{
		// empty
		{
			s:        "MULTIPOINT EMPTY",
			srid:     0,
			expected: orb.MultiPoint{},
		},
		{
			s:        "SRID=4326;MULTIPOINT EMPTY",
			srid:     4326,
			expected: orb.MultiPoint{},
		},
		// int
		{
			s:        "MULTIPOINT((1 2),(0.5 1.5))",
			srid:     0,
			expected: orb.MultiPoint{{1, 2}, {0.5, 1.5}},
		},
		{
			s:        "SRID=4326;MULTIPOINT((1 2),(0.5 1.5))",
			srid:     4326,
			expected: orb.MultiPoint{{1, 2}, {0.5, 1.5}},
		},
	}

	for _, tc := range cases {
		geom, srid, err := UnmarshalMultiPoint(tc.s)
		if err != nil {
			t.Fatal(err)
		}
		if !geom.Equal(tc.expected) {
			t.Log(geom)
			t.Log(tc.expected)
			t.Errorf("incorrect wkt unmarshalling")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Error("incorrect srid")
		}
	}
}

func TestUnmarshalLineString(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected orb.LineString
	}{
		{
			s:        "LINESTRING EMPTY",
			srid:     0,
			expected: orb.LineString{},
		},
		{
			s:        "SRID=4326;LINESTRING EMPTY",
			srid:     4326,
			expected: orb.LineString{},
		},

		{
			s:        "LINESTRING(1 2,0.5 1.5)",
			srid:     0,
			expected: orb.LineString{{1, 2}, {0.5, 1.5}},
		},
		{
			s:        "SRID=4326;LINESTRING(1 2,0.5 1.5)",
			srid:     4326,
			expected: orb.LineString{{1, 2}, {0.5, 1.5}},
		},
	}

	for _, tc := range cases {
		geom, srid, err := UnmarshalLineString(tc.s)
		if err != nil {
			t.Fatal(err)
		}
		if !geom.Equal(tc.expected) {
			t.Log(geom)
			t.Log(tc.expected)
			t.Errorf("incorrect wkt unmarshalling")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Error("incorrect srid")
		}
	}
}

func TestUnmarshalMultiLineString(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected orb.MultiLineString
	}{
		{
			s:        "MULTILINESTRING EMPTY",
			srid:     0,
			expected: orb.MultiLineString{},
		},
		{
			s:        "SRID=4326;MULTILINESTRING EMPTY",
			srid:     4326,
			expected: orb.MultiLineString{},
		},

		{
			s:        "MULTILINESTRING((1 2,3 4),(5 6,7 8))",
			srid:     0,
			expected: orb.MultiLineString{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
		},
		{
			s:        "SRID=4326;MULTILINESTRING((1 2,3 4),(5 6,7 8))",
			srid:     4326,
			expected: orb.MultiLineString{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
		},
	}

	for _, tc := range cases {
		geom, srid, err := UnmarshalMultiLineString(tc.s)
		if err != nil {
			t.Fatal(err)
		}
		if !geom.Equal(tc.expected) {
			t.Log(geom)
			t.Log(tc.expected)
			t.Errorf("incorrect wkt unmarshalling")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Error("incorrect srid")
		}
	}
}

func TestUnmarshalPolygon(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected orb.Polygon
	}{
		// empty
		{
			s:        "POLYGON EMPTY",
			srid:     0,
			expected: orb.Polygon{},
		},
		{
			s:        "SRID=4326;POLYGON EMPTY",
			srid:     4326,
			expected: orb.Polygon{},
		},
		// ring
		// origin: orb.Ring{{0, 0}, {1, 0}, {1, 1}, {0, 0}}
		// convert: orb.Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 0}}}
		{
			s:        "POLYGON((0 0,1 0,1 1,0 0))",
			srid:     0,
			expected: orb.Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 0}}},
		},
		{
			s:        "SRID=4326;POLYGON((0 0,1 0,1 1,0 0))",
			srid:     4326,
			expected: orb.Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 0}}},
		},
		// bound
		// origin: orb.Bound{Min: orb.Point{0, 0}, Max: orb.Point{1, 2}},
		// convert: orb.Polygon{{{0, 0}, {1, 0}, {1, 2}, {0, 2}, {0, 0}}}
		{
			s:        "POLYGON((0 0,1 0,1 2,0 2,0 0))",
			srid:     0,
			expected: orb.Polygon{{{0, 0}, {1, 0}, {1, 2}, {0, 2}, {0, 0}}},
		},
		{
			s:        "SRID=4326;POLYGON((0 0,1 0,1 2,0 2,0 0))",
			srid:     4326,
			expected: orb.Polygon{{{0, 0}, {1, 0}, {1, 2}, {0, 2}, {0, 0}}},
		},
		// polygon
		{
			s:        "POLYGON((1 2,3 4),(5 6,7 8))",
			srid:     0,
			expected: orb.Polygon{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
		},
		{
			s:        "SRID=4326;POLYGON((1 2,3 4),(5 6,7 8))",
			srid:     4326,
			expected: orb.Polygon{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
		},
	}

	for _, tc := range cases {
		geom, srid, err := UnmarshalPolygon(tc.s)
		if err != nil {
			t.Fatal(err)
		}
		if !geom.Equal(tc.expected) {
			t.Log(geom)
			t.Log(tc.expected)
			t.Errorf("incorrect wkt unmarshalling")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Error("incorrect srid")
		}
	}
}

func TestUnmarshalMutilPolygon(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected orb.MultiPolygon
	}{
		// empty
		{
			s:        "MULTIPOLYGON EMPTY",
			srid:     0,
			expected: orb.MultiPolygon{},
		},
		{
			s:        "SRID=4326;MULTIPOLYGON EMPTY",
			srid:     4326,
			expected: orb.MultiPolygon{},
		},
		// multi polygon
		{
			s:        "MULTIPOLYGON(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			srid:     0,
			expected: orb.MultiPolygon{{{{1, 2}, {3, 4}}}, {{{5, 6}, {7, 8}}, {{1, 2}, {5, 4}}}},
		},
		{
			s:        "SRID=4326;MULTIPOLYGON(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			srid:     4326,
			expected: orb.MultiPolygon{{{{1, 2}, {3, 4}}}, {{{5, 6}, {7, 8}}, {{1, 2}, {5, 4}}}},
		},
	}

	for _, tc := range cases {
		geom, srid, err := UnmarshalMultiPolygon(tc.s)
		if err != nil {
			t.Fatal(err)
		}
		if !geom.Equal(tc.expected) {
			t.Log(geom)
			t.Log(tc.expected)
			t.Errorf("incorrect wkt unmarshalling")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Error("incorrect srid")
		}
	}
}

func TestUnmarshalCollection(t *testing.T) {
	cases := []struct {
		s        string
		srid     int
		expected orb.Collection
	}{
		// empty
		{
			s:        "GEOMETRYCOLLECTION EMPTY",
			srid:     0,
			expected: orb.Collection{},
		},
		{
			s:        "SRID=4326;GEOMETRYCOLLECTION EMPTY",
			srid:     4326,
			expected: orb.Collection{},
		},
		// multi polygon
		{
			s:        "GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(3 4,5 6))",
			srid:     0,
			expected: orb.Collection{orb.Point{1, 2}, orb.LineString{{3, 4}, {5, 6}}},
		},
		{
			s:        "SRID=4326;GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(3 4,5 6))",
			srid:     4326,
			expected: orb.Collection{orb.Point{1, 2}, orb.LineString{{3, 4}, {5, 6}}},
		},
		{
			s:    "GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(3 4,5 6),MULTILINESTRING((1 2,3 4),(5 6,7 8)),POLYGON((0 0,1 0,1 1,0 0)),POLYGON((1 2,3 4),(5 6,7 8)),MULTIPOLYGON(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			srid: 0,
			expected: orb.Collection{
				orb.Point{1, 2},
				orb.LineString{{3, 4}, {5, 6}},
				orb.MultiLineString{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
				orb.Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 0}}},
				orb.Polygon{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
				orb.MultiPolygon{{{{1, 2}, {3, 4}}}, {{{5, 6}, {7, 8}}, {{1, 2}, {5, 4}}}},
			},
		},
		{
			s:    "SRID=4326;GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(3 4,5 6),MULTILINESTRING((1 2,3 4),(5 6,7 8)),POLYGON((0 0,1 0,1 1,0 0)),POLYGON((1 2,3 4),(5 6,7 8)),MULTIPOLYGON(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			srid: 4326,
			expected: orb.Collection{
				orb.Point{1, 2},
				orb.LineString{{3, 4}, {5, 6}},
				orb.MultiLineString{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
				orb.Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 0}}},
				orb.Polygon{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
				orb.MultiPolygon{{{{1, 2}, {3, 4}}}, {{{5, 6}, {7, 8}}, {{1, 2}, {5, 4}}}},
			},
		},
	}

	for _, tc := range cases {
		geom, srid, err := UnmarshalCollection(tc.s)
		if err != nil {
			// t.Fatal(err)
		}
		if !geom.Equal(tc.expected) {
			t.Log(geom)
			t.Log(tc.expected)
			t.Errorf("incorrect wkt unmarshalling")
		}
		if srid != tc.srid {
			t.Log(srid)
			t.Log(tc.srid)
			t.Error("incorrect srid")
		}
	}
}
