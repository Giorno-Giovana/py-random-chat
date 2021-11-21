package controllers

import (
	"junction-brella/internal/service"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type RoomController struct {
	log      *logrus.Entry
	registry *service.Registry
}

func NewRoomController(log *logrus.Entry, registry *service.Registry) *RoomController {
	return &RoomController{log: log, registry: registry}
}

func (c *RoomController) List(ctx echo.Context) error {
	list := c.registry.RoomService.List()

	return ctx.JSON(http.StatusOK, struct {
		List []service.RoomInfo `json:"list"`
	}{List: list})
}

func (rc *RoomController) Join(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		roomID := c.FormValue("rid")
		peerID := c.FormValue("pid")

		if roomID == "" || peerID == "" {
			// TODO: errors
			return
		}

		events, leave := rc.registry.RoomService.Join(
			service.RoomID(roomID),
			service.PeerID(peerID),
		)

		defer leave()

		closeC := make(chan error)

		go func() {
			for {
				var buffer string
				if err := websocket.Message.Receive(ws, &buffer); err != nil {
					closeC <- err
					return
				}

				rc.registry.RoomService.Broadcast(
					service.PeerID(peerID),
					service.RoomID(roomID),
					buffer,
				)
			}
		}()

		for {
			select {
			case event := <-events:
				if err := websocket.JSON.Send(ws, event); err != nil {
					rc.log.Error(err)
				}
			case <-closeC:
				return
			}
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}
