package service

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/danielmunro/otto-community-service/internal/mapper"
	"github.com/danielmunro/otto-community-service/internal/model"
	"github.com/danielmunro/otto-community-service/internal/repository"
	"github.com/google/uuid"
)

type ReplyService struct {
	userRepository  *repository.UserRepository
	postRepository  *repository.PostRepository
	replyRepository *repository.ReplyRepository
}

func CreateDefaultReplyService() *ReplyService {
	conn := db.CreateDefaultConnection()
	return CreateReplyService(
		repository.CreateReplyRepository(conn),
		repository.CreatePostRepository(conn),
		repository.CreateUserRepository(conn))
}

func CreateReplyService(
	replyRepository *repository.ReplyRepository,
	postRepository *repository.PostRepository,
	userRepository *repository.UserRepository) *ReplyService {
	return &ReplyService{
		userRepository,
		postRepository,
		replyRepository,
	}
}

func (r *ReplyService) CreateReply(reply *model.NewReply) (*model.Post, error) {
	user, err := r.userRepository.FindOneByUuid(uuid.MustParse(reply.User.Uuid))
	if err != nil {
		return nil, err
	}
	post, err := r.postRepository.FindOneByUuid(uuid.MustParse(reply.Post.Uuid))
	if err != nil {
		return nil, err
	}
	replyEntity := entity.CreateReply(user, post, reply)
	r.replyRepository.Create(replyEntity)
	post.Replies += 1
	r.postRepository.Save(post)
	return mapper.GetPostModelFromEntity(replyEntity), nil
}

func (r *ReplyService) GetRepliesForPost(postUuid uuid.UUID) ([]*model.Post, error) {
	post, err := r.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return nil, err
	}
	replies := r.replyRepository.FindRepliesForPost(post)
	return mapper.GetPostModelsFromEntities(replies), nil
}
