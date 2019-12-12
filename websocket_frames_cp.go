package goproxy

import (
	"net"

	"github.com/dtsvz/websocket"
)

const (
	maxFrameHeaderSize         = 2 + 8 + 4 // Fixed header + length + mask
	maxControlFramePayloadSize = 125

	defaultReadBufferSize  = 4096
	defaultWriteBufferSize = 4096
)

func cp_websocket_frames(src net.Conn, dst net.Conn, errorChan chan error) {

	srcWS := websocket.NewWSConn(src, false, 0, 0, nil, nil, nil)
	dstWS := websocket.NewWSConn(dst, false, 0, 0, nil, nil, nil)

	for {
		messageType, data, err := srcWS.ReadMessage()
		if err != nil {
			errorChan <- err
		}
		if err = dstWS.WriteMessage(messageType, data); err != nil {
			errorChan <- err
		}
	}
	// will be implemented
}
