package usecase_impl_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/di"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/infrastructure/record"
	"github.com/Tiratom/gin-study/usecase/usecase_impl"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 勉強がてらなのでusecaseのテストも一部だけ書いて完了とする。

func TestCreateTask_Do(t *testing.T) {
	conf, db := SetUpForDBTest(t)

	type args struct {
		p *gr.CreateTaskRequestParam
	}

	tests := []struct {
		name              string
		args              args
		want              *domain_obj.Task
		needCompareWant   bool // wantとgotの比較テストを行いたい場合にtrueを設定
		wantErr           bool
		wantErrMsgPartial string
		dbDataTestParam   *createdTaskDbDataTestWantParam // 作成後DBに保存されたデータを取り出して値比較を行いたい場合に値を設定する
	}{
		{
			name: "存在しないImportanceNameの場合",
			args: args{
				p: &gr.CreateTaskRequestParam{
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "NOT_EXIST_NAME",
					Deadline:       &timestamppb.Timestamp{Seconds: timestamp20210822150001},
				},
			},
			want:              nil,
			needCompareWant:   true,
			wantErr:           true,
			wantErrMsgPartial: "Cannot add or update a child row: a foreign key constraint fails",
		},
		{
			name: "正常パターン",
			args: args{
				p: &gr.CreateTaskRequestParam{
					Name:           "DUMMY_NAME",
					Details:        "DUMMY_DETAILS",
					ImportanceName: "HIGH",
					Deadline:       &timestamppb.Timestamp{Seconds: timestamp20210822150001},
				},
			},
			needCompareWant: false,
			wantErr:         false,
			dbDataTestParam: &createdTaskDbDataTestWantParam{
				importanceId: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BeforeEachForDBTest(t, conf, db)
			testStartTime := time.Now()

			c := &usecase_impl.CreateTask{
				// usecaseのテストは結合テストで行うのでモックではなく実物を使う
				Tr: di.InitializeTaskRepositoryInterface(),
			}
			got, err := c.Do(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask.Do() error = %v, \nwantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && !strings.Contains(err.Error(), tt.wantErrMsgPartial) {
				t.Errorf("エラーメッセージに想定の文字列が含まれていません\ngot = %v, \nwantToContain = %v", err, tt.wantErrMsgPartial)
				return
			}
			if tt.needCompareWant && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask.Do() = %v, \nwant %v", got, tt.want)
			}

			if tt.dbDataTestParam != nil {
				createdTaskDbDataTest(t, db, got, tt.args.p, tt.dbDataTestParam, testStartTime)
			}
		})
	}
}

type createdTaskDbDataTestWantParam struct {
	importanceId int64
}

// createdTaskDbDataTestはDBに実際に登録されたデータが想定通りかのテストを行う
// あくまでusecaseの実装から見たテストなので、idの固定値設定などは行わずまたテストのチェック項目からも外している
func createdTaskDbDataTest(t *testing.T, db *config.DB, createResult *domain_obj.Task, rp *gr.CreateTaskRequestParam, wantParam *createdTaskDbDataTestWantParam, testStartTime time.Time) {
	var taskSavedByTest *record.Task
	result := db.Gdb.Raw("SELECT * FROM gin_study.tasks WHERE id = ?;", createResult.Id).Scan(&taskSavedByTest)
	if result.Error != nil {
		t.Errorf("作成処理実施後のデータ存在チェックテストにおいてエラー発生; %v", result.Error)
	}

	// 作成依頼のパラメーターと実際に作成されたタスクの項目の値比較を行う。一通りチェックしたいので可変のフラグ用変数を置いてチェックしている
	isParamMismatch := false
	if taskSavedByTest.Name != rp.Name || taskSavedByTest.Details != rp.Details {
		isParamMismatch = true
	}
	if taskSavedByTest.ImportanceId != wantParam.importanceId { // HIGHを設定している想定
		isParamMismatch = true
	}
	if taskSavedByTest.Deadline != time20210822150001 {
		isParamMismatch = true
	}
	if taskSavedByTest.RegisteredTime.Unix() < testStartTime.Unix() {
		isParamMismatch = true
	}
	if taskSavedByTest.RegisteredTime != taskSavedByTest.UpdatedTime {
		isParamMismatch = true
	}
	if isParamMismatch {
		t.Errorf("作成処理実施後のデータの値が想定と異なります\nreturnValue = %v\nrequestParam = %v", taskSavedByTest, rp)
	}

	if taskSavedByTest.Version != 1 {
		t.Errorf("バージョン数が想定と異なります\ngot = %v\nwant = %v", taskSavedByTest.Version, 1)
	}

}
