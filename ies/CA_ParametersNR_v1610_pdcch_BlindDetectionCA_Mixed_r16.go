package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16 struct {
	pdcch_BlindDetectionCA1_r16  int64                                                                               `lb:1,ub:15,madatory`
	pdcch_BlindDetectionCA2_r16  int64                                                                               `lb:1,ub:15,madatory`
	supportedSpanArrangement_r16 CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16_supportedSpanArrangement_r16 `madatory`
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.pdcch_BlindDetectionCA1_r16, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionCA1_r16", err)
	}
	if err = w.WriteInteger(ie.pdcch_BlindDetectionCA2_r16, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionCA2_r16", err)
	}
	if err = ie.supportedSpanArrangement_r16.Encode(w); err != nil {
		return utils.WrapError("Encode supportedSpanArrangement_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_pdcch_BlindDetectionCA1_r16 int64
	if tmp_int_pdcch_BlindDetectionCA1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionCA1_r16", err)
	}
	ie.pdcch_BlindDetectionCA1_r16 = tmp_int_pdcch_BlindDetectionCA1_r16
	var tmp_int_pdcch_BlindDetectionCA2_r16 int64
	if tmp_int_pdcch_BlindDetectionCA2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger pdcch_BlindDetectionCA2_r16", err)
	}
	ie.pdcch_BlindDetectionCA2_r16 = tmp_int_pdcch_BlindDetectionCA2_r16
	if err = ie.supportedSpanArrangement_r16.Decode(r); err != nil {
		return utils.WrapError("Decode supportedSpanArrangement_r16", err)
	}
	return nil
}
