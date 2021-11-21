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
			r.Participants[pid] <- RoomEvent{Type: JoinedEvent, Peer: p}
		}
	}

	return r.Participants[pid]
}

func (r *Room) Broadcast(pid PeerID, msg string) {
	for p, c := range r.Participants {
		if p != pid {
			c <- RoomEvent{Type: Broadcast, Peer: pid, Message: msg}
		}
	}
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
	Broadcast
)

type RoomEvent struct {
	Type    RoomEventType `json:"type"`
	Peer    PeerID        `json:"peer,omitempty"`
	Message string        `json:"message,omitempty"`
}

type LeaveFunc func()

type RoomService interface {
	Join(RoomID, PeerID) (<-chan RoomEvent, LeaveFunc)

	Broadcast(PeerID, RoomID, string)
	List() []RoomInfo
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

func (rsi *roomServiceImpl) Broadcast(pid PeerID, rid RoomID, msg string) {
	rsi.mx.Lock()
	defer rsi.mx.Unlock()
	rsi.rooms[rid].Broadcast(pid, msg)
}

type RoomInfo struct {
	RoomID          RoomID `json:"room_id"`
	NumParticipants int    `json:"num_participants"`
}

func (rsi *roomServiceImpl) List() []RoomInfo {
	rsi.mx.Lock()
	defer rsi.mx.Unlock()

	var list []RoomInfo
	for id, room := range rsi.rooms {
		list = append(list, RoomInfo{RoomID: id, NumParticipants: len(room.Participants)})
	}

	return list
}
