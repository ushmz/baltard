package mysql_test

var (
	userId int
	uid    string
)

// func TestCreate(t *testing.T) {
// 	if user, err := userDao.Create("", ""); err != nil {
// 		t.Fatal(err)
// 	} else {
// 		userId = user.Id
// 		uid = user.Uid
// 	}
// }
//
// func TestFindById(t *testing.T) {
// 	_, err := userDao.FindById(userId)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
//
// func TestFindByUid(t *testing.T) {
// 	_, err := userDao.FindByUid(uid)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
//
// func TestAddCompletionCode(t *testing.T) {
// 	err := userDao.AddCompletionCode(userId, 999999)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
//
// func TestGetCompletionCodeById(t *testing.T) {
// 	_, err := userDao.GetCompletionCodeById(userId)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
