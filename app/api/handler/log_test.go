package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
)

var (
	timeLog = `{
		"user": 999,
		"time": 999,
		"task": 5,
		"condision": 5
	}`
	clickLog = `{
		"user": 999,
		"task_id": 5,
		"conditionId": 5,
		"time": 999,
		"page": 999,
		"rank": 999,
		"visible": false
	}`
)

func TestCreateTaskTimeLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logDaoMock := dao.NewMockLog(ctrl)
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodPost, "/v1/users/log/time", strings.NewReader(timeLog))

	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: mockDB}

	// if assert.NoError(t, h.CreateTaskTimeLog(c)) {
	// 	if diff := cmp.Diff(rec.Code, http.StatusCreated); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }
}

func TestCreateSerpClickLog(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodPost, "/v1/users/log/click", strings.NewReader(clickLog))

	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: mockDB}

	// if assert.NoError(t, h.CreateTaskTimeLog(c)) {
	// 	if diff := cmp.Diff(rec.Code, http.StatusCreated); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }
}
