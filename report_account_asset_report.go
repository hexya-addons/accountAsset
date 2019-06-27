package account_asset

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.AssetAssetReport().DeclareModel()

	h.AssetAssetReport().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Year",
			Required: false,
			ReadOnly: true,
		},
		"Date": models.DateField{
			ReadOnly: true,
		},
		"DepreciationDate": models.DateField{
			String:   "Depreciation Date",
			ReadOnly: true,
		},
		"AssetId": models.Many2OneField{
			RelationModel: h.AccountAssetAsset(),
			String:        "Asset",
			ReadOnly:      true,
		},
		"AssetCategoryId": models.Many2OneField{
			RelationModel: h.AccountAssetCategory(),
			String:        "Asset category",
			ReadOnly:      true,
		},
		"PartnerId": models.Many2OneField{
			RelationModel: h.Partner(),
			String:        "Partner",
			ReadOnly:      true,
		},
		"State": models.SelectionField{
			Selection: types.Selection{
				"draft": "Draft",
				"open":  "Running",
				"close": "Close",
			},
			String:   "Status",
			ReadOnly: true,
		},
		"DepreciationValue": models.FloatField{
			String:   "Amount of Depreciation Lines",
			ReadOnly: true,
		},
		"InstallmentValue": models.FloatField{
			String:   "Amount of Installment Lines",
			ReadOnly: true,
		},
		"MoveCheck": models.BooleanField{
			String:   "Posted",
			ReadOnly: true,
		},
		"InstallmentNbr": models.IntegerField{
			String:   "# of Installment Lines",
			ReadOnly: true,
		},
		"DepreciationNbr": models.IntegerField{
			String:   "# of Depreciation Lines",
			ReadOnly: true,
		},
		"GrossValue": models.FloatField{
			String:   "Gross Amount",
			ReadOnly: true,
		},
		"PostedValue": models.FloatField{
			String:   "Posted Amount",
			ReadOnly: true,
		},
		"UnpostedValue": models.FloatField{
			String:   "Unposted Amount",
			ReadOnly: true,
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			ReadOnly:      true,
		},
	})
	h.AssetAssetReport().Methods().Init().DeclareMethod(
		`Init`,
		func(rs m.AssetAssetReportSet) {
			//        tools.drop_view_if_exists(self._cr, 'asset_asset_report')
			//        self._cr.execute("""
			//            create or replace view asset_asset_report as (
			//                select
			//                    min(dl.id) as id,
			//                    dl.name as name,
			//                    dl.depreciation_date as depreciation_date,
			//                    a.date as date,
			//                    (CASE WHEN dlmin.id = min(dl.id)
			//                      THEN a.value
			//                      ELSE 0
			//                      END) as gross_value,
			//                    dl.amount as depreciation_value,
			//                    dl.amount as installment_value,
			//                    (CASE WHEN dl.move_check
			//                      THEN dl.amount
			//                      ELSE 0
			//                      END) as posted_value,
			//                    (CASE WHEN NOT dl.move_check
			//                      THEN dl.amount
			//                      ELSE 0
			//                      END) as unposted_value,
			//                    dl.asset_id as asset_id,
			//                    dl.move_check as move_check,
			//                    a.category_id as asset_category_id,
			//                    a.partner_id as partner_id,
			//                    a.state as state,
			//                    count(dl.*) as installment_nbr,
			//                    count(dl.*) as depreciation_nbr,
			//                    a.company_id as company_id
			//                from account_asset_depreciation_line dl
			//                    left join account_asset_asset a on (dl.asset_id=a.id)
			//                    left join (select min(d.id) as id,ac.id as ac_id from account_asset_depreciation_line as d inner join account_asset_asset as ac ON (ac.id=d.asset_id) group by ac_id) as dlmin on dlmin.ac_id=a.id
			//                group by
			//                    dl.amount,dl.asset_id,dl.depreciation_date,dl.name,
			//                    a.date, dl.move_check, a.state, a.category_id, a.partner_id, a.company_id,
			//                    a.value, a.id, a.salvage_value, dlmin.id
			//        )""")
		})
}
