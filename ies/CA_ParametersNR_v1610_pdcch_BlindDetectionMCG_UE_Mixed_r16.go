package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16 struct {
	pdcch_BlindDetectionMCG_UE1_r16 int64 `lb:0,ub:15,madatory`
	pdcch_BlindDetectionMCG_UE2_r16 int64 `lb:0,ub:15,madatory`
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.pdcch_BlindDetectionMCG_UE1_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionMCG_UE1_r16", err)
	}
	if err = w.WriteInteger(ie.pdcch_BlindDetectionMCG_UE2_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionMCG_UE2_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_pdcch_BlindDetectionMCG_UE1_r16 int64
	if tmp_int_pdcch_BlindDetectionMCG_UE1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionMCG_UE1_r16", err)
	}
	ie.pdcch_BlindDetectionMCG_UE1_r16 = tmp_int_pdcch_BlindDetectionMCG_UE1_r16
	var tmp_int_pdcch_BlindDetectionMCG_UE2_r16 int64
	if tmp_int_pdcch_BlindDetectionMCG_UE2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionMCG_UE2_r16", err)
	}
	ie.pdcch_BlindDetectionMCG_UE2_r16 = tmp_int_pdcch_BlindDetectionMCG_UE2_r16
	return nil
}
