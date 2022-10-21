package helper

import "net/mail"

const (
	EmailNotValid       = "email not valid"
	EmailUsed           = "email already used"
	UsernameUsed        = "username already used"
	AgeMinimum          = "too young to register"
	UserNotFound        = "user not found"
	PhotoNotFound       = "photo not found"
	CommentNotFound     = "comment not found"
	SocialMediaNotFound = "social media not found"
	PasswordInvalid     = "invalid password"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
