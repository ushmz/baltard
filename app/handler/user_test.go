package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"ratri/domain/model"
	"ratri/handler"
	mock "ratri/mock/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
)

var (
	task = model.TaskInfo{
		GroupId:     3,
		ConditionId: 2,
		TaskIds:     []int{5, 7},
	}
	userTests = []struct {
		name      string
		in        model.UserParam
		want      interface{}
		wantError bool
		err       error
	}{
		{"Want no error", model.UserParam{Uid: "test42"}, 200, false, nil},
	}

	completionTest = []struct {
		name      string
		in        interface{}
		want      interface{}
		wantError bool
		err       error
	}{
		{"Want no error", 999, 200, false, nil},
	}
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockUserUsecase(ctrl)
	for _, tt := range userTests {
		t.Run(tt.name, func(t *testing.T) {
			mck.EXPECT().FindByUid(tt.in.Uid).Return(model.User{}, false, nil)
			mck.EXPECT().CreateUser(tt.in.Uid).Return(model.User{}, nil)
			mck.EXPECT().AllocateTask().Return(task, nil)

			h := handler.NewUserHandler(mck)

			b, err := json.Marshal(tt.in)
			if err != nil {
				t.Fatal("Failed to marshal test case: %w\n", err)
			}
			req := httptest.NewRequest(
				http.MethodPost,
				"/users",
				bytes.NewBuffer(b),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err = h.CreateUser(c)

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

func TestGetCompletionCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockUserUsecase(ctrl)
	for _, tt := range completionTest {
		t.Run(tt.name, func(t *testing.T) {
			mck.EXPECT().GetCompletionCode(tt.in).Return(42424, nil)
			h := handler.NewUserHandler(mck)

			req := httptest.NewRequest(
				http.MethodGet,
				"/v1/users/code"+fmt.Sprintf("%v", tt.in),
				nil,
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set path parameter explicitly
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%v", tt.in))

			err := h.GetCompletionCode(c)

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
