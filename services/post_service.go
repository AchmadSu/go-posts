package services

import (
	"go-posts/dto"
	"go-posts/errs"
	"go-posts/models"
	"go-posts/repository"
	"go-posts/utils"
	"net/http"

	"gorm.io/gorm"
)

type PostService interface {
	GetPosts(page, pageSize int) ([]dto.PublicPost, int64, error)
	GetPostByID(id int) (dto.PublicPost, error)
	CreatePost(post dto.PublicPost) (models.Post, error)
	UpdatePost(post dto.PublicPost, id int) (models.Post, error)
	DeletePost(id int) error
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) GetPosts(page, pageSize int) ([]dto.PublicPost, int64, error) {
	offset := (page - 1) * pageSize
	posts, total, err := s.repo.GetAll(offset, pageSize)
	if err != nil {
		return []dto.PublicPost{}, 0, errs.New("post not found", http.StatusNotFound)
	}
	publicPosts := utils.ToPublicPosts(posts)
	return publicPosts, total, err
}

func (s *postService) GetPostByID(id int) (dto.PublicPost, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return dto.PublicPost{}, errs.New("post not found", http.StatusNotFound)
	}
	publicPosts := utils.ToPublicPost(post)
	return publicPosts, err
}

func (s *postService) CreatePost(post dto.PublicPost) (models.Post, error) {
	err := s.ValidatePost(post, 0)
	if err != nil {
		return models.Post{}, err
	}
	return s.repo.Create(post)
}

func (s *postService) UpdatePost(post dto.PublicPost, id int) (models.Post, error) {
	err := s.ValidatePost(post, id)
	if err != nil {
		return models.Post{}, err
	}
	return s.repo.Update(id, &post)
}

func (s *postService) DeletePost(id int) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errs.New("post not found", http.StatusNotFound)
	}

	result := s.repo.Delete(id)
	if result.Error != nil {
		return errs.New("Failed to delete post", http.StatusInternalServerError)
	}

	return nil
}

func (s *postService) ValidatePost(post dto.PublicPost, id int) error {

	if id != 0 {
		_, err := s.repo.FindByID(id)

		if err != nil {
			return errs.New("Post not found", http.StatusNotFound)
		}
	}

	if id == 0 && post.Title == "" {
		return errs.New("Title is required", http.StatusNotAcceptable)
	}

	if post.Title != "" {
		if len(post.Title) < 20 || len(post.Title) > 200 {
			return errs.New("Title minimum 20 characters and maximum 200 characters", http.StatusNotAcceptable)
		}

		postModel, existTitle := s.repo.FindByTitle(post.Title)
		if existTitle.Error != nil && existTitle.Error != gorm.ErrRecordNotFound {
			return existTitle.Error
		}

		if (existTitle.RowsAffected > 0 && id == 0) || (existTitle.RowsAffected > 0 && id != int(postModel.ID)) {
			return errs.New("Title already exists. Please try another title!", http.StatusNotAcceptable)
		}
	}

	if id == 0 && post.Category == "" {
		return errs.New("Category is required", http.StatusNotAcceptable)
	}

	if post.Category != "" {
		if len(post.Category) < 3 || len(post.Category) > 100 {
			return errs.New("Category minimum 3 characters and maximum 100 characters", http.StatusNotAcceptable)
		}
	}

	if id == 0 && post.Content == "" {
		return errs.New("Content is required", http.StatusNotAcceptable)
	}

	if post.Content != "" {
		if len(post.Content) < 200 {
			return errs.New("Content minimum 200 characters", http.StatusNotAcceptable)
		}
	}

	validStatus := map[string]bool{
		"publish": true,
		"draft":   true,
		"thrash":  true,
	}

	if id == 0 && post.Status == "" {
		return errs.New("Status is required", http.StatusNotAcceptable)
	}

	if post.Status != "" {
		if !validStatus[post.Status] {
			return errs.New("Status must be publish, draft, or thrash", http.StatusNotAcceptable)
		}
	}

	return nil
}
