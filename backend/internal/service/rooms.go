package service

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type PeerID string
type RoomID string

type Room struct {
	Participants map[PeerID]chan RoomEvent
}

func NewRoom() *Room {
	return &Room{Participants: make(map[PeerID]chan RoomEvent)}
}

func (r *Room) Add(pid PeerID) <-chan RoomEvent {
	for _, c := range r.Participants {
		c <- RoomEvent{Type: JoinedEvent, Peer: pid}
	}

	r.Participants[pid] = make(chan RoomEvent, len(r.Participants))
	for p := range r.Participants {
		if p != pid {
			r.Participants[pid] <- RoomEvent{JoinedEvent, p}
		}
	}

	return r.Participants[pid]
}

func (r *Room) Remove(pid PeerID) {
	close(r.Participants[pid])
	delete(r.Participants, pid)

	for _, c := range r.Participants {
		c <- RoomEvent{Type: LeftEvent, Peer: pid}
	}
}

type RoomEventType int

const (
	JoinedEvent RoomEventType = iota + 1
	LeftEvent
)

type RoomEvent struct {
	Type RoomEventType `json:"type"`
	Peer PeerID        `json:"peer"`
}

type LeaveFunc func()

type RoomService interface {
	Join(RoomID, PeerID) (<-chan RoomEvent, LeaveFunc)
}

func NewRoomService(log *logrus.Entry) RoomService {
	return &roomServiceImpl{
		rooms: make(map[RoomID]*Room),
		log:   log,
	}
}

type roomServiceImpl struct {
	mx    sync.Mutex
	rooms map[RoomID]*Room

	log *logrus.Entry
}

func (rsi *roomServiceImpl) Join(rid RoomID, pid PeerID) (<-chan RoomEvent, LeaveFunc) {
	rsi.mx.Lock()
	defer rsi.mx.Unlock()

	// create or get room
	room, found := rsi.rooms[rid]
	if !found {
		room = NewRoom()
		rsi.rooms[rid] = room
	}

	events := room.Add(pid)
	leave := func() {
		rsi.mx.Lock()
		defer rsi.mx.Unlock()
		rsi.rooms[rid].Remove(pid)
	}

	return events, leave
}
