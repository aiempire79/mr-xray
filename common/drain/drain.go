package drain

import "io"

//go:generate go run github.com/aiempire79/mr-xray/common/errors/errorgen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
