package dao

import (
	"testing"
)

func TestUserDAOImpl_Save(t *testing.T) {
	userDAO := &UserDAOImpl{}

	_, err := InitDB()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	user := &UserEntity{
		Username: "david",
		Password: "123456",
		Email:    "cajan2@163.com",
	}

	err = userDAO.Save(user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("new User ID is %d", user.ID)

}

func TestUserDAOImpl_SelectByEmail(t *testing.T) {

	userDAO := &UserDAOImpl{}

	_, err := InitDB()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	user, err := userDAO.SelectByEmail("cajan2@163.com")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("result uesrname is %s", user.Username)

}
