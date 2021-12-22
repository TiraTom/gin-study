package usecase_impl_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Tiratom/gin-study/di"
	"github.com/Tiratom/gin-study/domain/domain_obj"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/usecase/usecase_impl"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCreateTask_Do(t *testing.T) {
	conf, db := SetUpForDBTest(t)

	type args struct {
		p *gr.CreateTaskRequestParam
	}
	tests := []struct {
		name              string
		args              args
		want              *domain_obj.Task
		wantErr           bool
		wantErrMsgPartial string
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
			wantErr:           true,
			wantErrMsgPartial: "Cannot add or update a child row: a foreign key constraint fails",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BeforeEachForDBTest(t, conf, db)

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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask.Do() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
