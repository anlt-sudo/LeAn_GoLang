package repository

import (
	"database/sql"
	"fmt"
	"go-mysql/internal/model"

	"github.com/google/uuid"
)

type AlbumRepository struct {
	DB *sql.DB
}

func (r *AlbumRepository) AlbumsByArtist(name string) ([]model.Album, error){
	var albums []model.Album

	rows, err := r.DB.Query("SELECT id, title, artist, price, category_id FROM album WHERE artist = ?", name)

	if err != nil{
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
	}

	defer rows.Close()

	for rows.Next(){
		var alb model.Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price, &alb.CategoryID); err != nil {
			return nil, err
		}

		albums = append(albums, alb)
	}
	return albums, rows.Err()
}

func (r *AlbumRepository) AlbumsByCategory(catID string) ([]model.Album, error) {
	var albums []model.Album
	rows, err := r.DB.Query("SELECT id, title, artist, price, category_id FROM album WHERE category_id = ?", catID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var alb model.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price, &alb.CategoryID); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}
	return albums, nil
}

func (r *AlbumRepository) AlbumByID(id string) (model.Album, error){
	var alb model.Album

	row := r.DB.QueryRow("SELECT * FROM album WHERE id = ?", id)

	if err := row.Scan(&alb.ID, &alb.Price, &alb.Artist, &alb.Title, &alb.CategoryID); err != nil{
		if err == sql.ErrNoRows{
			return alb, fmt.Errorf("no album with id %s", id)
		}

		return  alb, nil
	}

	return alb, nil
}

func (r *AlbumRepository) AddAlbum(alb model.Album) (string, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return "", fmt.Errorf("begin tx: %w", err)
	}

	var exists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM category WHERE id = ?)", alb.CategoryID).Scan(&exists)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("check category: %w", err)
	}
	if !exists {
		tx.Rollback()
		return "", fmt.Errorf("category %s not found", alb.CategoryID)
	}

	id := uuid.New().String()
	_, err = tx.Exec(
		"INSERT INTO album (id, title, artist, price, category_id) VALUES (?, ?, ?, ?, ?)",
		id, alb.Title, alb.Artist, alb.Price, alb.CategoryID,
	)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("insert album: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return "", fmt.Errorf("commit tx: %w", err)
	}

	return id, nil
}

func (r *AlbumRepository) GetAllAlbums() ([]model.Album, error){
	rows, err := r.DB.Query("SELECT id, title, artist, price, category_id FROM album")
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	var albums []model.Album
	for rows.Next(){
		var a model.Album

		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price, &a.CategoryID); err != nil{
			return nil, err
		}

		albums = append(albums, a)
	}

	return albums, nil
}
