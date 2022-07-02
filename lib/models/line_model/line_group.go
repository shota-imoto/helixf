package line_model

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/shota-imoto/helixf/lib/utils/line"
)

type LineGroup struct {
	Id         uint   `gorm:"primaryKey" sql:"type:uint" json:"id,int"`
	GroupId    string `sql:"type:string,autoIncrement:false" json:"group_id,string"`
	GroupName  string `sql:"type:string" json:"group_name,string"`
	PictureUrl string `sql:"type:string" json:"picture_url,string"`
}

type LineGroupUserMap struct {
	Id          uint `gorm:"primaryKey" sql:"type:uint"`
	UserId      uint `gorm:"uniqueIndex:idx_user_line_group"`
	LineGroupId uint `gorm:"uniqueIndex:idx_user_line_group"`
}

func (group *LineGroup) GetFromLineDatabase() error {
	req, err := http.NewRequest("GET", line.GroupURL(group.GroupId), nil)

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("LINE_ACCESS_TOKEN"))
	client := new(http.Client)
	res, err := client.Do(req)

	reqBody, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(reqBody, &group)
	if err != nil {
		return err
	}
	return nil
}
