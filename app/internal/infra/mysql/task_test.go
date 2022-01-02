package mysql_test

import (
	"ratri/internal/domain/model"
	"testing"
)

var (
	testTaskInfo = []struct {
		name      string
		in        int
		want      interface{}
		wantError bool
		err       error
	}{

		{"Want no data error#1", 1, nil, true, model.NoSuchDataError{}},
		{"Want no data error#2", 2, nil, true, model.NoSuchDataError{}},
		{"Want no data error#3", 3, nil, true, model.NoSuchDataError{}},
		{"Want no data error#4", 4, nil, true, model.NoSuchDataError{}},
		{"Want no error#5", 5, &model.Task{}, false, nil},
		{"Want no error#6", 6, &model.Task{}, false, nil},
		{"Want no error#7", 7, &model.Task{}, false, nil},
		{"Want no error#8", 8, &model.Task{}, false, nil},
	}
)

func TestFetchTaskInfo(t *testing.T) {
	for _, tt := range testTaskInfo {
		t.Run(tt.name, func(t *testing.T) {
			task, err := taskDao.FetchTaskInfo(tt.in)

			if err != nil {
				if !tt.wantError {
					t.Fatal("Want no error, but got error: ", err)
				}
				if tt.wantError && err != tt.err {
					t.Fatalf("Want %#v, but got %#v", tt.err, err)
				}
			}

			// if reflect.DeepEqual(tt.want, *task) {
			// 	t.Fatalf("Want %#v, but got %#v:\n %#v", tt.want, task, err)
			// }

			if tt.want == task {
				t.Fatalf("Want %#v, but got %#v:\n %#v", tt.want, task, err)
			}

		})
	}
}

func TestUpdateTaskAllocation(t *testing.T) {
	if _, err := taskDao.UpdateTaskAllocation(); err != nil {
		t.Fatal(err)
	}
}

func TestGetTaskIdsByGroupId(t *testing.T) {
	// 1 - 6
	if _, err := taskDao.GetTaskIdsByGroupId(3); err != nil {
		t.Fatal(err)
	}
}

func TestGetConditionIdByGroupId(t *testing.T) {
	if _, err := taskDao.GetConditionIdByGroupId(1); err != nil {
		t.Fatal(err)
	}
}

func TestCreateTaskAnswer(t *testing.T) {
	p := model.Answer{
		UserId:      42,
		TaskId:      5,
		ConditionId: 5,
		Answer:      "",
		Reason:      "",
	}
	if err := taskDao.CreateTaskAnswer(&p); err != nil {
		t.Fatal(err)
	}
}
