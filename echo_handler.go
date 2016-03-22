package telnet


import (
	"github.com/reiver/go-oi"

	"io"
)


// EchoHandler is a simple TELNET server which "echos" back to the client any (non-command)
// data back to the TELNET client, it received from the TELNET client.
var EchoHandler Handler = internalEchoHandler{}


type internalEchoHandler struct{}


func (handler internalEchoHandler) ServeTELNET(ctx Context, w io.Writer, r io.Reader) {

	writer := NewDataWriter(w)
	reader := NewDataReader(r)


	var buffer [1]byte // Seems like the length of the buffer needs to be small, otherwise will have to wait for buffer to fill up.
	p := buffer[:]

	for {
		n, err := reader.Read(p)

		if n > 0 {
			oi.LongWrite(writer, p[:n])
		}

		if nil != err {
			break
		}
	}
}
