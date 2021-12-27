package handler_test

import (
	"testing"
)

var (
	userData = `{
		"uid": "test_user"
	}`
	// How do I get this value?
	// or, combine `TestCreateUser` and `TestGetCompletionCode` ?
	// If so, how do I get user ID from response body in *bytes.Buffer ?
	userId = "999"
)

func TestCreateUser(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userData))
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: database.New()}

	// if assert.NoError(t, h.CreateUser(c)) {
	// 	// Response body contain random string value
	// 	// Is there any way to test random value?
	// 	if diff := cmp.Diff(rec.Code, http.StatusOK); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }
}

func TestGetCompletionCode(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodPost, "/users/code/"+userId, nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: database.New()}

	// if assert.NoError(t, h.CreateUser(c)) {
	// 	// Response body contain random int value
	// 	// Is there any way to test random value?
	// 	if diff := cmp.Diff(rec.Code, http.StatusOK); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }

}
