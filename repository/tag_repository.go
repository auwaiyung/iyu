package repository

import (
	"github.com/lhlyu/iyu/common"
	"github.com/lhlyu/iyu/controller/vo"
	"github.com/lhlyu/iyu/repository/po"
)

// get all tags
func (d *dao) GetTagAll() []*po.YuTag {
	sql := "SELECT * FROM yu_tag ORDER BY is_delete,updated_at DESC,created_at DESC"
	var values []*po.YuTag
	if err := common.DB.Select(&values, sql); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return values
}

// get tag by name
func (d *dao) GetTagByName(name string) *po.YuTag {
	sql := "SELECT * FROM yu_tag WHERE `name` = ? limit 1"
	value := &po.YuTag{}
	if err := common.DB.Get(value, sql, name); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

func (d *dao) GetTagById(id int) *po.YuTag {
	sql := "SELECT * FROM yu_tag WHERE id = ? limit 1"
	value := &po.YuTag{}
	if err := common.DB.Get(value, sql, id); err != nil {
		common.Ylog.Debug(err)
		return nil
	}
	return value
}

// update tag
func (d *dao) UpdateTag(param *vo.TagVo) error {
	sql := "UPDATE yu_tag SET is_delete=?,`name` = ?,updated_at = NOW() WHERE id = ?"
	if _, err := common.DB.Exec(sql, param.IsDelete, param.Name, param.Id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// delete by id
func (d *dao) DeleteTagById(id int) error {
	sql := "DELETE FROM yu_tag WHERE id = ?"
	if _, err := common.DB.Exec(sql, id); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}

// add tag
func (d *dao) InsertTag(param *vo.TagVo) error {
	sql := "INSERT INTO yu_tag(`name`) VALUES(?)"
	if _, err := common.DB.Exec(sql, param.Name); err != nil {
		common.Ylog.Debug(err)
		return err
	}
	return nil
}
