package domain_obj

import (
	"reflect"
	"testing"

	infrastructure "github.com/Tiratom/gin-study/infrastructure/record"
)

func TestNewImportance(t *testing.T) {
	type args struct {
		ir *infrastructure.Importance
	}
	tests := []struct {
		name string
		args args
		want *Importance
	}{
		{
			name: "通常パターン",
			args: args{
				ir: &infrastructure.Importance{
					Id:    1,
					Name:  "HOGE",
					Level: 1,
				},
			},
			want: &Importance{
				Name:  "HOGE",
				Level: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImportance(tt.args.ir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImportance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImportance_String(t *testing.T) {
	type fields struct {
		Name  string
		Level int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "通常パターン",
			fields: fields{
				Name:  "HOGE",
				Level: 3,
			},
			want: "Name=HOGE&Level=3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Importance{
				Name:  tt.fields.Name,
				Level: tt.fields.Level,
			}
			if got := i.String(); got != tt.want {
				t.Errorf("Importance.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
