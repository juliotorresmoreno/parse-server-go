package models

import (
	"time"

	"github.com/go-xorm/xorm"

	"github.com/juliotorresmoreno/parse-server/db"
	"github.com/juliotorresmoreno/unravel-server/helper"
)

//User d
type User struct {
	ID        uint   `xorm:"id bigint not null autoincr pk"     json:"id"`
	Nombres   string `xorm:"varchar(100) not null"              json:"nombres"    valid:"required,alphaSpaces"`
	Apellidos string `xorm:"varchar(100) not null"              json:"apellidos"  valid:"required,alphaSpaces"`
	FullName  string `xorm:"varchar(200) not null"              json:"fullname"   valid:"required,alphaSpaces"`
	Email     string `xorm:"varchar(200) not null"              json:"email"      valid:"required,email"`
	Usuario   string `xorm:"varchar(100) not null unique index" json:"usuario"    valid:"required,username"`
	Passwd    string `xorm:"varchar(100) not null"              json:"-"          valid:"required,password,encrypt"`
	Recovery  string `xorm:"varchar(100) not null index"        json:"-"`
	Code      string `xorm:"varchar(400) not null"              json:"-"`
	Role      string `xorm:"varchar(20) not null"               json:"-"`

	CreateAt time.Time `xorm:"created" json:"create_at"`
	UpdateAt time.Time `xorm:"updated" json:"update_at"`
}

func (el *User) Save() error {
	conn, err := db.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	if el.ID == 0 {
		el.Passwd = helper.Encript(el.Passwd)
		_, err := conn.Insert(el)
		return err
	}
	if el.Passwd != "" {
		el.Passwd = helper.Encript(el.Passwd)
	}
	_, err = conn.ID(el.ID).MustCols().Update(el)
	return err
}

func (el User) TableName() string {
	return "users"
}

func GetUsers() (*xorm.Session, error) {
	var session *xorm.Session
	conn, err := db.NewConnection()
	if err != nil {
		return session, nil
	}
	defer conn.Close()
	session = conn.Table(User{}.TableName()).
		Where("")
	return session, nil
}

type Users []User

func (u Users) Exists(id uint) bool {
	user := User{}
	conn, err := db.NewConnection()
	if err != nil {
		return false
	}
	defer conn.Close()
	exists, err := conn.Where("usuario = ?", id).Exist(&user)
	if err != nil {
		return false
	}
	return exists
}

func (u Users) FindById(id uint) (User, error) {
	user := User{}
	conn, err := db.NewConnection()
	if err != nil {
		return user, err
	}
	defer conn.Close()
	_, err = conn.Where("usuario = ?", id).Get(&user)
	return user, err
}

func (u *Users) FindByUserName(username string) (User, error) {
	user := User{}
	conn, err := db.NewConnection()
	if err != nil {
		return user, err
	}
	defer conn.Close()
	_, err = conn.Where("usuario = ?", username).Get(&user)
	return user, err
}
