package repository_impl

import (
	"reflect"
	"testing"
	"time"

	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	"github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/infrastructure/record"
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

func TestTask_Create(t *testing.T) {
	type fields struct {
		db *config.DB
	}
	type args struct {
		p *domain_obj.Task
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
			name:   "重要度ラベルがDBにない値の場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.Task{
					Id:             "1",
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "NOT_EXIST_NAME",
					Deadline:       &time20210823000001,
					RegisteredAt:   &time20210823000002,
					UpdatedAt:      &time20210823000003,
					Version:        1,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "通常パターン",
			fields: fields{db},
			args: args{
				p: &domain_obj.Task{
					Id:             "123",
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "HIGH",
					Deadline:       &time20210823000001,
					RegisteredAt:   &time20210823000002,
					UpdatedAt:      &time20210823000003,
					Version:        1,
				},
			},
			want: &domain_obj.Task{
				Id:             "123",
				Name:           "DUMMY_NAME",
				Details:        "DUMMY_DETAILS",
				ImportanceName: "HIGH",
				Deadline:       &time20210823000001,
				RegisteredAt:   &time20210823000002,
				UpdatedAt:      &time20210823000003,
				Version:        1,
			},
			wantErr: false,
		},
		{
			name:   "作成に失敗した場合_ID重複パターン",
			fields: fields{db},
			args: args{
				p: &domain_obj.Task{
					Id:             "123",
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "HIGH",
					Deadline:       &time20210823000001,
					RegisteredAt:   &time20210823000002,
					UpdatedAt:      &time20210823000003,
					Version:        1,
				},
			},
			setUp:   setUp_Create_DuplicateId,
			want:    nil,
			wantErr: true,
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
			got, err := tr.Create(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Update(t *testing.T) {
	type fields struct {
		db *config.DB
	}
	type args struct {
		p *domain_obj.Task
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
			name:   "更新対象のデータがDBに存在しない場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.Task{
					Id:             "NOT_EXIST",
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "HIGH",
					Deadline:       &time20210823000001,
					RegisteredAt:   &time20210823000002,
					UpdatedAt:      &time20210823000003,
					Version:        1,
				},
			},
			want:    nil,
			wantErr: true,
			setUp:   nil,
		},
		{
			name:   "重要度ラベルがDBに存在しない値の場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.Task{
					Id:             "2",
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "NOT_EXIST",
					Deadline:       &time20210823000001,
					RegisteredAt:   &time20210823000002,
					UpdatedAt:      &time20210823000003,
					Version:        1,
				},
			},
			want:    nil,
			wantErr: true,
			setUp:   setUp_Update_TaskExist,
		},
		{
			name:   "通常パターン_全項目で更新",
			fields: fields{db},
			args: args{
				p: &domain_obj.Task{
					Id:             "2",
					Name:           "NEW_NAME",
					Details:        "NEW_DETAILS",
					ImportanceName: "LOW",
					Deadline:       &time20210923000001,
					RegisteredAt:   &time20210923000002,
					UpdatedAt:      &time20210923000003,
					Version:        2,
				},
			},
			want: &domain_obj.Task{
				Id:             "2",
				Name:           "NEW_NAME",
				Details:        "NEW_DETAILS",
				ImportanceName: "LOW",
				Deadline:       &time20210923000001,
				RegisteredAt:   &time20210923000002,
				UpdatedAt:      &time20210923000003,
				Version:        2,
			},
			wantErr: false,
			setUp:   setUp_Update_TaskExist,
		},
		{
			name:   "更新項目なし",
			fields: fields{db},
			args: args{
				p: &domain_obj.Task{
					Id:             "2",
					Name:           "taskName1",
					Details:        "details",
					ImportanceName: "HIGH",
					Deadline:       &time20210823000001,
					RegisteredAt:   &time20210823000002,
					UpdatedAt:      &time20210823000003,
					Version:        1,
				},
			},
			want: &domain_obj.Task{
				Id:             "2",
				Name:           "taskName1",
				Details:        "details",
				ImportanceName: "HIGH",
				Deadline:       &time20210823000001,
				RegisteredAt:   &time20210823000002,
				UpdatedAt:      &time20210823000003,
				Version:        1,
			},
			wantErr: false,
			setUp:   setUp_Update_TaskExist,
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
			got, err := tr.Update(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Delete(t *testing.T) {
	type fields struct {
		db *config.DB
	}
	type args struct {
		id string
	}

	conf, db := SetUpForDBTest(t)

	tests := []struct {
		name                     string
		fields                   fields
		args                     args
		wantErr                  bool
		setUp                    func(*config.DB) error // 各テストパターンの前処理
		doesExistTaskAfterDelete bool                   // 追加チェック項目：削除テスト実施後に対象タスクがDBに存在するかどうか
	}{
		{
			name:   "削除対象が存在しない場合",
			fields: fields{db},
			args: args{
				id: "NOT_EXIST",
			},
			wantErr:                  true,
			setUp:                    nil,
			doesExistTaskAfterDelete: false,
		},
		{
			name:   "削除対象が存在する場合",
			fields: fields{db},
			args: args{
				id: "2",
			},
			wantErr:                  false,
			setUp:                    setUp_Delete_TaskExist,
			doesExistTaskAfterDelete: false,
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
			err := tr.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			var taskAfterTest *record.Task
			result := tt.fields.db.Gdb.Raw("SELECT * FROM gin_study.tasks WHERE id = ?;", tt.args.id).Scan(&taskAfterTest)
			if result.Error != nil {
				t.Errorf("削除処理実施後のデータ存在チェックテストにおいてエラー発生: %v", result.Error)
			}

			if (taskAfterTest != nil) != tt.doesExistTaskAfterDelete {
				t.Errorf("削除処理実施後のデータ存在チェックに失敗 expected=%v actual=%v", tt.doesExistTaskAfterDelete, taskAfterTest != nil)
			}
		})
	}
}

func TestTask_Search(t *testing.T) {
	type fields struct {
		db *config.DB
	}
	type args struct {
		p *domain_obj.TaskSearchCondition
	}

	conf, db := SetUpForDBTest(t)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain_obj.Tasks
		wantErr bool
		setUp   func(*config.DB) error // 各テストパターンの前処理
	}{
		{
			name:   "検索結果が0件の場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.TaskSearchCondition{
					Name: "NOT_EXIST",
				},
			},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{},
			},
			wantErr: false,
			setUp:   nil,
		},
		{
			name:   "全項目を検索条件に指定した場合_対象データが存在する場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.TaskSearchCondition{
					Name:                  "taskName1",
					Details:               "details1",
					ImportanceName:        "VERY_HIGH",
					Deadline:              &time20210823000001,
					SearchTypeForDeadline: grpc.TimestampCompareBy_SAME.Enum(),
				},
			},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{
					{
						Id:             "1",
						Name:           "taskName1",
						Details:        "details1",
						ImportanceName: "VERY_HIGH",
						Deadline:       &time20210823000001,
						RegisteredAt:   &time20210823000002,
						UpdatedAt:      &time20210823000003,
						Version:        1,
					},
				},
			},
			wantErr: false,
			setUp:   setUp_Search_TaskExist,
		},
		{
			name:   "全項目を検索条件に指定した場合_対象データが存在しない場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.TaskSearchCondition{
					Name:                  "taskName1",
					Details:               "details1",
					ImportanceName:        "HIGH",
					Deadline:              &time20210823000001,
					SearchTypeForDeadline: grpc.TimestampCompareBy_SAME.Enum(),
				},
			},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{},
			},
			wantErr: false,
			setUp:   setUp_Search_TaskExist,
		},
		{
			name:   "「期限日時より前」を検索条件に指定した場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.TaskSearchCondition{
					Deadline:              &time20210823000001,
					SearchTypeForDeadline: grpc.TimestampCompareBy_BEFORE.Enum(),
				},
			},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{
					{
						Id:             "1",
						Name:           "taskName1",
						Details:        "details",
						ImportanceName: "VERY_HIGH",
						Deadline:       &time20210822000001,
						RegisteredAt:   &time20210823000001,
						UpdatedAt:      &time20210923000001,
						Version:        3,
					},
				},
			},
			wantErr: false,
			setUp:   setUp_Search_VariableDeadlineTaskExist,
		},
		{
			name:   "「期限日時より後」を検索条件に指定した場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.TaskSearchCondition{
					Deadline:              &time20210823000001,
					SearchTypeForDeadline: grpc.TimestampCompareBy_AFTER.Enum(),
				},
			},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{
					{
						Id:             "3",
						Name:           "taskName3",
						Details:        "details",
						ImportanceName: "MEDIUM",
						Deadline:       &time20210824000001,
						RegisteredAt:   &time20210823000003,
						UpdatedAt:      &time20210923000003,
						Version:        5,
					},
					{
						Id:             "4",
						Name:           "taskName4",
						Details:        "details",
						ImportanceName: "LOW",
						Deadline:       &time20210825000001,
						RegisteredAt:   &time20210823000004,
						UpdatedAt:      &time20210923000004,
						Version:        6,
					},
				},
			},
			wantErr: false,
			setUp:   setUp_Search_VariableDeadlineTaskExist,
		},
		{
			name:   "「期限日時が一致」を検索条件に指定した場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.TaskSearchCondition{
					Deadline:              &time20210823000001,
					SearchTypeForDeadline: grpc.TimestampCompareBy_SAME.Enum(),
				},
			},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{
					{
						Id:             "2",
						Name:           "taskName2",
						Details:        "details",
						ImportanceName: "HIGH",
						Deadline:       &time20210823000001,
						RegisteredAt:   &time20210823000002,
						UpdatedAt:      &time20210923000002,
						Version:        4,
					},
				},
			},
			wantErr: false,
			setUp:   setUp_Search_VariableDeadlineTaskExist,
		},
		{
			name:   "空文字を検索条件に指定した場合",
			fields: fields{db},
			args: args{
				p: &domain_obj.TaskSearchCondition{
					Name:                  "",
					Details:               "",
					ImportanceName:        "",
					Deadline:              &time.Time{},
					SearchTypeForDeadline: nil,
				},
			},
			want: &domain_obj.Tasks{
				Value: []*domain_obj.Task{
					{
						Id:             "1",
						Name:           "taskName1",
						Details:        "details1",
						ImportanceName: "VERY_HIGH",
						Deadline:       &time20210823000001,
						RegisteredAt:   &time20210823000002,
						UpdatedAt:      &time20210823000003,
						Version:        1,
					},
					{
						Id:             "2",
						Name:           "taskName2",
						Details:        "details2",
						ImportanceName: "HIGH",
						Deadline:       &time20210923000001,
						RegisteredAt:   &time20210923000002,
						UpdatedAt:      &time20210923000003,
						Version:        2,
					},
				},
			},
			wantErr: false,
			setUp:   setUp_Search_TaskExist,
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
			got, err := tr.Search(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Task.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Task.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

// setUp_GetAllTasks_MultipleTasksは、GetAllTasksのテスト用に複数タスクDBに存在する状態を用意する。
func setUp_GetAllTasks_MultipleTasks(db *config.DB) error {
	return db.Gdb.Exec(
		`INSERT INTO gin_study.tasks
		(id,name,importance_id,details,registered_time,deadline,isDone,updated_time,version)
		VALUES
		('1', 'taskName1', 2, 'details1', '2021-08-23 00:00:01', '2021-08-23 00:00:02', true,  '2021-08-23 00:00:03', '1'),
		('2', 'taskName2', 3, 'details2', '2021-09-23 00:00:01', '2021-09-23 00:00:02', false, '2021-09-23 00:00:03', '2');
	`).Error
}

// setUp_GetById_TaskExistは、GetByIdのテスト用に検索対象タスクがDBに存在する状態を用意する。
func setUp_GetById_TaskExist(db *config.DB) error {
	return db.Gdb.Exec(
		`INSERT INTO gin_study.tasks
		(id,name,importance_id,details,deadline,registered_time,isDone,updated_time,version)
		VALUES
		('1', 'taskName1', 2, 'details1', '2021-08-23 00:00:01', '2021-08-23 00:00:02', true,  '2021-08-23 00:00:03', '1'),
		('2', 'taskName2', 3, 'details2', '2021-09-23 00:00:01', '2021-09-23 00:00:02', false, '2021-09-23 00:00:03', '2');
	`).Error
}

// setUp_Create_DuplicateIdは、Createのテスト用にこれから作成したいタスクと同じIDが既にDBに存在する状態を用意する。
func setUp_Create_DuplicateId(db *config.DB) error {
	return db.Gdb.Exec(
		`INSERT INTO gin_study.tasks
		(id,name,importance_id,details,deadline,registered_time,isDone,updated_time,version)
		VALUES
		('123', 'taskName1', 2, 'details1', '2021-08-23 00:00:01', '2021-08-23 00:00:02', true,  '2021-08-23 00:00:03', '1');
	`).Error
}

// setUp_Update_TaskExistは、Updateのテスト用に更新対象タスクがDBに存在する状態を用意する。
func setUp_Update_TaskExist(db *config.DB) error {
	return db.Gdb.Exec(
		`INSERT INTO gin_study.tasks
		(id,name,importance_id,details,deadline,registered_time,isDone,updated_time,version)
		VALUES
		('2', 'taskName1', 2, 'details1', '2021-08-23 00:00:01', '2021-08-23 00:00:02', true,  '2021-08-23 00:00:03', '1');
	`).Error
}

// setUp_Delete_TaskExistは、Deleteのテスト用に更新対象タスクがDBに存在する状態を用意する。
func setUp_Delete_TaskExist(db *config.DB) error {
	return db.Gdb.Exec(
		`INSERT INTO gin_study.tasks
		(id,name,importance_id,details,deadline,registered_time,isDone,updated_time,version)
		VALUES
		('2', 'taskName1', 2, 'details1', '2021-08-23 00:00:01', '2021-08-23 00:00:02', true,  '2021-08-23 00:00:03', '1');
	`).Error
}

// setUp_Search_TaskExistは、Searchのテスト用に検索対象タスクがDBに存在する状態を用意する。
func setUp_Search_TaskExist(db *config.DB) error {
	return db.Gdb.Exec(
		`INSERT INTO gin_study.tasks
		(id,name,importance_id,details,deadline,registered_time,isDone,updated_time,version)
		VALUES
		('1', 'taskName1', 2, 'details1', '2021-08-23 00:00:01', '2021-08-23 00:00:02', true,  '2021-08-23 00:00:03', '1'),
		('2', 'taskName2', 3, 'details2', '2021-09-23 00:00:01', '2021-09-23 00:00:02', false, '2021-09-23 00:00:03', '2');
	`).Error
}

// setUp_Search_TaskExistは、Searchのテスト用にさまざまな期限日時を持ったタスクがDBに存在する状態を用意する。
func setUp_Search_VariableDeadlineTaskExist(db *config.DB) error {
	return db.Gdb.Exec(
		`INSERT INTO gin_study.tasks
		(id,name,importance_id,details,deadline,registered_time,isDone,updated_time,version)
		VALUES
		('1', 'taskName1', 2, 'details', '2021-08-22 00:00:01', '2021-08-23 00:00:01', true,  '2021-09-23 00:00:01', '3'),
		('2', 'taskName2', 3, 'details', '2021-08-23 00:00:01', '2021-08-23 00:00:02', false, '2021-09-23 00:00:02', '4'),
		('3', 'taskName3', 1, 'details', '2021-08-24 00:00:01', '2021-08-23 00:00:03', true,  '2021-09-23 00:00:03', '5'),
		('4', 'taskName4', 4, 'details', '2021-08-25 00:00:01', '2021-08-23 00:00:04', true,  '2021-09-23 00:00:04', '6');
	`).Error
}
