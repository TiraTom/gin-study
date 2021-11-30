package domain_obj

import "testing"

func TestImportanceID_IsValid(t *testing.T) {
	type fields struct {
		Id int64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ゼロ値の場合",
			fields: fields{
				Id: 0,
			},
			want: false,
		},
		{
			name: "マイナスの値の場合",
			fields: fields{
				Id: -3,
			},
			want: false,
		},
		{
			name: "正の値の場合",
			fields: fields{
				Id: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ImportanceID{
				Id: tt.fields.Id,
			}
			if got := i.IsValid(); got != tt.want {
				t.Errorf("ImportanceID.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
