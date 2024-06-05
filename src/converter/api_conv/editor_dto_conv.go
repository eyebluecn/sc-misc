package api_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/do"
)

// 领域模型转为传输模型
func ConvertEditorDTO(thing *do.Editor) *sc_misc_api.EditorDTO {
	if thing == nil {
		return nil
	}

	return &sc_misc_api.EditorDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}

// 数据库模型转换为领域模型
func ConvertEditorDTOs(things []*do.Editor) []*sc_misc_api.EditorDTO {
	if things == nil {
		return nil
	}
	var editorDTOs []*sc_misc_api.EditorDTO
	for _, item := range things {
		editorDTOs = append(editorDTOs, ConvertEditorDTO(item))
	}
	return editorDTOs
}
