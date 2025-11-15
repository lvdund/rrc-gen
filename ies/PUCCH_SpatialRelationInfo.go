package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_SpatialRelationInfo struct {
	pucch_SpatialRelationInfoId  PUCCH_SpatialRelationInfoId               `madatory`
	servingCellId                *ServCellIndex                            `optional`
	referenceSignal              PUCCH_SpatialRelationInfo_referenceSignal `madatory`
	pucch_PathlossReferenceRS_Id PUCCH_PathlossReferenceRS_Id              `madatory`
	p0_PUCCH_Id                  P0_PUCCH_Id                               `madatory`
	closedLoopIndex              PUCCH_SpatialRelationInfo_closedLoopIndex `madatory`
}

func (ie *PUCCH_SpatialRelationInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.servingCellId != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pucch_SpatialRelationInfoId.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_SpatialRelationInfoId", err)
	}
	if ie.servingCellId != nil {
		if err = ie.servingCellId.Encode(w); err != nil {
			return utils.WrapError("Encode servingCellId", err)
		}
	}
	if err = ie.referenceSignal.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignal", err)
	}
	if err = ie.pucch_PathlossReferenceRS_Id.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.p0_PUCCH_Id.Encode(w); err != nil {
		return utils.WrapError("Encode p0_PUCCH_Id", err)
	}
	if err = ie.closedLoopIndex.Encode(w); err != nil {
		return utils.WrapError("Encode closedLoopIndex", err)
	}
	return nil
}

func (ie *PUCCH_SpatialRelationInfo) Decode(r *uper.UperReader) error {
	var err error
	var servingCellIdPresent bool
	if servingCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pucch_SpatialRelationInfoId.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_SpatialRelationInfoId", err)
	}
	if servingCellIdPresent {
		ie.servingCellId = new(ServCellIndex)
		if err = ie.servingCellId.Decode(r); err != nil {
			return utils.WrapError("Decode servingCellId", err)
		}
	}
	if err = ie.referenceSignal.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignal", err)
	}
	if err = ie.pucch_PathlossReferenceRS_Id.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_PathlossReferenceRS_Id", err)
	}
	if err = ie.p0_PUCCH_Id.Decode(r); err != nil {
		return utils.WrapError("Decode p0_PUCCH_Id", err)
	}
	if err = ie.closedLoopIndex.Decode(r); err != nil {
		return utils.WrapError("Decode closedLoopIndex", err)
	}
	return nil
}
