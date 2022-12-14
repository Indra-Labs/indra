package transport

import (
	"github.com/indra-labs/indra"
	log2 "github.com/indra-labs/indra/pkg/log"
	"github.com/indra-labs/indra/pkg/slice"
)

var (
	log   = log2.GetLogger(indra.PathBase)
	check = log.E.Chk
)

type Sim chan slice.Bytes

func NewSim(bufs int) Sim { return make(Sim, bufs) }
func (d Sim) Send(b slice.Bytes) {
	d <- b
}
func (d Sim) Receive() <-chan slice.Bytes {
	return d
}
