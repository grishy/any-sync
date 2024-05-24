package handshake

import (
	"encoding/binary"
	"io"
	"slices"
	"sync"

	"github.com/anyproto/any-sync/net/internal/secureservice/handshake/handshakeproto"
	handshake2 "github.com/anyproto/any-sync/net/secureservice/handshake"
)

const headerSize = 5 // 1 byte for type + 4 byte for uint32 size

const (
	msgTypeCred  = byte(1)
	msgTypeAck   = byte(2)
	msgTypeProto = byte(3)

	sizeLimit = 200 * 1024 // 200 Kb
)

var (
	credMsgTypes     = []byte{msgTypeCred, msgTypeAck}
	protoMsgTypes    = []byte{msgTypeProto, msgTypeAck}
	protoMsgTypesAck = []byte{msgTypeAck}
)

var handshakePool = &sync.Pool{New: func() any {
	return &handshake{
		remoteCred:  &handshakeproto.Credentials{},
		remoteAck:   &handshakeproto.Ack{},
		localAck:    &handshakeproto.Ack{},
		remoteProto: &handshakeproto.Proto{},
		buf:         make([]byte, 0, 1024),
	}
}}

func newHandshake() *handshake {
	return handshakePool.Get().(*handshake)
}

type handshake struct {
	conn        io.ReadWriteCloser
	remoteCred  *handshakeproto.Credentials
	remoteProto *handshakeproto.Proto
	remoteAck   *handshakeproto.Ack
	localAck    *handshakeproto.Ack
	buf         []byte
}

func (h *handshake) writeCredentials(cred *handshakeproto.Credentials) (err error) {
	h.buf = slices.Grow(h.buf, cred.Size()+headerSize)[:cred.Size()+headerSize]
	n, err := cred.MarshalToSizedBuffer(h.buf[headerSize:])
	if err != nil {
		return err
	}
	return h.writeData(msgTypeCred, n)
}

func (h *handshake) writeProto(proto *handshakeproto.Proto) (err error) {
	h.buf = slices.Grow(h.buf, proto.Size()+headerSize)[:proto.Size()+headerSize]
	n, err := proto.MarshalToSizedBuffer(h.buf[headerSize:])
	if err != nil {
		return err
	}
	return h.writeData(msgTypeProto, n)
}

func (h *handshake) tryWriteErrAndClose(err error) {
	if err == handshake2.ErrUnexpectedPayload {
		// if we got unexpected message - just close the connection
		_ = h.conn.Close()
		return
	}
	var ackErr handshakeproto.Error
	if he, ok := err.(handshake2.HandshakeError); ok {
		ackErr = he.E
	} else {
		ackErr = handshakeproto.Error_Unexpected
	}
	_ = h.writeAck(ackErr)
	_ = h.conn.Close()
}

func (h *handshake) writeAck(ackErr handshakeproto.Error) (err error) {
	h.localAck.Error = ackErr
	h.buf = slices.Grow(h.buf, h.localAck.Size()+headerSize)[:h.localAck.Size()+headerSize]
	n, err := h.localAck.MarshalTo(h.buf[headerSize:])
	if err != nil {
		return err
	}
	return h.writeData(msgTypeAck, n)
}

func (h *handshake) writeData(tp byte, size int) (err error) {
	h.buf[0] = tp
	binary.LittleEndian.PutUint32(h.buf[1:headerSize], uint32(size))
	_, err = h.conn.Write(h.buf[:size+headerSize])
	return err
}

type message struct {
	cred  *handshakeproto.Credentials
	proto *handshakeproto.Proto
	ack   *handshakeproto.Ack
}

func (h *handshake) readMsg(allowedTypes ...byte) (msg message, err error) {
	h.buf = slices.Grow(h.buf, headerSize)[:headerSize]
	if _, err = io.ReadFull(h.conn, h.buf[:headerSize]); err != nil {
		return
	}
	tp := h.buf[0]
	if !slices.Contains(allowedTypes, tp) {
		err = handshake2.ErrUnexpectedPayload
		return
	}
	size := binary.LittleEndian.Uint32(h.buf[1:headerSize])
	if size > sizeLimit {
		err = handshake2.ErrGotUnexpectedMessage
		return
	}
	h.buf = slices.Grow(h.buf, int(size))[:size]
	if _, err = io.ReadFull(h.conn, h.buf[:size]); err != nil {
		return
	}
	switch tp {
	case msgTypeCred:
		if err = h.remoteCred.Unmarshal(h.buf[:size]); err != nil {
			return
		}
		msg.cred = h.remoteCred
	case msgTypeAck:
		if err = h.remoteAck.Unmarshal(h.buf[:size]); err != nil {
			return
		}
		msg.ack = h.remoteAck
	case msgTypeProto:
		if err = h.remoteProto.Unmarshal(h.buf[:size]); err != nil {
			return
		}
		msg.proto = h.remoteProto
	}
	return
}

func (h *handshake) release() {
	h.buf = h.buf[:0]
	h.conn = nil
	h.localAck.Error = 0
	h.remoteAck.Error = 0
	h.remoteCred.Type = 0
	h.remoteCred.Payload = h.remoteCred.Payload[:0]
	h.remoteProto.Proto = 0
	handshakePool.Put(h)
}
