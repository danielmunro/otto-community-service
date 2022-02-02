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
	imageRepository  *repository.ImageRepository
	likeRepository   *repository.LikeRepository
}

func CreateDefaultPostService() *PostService {
	conn := db.CreateDefaultConnection()
	return CreatePostService(
		repository.CreatePostRepository(conn),
		repository.CreateUserRepository(conn),
		repository.CreateFollowRepository(conn),
		repository.CreateImageRepository(conn),
		repository.CreateLikeRepository(conn),
	)
}

func CreatePostService(
	postRepository *repository.PostRepository,
	userRepository *repository.UserRepository,
	followRepository *repository.FollowRepository,
	imageRepository *repository.ImageRepository,
	likeRepository *repository.LikeRepository) *PostService {
	return &PostService{
		userRepository,
		postRepository,
		followRepository,
		imageRepository,
		likeRepository,
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
	user, err := p.userRepository.FindOneByUuid(uuid.MustParse(newPost.User.Uuid))
	if err != nil {
		return nil, err
	}
	post := entity.CreatePost(user, newPost)
	p.postRepository.Save(post)
	var imageEntities []*entity.Image
	for _, newImage := range newPost.Images {
		imageEntity := entity.CreateImage(user, post, &newImage)
		p.imageRepository.Create(imageEntity)
		imageEntities = append(imageEntities, imageEntity)
	}
	post.Images = imageEntities
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

func (p *PostService) GetNewPosts(username string, limit int) []*model.Post {
	user, _ := p.userRepository.FindOneByUsername(username)
	followPosts := p.postRepository.FindByUserFollows(username, limit)
	userPosts := p.postRepository.FindByUser(user, limit)
	allPosts := append(followPosts, userPosts...)
	return mapper.GetPostModelsFromEntities(removeDuplicatePosts(allPosts))
}

func (p *PostService) GetPostsForUser(username string, limit int) ([]*model.Post, error) {
	user, err := p.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}
	return mapper.GetPostModelsFromEntities(p.postRepository.FindByUser(user, limit)), nil
}

func (p *PostService) GetPostsForUserFollows(username string, limit int) ([]*model.Post, error) {
	_, err := p.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}
	posts := p.postRepository.FindByUserFollows(username, limit)
	return mapper.GetPostModelsFromEntities(posts), nil
}

func (p *PostService) GetAllPosts(limit int) []*model.Post {
	posts := p.postRepository.FindAll(limit)
	return mapper.GetPostModelsFromEntities(posts)
}

func (p *PostService) GetPosts(username *string, limit int) ([]*model.Post, error) {
	var selfPosts []*entity.Post
	var followingPosts []*entity.Post
	var publicPosts []*entity.Post
	remaining := constants.UserPostsDefaultPageSize
	if username != nil {
		user, _ := p.userRepository.FindOneByUsername(*username)
		selfPosts = p.postRepository.FindByUser(user, limit)
		remaining -= len(selfPosts)
	}
	if remaining > 0 && username != nil {
		followingPosts = p.postRepository.FindByUserFollows(*username, remaining)
		remaining -= len(followingPosts)
	}
	if remaining > 0 {
		publicPosts = p.postRepository.FindAll(remaining)
	}
	allPosts := append(selfPosts, followingPosts...)
	allPosts = append(allPosts, publicPosts...)
	sort.SliceStable(allPosts, func(i, j int) bool {
		return allPosts[i].CreatedAt.After(allPosts[j].CreatedAt)
	})
	fullList := removeDuplicatePosts(allPosts)
	postIds := p.getPostIDs(fullList)
	postLikes := p.likeRepository.FindLikesForPosts(postIds)
	likedPosts := make(map[uint]bool)
	for _, postLike := range postLikes {
		likedPosts[postLike.PostID] = true
	}
	fullListModels := mapper.GetPostModelsFromEntities(fullList)
	for i, item := range fullList {
		if likedPosts[item.ID] {
			fullListModels[i].SelfLiked = true
		}
	}
	return fullListModels, nil
}

func (p *PostService) getPostIDs(posts []*entity.Post) []uint {
	postIDs := make([]uint, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}
	return postIDs
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

func removeDuplicatePosts(posts []*entity.Post) []*entity.Post {
	var dedup []*entity.Post
	allKeys := make(map[uint]bool)
	for _, item := range posts {
		if value := allKeys[item.ID]; !value {
			allKeys[item.ID] = true
			dedup = append(dedup, item)
		}
	}
	return dedup
}
