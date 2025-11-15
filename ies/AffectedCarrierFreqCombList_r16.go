package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreqCombList_r16 struct {
	Value []AffectedCarrierFreqComb_r16 `lb:1,ub:maxCombIDC_r16,madatory`
}

func (ie *AffectedCarrierFreqCombList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*AffectedCarrierFreqComb_r16]([]*AffectedCarrierFreqComb_r16{}, uper.Constraint{Lb: 1, Ub: maxCombIDC_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode AffectedCarrierFreqCombList_r16", err)
	}
	return nil
}

func (ie *AffectedCarrierFreqCombList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*AffectedCarrierFreqComb_r16]([]*AffectedCarrierFreqComb_r16{}, uper.Constraint{Lb: 1, Ub: maxCombIDC_r16}, false)
	fn := func() *AffectedCarrierFreqComb_r16 {
		return new(AffectedCarrierFreqComb_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode AffectedCarrierFreqCombList_r16", err)
	}
	ie.Value = []AffectedCarrierFreqComb_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
