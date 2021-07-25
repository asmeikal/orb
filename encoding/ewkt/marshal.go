package ewkt

import (
	"bytes"
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/internal/wkt"
)

func MarshalString(g orb.Geometry, srid int) string {
	buf := bytes.NewBuffer(nil)

	wkt.Marshal(buf, g)
	return fmt.Sprintf("SRID=%d;%s", srid, buf.String())
}
