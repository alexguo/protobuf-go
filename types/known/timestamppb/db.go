package timestamppb

import (
	"fmt"
	time "time"

	"github.com/golang/protobuf/jsonpb"
)

func toTime(ts *Timestamp) time.Time {
	var t time.Time
	if ts == nil {
		t = time.Unix(0, 0).UTC() // treat nil like the empty Timestamp
	} else {
		t = time.Unix(ts.Seconds, int64(ts.Nanos)).UTC()
	}
	return t
}

func (ts *Timestamp) Scan(v interface{}) error {
	t := v.(time.Time)
	seconds := t.Unix()
	nanos := int32(t.Sub(time.Unix(seconds, 0)))
	*ts = Timestamp{
		Seconds: seconds,
		Nanos:   nanos,
	}
	return nil
}

func (ts Timestamp) MarshalJSONPB(m *jsonpb.Marshaler) ([]byte, error) {
	t := toTime(&ts)
	return []byte(
		fmt.Sprintf("\"%s\"", t.Format(time.RFC3339Nano)),
	), nil
}
