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
	"log"
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

func (p *PostService) GetPostsForUser(username string, viewerUuid *uuid.UUID, limit int) ([]*model.Post, error) {
	user, err := p.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}
	postEntities := p.postRepository.FindByUser(user, limit)
	var fullListModels []*model.Post
	if viewerUuid != nil {
		viewer, _ := p.userRepository.FindOneByUuid(*viewerUuid)
		fullListModels = p.populateModelsWithLikes(postEntities, viewer)
	} else {
		fullListModels = mapper.GetPostModelsFromEntities(postEntities)
	}
	return fullListModels, nil
}

func (p *PostService) GetPostsForUserFollows(username string, viewerUserUuid uuid.UUID, limit int) ([]*model.Post, error) {
	_, err := p.userRepository.FindOneByUsername(username)
	if err != nil {
		return nil, err
	}
	viewer, err := p.userRepository.FindOneByUuid(viewerUserUuid)
	if err != nil {
		return nil, err
	}
	posts := p.postRepository.FindByUserFollows(username, limit)
	postModels := p.populateModelsWithLikes(posts, viewer)
	return postModels, nil
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
	var user *entity.User
	if username != nil {
		user, _ = p.userRepository.FindOneByUsername(*username)
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
	postsWithShares := p.countPostsWithShares(fullList)
	log.Print("posts with shares, count :: ", postsWithShares)
	if user != nil {
		return p.populateModelsWithLikes(fullList, user), nil
	}
	return mapper.GetPostModelsFromEntities(fullList), nil
}

func (p *PostService) countPostsWithShares(posts []*entity.Post) int {
	amount := 0
	for _, post := range posts {
		if post.SharePostID != 0 {
			amount += 1
		}
	}
	return amount
}

func (p *PostService) populateModelsWithLikes(posts []*entity.Post, viewer *entity.User) []*model.Post {
	postIds := p.getPostIDs(posts)
	postLikes := p.likeRepository.FindLikesForPosts(postIds, viewer)
	likedPosts := make(map[uint]bool)
	for _, postLike := range postLikes {
		likedPosts[postLike.PostID] = true
	}
	fullListModels := mapper.GetPostModelsFromEntities(posts)
	for i, item := range posts {
		if likedPosts[item.ID] {
			fullListModels[i].SelfLiked = true
		}
	}
	return fullListModels
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
