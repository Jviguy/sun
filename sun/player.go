package sun

import (
	"github.com/sandertv/gophertunnel/minecraft"
)

type Player struct {
	conn   *minecraft.Conn
	remote *Remote
	Sun    *Sun
}
