package service

import (
	"baseProject/config"
	"baseProject/entity"
	"baseProject/model"
	"baseProject/repository"
	"crypto/md5"
	"fmt"
	"time"
)

type genderType int

const (
	Male genderType = iota + 1
	Female
)

func (g genderType) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}

type (
	IUserService interface {
		RegisterUser(req *entity.UserRegisterReq) *entity.UserResp
		LoadUserById(req *entity.UserLoaderReq) *entity.UserResp
		LoadUsers(req *entity.UsersLoaderReq) []entity.UserResp
	}
	UserService struct {
		conf     *config.Config
		userRepo repository.IUserRepository
	}
)

func NewUserService(conf *config.Config, userRepo repository.IUserRepository) *UserService {
	return &UserService{
		conf,
		userRepo,
	}
}

func (us *UserService) birthdayStr2Time(birthdayStr string) time.Time {
	dt, _ := time.Parse("2006-01-02", birthdayStr)
	return dt
}

func (us *UserService) encryptPassword(rawPassword string) string {
	b := []byte(rawPassword + us.conf.SecretKey)
	return fmt.Sprintf("%x", md5.Sum(b))
}

func (us *UserService) getAgeByBirthday(birthday time.Time) int {
	return time.Now().Year() - birthday.Year()
}

func (us *UserService) model2resp(userModel *model.UserModel) *entity.UserResp {
	var userResp entity.UserResp
	userResp.Id = int(userModel.ID)
	userResp.Username = userModel.Username
	userResp.Nickname = userModel.Nickname
	userResp.Gender = userModel.Gender
	userResp.GenderText = fmt.Sprintf("%v", genderType(userModel.Gender))
	userResp.Age = us.getAgeByBirthday(userModel.Birthday)
	return &userResp
}

func (us *UserService) RegisterUser(req *entity.UserRegisterReq) *entity.UserResp {
	userModel := us.userRepo.CreateUser(
		req.Username,
		us.encryptPassword(req.Password),
		req.Nickname,
		req.Gender,
		us.birthdayStr2Time(req.Birthday),
	)
	return us.model2resp(userModel)
}

func (us *UserService) LoadUserById(req *entity.UserLoaderReq) *entity.UserResp {
	userModel := us.userRepo.GetUserById(req.UserId)
	return us.model2resp(userModel)
}

func (us *UserService) LoadUsers(req *entity.UsersLoaderReq) []entity.UserResp {
	perPage := 20
	var usersResp []entity.UserResp
	for _, userModel := range us.userRepo.GetUsers(req.MaxUserId, perPage) {
		userResp := us.model2resp(&userModel)
		usersResp = append(usersResp, *userResp)
	}
	return usersResp
}