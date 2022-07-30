package line_test

import (
	"testing"
	"time"

	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/utils/line"
)

func TestSendConfirm(t *testing.T) {
	date := time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC)
	schedule := regular_schedule.RegularSchedule{Date: date}
	group := line_model.LineGroup{GroupId: "Caa245c3d70b26b44b475553ab3ed017e"}
	test := struct {
		Id      string
		Message string
	}{
		Id:      group.GroupId,
		Message: "Can you attend schedule at 1/2 " + date.Weekday().String(),
	}

	// 検証用のLineグループに実際に通知する場合はこちら。テスト自体はエラーになる点は留意。
	// bot, _ := linebot.New(os.Getenv("LINE_SECRET"), os.Getenv("LINE_ACCESS_TOKEN"))
	// wrapper := line.LineBotWrapper{Bot: bot}
	wrapper := line.DummyBotWrapper{}
	messager := line.Messager{GroupId: group.GroupId, RegularSchedule: schedule, Wrapper: &wrapper}

	messager.SendConfirm()

	if wrapper.PushedId != test.Id {
		t.Errorf("FAIL GroupId expected: %s, got: %s", test.Id, wrapper.PushedId)
	}

	if wrapper.PushedMessages[0] != test.Message {
		t.Errorf("FAIL PushedMessage expected: %s, got: %s", test.Message, wrapper.PushedMessages[0])
	}
}
