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
	taskTests = []struct {
		name      string
		in        int
		want      interface{}
		wantError bool
		err       error
	}{
		{"Want 404 error#1", 1, 404, true, &echo.HTTPError{}},
		{"Want 404 error#2", 2, 404, true, &echo.HTTPError{}},
		{"Want 404 error#3", 3, 404, true, &echo.HTTPError{}},
		{"Want 404 error#4", 4, 404, true, &echo.HTTPError{}},
		{"Want no error#1", 5, 200, false, nil},
		{"Want no error#2", 6, 200, false, nil},
		{"Want no error#3", 7, 200, false, nil},
		{"Want no error#4", 8, 200, false, nil},
	}

	answerTests = []struct {
		name      string
		in        *model.Answer
		want      interface{}
		wantError bool
		err       error
	}{
		{"Want no error", &model.Answer{
			UserID:      42,
			TaskID:      5,
			ConditionID: 3,
			Answer:      "",
			Reason:      "",
		}, 201, false, nil},
	}
)

func TestFetchTaskInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockTaskUsecase(ctrl)
	for _, tt := range taskTests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.in > 4 && tt.in < 9 {
				mck.EXPECT().FetchTaskInfo(tt.in).Return(&model.Task{}, nil)
			} else {
				mck.EXPECT().FetchTaskInfo(tt.in).Return(&model.Task{}, model.ErrNoSuchData)
			}
			h := handler.NewTaskHandler(mck)

			req := httptest.NewRequest(
				http.MethodGet,
				"/v1/task/"+fmt.Sprintf("%v", tt.in),
				nil,
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Set path parameter explicitly
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%v", tt.in))

			err := h.FetchTaskInfo(c)

			// Throw t.Fatal if unexpected error has occurred.
			if !tt.wantError && err != nil {
				t.Fatalf("Want no error, but got %#v", err)
			}

			// Throw t.Fatal if different error has occurred.
			// if tt.wantError && !(err == tt.err) {
			// 	t.Fatalf("Want %#v, but got %#v", tt.err, err)
			// }

			if tt.wantError {
				t.Logf("%+v", err)
			}

			// Throw t.Fatal if expected value is different from result.
			if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
				t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
			}
		})
	}
}

func TestSubmitTaskAnswer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	mck := mock.NewMockTaskUsecase(ctrl)
	for _, tt := range answerTests {
		t.Run(tt.name, func(t *testing.T) {
			mck.EXPECT().CreateTaskAnswer(tt.in).Return(nil)
			h := handler.NewTaskHandler(mck)

			b, err := json.Marshal(tt.in)
			if err != nil {
				t.Fatal("Failed to marshal test case: %w\n", err)
			}

			req := httptest.NewRequest(
				http.MethodPost,
				"/v1/task/answer",
				bytes.NewBuffer(b),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err = h.SubmitTaskAnswer(c)

			// Throw t.Fatal if unexpected error has occurred.
			if !tt.wantError && err != nil {
				t.Fatalf("Want no error, but got %#v", err)
			}

			// Throw t.Fatal if different error has occurred.
			// if tt.wantError && !(err == tt.err) {
			// 	t.Fatalf("Want %#v, but got %#v", tt.err, err)
			// }
			if tt.wantError {
				t.Logf("%+v", err)
			}

			// Throw t.Fatal if expected value is different from result.
			if diff := cmp.Diff(tt.want, rec.Code); !tt.wantError && diff != "" {
				t.Fatalf("Want %d, but got %d\n%v", tt.want, rec.Code, diff)
			}
		})
	}
}
