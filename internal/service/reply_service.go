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

func (r *ReplyService) CreateReply(reply *model.NewReply) (*model.Reply, error) {
	user, err := r.userRepository.FindOneByUuid(reply.User.Uuid)
	if err != nil {
		return nil, err
	}
	post, err := r.postRepository.FindOneByUuid(*reply.Post.Uuid)
	if err != nil {
		return nil, err
	}
	replyEntity := entity.CreateReply(user, post, reply)
	r.replyRepository.Create(replyEntity)
	return mapper.GetReplyModelFromEntity(replyEntity), nil
}

func (r *ReplyService) GetRepliesForPost(postUuid uuid.UUID) ([]*model.Reply, error) {
	post, err := r.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return nil, err
	}
	replies := r.replyRepository.FindRepliesForPost(post)
	return mapper.GetReplyModelsFromEntities(replies), nil
}
