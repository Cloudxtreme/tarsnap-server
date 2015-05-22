package tarsnap

import (
	"bufio"
	"encoding/binary"
	"net"
)

type Transport struct {

	// Net Conn
	conn net.Conn

	// Read Buffer
	r *bufio.Reader

	// Write buffer
	w *bufio.Writer

	// Traffic statistics
	BytesIn  int
	BytesOut int
}

func (t *Transport) Close() (err error) {

	return t.conn.Close()
}

func (t *Transport) WriteByte(b byte) (err error) {

	if err = t.w.WriteByte(b); err != nil {

		return
	}

	if err = t.w.Flush(); err != nil {

		return
	}
	t.BytesOut += 1

	return
}

func (t *Transport) ReadByte() (b byte, err error) {

	b, err = t.r.ReadByte()
	if err != nil {

		return
	}
	t.BytesIn += 1

	return
}

func (t *Transport) ReadBytes() (packet []byte, err error) {

	for {

		var b byte
		err = binary.Read(t.r, binary.BigEndian, &b)
		if err != nil {

			return
		}
		packet = append(packet, b)

		if t.r.Buffered() == 0 {

			break
		}
	}
	t.BytesIn = len(packet)

	return
}

func (t *Transport) WriteBytes(data []byte) (err error) {

	if err = binary.Write(t.w, binary.BigEndian, data); err != nil {

		return
	}

	if err = t.w.Flush(); err != nil {

		return
	}

	t.BytesOut = len(data)

	return
}

func (t *Transport) ReadPacket() (p *Packet, err error) {

	b, err := t.ReadBytes()
	if err != nil {

		return
	}

	p = new(Packet)
	if err = UnMarshal(b, p); err != nil {

		return
	}

	return
}

func (t *Transport) WritePacket(data []byte) (err error) {

	if err = t.WriteBytes(data); err != nil {

		return
	}

	return
}
