package repository

import (
	"go-posts/dto"
	"go-posts/models"
	"go-posts/utils"
	"time"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll(offset, limit int) ([]models.Post, int64, error)
	FindByID(id int) (models.Post, error)
	FindByTitle(title string) (models.Post, *gorm.DB)
	Create(input dto.PublicPost) (models.Post, error)
	Update(id int, input *dto.PublicPost) (models.Post, error)
	Delete(id int) *gorm.DB
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) GetAll(offset, limit int) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	r.db.Model(&models.Post{}).Count(&total)

	err := r.db.Limit(limit).Offset(offset).Find(&posts).Error
	return posts, total, err
}

func (r *postRepository) FindByID(id int) (models.Post, error) {
	var post models.Post
	err := r.db.First(&post, "id = ?", id).Error
	return post, err
}

func (r *postRepository) FindByTitle(title string) (models.Post, *gorm.DB) {
	var post models.Post
	result := r.db.First(&post, "title = ?", title)
	return post, result
}

func (r *postRepository) Create(input dto.PublicPost) (models.Post, error) {
	post := models.Post{
		Title:       utils.CapitalizeWord(input.Title),
		Content:     input.Content,
		Category:    utils.CapitalizeWord(input.Category),
		Status:      utils.CapitalizeWord(input.Status),
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	err := r.db.Create(&post).Error
	return post, err
}

func (r *postRepository) Update(id int, input *dto.PublicPost) (models.Post, error) {
	var post models.Post
	trx := r.db.Begin()
	if trx.Error != nil {
		trx.Rollback()
		return models.Post{}, trx.Error
	}

	if err := trx.First(&post, id).Error; err != nil {
		trx.Rollback()
		return models.Post{}, err
	}

	data := map[string]any{
		"Title":       utils.CapitalizeWord(input.Title),
		"Content":     input.Content,
		"Category":    utils.CapitalizeWord(input.Category),
		"Status":      utils.CapitalizeWord(input.Status),
		"UpdatedDate": time.Now(),
	}

	if err := utils.AssignedKeyModel(&post, data); err != nil {
		trx.Rollback()
		return models.Post{}, err
	}

	if err := trx.Save(&post).Error; err != nil {
		trx.Rollback()
		return models.Post{}, err
	}

	if err := trx.Commit().Error; err != nil {
		return models.Post{}, err
	}

	return models.Post{}, nil
}

func (r *postRepository) Delete(id int) *gorm.DB {
	return r.db.Delete(&models.Post{}, id)
}
