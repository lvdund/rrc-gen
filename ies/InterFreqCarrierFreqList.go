package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqList struct {
	Value []InterFreqCarrierFreqInfo `lb:1,ub:maxFreq,madatory`
}

func (ie *InterFreqCarrierFreqList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo]([]*InterFreqCarrierFreqInfo{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqCarrierFreqList", err)
	}
	return nil
}

func (ie *InterFreqCarrierFreqList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo]([]*InterFreqCarrierFreqInfo{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *InterFreqCarrierFreqInfo {
		return new(InterFreqCarrierFreqInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqCarrierFreqList", err)
	}
	ie.Value = []InterFreqCarrierFreqInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
