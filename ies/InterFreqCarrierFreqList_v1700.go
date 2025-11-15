package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqList_v1700 struct {
	Value []InterFreqCarrierFreqInfo_v1700 `lb:1,ub:maxFreq,madatory`
}

func (ie *InterFreqCarrierFreqList_v1700) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1700]([]*InterFreqCarrierFreqInfo_v1700{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqCarrierFreqList_v1700", err)
	}
	return nil
}

func (ie *InterFreqCarrierFreqList_v1700) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1700]([]*InterFreqCarrierFreqInfo_v1700{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *InterFreqCarrierFreqInfo_v1700 {
		return new(InterFreqCarrierFreqInfo_v1700)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqCarrierFreqList_v1700", err)
	}
	ie.Value = []InterFreqCarrierFreqInfo_v1700{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
