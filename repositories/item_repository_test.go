package repositories

import (
	"gin-fleamarket/models"
	"testing"
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
				{Name: "テストアイテム1", Price: 1000, Description: "", SoldOut: false, ID: 1},
				{Name: "テストアイテム2", Price: 2000, Description: "テスト2", SoldOut: true, ID: 2},
				{Name: "テストアイテム3", Price: 3000, Description: "テスト3", SoldOut: false, ID: 3},
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
