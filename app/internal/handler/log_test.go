package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ratri/internal/domain/model"
	"ratri/internal/handler"
	mock "ratri/internal/mock/usecase"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
)

var (
	serpViewLogs = []struct {
		name      string
		in        model.SerpViewingLogParam
		want      int
		wantError bool
		err       error
	}{
		{
			"name",
			model.SerpViewingLogParam{UserId: 999, TaskId: 5, ConditionId: 3},
			201,
			false,
			nil,
		},
	}
	pageViewLogs = []struct {
		name      string
		in        model.PageViewingLogParam
		want      int
		wantError bool
		err       error
	}{
		{
			"name",
			model.PageViewingLogParam{UserId: 999, TaskId: 5, ConditionId: 3, PageId: 356},
			201,
			false,
			nil,
		},
	}

	serpEventLogs = []struct {
		name      string
		in        model.SearchPageEventLogParam
		want      int
		wantError bool
		err       error
	}{
		{
			"name",
			model.SearchPageEventLogParam{
				User:        42,
				TaskId:      5,
				ConditionId: 3,
				Time:        42,
				Page:        1,
				Rank:        3,
				IsVisible:   true,
				Event:       "click",
			},
			201,
			false,
			nil,
		},
	}
	searchSessionLogs = []struct {
		name      string
		in        model.SearchSession
		want      int
		wantError bool
		err       error
	}{
		{
			"name",
			model.SearchSession{UserId: 42, TaskId: 5, ConditionId: 3},
			201,
			false,
			nil,
		},
	}
)

func TestCreateTaskTimeLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockLog(ctrl)
	for _, tt := range serpViewLogs {
		t.Run(tt.name, func(t *testing.T) {
			mck.EXPECT().CumulateSerpViewingTime(&tt.in).Return(nil)

			b, err := json.Marshal(tt.in)
			if err != nil {
				t.Fatal("Failed to marshal test case: %w\n", err)
			}

			req := httptest.NewRequest(
				http.MethodPost,
				"/v1/logs/time",
				bytes.NewBuffer(b),
			)

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := handler.NewLogHandler(mck)

			err = h.CumulateSerpViewingTime(c)

			// Throw t.Fatal if unexpected error has occurred.
			if !tt.wantError && err != nil {
				t.Fatalf("Want no error, but got %#v", err)
			}

			// Throw t.Fatal if different error has occurred.
			if tt.wantError && !(err == tt.err) {
				t.Fatalf("Want %#v, but got %#v", tt.err, err)
			}

			// Throw t.Fatal if expected value is different from result.
			if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
				t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
			}
		})
	}
}

func TestCumulateSerpViewingTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mock := mock.NewMockLog(ctrl)

	for _, tt := range serpViewLogs {

		mock.EXPECT().CumulateSerpViewingTime(&tt.in).Return(nil)

		b, err := json.Marshal(tt.in)
		if err != nil {
			t.Fatal("Failed to marshal test case: %w\n", err)
		}

		req := httptest.NewRequest(
			http.MethodPost,
			"/v1/logs/click",
			bytes.NewBuffer(b),
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := handler.NewLogHandler(mock)

		err = h.CumulateSerpViewingTime(c)

		// Throw t.Fatal if unexpected error has occurred.
		if !tt.wantError && err != nil {
			t.Fatalf("Want no error, but got %#v", err)
		}

		// Throw t.Fatal if different error has occurred.
		if tt.wantError && !(err == tt.err) {
			t.Fatalf("Want %#v, but got %#v", tt.err, err)
		}

		// Throw t.Fatal if expected value is different from result.
		if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
			t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
		}
	}
}

func TestCumulatePageViewingTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mock := mock.NewMockLog(ctrl)

	for _, tt := range pageViewLogs {
		mock.EXPECT().CumulatePageViewingTime(&tt.in).Return(nil)
		b, err := json.Marshal(tt.in)
		if err != nil {
			t.Fatal("Failed to marshal test case: %w\n", err)
		}

		req := httptest.NewRequest(
			http.MethodPost,
			"/v1/logs/click",
			bytes.NewBuffer(b),
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := handler.NewLogHandler(mock)

		err = h.CumulatePageViewingTime(c)

		// Throw t.Fatal if unexpected error has occurred.
		if !tt.wantError && err != nil {
			t.Fatalf("Want no error, but got %#v", err)
		}

		// Throw t.Fatal if different error has occurred.
		if tt.wantError && !(err == tt.err) {
			t.Fatalf("Want %#v, but got %#v", tt.err, err)
		}

		// Throw t.Fatal if expected value is different from result.
		if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
			t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
		}
	}
}

func TestCreateSerpEventLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLog(ctrl)

	for _, tt := range serpEventLogs {
		mock.EXPECT().StoreSerpEventLog(&tt.in).Return(nil)

		b, err := json.Marshal(tt.in)
		if err != nil {
			t.Fatal("Failed to marshal test case: %w\n", err)
		}

		e := echo.New()
		req := httptest.NewRequest(
			http.MethodPost,
			"/v1/logs/click",
			bytes.NewBuffer(b),
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := handler.NewLogHandler(mock)

		err = h.CreateSerpEventLog(c)

		// Throw t.Fatal if unexpected error has occurred.
		if !tt.wantError && err != nil {
			t.Fatalf("Want no error, but got %#v", err)
		}

		// Throw t.Fatal if different error has occurred.
		if tt.wantError && !(err == tt.err) {
			t.Fatalf("Want %#v, but got %#v", tt.err, err)
		}

		// Throw t.Fatal if expected value is different from result.
		if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
			t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
		}
	}
}

func TestStoreSearchSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock.NewMockLog(ctrl)

	for _, tt := range searchSessionLogs {
		mock.EXPECT().StoreSearchSeeion(&tt.in).Return(nil)

		b, err := json.Marshal(tt.in)
		if err != nil {
			t.Fatal("Failed to marshal test case: %w\n", err)
		}

		e := echo.New()
		req := httptest.NewRequest(
			http.MethodPost,
			"/v1/logs/session",
			bytes.NewBuffer(b),
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := handler.NewLogHandler(mock)

		err = h.StoreSearchSeeion(c)

		// Throw t.Fatal if unexpected error has occurred.
		if !tt.wantError && err != nil {
			t.Fatalf("Want no error, but got %#v", err)
		}

		// Throw t.Fatal if different error has occurred.
		if tt.wantError && !(err == tt.err) {
			t.Fatalf("Want %#v, but got %#v", tt.err, err)
		}

		// Throw t.Fatal if expected value is different from result.
		if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
			t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
		}
	}
}
