package util

import (
	"github.com/danielmunro/otto-community-service/internal/model"
)

func CombinePosts(posts1 []*model.Post, posts2 []*model.Post) []*model.Post {
	p1len := len(posts1)
	p2len := len(posts2)
	allPostsCount := p1len + p2len
	var posts = make([]*model.Post, allPostsCount)
	if allPostsCount == 0 {
		return posts
	}
	p1 := 0
	p2 := 0
	for i := 0; i < allPostsCount; i++ {
		if p1 >= p1len {
			posts[i] = posts2[p2]
			p2++
			continue
		}
		if p2 >= p2len {
			posts[i] = posts1[p1]
			p1++
			continue
		}
		if posts1[p1].CreatedAt.Unix() > posts2[p2].CreatedAt.Unix() {
			posts[i] = posts1[p1]
			p1++
		} else {
			posts[i] = posts2[p2]
			p2++
		}
	}
	return posts
}
