package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC struct {
	pdcch_BlindDetectionMCG_UE int64 `lb:1,ub:15,madatory`
	pdcch_BlindDetectionSCG_UE int64 `lb:1,ub:15,madatory`
}

func (ie *Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.pdcch_BlindDetectionMCG_UE, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionMCG_UE", err)
	}
	if err = w.WriteInteger(ie.pdcch_BlindDetectionSCG_UE, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionSCG_UE", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_pdcch_BlindDetectionMCG_UE int64
	if tmp_int_pdcch_BlindDetectionMCG_UE, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionMCG_UE", err)
	}
	ie.pdcch_BlindDetectionMCG_UE = tmp_int_pdcch_BlindDetectionMCG_UE
	var tmp_int_pdcch_BlindDetectionSCG_UE int64
	if tmp_int_pdcch_BlindDetectionSCG_UE, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionSCG_UE", err)
	}
	ie.pdcch_BlindDetectionSCG_UE = tmp_int_pdcch_BlindDetectionSCG_UE
	return nil
}
