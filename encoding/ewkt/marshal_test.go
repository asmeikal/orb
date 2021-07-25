package ewkt

import (
	"github.com/paulmach/orb"
	"testing"
)

func TestMarshalString(t *testing.T) {
	cases := []struct {
		name     string
		geo      orb.Geometry
		srid     int
		expected string
	}{
		{
			name:     "point",
			geo:      orb.Point{1, 2},
			srid:     4326,
			expected: "SRID=4326;POINT(1 2)",
		},
		{
			name:     "multipoint",
			geo:      orb.MultiPoint{{1, 2}, {0.5, 1.5}},
			srid:     4326,
			expected: "SRID=4326;MULTIPOINT((1 2),(0.5 1.5))",
		},
		{
			name:     "multipoint empty",
			geo:      orb.MultiPoint{},
			srid:     4326,
			expected: "SRID=4326;MULTIPOINT EMPTY",
		},
		{
			name:     "linestring",
			geo:      orb.LineString{{1, 2}, {0.5, 1.5}},
			srid:     4326,
			expected: "SRID=4326;LINESTRING(1 2,0.5 1.5)",
		},
		{
			name:     "linestring empty",
			geo:      orb.LineString{},
			srid:     4326,
			expected: "SRID=4326;LINESTRING EMPTY",
		},
		{
			name:     "multilinestring",
			geo:      orb.MultiLineString{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
			srid:     4326,
			expected: "SRID=4326;MULTILINESTRING((1 2,3 4),(5 6,7 8))",
		},
		{
			name:     "multilinestring empty",
			geo:      orb.MultiLineString{},
			srid:     4326,
			expected: "SRID=4326;MULTILINESTRING EMPTY",
		},
		{
			name:     "ring",
			geo:      orb.Ring{{0, 0}, {1, 0}, {1, 1}, {0, 0}},
			srid:     4326,
			expected: "SRID=4326;POLYGON((0 0,1 0,1 1,0 0))",
		},
		{
			name:     "polygon",
			geo:      orb.Polygon{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
			srid:     4326,
			expected: "SRID=4326;POLYGON((1 2,3 4),(5 6,7 8))",
		},
		{
			name:     "polygon empty",
			geo:      orb.Polygon{},
			srid:     4326,
			expected: "SRID=4326;POLYGON EMPTY",
		},
		{
			name:     "multipolygon",
			geo:      orb.MultiPolygon{{{{1, 2}, {3, 4}}}, {{{5, 6}, {7, 8}}, {{1, 2}, {5, 4}}}},
			srid:     4326,
			expected: "SRID=4326;MULTIPOLYGON(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
		},
		{
			name:     "multipolygon empty",
			geo:      orb.MultiPolygon{},
			srid:     4326,
			expected: "SRID=4326;MULTIPOLYGON EMPTY",
		},
		{
			name:     "collection",
			geo:      orb.Collection{orb.Point{1, 2}, orb.LineString{{3, 4}, {5, 6}}},
			srid:     4326,
			expected: "SRID=4326;GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(3 4,5 6))",
		},
		{
			name:     "collection empty",
			geo:      orb.Collection{},
			srid:     4326,
			expected: "SRID=4326;GEOMETRYCOLLECTION EMPTY",
		},
		{
			name:     "bound",
			geo:      orb.Bound{Min: orb.Point{0, 0}, Max: orb.Point{1, 2}},
			srid:     4326,
			expected: "SRID=4326;POLYGON((0 0,1 0,1 2,0 2,0 0))",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := MarshalString(tc.geo, tc.srid)
			if v != tc.expected {
				t.Log(v)
				t.Log(tc.expected)
				t.Errorf("incorrect ewkt marshalling")
			}
		})
	}
}
