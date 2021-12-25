package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ratri/internal/domain/model"
	mock "ratri/internal/mock/usecase"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	serpViewLog = model.SerpViewingLogParam{
		UserId:      999,
		TaskId:      5,
		ConditionId: 3,
	}
	serpViewLogBlob = `{
		"user":      999,
		"task":      5,
		"condition": 3
	}`
	serpEventLog = model.SearchPageEventLogParam{
		User:        42,
		TaskId:      5,
		ConditionId: 3,
		Time:        42,
		Page:        1,
		Rank:        3,
		IsVisible:   true,
		Event:       "click",
	}
	serpEventLogBlob = `{
		"user":      42,
		"task":      5,
		"condition": 3,
		"time":      42,
		"page":      1,
		"rank":      3,
		"visible":   true,
		"event":     "click"
	}`
	searchSessionLog = model.SearchSession{
		UserId:      42,
		TaskId:      5,
		ConditionId: 3,
	}
	searchSessionLogBlob = `{
		"user":      42,
		"task":      5,
		"condition": 3
	}`
)

func TestCreateTaskTimeLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLog(ctrl)
	mock.EXPECT().CumulateSerpViewingTime(&serpViewLog).Return(nil)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/logs/time",
		strings.NewReader(serpViewLogBlob),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewLogHandler(mock)

	if assert.NoError(t, h.CumulateSerpViewingTime(c)) {
		if diff := cmp.Diff(rec.Code, http.StatusCreated); diff != "" {
			t.Fatal("Status code does not match.\n" + diff)
		}
	}
}

func TestCreateSerpClickLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLog(ctrl)
	mock.EXPECT().StoreSerpEventLog(&serpEventLog).Return(nil)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/logs/click",
		strings.NewReader(serpEventLogBlob),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewLogHandler(mock)

	if assert.NoError(t, h.CreateSerpEventLog(c)) {
		if diff := cmp.Diff(rec.Code, http.StatusCreated); diff != "" {
			t.Fatal("Status code does not match.\n" + diff)
		}
	}
}

func TestStoreSearchSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLog(ctrl)
	mock.EXPECT().StoreSearchSeeion(&searchSessionLog).Return(nil)

	e := echo.New()
	req := httptest.NewRequest(
		http.MethodPost,
		"/v1/logs/session",
		strings.NewReader(searchSessionLogBlob),
	)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewLogHandler(mock)

	if assert.NoError(t, h.StoreSearchSeeion(c)) {
		if diff := cmp.Diff(rec.Code, http.StatusCreated); diff != "" {
			t.Fatal("Status code does not match.\n" + diff)
		}
	}
}
