package handler_test

import (
	"testing"
)

/*
 * Functions in `handler/serp.go` have huge response body,
 * so, is it better to read ideal request body from external text file?
 */

func TestFetchSerpWithDistributionByID(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/v1/serp/5/pct", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: database.New()}

	// if assert.NoError(t, h.FetchSerpWithDistributionByID(c)) {
	// 	if diff := cmp.Diff(rec.Code, http.StatusOK); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }
}

func TestFetchSerpWithIconByID(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/v1/serp/5/icon", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: database.New()}

	// if assert.NoError(t, h.FetchSerpWithDistributionByID(c)) {
	// 	if diff := cmp.Diff(rec.Code, http.StatusOK); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }
}

func TestFetchSerpByID(t *testing.T) {
	// e := echo.New()
	// req := httptest.NewRequest(http.MethodGet, "/v1/serp/5", nil)
	// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// h := &Handler{DB: database.New()}

	// if assert.NoError(t, h.FetchSerpWithDistributionByID(c)) {
	// 	if diff := cmp.Diff(rec.Code, http.StatusOK); diff != "" {
	// 		t.Errorf("Status code does not match.\n%v", diff)
	// 	}
	// }
}
