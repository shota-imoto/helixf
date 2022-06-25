package line_model_test

import (
	"testing"

	"github.com/shota-imoto/helixf/lib/models/helixf_user"
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

func TestCheckMember(t *testing.T) {
	// 開発用グループに開発用ユーザーが入っている前提。
	// TDDのためとりあえず上記の前提でテストを書く。最終的にはmock書くまでもないので開発後落ちるようになったら消す

	group := line_model.LineGroup{GroupId: "Caa245c3d70b26b44b475553ab3ed017e"}
	user := helixf_user.User{LineId: "Ub3946af31a9de0bbba4952de1fdacc23"}

	result, err := group.CheckMember(user)

	if err != nil {
		t.Error(err)
	}

	if !result {
		t.Errorf("belongs to group: %v", user.LineId)
	}

	user = helixf_user.User{LineId: "Ub3946af31a9de0bbba4952de1fdacc22"}

	result, err = group.CheckMember(user)

	if err != nil {
		t.Error(err)
	}

	if result {
		t.Errorf("not belongs to group: %v", user.LineId)
	}
}
