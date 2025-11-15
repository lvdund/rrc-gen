package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_EUTRA_AnchorCarrierFreqList_r16 struct {
	Value []ARFCN_ValueEUTRA `lb:1,ub:maxFreqSL_EUTRA_r16,madatory`
}

func (ie *SL_EUTRA_AnchorCarrierFreqList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreqSL_EUTRA_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SL_EUTRA_AnchorCarrierFreqList_r16", err)
	}
	return nil
}

func (ie *SL_EUTRA_AnchorCarrierFreqList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ARFCN_ValueEUTRA]([]*ARFCN_ValueEUTRA{}, uper.Constraint{Lb: 1, Ub: maxFreqSL_EUTRA_r16}, false)
	fn := func() *ARFCN_ValueEUTRA {
		return new(ARFCN_ValueEUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SL_EUTRA_AnchorCarrierFreqList_r16", err)
	}
	ie.Value = []ARFCN_ValueEUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
