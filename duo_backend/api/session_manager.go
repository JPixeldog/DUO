package api

import (
	"context"
	"fmt"
	"sync"
	"time"

	db "github.com/duo/db/sqlc"
	"github.com/duo/pb"
	"github.com/google/uuid"
)

type UserStream struct {
	Stream   pb.DUOService_JoinSessionServer
	UserId   uuid.UUID
	Username string
}

func NewUserStream(stream pb.DUOService_JoinSessionServer, userId uuid.UUID, username string) *UserStream {
	return &UserStream{
		Stream:   stream,
		Username: username,
		UserId:   userId,
	}
}

type SessionManager struct {
	SessionStreams map[int][]UserStream
	Mu             sync.Mutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		SessionStreams: make(map[int][]UserStream),
		Mu:             sync.Mutex{},
	}
}

func (sm *SessionManager) SendTestMessagesToAll() {
	go func() {
		i := 0
		for {
			for sessionId, _ := range sm.SessionStreams {
				sm.SendMessage(sessionId, &pb.SessionStream{
					SessionId: int32(sessionId),
					GameState: &pb.GameState{},
					SessionState: &pb.SessionState{
						CurrentPlayers: fmt.Sprint(i),
					},
				})
			}
			i++
			time.Sleep(1 * time.Second)
		}
	}()
}

func (sm *SessionManager) CreateSession(userUUID uuid.UUID, pin string, store db.Store) (*db.GameSession, error) {
	//TODO move to route
	// _, getErr := store.GetSessionByOwnerUUID(context.Background(), stream.UserId)
	// if getErr != sql.ErrNoRows {
	// 	return fmt.Errorf("user already owns a session")
	// }

	dbSession, createErr := store.CreateSession(context.Background(), db.CreateSessionParams{
		OwnerID: userUUID,
		Pin:     pin,
	})
	if createErr != nil {
		return nil, createErr
	}

	sm.Mu.Lock()

	sm.SessionStreams[int(dbSession.ID)] = []UserStream{}

	sm.Mu.Unlock()

	return &dbSession, nil
}

func (sm *SessionManager) GetPlayerIdsInSession(sessionId int) ([]uuid.UUID, error) {
	sm.Mu.Lock()
	defer sm.Mu.Unlock()

	//Session must exist
	if sm.SessionStreams[sessionId] == nil {
		return []uuid.UUID{}, fmt.Errorf("session does not exist")
	}

	var userIds []uuid.UUID
	for _, s := range sm.SessionStreams[sessionId] {
		userIds = append(userIds, s.UserId)
	}

	return userIds, nil
}

func (sm *SessionManager) GetSession(sessionId int) ([]UserStream, error) {
	return sm.SessionStreams[sessionId], nil
}

func (sm *SessionManager) SendMessage(sessionId int, message *pb.SessionStream) []error {
	sm.Mu.Lock()
	defer sm.Mu.Unlock()

	//Session must exist
	if sm.SessionStreams[sessionId] == nil {
		return []error{fmt.Errorf("session does not exist")}
	}

	var errs []error

	for _, s := range sm.SessionStreams[sessionId] {
		err := s.Stream.Send(message)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (sm *SessionManager) DeleteSession(sessionId int, store db.Store) error {
	_, delErr := store.DeleteSessionByID(context.Background(), int32(sessionId))
	if delErr != nil {
		return delErr
	}

	sm.Mu.Lock()
	defer sm.Mu.Unlock()

	for _, s := range sm.SessionStreams[sessionId] {
		_ = s.Stream.Context().Done() //TODO check if necessary
	}
	delete(sm.SessionStreams, sessionId)

	return nil
}

func (sm *SessionManager) GetSessionStreams(sessionId int) ([]UserStream, error) {
	return sm.SessionStreams[sessionId], nil
}

func (sm *SessionManager) AddStreamToSession(sessionId int, stream UserStream) error {
	sm.Mu.Lock()
	defer sm.Mu.Unlock()

	//Stream must exist
	if sm.SessionStreams[sessionId] == nil {
		return fmt.Errorf("session does not exist")
	}

	//Stream must be unique
	for _, s := range sm.SessionStreams[sessionId] {
		if s.UserId == stream.UserId {
			return fmt.Errorf("stream already exists in session")
		}
	}

	sm.SessionStreams[sessionId] = append(sm.SessionStreams[sessionId], stream)

	return nil
}

func (sm *SessionManager) RemoveStreamFromSession(sessionId int, userId uuid.UUID) error {
	sm.Mu.Lock()
	defer sm.Mu.Unlock()

	//Stream must exist
	if sm.SessionStreams[sessionId] == nil {
		return fmt.Errorf("session does not exist")
	}

	newStreams := []UserStream{}
	for _, s := range sm.SessionStreams[sessionId] {
		if s.UserId != userId {
			newStreams = append(newStreams, s)
		}
	}
	sm.SessionStreams[sessionId] = newStreams

	return nil
}
