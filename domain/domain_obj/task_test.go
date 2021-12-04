package domain_obj

import (
	"reflect"
	"testing"
	"time"

	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/infrastructure/record"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTask_ToDto(t *testing.T) {
	type fields struct {
		Id             string
		Name           string
		Details        string
		ImportanceName string
		RegisteredAt   *time.Time
		Deadline       *time.Time
		UpdatedAt      *time.Time
		Version        uint
	}

	testVs := getDummyValues(t)

	tests := []struct {
		name    string
		fields  fields
		want    *gr.Task
		wantErr bool
	}{
		{
			name: "値が全て存在するパターン",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        2,
			},
			want: &gr.Task{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt: &timestamppb.Timestamp{
					Seconds: testVs.dummyTimestampSec1,
				},
				Deadline: &timestamppb.Timestamp{
					Seconds: testVs.dummyTimestampSec2,
				},
				UpdatedAt: &timestamppb.Timestamp{
					Seconds: testVs.dummyTimestampSec3,
				},
			},
			wantErr: false,
		},
		{
			name: "値が存在しないパターンでも例外が起きない",
			fields: fields{
				Id:             "",
				Name:           "",
				Details:        "",
				ImportanceName: "",
				RegisteredAt:   nil,
				Deadline:       nil,
				UpdatedAt:      nil,
				Version:        0,
			},
			want: &gr.Task{
				Id:             "",
				Name:           "",
				Details:        "",
				ImportanceName: "",
				RegisteredAt:   nil,
				Deadline:       nil,
				UpdatedAt:      nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				Details:        tt.fields.Details,
				ImportanceName: tt.fields.ImportanceName,
				RegisteredAt:   tt.fields.RegisteredAt,
				Deadline:       tt.fields.Deadline,
				UpdatedAt:      tt.fields.UpdatedAt,
				Version:        tt.fields.Version,
			}
			got, err := tr.ToDto()
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.ToDto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.ToDto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_ToRecord(t *testing.T) {
	type fields struct {
		Id             string
		Name           string
		Details        string
		ImportanceName string
		RegisteredAt   *time.Time
		Deadline       *time.Time
		UpdatedAt      *time.Time
		Version        uint
	}
	type args struct {
		i int64
	}

	testVs := getDummyValues(t)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *record.Task
	}{
		{
			name: "値が全て存在する場合",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        2,
			},
			args: args{
				i: 3,
			},
			want: &record.Task{
				Id:           "DUMMY_ID1",
				Name:         "DUMMY_NAME1",
				Details:      "DUMMY_DETAILS1",
				ImportanceId: 3,
				RegisteredAt: testVs.dummyTime1,
				Deadline:     testVs.dummyTime2,
				UpdatedAt:    testVs.dummyTime3,
				Version:      2,
			},
		},
		{
			name: "存在しない値がある場合でもpanicが起きない",
			fields: fields{
				Id:             "",
				Name:           "",
				Details:        "",
				ImportanceName: "",
				RegisteredAt:   nil,
				Deadline:       nil,
				UpdatedAt:      nil,
				Version:        0,
			},
			args: args{
				i: 0,
			},
			want: &record.Task{
				Id:           "",
				Name:         "",
				Details:      "",
				ImportanceId: 0,
				RegisteredAt: *new(time.Time),
				Deadline:     *new(time.Time),
				UpdatedAt:    *new(time.Time),
				Version:      0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				Details:        tt.fields.Details,
				ImportanceName: tt.fields.ImportanceName,
				RegisteredAt:   tt.fields.RegisteredAt,
				Deadline:       tt.fields.Deadline,
				UpdatedAt:      tt.fields.UpdatedAt,
				Version:        tt.fields.Version,
			}
			if got := tr.ToRecord(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.ToRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTask(t *testing.T) {
	type args struct {
		tr *record.TaskAndImportance
	}

	testVs := getDummyValues(t)

	tests := []struct {
		name string
		args args
		want *Task
	}{
		{
			name: "値が全て存在する場合",
			args: args{
				tr: &record.TaskAndImportance{
					Id:              "DUMMY_ID1",
					Name:            "DUMMY_NAME1",
					Details:         "DUMMY_DETAILS1",
					ImportanceId:    2,
					ImportanceName:  "DUMMY_IMPORTANCE_NAME1",
					ImportanceLevel: 3,
					RegisteredAt:    testVs.dummyTime1,
					Deadline:        testVs.dummyTime2,
					UpdatedAt:       testVs.dummyTime3,
					Version:         4,
				},
			},
			want: &Task{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        4,
			},
		},
		{
			name: "存在しない値がある場合もpanicが起きない",
			args: args{
				tr: &record.TaskAndImportance{
					Id:              "",
					Name:            "",
					Details:         "",
					ImportanceId:    0,
					ImportanceName:  "",
					ImportanceLevel: 0,
					RegisteredAt:    *new(time.Time),
					Deadline:        *new(time.Time),
					UpdatedAt:       *new(time.Time),
					Version:         0,
				},
			},
			want: &Task{
				Id:             "",
				Name:           "",
				Details:        "",
				ImportanceName: "",
				RegisteredAt:   new(time.Time),
				Deadline:       new(time.Time),
				UpdatedAt:      new(time.Time),
				Version:        0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTask(tt.args.tr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTaskToCreate(t *testing.T) {
	type args struct {
		p *gr.CreateTaskRequestParam
	}

	testVs := getDummyValues(t)

	// 現在時刻設定処理部分はテスト用に書き換える
	nowTimeFunc = func() time.Time {
		return testVs.dummyNowTime
	}

	// ID発行処理もテスト用に書き換える
	dummyUuidStr := "e09d842c-04d2-4f75-851a-f4a49e578329"
	newUuidFunc = func() uuid.UUID {
		id, err := uuid.Parse(dummyUuidStr)
		if err != nil {
			t.Errorf("テスト用に書き換えたuuid発行処理でエラー %v", err)
		}
		return id
	}

	tests := []struct {
		name    string
		args    args
		want    *Task
		wantErr bool
	}{
		{
			name: "通常パターン",
			args: args{
				p: &gr.CreateTaskRequestParam{
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "DUMMY_IMPORTANCE_NAME",
					Deadline: &timestamppb.Timestamp{
						Seconds: testVs.dummyTimestampSec1,
					},
				},
			},
			want: &Task{
				Id:             dummyUuidStr,
				Name:           "DUMMY_NAME",
				Details:        "DUMMY_DETAILS",
				ImportanceName: "DUMMY_IMPORTANCE_NAME",
				RegisteredAt:   &testVs.dummyNowTime,
				Deadline:       &testVs.dummyTime1,
				UpdatedAt:      &testVs.dummyNowTime,
				Version:        1,
			},
			wantErr: false,
		},
		// バリデーションで弾けるパターンはチェックしていない
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTaskToCreate(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTaskToCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskToCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTaskToUpdate(t *testing.T) {
	type args struct {
		o *Task
		p *gr.UpdateTaskRequestParam
	}

	testVs := getDummyValues(t)

	// 2021/11/30 05:51:08(UTC)
	newDeadlineTimestamp := int64(1638251468)
	newDeadline, err := time.Parse("2006/01/02 15:04:05", "2021/11/30 05:51:08")
	if err != nil {
		t.Errorf("時刻の変換処理でエラー発生 %s", err)
	}

	dummyNowTime := time.Date(2021, 8, 26, 14, 16, 18, 0, time.UTC)
	// 現在時刻設定処理部分はテスト用に書き換える
	nowTimeFunc = func() time.Time {
		return dummyNowTime
	}

	tests := []struct {
		name    string
		args    args
		want    *Task
		wantErr bool
	}{
		{
			name: "すべての項目で更新を行うパターン",
			args: args{
				o: &Task{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					RegisteredAt:   &testVs.dummyTime1,
					Deadline:       &testVs.dummyTime2,
					UpdatedAt:      &testVs.dummyTime3,
					Version:        4,
				},
				p: &gr.UpdateTaskRequestParam{
					Id:             "NEW_ID",
					Name:           "NEW_NAME",
					Details:        "NEW_DETAILS",
					ImportanceName: "NEW_IMPORTANCE_NAME",
					Deadline: &timestamppb.Timestamp{
						Seconds: newDeadlineTimestamp,
					},
				},
			},
			want: &Task{
				Id:             "DUMMY_ID1",
				Name:           "NEW_NAME",
				Details:        "NEW_DETAILS",
				ImportanceName: "NEW_IMPORTANCE_NAME",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &newDeadline,
				UpdatedAt:      &dummyNowTime,
				Version:        5,
			},
			wantErr: false,
		},
		{
			name: "一部の項目で更新を行うパターン",
			args: args{
				o: &Task{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					RegisteredAt:   &testVs.dummyTime1,
					Deadline:       &testVs.dummyTime2,
					UpdatedAt:      &testVs.dummyTime3,
					Version:        4,
				},
				p: &gr.UpdateTaskRequestParam{
					Deadline: &timestamppb.Timestamp{
						Seconds: newDeadlineTimestamp,
					},
				},
			},
			want: &Task{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &newDeadline,
				UpdatedAt:      &dummyNowTime,
				Version:        5,
			},
			wantErr: false,
		},
		{
			// 値に変更がないかはdomain_obj.TaskのIsNeededToUpdate()で事前にチェックする想定のため、変更がなくてもVerUpした新規オブジェクトが返却される動作でOKとしている
			name: "項目に値が設定されていない場合",
			args: args{
				o: &Task{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					RegisteredAt:   &testVs.dummyTime1,
					Deadline:       &testVs.dummyTime2,
					UpdatedAt:      &testVs.dummyTime3,
					Version:        4,
				},
				p: &gr.UpdateTaskRequestParam{
					Id:             "",
					Name:           "",
					Details:        "",
					ImportanceName: "",
					Deadline:       nil,
				},
			},
			want: &Task{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &dummyNowTime,
				Version:        5,
			},
			wantErr: false,
		},
		{
			// 値に変更がないかはdomain_obj.TaskのIsNeededToUpdate()で事前にチェックする想定のため、変更がなくてもVerUpした新規オブジェクトが返却される動作でOKとしている
			name: "値に変更がない場合_全項目送られた場合",
			args: args{
				o: &Task{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					RegisteredAt:   &testVs.dummyTime1,
					Deadline:       &testVs.dummyTime2,
					UpdatedAt:      &testVs.dummyTime3,
					Version:        4,
				},
				p: &gr.UpdateTaskRequestParam{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					Deadline: &timestamppb.Timestamp{
						Seconds: testVs.dummyTimestampSec2,
					},
				},
			},
			want: &Task{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &dummyNowTime,
				Version:        5,
			},
			wantErr: false,
		},
		{
			// 値に変更がないかはdomain_obj.TaskのIsNeededToUpdate()で事前にチェックする想定のため、変更がなくてもVerUpした新規オブジェクトが返却される動作でOKとしている
			name: "値に変更がない場合_一部の項目だけ送られた場合",
			args: args{
				o: &Task{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					RegisteredAt:   &testVs.dummyTime1,
					Deadline:       &testVs.dummyTime2,
					UpdatedAt:      &testVs.dummyTime3,
					Version:        4,
				},
				p: &gr.UpdateTaskRequestParam{
					Details: "DUMMY_DETAILS1",
				},
			},
			want: &Task{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyNowTime,
				Version:        5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTaskToUpdate(tt.args.o, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTaskToUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskToUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_IsNeededToUpdate(t *testing.T) {
	type fields struct {
		Id             string
		Name           string
		Details        string
		ImportanceName string
		RegisteredAt   *time.Time
		Deadline       *time.Time
		UpdatedAt      *time.Time
		Version        uint
	}
	type args struct {
		p *gr.UpdateTaskRequestParam
	}

	testVs := getDummyValues(t)
	// 2021/11/30 05:51:08(UTC)
	newDeadlineTimestamp := int64(1638251468)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Nameが更新対象の場合",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        4,
			},
			args: args{
				p: &gr.UpdateTaskRequestParam{
					Id:             "DUMMY_ID1",
					Name:           "NEW_NAME",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					Deadline: &timestamppb.Timestamp{
						Seconds: testVs.dummyTimestampSec2,
					},
				},
			},
			want: true,
		},
		{
			name: "Detailsが更新対象の場合",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        4,
			},
			args: args{
				p: &gr.UpdateTaskRequestParam{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "NEW_DETAILS",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					Deadline: &timestamppb.Timestamp{
						Seconds: testVs.dummyTimestampSec2,
					},
				},
			},
			want: true,
		},
		{
			name: "ImportanceNameが更新対象の場合",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        4,
			},
			args: args{
				p: &gr.UpdateTaskRequestParam{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "NEW_IMPORTANCE_NAME1",
					Deadline: &timestamppb.Timestamp{
						Seconds: testVs.dummyTimestampSec2,
					},
				},
			},
			want: true,
		},
		{
			name: "Deadlineが更新対象の場合",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        4,
			},
			args: args{
				p: &gr.UpdateTaskRequestParam{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					Deadline: &timestamppb.Timestamp{
						Seconds: newDeadlineTimestamp,
					},
				},
			},
			want: true,
		},
		{
			name: "値に変更がない場合",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        4,
			},
			args: args{
				p: &gr.UpdateTaskRequestParam{
					Id:             "DUMMY_ID1",
					Name:           "DUMMY_NAME1",
					Details:        "DUMMY_DETAILS1",
					ImportanceName: "DUMMY_IMPORTANCE_NAME1",
					Deadline: &timestamppb.Timestamp{
						Seconds: testVs.dummyTimestampSec2,
					},
				},
			},
			want: false,
		},
		{
			name: "値に変更がない場合_項目が空の場合",
			fields: fields{
				Id:             "DUMMY_ID1",
				Name:           "DUMMY_NAME1",
				Details:        "DUMMY_DETAILS1",
				ImportanceName: "DUMMY_IMPORTANCE_NAME1",
				RegisteredAt:   &testVs.dummyTime1,
				Deadline:       &testVs.dummyTime2,
				UpdatedAt:      &testVs.dummyTime3,
				Version:        4,
			},
			args: args{
				p: &gr.UpdateTaskRequestParam{
					Id: "DUMMY_ID1",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				Id:             tt.fields.Id,
				Name:           tt.fields.Name,
				Details:        tt.fields.Details,
				ImportanceName: tt.fields.ImportanceName,
				RegisteredAt:   tt.fields.RegisteredAt,
				Deadline:       tt.fields.Deadline,
				UpdatedAt:      tt.fields.UpdatedAt,
				Version:        tt.fields.Version,
			}
			if got := tr.IsNeededToUpdate(tt.args.p); got != tt.want {
				t.Errorf("Task.IsNeededToUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestValuesSet struct {
	// 2021/11/29 00:11:12(UTC)
	dummyTime1         time.Time
	dummyTimestampSec1 int64
	// 2021/11/30 00:11:12(UTC)
	dummyTime2         time.Time
	dummyTimestampSec2 int64
	// 2021/11/30 00:11:12(UTC)
	dummyTime3         time.Time
	dummyTimestampSec3 int64

	// 2021/08/26 14:16:18(UTC)
	dummyNowTime time.Time
}

func getDummyValues(t *testing.T) *TestValuesSet {
	// 2021/11/29 00:11:12(UTC)
	dummyTime1, err := time.Parse("2006/01/02 15:04:05", "2021/11/29 00:11:12")
	if err != nil {
		t.Errorf("時刻の変換処理でエラー発生 %s", err)
	}
	dummyTimestampSec1 := int64(1638144672)

	// 2021/11/30 00:11:12(UTC)
	dummyTime2 := dummyTime1.AddDate(0, 0, 1)
	dummyTimestampSec2 := int64(1638231072)

	// 2021/11/30 00:11:12(UTC)
	dummyTime3 := dummyTime2.AddDate(0, 0, 1)
	dummyTimestampSec3 := int64(1638317472)

	dummyNowTime := time.Date(2021, 8, 26, 14, 16, 18, 0, time.UTC)
	// 現在時刻設定処理部分はテスト用に書き換える
	nowTimeFunc = func() time.Time {
		return dummyNowTime
	}

	return &TestValuesSet{
		dummyTime1:         dummyTime1,
		dummyTimestampSec1: dummyTimestampSec1,
		dummyTime2:         dummyTime2,
		dummyTimestampSec2: dummyTimestampSec2,
		dummyTime3:         dummyTime3,
		dummyTimestampSec3: dummyTimestampSec3,
		dummyNowTime:       dummyNowTime,
	}
}
