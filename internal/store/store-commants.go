package store

import (
	"context"
	"database/sql"
	"errors"
)

type Comment struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	UserID    int64  `json:"user_id"`
	ProductID int64  `json:"product_id"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
	User      User   `json:"user"`
}

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) GetByProductID(ctx context.Context, productID int64) ([]Comment, error) {
	query := `
		SELECT  comments.id, comments.content, comments.user_id, comments.product_id, comments.created_at, comments.updated_at, users.username, users.email, users.id FROM comments JOIN users on users.id = comments.user_id where comments.product_id = $1 order by created_at desc
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
		err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.ProductID, &comment.CreatedAt, &comment.UpdateAt, &comment.User.Username, &comment.User.Email, &comment.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (s *CommentStore) GetByID(ctx context.Context, id int64) (*Comment, error) {
	query := `SELECT id, content, user_id, product_id, created_at, updated_at FROM comments WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	comment := &Comment{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.ProductID, &comment.CreatedAt, &comment.UpdateAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return comment, nil
}

func (s *CommentStore) Create(ctx context.Context, comment *Comment) error {
	query := `INSERT INTO comments (content, user_id, product_id) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, comment.Content, comment.UserID, comment.ProductID).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdateAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentStore) Update(ctx context.Context, comment *Comment) error {
	query := `UPDATE comments SET content = $1, user_id = $2, product_id = $3, updated_at = now() WHERE id = $4 RETURNING updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, comment.Content, comment.UserID, comment.ProductID, comment.ID).Scan(&comment.UpdateAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil

}

func (s *CommentStore) Delete(ctx context.Context, commentID, productID int64) error {
	query := `DELETE FROM comments WHERE id = $1 AND product_id = $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, commentID, productID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
