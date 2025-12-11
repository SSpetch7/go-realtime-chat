package service

import (
	"context"
	m "realtime_chat_server/internal/model"
	"realtime_chat_server/internal/repository"
	"realtime_chat_server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type userService struct {
	userRepo repository.UserRepository
	timeout  time.Duration
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo, time.Duration(2) * time.Second}
}

func (s userService) Register(c context.Context, req *m.RegisterReq) (*m.RegisterRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)

	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	newUser := &m.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	res, err := s.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	user := m.RegisterRes{
		ID:       strconv.Itoa(int(res.ID)),
		Username: res.Username,
		Email:    res.Email,
	}

	return &user, nil
}

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s userService) Login(c context.Context, req *m.LoginReq) (*m.LoginRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)

	defer cancel()

	u, err := s.userRepo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return &m.LoginRes{}, err
	}

	err = util.VerifyPassword(req.Password, u.Password)

	if err != nil {
		return &m.LoginRes{}, err
	}

	jwtSecretKey := viper.GetString("env.jwtSecretKey")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
	})

	accessToken, err := token.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return &m.LoginRes{}, err
	}

	res := &m.LoginRes{
		AccessToken: accessToken,
		ID:          strconv.Itoa(int(u.ID)),
		Username:    u.Username,
	}

	return res, nil

}
