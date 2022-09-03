package model

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mcholismalik/boilerplate-golang/internal/abstraction"
	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/date"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	Name         string `json:"name" validate:"required" gorm:"size:50;not null"`
	Email        string `json:"email" validate:"required,email" gorm:"index:idx_user_email;unique;size:150;not null"`
	PasswordHash string `json:"-"`
	Password     string `json:"password" validate:"required" gorm:"-"`
}

type UserModel struct {
	// abstraction
	abstraction.Entity

	// entity
	UserEntity

	// context
	Context abstraction.Context `json:"-" gorm:"-"`
}

func (UserModel) TableName() string {
	return "users"
}

func (m *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	m.HashPassword()
	m.Password = ""
	return
}

func (m *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}

func (m *UserModel) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserModel) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    m.ID,
		"email": m.Email,
		"name":  m.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
