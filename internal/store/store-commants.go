package store

import (
	"context"
	"database/sql"
	"errors"
	"math"
)

type Comment struct {
	ID        int64   `json:"id"`
	Rating    int `json:"rating"`
	Content   string  `json:"content"`
	UserID    int64   `json:"user_id"`
	ProductID int64   `json:"product_id"`
	CreatedAt string  `json:"created_at"`
	UpdateAt  string  `json:"updated_at"`
	User      User    `json:"user"`
}

type MetaCommentPaginated struct {
	Total     int       `json:"total"`
	Offset    int       `json:"offset"`
	Limit     int       `json:"limit"`
	TotalPage int       `json:"total_page"`
	Coment    []Comment `json:"comment"`
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
	query := `INSERT INTO comments (content, user_id, product_id, rating) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, comment.Content, comment.UserID, comment.ProductID, comment.Rating).Scan(&comment.ID, &comment.CreatedAt, &comment.UpdateAt)
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

func (s *CommentStore) GetComments(ctx context.Context, slugProduct string, fq PaginatedFeedQuery) (MetaCommentPaginated, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return MetaCommentPaginated{}, err
	}
	defer tx.Rollback()

	comments, err := s.GetCommentBySlugProduct(ctx, tx, slugProduct, fq)
	if err != nil {
		return MetaCommentPaginated{}, err
	}

	total, err := s.CountCommentBySlugProduct(ctx, tx, slugProduct)
	if err != nil {
		return MetaCommentPaginated{}, err
	}

	meta := MetaCommentPaginated{
		Total:     total,
		Offset:    fq.Offset,
		Limit:     fq.Limit,
		Coment:    comments,
		TotalPage: int(math.Ceil(float64(total) / float64(fq.Limit))), // Membulatkan ke atas
	}

	if err := tx.Commit(); err != nil {
		return MetaCommentPaginated{}, err
	}

	return meta, nil

}

func (s *CommentStore) GetCommentBySlugProduct(ctx context.Context, tx *sql.Tx, slugProduct string, fq PaginatedFeedQuery) ([]Comment, error) {
	query := `SELECT comments.id, comments.content, comments.rating, comments.user_id, comments.product_id, comments.created_at, comments.updated_at, users.username, users.email, users.picture, users.id FROM comments JOIN users on users.id = comments.user_id where comments.product_id = (select id from products where slug = $1) order by comments.id desc LIMIT $2 OFFSET $3`

	skip := fq.Offset * (fq.Limit)

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := tx.QueryContext(ctx, query, slugProduct, fq.Limit, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := make([]Comment, 0)
	for rows.Next() {
		var comment Comment
		comment.User = User{}
		err := rows.Scan(&comment.ID, &comment.Content, &comment.Rating, &comment.UserID, &comment.ProductID, &comment.CreatedAt, &comment.UpdateAt, &comment.User.Username, &comment.User.Email, &comment.User.Picture, &comment.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (s *CommentStore) CountCommentBySlugProduct(ctx context.Context, tx *sql.Tx, slugProduct string) (int, error) {
	query := `SELECT COUNT(*) FROM comments WHERE product_id = (select id from products where slug = $1)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var total int
	err := tx.QueryRowContext(ctx, query, slugProduct).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
