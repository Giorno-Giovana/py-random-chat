package controllers

import (
	"junction-brella/internal/model/core"
	"junction-brella/internal/service"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type TinderController struct {
	log      *logrus.Entry
	registry *service.Registry
}

func NewTinderController(log *logrus.Entry, registry *service.Registry) *TinderController {
	return &TinderController{log: log, registry: registry}
}

func (rc *TinderController) Next(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		uid := c.FormValue("uid")
		paricipationMode := c.FormValue("participation_mode")

		meetingC := make(chan service.Meeting)
		closeC := closeGuard(ws)

		go func() {
			meetingC <- rc.registry.TinderService.Next(
				core.UserID(uid),
				service.ParticipationModeFromString(paricipationMode),
			)
		}()

		select {
		case <-closeC:
			return
		case m := <-meetingC:
			websocket.JSON.Send(ws, m)
		}

	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func closeGuard(ws *websocket.Conn) <-chan error {
	closeC := make(chan error)
	go func() {
		for {
			var buffer []byte
			if _, err := ws.Read(buffer); err != nil {
				closeC <- err
				return
			}
		}
	}()

	return closeC
}
