package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionMCG_SCG_r17 struct {
	pdcch_BlindDetectionMCG_UE_r17 int64 `lb:1,ub:15,madatory`
	pdcch_BlindDetectionSCG_UE_r17 int64 `lb:1,ub:15,madatory`
}

func (ie *PDCCH_BlindDetectionMCG_SCG_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.pdcch_BlindDetectionMCG_UE_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionMCG_UE_r17", err)
	}
	if err = w.WriteInteger(ie.pdcch_BlindDetectionSCG_UE_r17, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionSCG_UE_r17", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetectionMCG_SCG_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_pdcch_BlindDetectionMCG_UE_r17 int64
	if tmp_int_pdcch_BlindDetectionMCG_UE_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionMCG_UE_r17", err)
	}
	ie.pdcch_BlindDetectionMCG_UE_r17 = tmp_int_pdcch_BlindDetectionMCG_UE_r17
	var tmp_int_pdcch_BlindDetectionSCG_UE_r17 int64
	if tmp_int_pdcch_BlindDetectionSCG_UE_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionSCG_UE_r17", err)
	}
	ie.pdcch_BlindDetectionSCG_UE_r17 = tmp_int_pdcch_BlindDetectionSCG_UE_r17
	return nil
}
