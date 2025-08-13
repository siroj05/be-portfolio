package repository

import (
	"context"
	"database/sql"

	"github.com/siroj05/portfolio/internal/dto"
)

// MessagesRepository adalah struct yang digunakan untuk mengelola operasi database terkait pesan (messages).
// Struct ini menyimpan objek *sql.DB sebagai koneksi ke database yang digunakan untuk menjalankan query.
type MessagesRepository struct {
	db *sql.DB
}

// NewMessagesRepository membuat instance baru dari MessagesRepository.
// Fungsi ini menerima parameter *sql.DB yang digunakan untuk koneksi ke database,
// lalu mengembalikan pointer ke MessagesRepository yang sudah diinisialisasi.
// Fungsi ini biasanya digunakan untuk dependency injection pada layer service atau handler.
func NewMessagesRepository(db *sql.DB) *MessagesRepository {
	return &MessagesRepository{
		db: db,
	}
}

/*
* CreateMessage adalah method yang digunakan untuk menyimpan pesan baru ke dalam database.
 */

func (r *MessagesRepository) Create(ctx context.Context, req dto.MessageDto) error {
	query := "INSERT INTO messages (email, messages, is_read) VALUES (?, ?, ?)"

	_, err := r.db.ExecContext(ctx, query, req.Email, req.Messages, req.IsRead)

	if err != nil {
		return err
	}

	return nil
}

func (r *MessagesRepository) GetAll(ctx context.Context) ([]dto.MessageDto, error) {
	query := "SELECT * FROM messages"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var req = make([]dto.MessageDto, 0)
	for rows.Next() {
		var m dto.MessageDto
		rows.Scan(&m.ID, &m.Email, &m.Messages, &m.IsRead, &m.CreatedAt)
		req = append(req, m)
	}

	return req, nil
}
