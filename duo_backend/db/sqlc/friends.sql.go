// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: friends.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addFriendRequest = `-- name: AddFriendRequest :one
INSERT INTO friend_requests (requester_id, requestee_id, requester_name, status)
VALUES ($1, $2, $3, 'pending')
ON CONFLICT (requester_id, requestee_id) DO NOTHING
RETURNING requester_id, requestee_id, requester_name, status, created_at
`

type AddFriendRequestParams struct {
	RequesterID   uuid.UUID `json:"requester_id"`
	RequesteeID   uuid.UUID `json:"requestee_id"`
	RequesterName string    `json:"requester_name"`
}

func (q *Queries) AddFriendRequest(ctx context.Context, arg AddFriendRequestParams) (FriendRequest, error) {
	row := q.db.QueryRowContext(ctx, addFriendRequest, arg.RequesterID, arg.RequesteeID, arg.RequesterName)
	var i FriendRequest
	err := row.Scan(
		&i.RequesterID,
		&i.RequesteeID,
		&i.RequesterName,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const addFriendship = `-- name: AddFriendship :one
INSERT INTO friendships (user1_id, user2_id)
VALUES (
    CASE WHEN $1::uuid < $2::uuid THEN $1::uuid ELSE $2::uuid END,
    CASE WHEN $1::uuid < $2::uuid THEN $2::uuid ELSE $1::uuid END
)
ON CONFLICT (user1_id, user2_id) DO NOTHING
RETURNING user1_id, user2_id
`

type AddFriendshipParams struct {
	Column1 uuid.UUID `json:"column_1"`
	Column2 uuid.UUID `json:"column_2"`
}

func (q *Queries) AddFriendship(ctx context.Context, arg AddFriendshipParams) (Friendship, error) {
	row := q.db.QueryRowContext(ctx, addFriendship, arg.Column1, arg.Column2)
	var i Friendship
	err := row.Scan(&i.User1ID, &i.User2ID)
	return i, err
}

const deleteFriendRequest = `-- name: DeleteFriendRequest :one
DELETE FROM friend_requests
WHERE requester_id = $1 AND requestee_id = $2
RETURNING requester_id, requestee_id, requester_name, status, created_at
`

type DeleteFriendRequestParams struct {
	RequesterID uuid.UUID `json:"requester_id"`
	RequesteeID uuid.UUID `json:"requestee_id"`
}

func (q *Queries) DeleteFriendRequest(ctx context.Context, arg DeleteFriendRequestParams) (FriendRequest, error) {
	row := q.db.QueryRowContext(ctx, deleteFriendRequest, arg.RequesterID, arg.RequesteeID)
	var i FriendRequest
	err := row.Scan(
		&i.RequesterID,
		&i.RequesteeID,
		&i.RequesterName,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFriendship = `-- name: DeleteFriendship :one
DELETE FROM friendships
WHERE CASE WHEN $1::uuid < $2::uuid THEN user1_id ELSE user2_id END = $1::uuid
AND CASE WHEN $1::uuid < $2::uuid THEN user2_id ELSE user1_id END = $2::uuid
RETURNING user1_id, user2_id
`

type DeleteFriendshipParams struct {
	Column1 uuid.UUID `json:"column_1"`
	Column2 uuid.UUID `json:"column_2"`
}

func (q *Queries) DeleteFriendship(ctx context.Context, arg DeleteFriendshipParams) (Friendship, error) {
	row := q.db.QueryRowContext(ctx, deleteFriendship, arg.Column1, arg.Column2)
	var i Friendship
	err := row.Scan(&i.User1ID, &i.User2ID)
	return i, err
}

const getFriendsByUserId = `-- name: GetFriendsByUserId :many
SELECT user1_id, user2_id, uuid, username, user_status, score, public_key FROM friendships JOIN duouser ON (
    (friendships.user1_id = duouser.uuid AND friendships.user2_id = $1)
    OR
    (friendships.user2_id = duouser.uuid AND friendships.user1_id = $1)
)
WHERE user1_id = $1 OR user2_id = $1
`

type GetFriendsByUserIdRow struct {
	User1ID    uuid.UUID  `json:"user1_id"`
	User2ID    uuid.UUID  `json:"user2_id"`
	Uuid       uuid.UUID  `json:"uuid"`
	Username   string     `json:"username"`
	UserStatus Userstatus `json:"user_status"`
	Score      int32      `json:"score"`
	PublicKey  string     `json:"public_key"`
}

func (q *Queries) GetFriendsByUserId(ctx context.Context, user2ID uuid.UUID) ([]GetFriendsByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getFriendsByUserId, user2ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFriendsByUserIdRow{}
	for rows.Next() {
		var i GetFriendsByUserIdRow
		if err := rows.Scan(
			&i.User1ID,
			&i.User2ID,
			&i.Uuid,
			&i.Username,
			&i.UserStatus,
			&i.Score,
			&i.PublicKey,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOpenFriendRequestByRequesteeId = `-- name: GetOpenFriendRequestByRequesteeId :many
SELECT requester_id, requestee_id, requester_name, status, created_at FROM friend_requests
WHERE requestee_id = $1
AND status = 'pending'
`

func (q *Queries) GetOpenFriendRequestByRequesteeId(ctx context.Context, requesteeID uuid.UUID) ([]FriendRequest, error) {
	rows, err := q.db.QueryContext(ctx, getOpenFriendRequestByRequesteeId, requesteeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FriendRequest{}
	for rows.Next() {
		var i FriendRequest
		if err := rows.Scan(
			&i.RequesterID,
			&i.RequesteeID,
			&i.RequesterName,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFriendRequestStatus = `-- name: UpdateFriendRequestStatus :one
UPDATE friend_requests
SET status = $3
WHERE requester_id = $1 AND requestee_id = $2
RETURNING requester_id, requestee_id, requester_name, status, created_at
`

type UpdateFriendRequestStatusParams struct {
	RequesterID uuid.UUID `json:"requester_id"`
	RequesteeID uuid.UUID `json:"requestee_id"`
	Status      string    `json:"status"`
}

func (q *Queries) UpdateFriendRequestStatus(ctx context.Context, arg UpdateFriendRequestStatusParams) (FriendRequest, error) {
	row := q.db.QueryRowContext(ctx, updateFriendRequestStatus, arg.RequesterID, arg.RequesteeID, arg.Status)
	var i FriendRequest
	err := row.Scan(
		&i.RequesterID,
		&i.RequesteeID,
		&i.RequesterName,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
