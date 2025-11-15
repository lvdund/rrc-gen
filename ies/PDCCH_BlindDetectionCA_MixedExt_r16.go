package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionCA_MixedExt_r16 struct {
	pdcch_BlindDetectionCA1_r16 int64 `lb:1,ub:15,madatory`
	pdcch_BlindDetectionCA2_r16 int64 `lb:1,ub:15,madatory`
}

func (ie *PDCCH_BlindDetectionCA_MixedExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.pdcch_BlindDetectionCA1_r16, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionCA1_r16", err)
	}
	if err = w.WriteInteger(ie.pdcch_BlindDetectionCA2_r16, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger pdcch_BlindDetectionCA2_r16", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetectionCA_MixedExt_r16) Decode(r *uper.UperReader) error {
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
	return nil
}
