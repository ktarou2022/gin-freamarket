package repositories

import (
	"gin-fleamarket/models"
	"testing"
	"time"

	"gorm.io/gorm"
)

func Test_itemMemoryRepository_Delate(t *testing.T) {
	tests := []struct {
		name string 
		itemId  uint
		wantErr bool
	}{
		{
			name:    "削除が成功する",
			itemId:  1,
			wantErr: false,
		},
		{
			name:    "存在しないIDの場合、削除が失敗する",
			itemId:  4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			items := []models.Item{
				{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, Model: gorm.Model{
					ID: 1,
					CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
					DeletedAt: gorm.DeletedAt{},
				 },
				},
				{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true, Model: gorm.Model{
					ID: 2,
					CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
					DeletedAt: gorm.DeletedAt{},
				 }},
				{Name: "テストアイテム3", Price: 3000, Description: "テスト3", SoldOut: false, Model: gorm.Model{
					ID: 3,
					CreatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
					UpdatedAt: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
					DeletedAt: gorm.DeletedAt{},
				 }},
			}

			r := NewItemMemoryRepository(items)
			gotErr := r.Delate(tt.itemId)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Delate() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Delate() succeeded unexpectedly")
			}
		})
	}
}
