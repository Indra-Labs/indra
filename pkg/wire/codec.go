package wire

import (
	"fmt"

	"github.com/indra-labs/indra"
	log2 "github.com/indra-labs/indra/pkg/log"
	"github.com/indra-labs/indra/pkg/slice"
	"github.com/indra-labs/indra/pkg/types"
	"github.com/indra-labs/indra/pkg/wire/cipher"
	"github.com/indra-labs/indra/pkg/wire/confirm"
	"github.com/indra-labs/indra/pkg/wire/delay"
	"github.com/indra-labs/indra/pkg/wire/exit"
	"github.com/indra-labs/indra/pkg/wire/forward"
	"github.com/indra-labs/indra/pkg/wire/layer"
	"github.com/indra-labs/indra/pkg/wire/magicbytes"
	"github.com/indra-labs/indra/pkg/wire/response"
	"github.com/indra-labs/indra/pkg/wire/reverse"
	"github.com/indra-labs/indra/pkg/wire/session"
	"github.com/indra-labs/indra/pkg/wire/token"
)

var (
	log   = log2.GetLogger(indra.PathBase)
	check = log.E.Chk
)

func EncodeOnion(on types.Onion) (b slice.Bytes) {
	b = make(slice.Bytes, on.Len())
	var sc slice.Cursor
	c := &sc
	on.Encode(b, c)
	return
}

func PeelOnion(b slice.Bytes, c *slice.Cursor) (on types.Onion, e error) {
	switch b[*c:c.Inc(magicbytes.Len)].String() {
	case cipher.MagicString:
		o := &cipher.OnionSkin{}
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case confirm.MagicString:
		o := &confirm.OnionSkin{}
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case delay.MagicString:
		o := &delay.OnionSkin{}
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case exit.MagicString:
		o := &exit.OnionSkin{}
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case forward.MagicString:
		o := &forward.OnionSkin{}
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case layer.MagicString:
		var o layer.OnionSkin
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = &o
	case reverse.MagicString:
		o := &reverse.OnionSkin{}
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case response.MagicString:
		o := response.New()
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case session.MagicString:
		o := &session.OnionSkin{}
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	case token.MagicString:
		o := token.NewOnionSkin()
		if e = o.Decode(b, c); check(e) {
			return
		}
		on = o
	default:
		e = fmt.Errorf("message magic not found")
		check(e)
		log.I.S(b.ToBytes())
		return
	}
	return
}
