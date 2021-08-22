package db

import (
	"database/sql"
	"log"

	posts "github.com/burxtx/car/posts"
	_ "github.com/lib/pq"
)

type Config struct {
	Host            string `mapstructure:"host"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
	Debug           bool   `mapstructure:"debug"`
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(cfg Config) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", cfg.Host)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db: db,
	}, nil
}

func (r *PostgresRepository) Close() {
	if err := r.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (r *PostgresRepository) Create(post *posts.Post) error {
	return nil
}

func (r *PostgresRepository) Find(ID string) (*posts.Post, error) {
	return nil, nil
}

func (r *PostgresRepository) List() []*posts.Post {
	return nil
}
