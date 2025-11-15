package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqList_v1720 struct {
	Value []InterFreqCarrierFreqInfo_v1720 `lb:1,ub:maxFreq,madatory`
}

func (ie *InterFreqCarrierFreqList_v1720) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1720]([]*InterFreqCarrierFreqInfo_v1720{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreqCarrierFreqList_v1720", err)
	}
	return nil
}

func (ie *InterFreqCarrierFreqList_v1720) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*InterFreqCarrierFreqInfo_v1720]([]*InterFreqCarrierFreqInfo_v1720{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *InterFreqCarrierFreqInfo_v1720 {
		return new(InterFreqCarrierFreqInfo_v1720)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode InterFreqCarrierFreqList_v1720", err)
	}
	ie.Value = []InterFreqCarrierFreqInfo_v1720{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
