package unmarshal

import (
	"io"
	"unsafe"

	json "github.com/json-iterator/go"

	"github.com/grafana/loki/pkg/loghttp"
	"github.com/grafana/loki/pkg/logproto"
)

// DecodePushRequest directly decodes json to a logproto.PushRequest
func DecodePushRequest(b io.Reader, r *logproto.PushRequest) error {
	var request loghttp.PushRequest

	if err := json.NewDecoder(b).Decode(&request); err != nil {
		return err
	}
	*r = NewPushRequest(request)

	return nil
}

// NewPushRequest constructs a logproto.PushRequest from a PushRequest
func NewPushRequest(r loghttp.PushRequest) logproto.PushRequest {
	ret := logproto.PushRequest{
		Streams: make([]logproto.Stream, len(r.Streams)),
	}

	for i, s := range r.Streams {
		ret.Streams[i] = NewStream(s)
	}

	return ret
}

// NewStream constructs a logproto.Stream from a Stream
func NewStream(s *loghttp.Stream) logproto.Stream {
	return logproto.Stream{
		Entries: *(*[]logproto.Entry)(unsafe.Pointer(&s.Entries)),
		Labels:  s.Labels.String(),
	}
}
