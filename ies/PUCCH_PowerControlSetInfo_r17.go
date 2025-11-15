package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PowerControlSetInfo_r17 struct {
	pucch_PowerControlSetInfoId_r17  PUCCH_PowerControlSetInfoId_r17                         `madatory`
	p0_PUCCH_Id_r17                  P0_PUCCH_Id                                             `madatory`
	pucch_ClosedLoopIndex_r17        PUCCH_PowerControlSetInfo_r17_pucch_ClosedLoopIndex_r17 `madatory`
	pucch_PathlossReferenceRS_Id_r17 PUCCH_PathlossReferenceRS_Id_r17                        `madatory`
}

func (ie *PUCCH_PowerControlSetInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pucch_PowerControlSetInfoId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_PowerControlSetInfoId_r17", err)
	}
	if err = ie.p0_PUCCH_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode p0_PUCCH_Id_r17", err)
	}
	if err = ie.pucch_ClosedLoopIndex_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_ClosedLoopIndex_r17", err)
	}
	if err = ie.pucch_PathlossReferenceRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_PathlossReferenceRS_Id_r17", err)
	}
	return nil
}

func (ie *PUCCH_PowerControlSetInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pucch_PowerControlSetInfoId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_PowerControlSetInfoId_r17", err)
	}
	if err = ie.p0_PUCCH_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode p0_PUCCH_Id_r17", err)
	}
	if err = ie.pucch_ClosedLoopIndex_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_ClosedLoopIndex_r17", err)
	}
	if err = ie.pucch_PathlossReferenceRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_PathlossReferenceRS_Id_r17", err)
	}
	return nil
}
