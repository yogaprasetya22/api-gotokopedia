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
		SELECT  comments.id, comments.content, comments.user_id, comments.product_id, comments.created_at, users.username, users.email, users.id FROM comments JOIN users on users.id = comments.user_id where comments.product_id = $1 order by created_at desc
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
