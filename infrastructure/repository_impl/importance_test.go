package repository_impl

import (
	"reflect"
	"testing"

	"github.com/Tiratom/gin-study/config"
	"github.com/Tiratom/gin-study/domain/domain_obj"
)

func TestImportance_GetAll(t *testing.T) {
	type fields struct {
		db *config.DB
	}

	conf, db := SetUpForInfrastructureDBTest(t)

	tests := []struct {
		name    string
		fields  fields
		want    []*domain_obj.Importance
		wantErr bool
	}{
		{
			name:   "初期状態",
			fields: fields{db},
			want: []*domain_obj.Importance{
				{
					Name:  "MEDIUM",
					Level: 2,
				},
				{
					Name:  "VERY_HIGH",
					Level: 4,
				},
				{
					Name:  "HIGH",
					Level: 3,
				},
				{
					Name:  "LOW",
					Level: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BeforeEachForDBTest(t, conf, tt.fields.db)

			i := &Importance{
				db: tt.fields.db,
			}
			got, err := i.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Importance.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Importance.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
