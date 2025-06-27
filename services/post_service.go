package services

import (
	"post_blog/models"
	"post_blog/repositories"
)

type PostService interface {
	Create(post models.BlogCreateUpdate) (uint, error)
	GetByID(id int) (*models.BlogPost, error)
	Update(id int, update models.BlogCreateUpdate) error
	Delete(id int) error
	List(page int, limit int) ([]models.BlogPost, int64, error)
	DeleteMany(ids []int) error
}

type postService struct {
	repo repositories.PostRepo
}

func NewPostService(repo repositories.PostRepo) PostService {
	return &postService{repo}
}

func (s *postService) Create(post models.BlogCreateUpdate) (uint, error) {
	return s.repo.Create(post)
}

func (s *postService) GetByID(id int) (*models.BlogPost, error) {
	return s.repo.GetByID(id)
}

func (s *postService) Update(id int, update models.BlogCreateUpdate) error {
	return s.repo.Update(id, update)
}

func (s *postService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *postService) List(page int, limit int) ([]models.BlogPost, int64, error) {
	return s.repo.List(page, limit)
}

func (s *postService) DeleteMany(ids []int) error {
	return s.repo.DeleteMany(ids)
}
