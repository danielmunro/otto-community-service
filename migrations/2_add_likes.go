package main

import (
	"github.com/danielmunro/otto-community-service/internal/db"
	"github.com/danielmunro/otto-community-service/internal/entity"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	conn := db.CreateDefaultConnection()
	conn.AutoMigrate(
		&entity.PostLike{},
		&entity.ReplyLike{},
	)
	conn.Model(&entity.PostLike{}).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")
	conn.Model(&entity.ReplyLike{}).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("reply_id", "replies(id)", "RESTRICT", "RESTRICT")
}
