package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
	"regexp"
	//"time"
	"fmt"
)


type Users struct {
	Id       int
	Username string
	Password string
	Salt     string
	Email    string
}

func (this *Users) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(new(Users))
}

func AddUser(username string, password string) (error) {
	o := orm.NewOrm()
	o.Using("default")
	user := new(Users)
	user.Username = username
	user.Salt = com.RandString(10)
	user.Password = com.Md5(password + user.Salt)
	o.Insert(user)
	return nil
}

func FindUser(username string) (Users, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := Users{Username: username}
	err := o.Read(&user, "username")

	return user, err
}

func ChangeUsername(oldUsername string, newUsername string) error {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("users").Filter("username", oldUsername).Update(orm.Params{
		"username": newUsername,
	})
	return err
}

func ChangeEmail(username string, email string) error {
	o := orm.NewOrm()
	o.Using("default")

	reg := regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`)
	result := reg.MatchString(email)
	if !result {
		return errors.New("not a email")
	}

	num, err := o.QueryTable("users").Filter("username", username).Update(orm.Params{"email": email})

	if nil != err {
		return err
	} else if 0 == num {
		return errors.New("not update")
	} else {
		return nil
	}
}

func SetPassword(username string, password string) error {
	o := orm.NewOrm()
	o.Using("default")
	salt := com.RandString(10)

	num, err := o.QueryTable("users").Filter("username", username).Update(orm.Params{
		"salt":     salt,
		"password": com.Md5(password + salt),
	})
	if 0 == num {
		return errors.New("item not exist")
	}

	return err
}
func ChangePassword(username string, oldPassword string, newPassword string) error {
	o := orm.NewOrm()
	o.Using("default")
	salt := com.RandString(10)

	user := Users{Username: username}
	err := o.Read(&user, "username")
	if nil != err {
		return err
	} else {
		if user.Password == com.Md5(oldPassword+user.Salt) {
			_, err := o.QueryTable("users").Filter("username", username).Update(orm.Params{
				"salt":     salt,
				"password": com.Md5(newPassword + salt),
			})
			return err
		} else {
			return errors.New("verification failed")
		}
	}
}

func ValidateUser(user Users,password string) error {
	u,_ :=FindUser(user.Username)
	if u.Username == "" {
		return errors.New("用户名或密码错误！")
	}
	fmt.Println(u.Password)
	fmt.Println(com.Md5(password + u.Salt))
	if u.Password != com.Md5(password + u.Salt) {
		return errors.New("密码错误！")
	}
	return nil
}


