package account_asset

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AccountMove().DeclareModel()

	h.AccountMove().AddFields(map[string]models.FieldDefinition{
		"AssetDepreciationIds": models.One2ManyField{
			RelationModel: h.AccountAssetDepreciationLine(),
			ReverseFK:     "",
			String:        "Assets Depreciation Lines",
			OnDelete:      `restrict`,
		},
	})
	h.AccountMove().Methods().ButtonCancel().DeclareMethod(
		`ButtonCancel`,
		func(rs m.AccountMoveSet) {
			//        for move in self:
			//            for line in move.asset_depreciation_ids:
			//                line.move_posted_check = False
			//        return super(AccountMove, self).button_cancel()
		})
	h.AccountMove().Methods().Post().DeclareMethod(
		`Post`,
		func(rs m.AccountMoveSet) {
			//        for move in self:
			//            for depreciation_line in move.asset_depreciation_ids:
			//                depreciation_line.post_lines_and_close_asset()
			//        return super(AccountMove, self).post()
		})
}
