package repositories

import (
	"post_blog/models"

	"gorm.io/gorm"
)

type PostRepo interface {
	Create(post models.BlogCreateUpdate) (uint, error)
	GetByID(id int) (*models.BlogPost, error)
	Update(id int, update models.BlogCreateUpdate) error
	Delete(id int) error
	List(page int, limit int) ([]models.BlogPost, int64, error)
	DeleteMany(ids []int) error
}
type postRepo struct {
	db *gorm.DB
}

func NewPostRepo(db *gorm.DB) PostRepo {
	return &postRepo{db}
}
func (r *postRepo) Create(post models.BlogCreateUpdate) (uint, error) {
	item := models.BlogPost{
		Title:   post.Title,
		Content: post.Content,
		Author:  post.Author,
	}
	result := r.db.Create(&item)
	return item.Id, result.Error
}

func (r *postRepo) GetByID(id int) (*models.BlogPost, error) {
	var item models.BlogPost
	if err := r.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *postRepo) Update(id int, update models.BlogCreateUpdate) error {
	return r.db.Table("blog_post_management").Where("id = ?", id).Updates(&update).Error
}

func (r *postRepo) Delete(id int) error {
	return r.db.Table("blog_post_management").Where("id = ?", id).Delete(nil).Debug().Error
}

func (r *postRepo) List(page int, limit int) ([]models.BlogPost, int64, error) {
	var items []models.BlogPost
	var total int64
	offset := (page - 1) * limit
	r.db.Model(&models.BlogPost{}).Count(&total)
	err := r.db.Order("id desc").Offset(offset).Limit(limit).Find(&items).Error
	return items, total, err
}

func (r *postRepo) DeleteMany(ids []int) error {
	return r.db.Table("blog_post_management").Where("id IN ?", ids).Delete(nil).Error
}
