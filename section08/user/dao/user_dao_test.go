package dao

import (
	"strings"
	"testing"
)

func TestUserDAOImpl_Save(t *testing.T) {
	userDAO := &UserDAOImpl{}

	//_, err := InitDB()
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}
	user := &UserEntity{
		Username: "david",
		Password: "123456",
		Email:    "cajan2@163.com",
	}

	err := userDAO.Save(user)
	if err != nil {
		if !strings.Contains(err.Error(), "Duplicate entry") {
			t.Error(err)
			t.FailNow()
		} else {
			t.Logf("%s", err.Error())
			return
		}
	}
	t.Logf("new User ID is %d", user.ID)

}

func TestUserDAOImpl_SelectByEmail(t *testing.T) {

	userDAO := &UserDAOImpl{}

	//_, err := InitDB()
	//if err != nil {
	//	t.Error(err)
	//	t.FailNow()
	//}

	user, err := userDAO.SelectByEmail("cajan2@163.com")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("result uesrname is %s", user.Username)

}
