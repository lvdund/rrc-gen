package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqList_v1730 struct {
	Value []InterFreqCarrierFreqInfo_v1730 `lb:1,ub:maxFreq,madatory`
}

func (ie *InterFreqCarrierFreqList_v1730) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1730]([]*InterFreqCarrierFreqInfo_v1730{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqCarrierFreqList_v1730", err)
	}
	return nil
}

func (ie *InterFreqCarrierFreqList_v1730) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1730]([]*InterFreqCarrierFreqInfo_v1730{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *InterFreqCarrierFreqInfo_v1730 {
		return new(InterFreqCarrierFreqInfo_v1730)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqCarrierFreqList_v1730", err)
	}
	ie.Value = []InterFreqCarrierFreqInfo_v1730{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
