package service

import (
	"junction-brella/internal/model/core"
	"sync"

	"github.com/sirupsen/logrus"
)

type ParticipationMode int

const (
	Unknown ParticipationMode = iota
	OnlineMode
	OfflineMode
)

func ParticipationModeFromString(s string) ParticipationMode {
	switch s {
	case "offline":
		return OfflineMode
	case "online":
		return OnlineMode
	default:
		return 0
	}
}

type Participant struct {
	UserID core.UserID `json:"user_id"`

	ParticipationMode ParticipationMode `json:"participation_mode"`
}

type Meeting struct {
	Participants      [2]Participant    `json:"participants"`
	ParticipationMode ParticipationMode `json:"participation_mode"`
}

type TinderService interface {
	Next(core.UserID, ParticipationMode) Meeting
}

func NewTinderService(log *logrus.Entry) TinderService {
	return &tinderServiceImpl{
		log:          log,
		participants: make(chan match),
	}
}

type match struct {
	participant *Participant
	backchan    chan<- *Participant
}

type tinderServiceImpl struct {
	mx sync.Mutex

	participants chan match

	online  map[core.UserID]chan Participant
	matches map[core.UserID]map[core.UserID]bool

	log *logrus.Entry
}

func (ts *tinderServiceImpl) Join(uid core.UserID) (<-chan Participant, LeaveFunc) {
	ts.mx.Lock()
	defer ts.mx.Unlock()

	matches := make(chan Participant)

	ts.online[uid] = matches
	ts.matches[uid] = make(map[core.UserID]bool)

	leave := func() {
		ts.mx.Lock()
		defer ts.mx.Unlock()

		close(ts.online[uid])
		delete(ts.online, uid)
	}

	return matches, leave
}

func (ts *tinderServiceImpl) Next(uid core.UserID, mode ParticipationMode) Meeting {
	p := Participant{UserID: uid, ParticipationMode: mode}

	var m match
	select {
	case m = <-ts.participants:
		m.backchan <- &p

	default:
		response := make(chan *Participant)
		ts.participants <- match{participant: &p, backchan: response}
		m.participant = <-response
	}

	return Meeting{
		Participants:      [2]Participant{*m.participant, p},
		ParticipationMode: mode,
	}
}
