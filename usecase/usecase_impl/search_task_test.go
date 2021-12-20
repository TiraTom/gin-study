package usecase_impl_test

import (
	"fmt"
	"testing"

	"github.com/Tiratom/gin-study/domain/domain_obj"
	gr "github.com/Tiratom/gin-study/grpc"
	"github.com/Tiratom/gin-study/mock/mock_repository_interface"
	"github.com/Tiratom/gin-study/usecase/usecase_impl"
	"github.com/golang/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSearchTask_Do(t *testing.T) {
	type args struct {
		p *gr.GetTaskByConditionRequestParam
	}
	tests := []struct {
		name string
		args args
		// want            *domain_obj.Tasks　　// mockからの戻り値をそのまま返すだけなのでテスト対象から外す
		wantErr         bool
		prepareMockFunc func(*mock_repository_interface.MockTask) // mock準備用関数
	}{
		{
			name: "全ての項目で検索条件を指定している場合",
			args: args{
				p: &gr.GetTaskByConditionRequestParam{
					Name:                  "Name",
					Details:               "Details",
					ImportanceName:        "HIGH",
					Deadline:              &timestamppb.Timestamp{Seconds: timestamp20210822150001},
					SearchTypeForDeadline: gr.TimestampCompareBy_AFTER,
				},
			},
			wantErr: false,
			prepareMockFunc: func(mTaskR *mock_repository_interface.MockTask) {
				mTaskR.EXPECT().Search(&domain_obj.TaskSearchCondition{
					Name:                  "Name",
					Details:               "Details",
					ImportanceName:        "HIGH",
					Deadline:              &time20210822150001,
					SearchTypeForDeadline: &searchTypeForDeadlineAfter,
				}).Return(&domain_obj.Tasks{
					Value: []*domain_obj.Task{{
						Id:             "Id",
						Name:           "Name",
						Details:        "Details",
						ImportanceName: "HIGH",
						Deadline:       &time20210822150001,
						RegisteredAt:   &time20210822150002,
						UpdatedAt:      &time20210822150003,
						Version:        1,
					}},
				}, nil)
			},
		},
		{
			name: "検索条件が空の場合",
			args: args{
				p: &gr.GetTaskByConditionRequestParam{
					Name:                  "",
					Details:               "",
					ImportanceName:        "",
					Deadline:              nil,
					SearchTypeForDeadline: searchTypeForDeadlineNone,
				},
			},
			wantErr: false,
			prepareMockFunc: func(mTaskR *mock_repository_interface.MockTask) {
				mTaskR.EXPECT().Search(&domain_obj.TaskSearchCondition{
					Name:                  "",
					Details:               "",
					ImportanceName:        "",
					Deadline:              nil,
					SearchTypeForDeadline: nil,
				}).Return(&domain_obj.Tasks{
					Value: []*domain_obj.Task{},
				}, nil)
			},
		},
		{
			name: "検索処理でエラーが発生した場合_エラーを握りつぶしていないことを確認",
			args: args{
				p: &gr.GetTaskByConditionRequestParam{
					Name:                  "Name",
					Details:               "Details",
					ImportanceName:        "HIGH",
					Deadline:              &timestamppb.Timestamp{Seconds: timestamp20210822150001},
					SearchTypeForDeadline: gr.TimestampCompareBy_AFTER,
				},
			},
			wantErr: true,
			prepareMockFunc: func(mTaskR *mock_repository_interface.MockTask) {
				mTaskR.EXPECT().Search(&domain_obj.TaskSearchCondition{
					Name:                  "Name",
					Details:               "Details",
					ImportanceName:        "HIGH",
					Deadline:              &time20210822150001,
					SearchTypeForDeadline: &searchTypeForDeadlineAfter,
				}).Return(&domain_obj.Tasks{
					Value: []*domain_obj.Task{},
				}, fmt.Errorf("エラー"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mTaskR := mock_repository_interface.NewMockTask(ctrl)
			tt.prepareMockFunc(mTaskR)

			s := &usecase_impl.SearchTask{
				Tr: mTaskR,
			}

			_, err := s.Do(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchTask.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
