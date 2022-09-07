package line_service

import (
	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/line_model"
)

func FindOrCreateGroupByGroupId(group_id string) (line_model.LineGroup, error) {
	group := line_model.LineGroup{GroupId: group_id}
	db.Db.Where("group_id = ?", group_id).First(&group)

	if group.Id != 0 {
		return group, nil
	}
	err := group.GetFromLineDatabase()

	if err == nil {
		db.Db.FirstOrCreate(&group)
	}

	return group, err
}

func GetListGroups(user helixf_user.User) ([]line_model.LineGroup, error) {
	var groups []line_model.LineGroup
	result := db.Db.Joins("join line_group_user_maps on line_group_user_maps.line_group_id = line_groups.id", db.Db.Where(&line_model.LineGroupUserMap{UserId: user.Id})).Find(&groups)
	if result.Error != nil {
		return groups, result.Error
	}

	return groups, nil
}

func GetGroup(user helixf_user.User, group_id string) (line_model.LineGroup, error) {
	var group line_model.LineGroup
	result := db.Db.Joins("join line_group_user_maps on line_group_user_maps.line_group_id = line_groups.id", db.Db.Where(&line_model.LineGroupUserMap{UserId: user.Id})).First(&group, group_id)
	if result.Error != nil {
		return group, result.Error
	}
	return group, nil
}

func GetGroupWithTemplate(user helixf_user.User, group_id string) (line_model.LineGroup, error) {
	var group line_model.LineGroup
	result := db.Db.Preload("RegularScheduleTemplates").Joins("join line_group_user_maps on line_group_user_maps.line_group_id = line_groups.id", db.Db.Where(&line_model.LineGroupUserMap{UserId: user.Id})).First(&group, group_id)
	if result.Error != nil {
		return group, result.Error
	}
	return group, nil
}

func JoinGroup(group line_model.LineGroup, user helixf_user.User) {
	group_map := line_model.LineGroupUserMap{LineGroupId: group.Id, UserId: user.Id}
	db.Db.FirstOrCreate(&group_map)
}
