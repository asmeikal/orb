package wkt

import (
	"bytes"
	"github.com/paulmach/orb/encoding/internal/wkt"

	"github.com/paulmach/orb"
)

// MarshalString returns a WKT representation of the Geometry if possible.
func MarshalString(g orb.Geometry) string {
	buf := bytes.NewBuffer(nil)

	wkt.Marshal(buf, g)
	return buf.String()
}
