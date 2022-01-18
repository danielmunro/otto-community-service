package service

import (
	"errors"
	"github.com/danielmunro/otto-community-service/internal/constants"
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/google/uuid"
	"sort"
	"time"
)

type PostService struct {
	userRepository   *repository.UserRepository
	postRepository   *repository.PostRepository
	followRepository *repository.FollowRepository
}

func CreateDefaultPostService() *PostService {
	conn := db.CreateDefaultConnection()
	return CreatePostService(
		repository.CreatePostRepository(conn),
		repository.CreateUserRepository(conn),
		repository.CreateFollowRepository(conn))
}

func CreatePostService(
	postRepository *repository.PostRepository,
	userRepository *repository.UserRepository,
	followRepository *repository.FollowRepository) *PostService {
	return &PostService{
		userRepository,
		postRepository,
		followRepository,
	}
}

func (p *PostService) GetPost(viewerUuid *uuid.UUID, postUuid uuid.UUID) (*model.Post, error) {
	post, err := p.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return nil, err
	}
	if post.User == nil {
		return nil, errors.New(constants.ErrorMessageUserNotFound)
	}
	if !p.canSee(viewerUuid, post) {
		return nil, errors.New("not accessible")
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

func (p *PostService) DeletePost(postUuid uuid.UUID, userUuid uuid.UUID) error {
	post, err := p.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return err
	}
	if *post.User.Uuid != userUuid {
		return errors.New("access denied")
	}
	now := time.Now()
	post.DeletedAt = &now
	p.postRepository.Save(post)
	return nil
}

func (p *PostService) GetNewPosts(userUuid uuid.UUID, limit int) []*model.Post {
	followPosts, _ := p.GetPostsForUserFollows(userUuid, limit)
	userPosts, _ := p.GetPostsForUser(userUuid, limit)
	allPosts := append(followPosts, userPosts...)
	return removeDuplicatePosts(allPosts)
}

func (p *PostService) GetPostsForUser(userUuid uuid.UUID, limit int) ([]*model.Post, error) {
	user, err := p.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		return nil, err
	}
	return mapper.GetPostModelsFromEntities(p.postRepository.FindByUser(user, limit)), nil
}

func (p *PostService) GetPostsForUserFollows(userUuid uuid.UUID, limit int) ([]*model.Post, error) {
	_, err := p.userRepository.FindOneByUuid(userUuid.String())
	if err != nil {
		return nil, err
	}
	posts := p.postRepository.FindByUserFollows(userUuid, limit)
	return mapper.GetPostModelsFromEntities(posts), nil
}

func (p *PostService) GetAllPosts(limit int) []*model.Post {
	posts := p.postRepository.FindAll(limit)
	return mapper.GetPostModelsFromEntities(posts)
}

func (p *PostService) GetPosts(userUuid uuid.UUID, limit int) ([]*model.Post, error) {
	selfPosts, _ := p.GetPostsForUser(userUuid, limit)
	remaining := limit - len(selfPosts)
	var allPosts []*model.Post
	if remaining > 0 {
		friendsPosts, err := p.GetPostsForUserFollows(userUuid, remaining)
		if err != nil {
			return nil, err
		}
		allPosts = append(selfPosts, friendsPosts...)
		remaining -= len(friendsPosts)
		if remaining > 0 {
			otherPosts := p.GetAllPosts(remaining)
			allPosts = append(allPosts, otherPosts...)
		}
	} else {
		allPosts = selfPosts
	}
	sort.SliceStable(allPosts, func(i, j int) bool {
		return allPosts[i].CreatedAt.After(allPosts[j].CreatedAt)
	})
	return removeDuplicatePosts(allPosts), nil
}

func (p *PostService) canSee(viewerUuid *uuid.UUID, post *entity.Post) bool {
	if post.Visibility == model.PUBLIC {
		return true
	}
	if viewerUuid == nil {
		return false
	}
	if post.Visibility == model.PRIVATE && viewerUuid != post.User.Uuid {
		return false
	}
	follow := p.followRepository.FindByUserAndFollowing(*post.User.Uuid, *viewerUuid)
	if follow == nil {
		return false
	}

	return true
}

func removeDuplicatePosts(posts []*model.Post) []*model.Post {
	var dedup []*model.Post
	allKeys := make(map[string]bool)
	for _, item := range posts {
		if value := allKeys[item.Uuid.String()]; !value {
			allKeys[item.Uuid.String()] = true
			dedup = append(dedup, item)
		}
	}
	return dedup
}
