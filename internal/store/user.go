package store

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	UserNotValid     = errors.New("Incorrect parametrs!")
	UserAlreadyExist = errors.New("user already exist!")
	AuthFiled        = errors.New("auth failed")
)

// User ...
type User struct {
	gorm.Model
	FirstName string `json:"firstname";gorm:"not null"`
	LastName  string `json:"lastname";gorm:"not null;`
	Email     string `json:"email";gorm:"unique,not null"`
	Phone     string `json:"phone""`
	Password  string `json:"password,omitempty"`
	Admin     string `json:"admin,omitempty";gorm:"default:false"`
	Location  string
	Position  string
	Avatart   string
}

/** CRUD Methods **/
func CreateOrUpdate(u *User) error {

	if u.Valid() != true {
		return UserNotValid
	}

	if IsUserExist(u) {
		return UserAlreadyExist
	}
	u.EncodePass()
	u.Email = NormalizeEmail(u.Email)

	if u.ID > 0 {
		return x.Create(u).Error
	}

	return x.Save(u).Error
}

func FindUser(page int, limit int) []*User {
	var users []*User
	offset := page - 1
	if err := x.Offset(offset * limit).Limit(limit).Find(&users).Error; err != nil {
		// Пишем ошибку в лог
		return nil
	}
	return users
}

// func UpdateUser(u *User) error {
// 	u.EncodePass()
// 	return x.Save(u).Error
// }

func DeleteUser(u *User) error {
	return x.Unscoped().Delete(&User{}, "email LIKE ?", u.Email).Error
}

func DeleteUserByID(id uint) error {
	if id > 0 {
		return x.Unscoped().Delete(&User{}, "id = ?", id).Error
	}
	return errors.New("id is not valid")
}

func FindByEmail(email string) (*User, error) {
	u := new(User)
	if err := x.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func FindByID(id uint) *User {
	u := new(User)
	x.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&u)
	return u
}

func (u *User) TimeCreateAt() string {
	return u.CreatedAt.Format("02.01.2006 15:04")
}
func (u *User) TimeUpdateAt() string {
	return u.UpdatedAt.Format("02.01.2006 15:04")
}

// Valid ...
func (u *User) Valid() bool {
	// TODO Сделать номальную валидацию
	return strings.Contains(u.Email, "@") && len(u.FirstName) > 0 && len(u.LastName) > 0
}

func IsUserExist(u *User) bool {
	return x.Debug().Where("id = ? OR email = ?", u.ID, strings.ToLower(u.Email)).Find(&User{}).RowsAffected > 0
}

// Работа с паролями
func (u *User) EncodePass() {
	if password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0); err == nil {
		u.Password = fmt.Sprintf("%s", password)
	}
	// Логировать ошибку
}
func (u *User) ValidatePass(pwd string) bool {
	if len(pwd) == 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	return err == nil
}

func (u *User) DisplayName() string {
	if u.ID == 0 {
		return "Новый пользователь"
	}
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPhoneNumber() string {
	return u.Phone
}

func NormalizeEmail(email string) string {
	return strings.ToLower(email)
}

func UserCount() int64 {
	var count int64
	fmt.Printf("DB: %+v\n", GetEnginie())
	fmt.Println(x.Debug().Where("id > 0").Find(&User{}).RowsAffected)
	x.Debug().Model(&User{}).Count(&count)
	return count
}
