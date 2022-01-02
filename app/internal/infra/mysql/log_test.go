package mysql_test

import (
	"ratri/internal/domain/model"
	"testing"
)

// TestStoreTaskTimeLog : [Deprecated] Logging task time.
func TestStoreTaskTimeLog(t *testing.T) {
	p := model.SerpViewingLogParamWithTime{
		UserId:      1,
		TaskId:      5,
		ConditionId: 5,
		TimeOnPage:  42,
	}

	if err := logDao.StoreTaskTimeLog(&p); err != nil {
		t.Fatal(err)
	}
}

func TestCumulateSerpViewingTime(t *testing.T) {
	p := model.SerpViewingLogParam{
		UserId:      42,
		TaskId:      5,
		ConditionId: 5,
	}

	if err := logDao.CumulateSerpViewingTime(&p); err != nil {
		t.Fatal(err)
	}
}

func TestCumulatePageViewingTime(t *testing.T) {
	p := model.PageViewingLogParam{
		UserId:      42,
		TaskId:      5,
		ConditionId: 5,
		PageId:      432,
	}

	if err := logDao.CumulatePageViewingTime(&p); err != nil {
		t.Fatal(err)
	}
}

func TestStoreSerpEventLog(t *testing.T) {
	p := model.SearchPageEventLogParam{
		User:        42,
		TaskId:      5,
		ConditionId: 5,
		Time:        142,
		Page:        2,
		Rank:        5,
		IsVisible:   true,
		Event:       "click",
	}
	if err := logDao.StoreSerpEventLog(&p); err != nil {
		t.Fatal(err)
	}
}

func TestStoreSearchSession(t *testing.T) {
	p := model.SearchSession{
		UserId:      42,
		TaskId:      5,
		ConditionId: 5,
	}

	if err := logDao.StoreSearchSeeion(&p); err != nil {
		t.Fatal(err)
	}
}
