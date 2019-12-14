package goproxy

import (
	"errors"
	"net"

	"github.com/dtsvz/websocket"
)

const (
	maxFrameHeaderSize         = 2 + 8 + 4 // Fixed header + length + mask
	maxControlFramePayloadSize = 125

	defaultReadBufferSize  = 4096
	defaultWriteBufferSize = 4096
)

func (proxy *ProxyHttpServer) cp_websocket_frames(src net.Conn, dst net.Conn, ctx *ProxyCtx, errorChan chan error, which string, cts bool) {

	srcWS := websocket.NewWSConn(src, cts, 0, 0, nil, nil, nil)
	dstWS := websocket.NewWSConn(dst, !cts, 0, 0, nil, nil, nil)

	for {
		messageType, data, err := srcWS.ReadMessage()
		if messageType == -1 {
			errorChan <- errors.New("Failed")
			return
		}
		proxy.filterWebsocketMessages(ctx, messageType, data, cts)
		if err != nil {
			errorChan <- err
			return
		}
		if err = dstWS.WriteMessage(messageType, data); err != nil {
			errorChan <- err
			return
		}
	}
}
