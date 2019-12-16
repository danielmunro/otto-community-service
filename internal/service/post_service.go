package service

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/danielmunro/otto-community-service/internal/util"
	"github.com/google/uuid"
	"log"
)

type PostService struct {
	userRepository *repository.UserRepository
	postRepository *repository.PostRepository
}

func CreateDefaultPostService() *PostService {
	conn := db.CreateDefaultConnection()
	return CreatePostService(
		repository.CreatePostRepository(conn),
		repository.CreateUserRepository(conn))
}

func CreatePostService(postRepository *repository.PostRepository, userRepository *repository.UserRepository) *PostService {
	return &PostService{
		userRepository,
		postRepository,
	}
}

func (p *PostService) GetPost(postUuid uuid.UUID) (*model.Post, error) {
	post, err := p.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return nil, err
	}
	if post.User == nil {
		return nil, errors.New(constants.ErrorMessageUserNotFound)
	}
	return mapper.GetPostModelFromEntity(post), nil
}

func (p *PostService) CreatePost(newPost *model.NewPost) (*model.Post, error) {
	user, err := p.userRepository.FindOneByUuid(newPost.User.Uuid)
	if err != nil {
		return nil, err
	}
	post := entity.CreatePost(user, newPost)
	p.postRepository.Save(post)
	return mapper.GetPostModelFromEntity(post), nil
}

func (p *PostService) GetNewPosts(userUuid uuid.UUID) []*model.Post {
	posts, _ := p.GetPostsForUserFollows(userUuid)
	return util.CombinePosts(
		p.GetPostsForUser(userUuid),
		posts)
}

func (p *PostService) GetPostsForUser(userUuid uuid.UUID) []*model.Post {
	user, err := p.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		log.Fatal(err) // return an error!
	}
	return mapper.GetPostModelsFromEntities(p.postRepository.FindByUser(user))
}

func (p *PostService) GetPostsForUserFollows(userUuid uuid.UUID) ([]*model.Post, error) {
	_, err := p.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		return nil, err
	}
	posts := p.postRepository.FindByUserFollows(userUuid)
	return mapper.GetPostModelsFromEntities(posts), nil
}
