package repository_impl

import (
	"reflect"
	"testing"

	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
)

func TestTask_GetAll(t *testing.T) {
	type fields struct {
		db *config.DB
	}

	conf, db := SetUpForDBTest(t)

	tests := []struct {
		name    string
		fields  fields
		want    *domain_obj.Tasks
		wantErr bool
		setUp   func(*config.DB) error // 各テストパターンの前処理
	}{
		{
			name:   "データが空の初期状態の場合",
			fields: fields{db},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{},
			},
			wantErr: false,
			setUp:   nil,
		},
		{
			name:   "データが2件存在する場合",
			fields: fields{db},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{
					{
						Id:             "1",
						Name:           "taskName1",
						Details:        "details1",
						ImportanceName: "VERY_HIGH",
						RegisteredAt:   &time20210823000001,
						Deadline:       &time20210823000002,
						UpdatedAt:      &time20210823000003,
						Version:        1,
					},
					{
						Id:             "2",
						Name:           "taskName2",
						Details:        "details2",
						ImportanceName: "HIGH",
						RegisteredAt:   &time20210923000001,
						Deadline:       &time20210923000002,
						UpdatedAt:      &time20210923000003,
						Version:        2,
					},
				},
			},
			wantErr: false,
			setUp:   setUp_GetAllTasks_MultipleTasks,
		},
	}
	for _, tt := range tests {
		BeforeEachForDBTest(t, conf, tt.fields.db)

		if tt.setUp != nil {
			err := tt.setUp(db)
			if err != nil {
				t.Errorf("テスト用前処理でエラーが発生しました: %v", err)
				t.FailNow()
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				db: tt.fields.db,
			}
			got, err := tr.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_GetById(t *testing.T) {
	type fields struct {
		db *config.DB
	}
	type args struct {
		id string
	}

	conf, db := SetUpForDBTest(t)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain_obj.Task
		wantErr bool
		setUp   func(*config.DB) error // 各テストパターンの前処理
	}{
		{
			name:   "データが空の初期状態の場合",
			fields: fields{db},
			args: args{
				id: "1",
			},
			want:    nil,
			wantErr: true,
			setUp:   nil,
		},
		{
			name:   "DBにデータはあるが対象データは存在しない場合",
			fields: fields{db},
			args: args{
				id: "9",
			},
			want:    nil,
			wantErr: true,
			setUp:   setUp_GetById_TaskExist,
		},
		{
			name:   "対象データが存在する場合",
			fields: fields{db},
			args: args{
				id: "1",
			},
			want: &domain_obj.Task{
				Id:             "1",
				Name:           "taskName1",
				Details:        "details1",
				ImportanceName: "VERY_HIGH",
				Deadline:       &time20210823000001,
				RegisteredAt:   &time20210823000002,
				UpdatedAt:      &time20210823000003,
				Version:        1,
			},
			wantErr: false,
			setUp:   setUp_GetById_TaskExist,
		},
	}
	for _, tt := range tests {
		BeforeEachForDBTest(t, conf, tt.fields.db)

		if tt.setUp != nil {
			err := tt.setUp(db)
			if err != nil {
				t.Errorf("テスト用前処理でエラーが発生しました: %v", err)
				t.FailNow()
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			tr := &Task{
				db: tt.fields.db,
			}
			got, err := tr.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}
