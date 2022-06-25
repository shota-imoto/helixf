package line_service_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/service/line_service"
)

func TestMain(m *testing.M) {
	defer func() {
		db.Db.Rollback()
	}()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestFindOrCreateUserByIdToken(t *testing.T) {
	var before_count, after_count int64

	userName := "test"
	lineId := "line_id"
	parser := &helixf_user.ParseStruct{Parser: &helixf_user.DummyParser{Name: userName, LineId: lineId}}

	fmt.Println("レコード未作成のユーザーの場合")
	db.Db.Model(helixf_user.User{}).Where("name = ?", userName).Count(&before_count)
	_, err := line_service.FindOrCreateUserByIdToken(parser)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	db.Db.Model(helixf_user.User{}).Where("name = ?", userName).Count(&after_count)
	count_diff := after_count - before_count
	if count_diff != 1 {
		t.Errorf("record count is wrong: %v", count_diff)
	}

	before_count = after_count
	fmt.Println("レコード未作成のユーザーの場合")
	_, err = line_service.FindOrCreateUserByIdToken(parser)

	if err != nil {
		t.Error(err)
	}
	db.Db.Model(helixf_user.User{}).Where("name = ?", userName).Count(&after_count)

	count_diff = after_count - before_count

	if count_diff != 0 {
		t.Errorf("record count is wrong: %v", count_diff)
	}
}
