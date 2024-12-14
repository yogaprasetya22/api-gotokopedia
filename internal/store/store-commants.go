package store

import (
	"context"
	"database/sql"
)

type Comment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	ProductID int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetByProductID(ctx context.Context, productID int64) ([]Comment, error) {
	query := `
		SELECT  COMMENTS.ID, COMMENTS.CONTENT, COMMENTS.USER_ID, COMMENTS.PRODUCT_ID, COMMENTS.CREATED_AT, USERS.USERNAME, USERS.EMAIL, USERS.ID FROM COMMENTS JOIN USERS ON USERS.ID = COMMENTS.USER_ID WHERE COMMENTS.PRODUCT_ID = $1 ORDER BY CREATED_AT DESC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]Comment, 0)
	for rows.Next() {
		var comment Comment
		comment.User = User{}
		err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.ProductID, &comment.CreatedAt, &comment.User.Username, &comment.User.Email, &comment.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `INSERT INTO comments (content, user_id, product_id) VALUES ($1, $2, $3) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, comment.Content, comment.UserID, comment.ProductID).Scan(&comment.ID, &comment.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
