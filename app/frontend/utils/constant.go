package utils

const ServiceName = "frontend"

type SessionUserIdKey string

const (
	UserIdKey   = SessionUserIdKey("user_id")
	UserRoleKey = SessionUserIdKey("user_role")
)
