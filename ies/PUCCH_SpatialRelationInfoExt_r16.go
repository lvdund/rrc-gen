package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SpatialRelationInfoExt_r16 struct {
	pucch_SpatialRelationInfoId_v1610  *PUCCH_SpatialRelationInfoId_v1610  `optional`
	pucch_PathlossReferenceRS_Id_v1610 *PUCCH_PathlossReferenceRS_Id_v1610 `optional`
}

func (ie *PUCCH_SpatialRelationInfoExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pucch_SpatialRelationInfoId_v1610 != nil, ie.pucch_PathlossReferenceRS_Id_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pucch_SpatialRelationInfoId_v1610 != nil {
		if err = ie.pucch_SpatialRelationInfoId_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_SpatialRelationInfoId_v1610", err)
		}
	}
	if ie.pucch_PathlossReferenceRS_Id_v1610 != nil {
		if err = ie.pucch_PathlossReferenceRS_Id_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_PathlossReferenceRS_Id_v1610", err)
		}
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfoExt_r16) Decode(r *uper.UperReader) error {
	var err error
	var pucch_SpatialRelationInfoId_v1610Present bool
	if pucch_SpatialRelationInfoId_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_PathlossReferenceRS_Id_v1610Present bool
	if pucch_PathlossReferenceRS_Id_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pucch_SpatialRelationInfoId_v1610Present {
		ie.pucch_SpatialRelationInfoId_v1610 = new(PUCCH_SpatialRelationInfoId_v1610)
		if err = ie.pucch_SpatialRelationInfoId_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_SpatialRelationInfoId_v1610", err)
		}
	}
	if pucch_PathlossReferenceRS_Id_v1610Present {
		ie.pucch_PathlossReferenceRS_Id_v1610 = new(PUCCH_PathlossReferenceRS_Id_v1610)
		if err = ie.pucch_PathlossReferenceRS_Id_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_PathlossReferenceRS_Id_v1610", err)
		}
	}
	return nil
}
