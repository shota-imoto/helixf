package line_model_test

import (
	"testing"

	"github.com/shota-imoto/helixf/lib/models/line_model"
)

func TestGetFromLineDatabase(t *testing.T) {
	// 開発用グループの取得
	// TDDのためとりあえず上記の前提でテストを書く。最終的にはmock書くまでもないので開発後落ちるようになったら消す
	group := line_model.LineGroup{GroupId: "Caa245c3d70b26b44b475553ab3ed017e"}

	err := group.GetFromLineDatabase()

	if err != nil {
		t.Error(err)
	}

	if group.GroupName != "APIテスト" {
		t.Errorf("Name is wrong: %v", group.GroupName)
	}
}
