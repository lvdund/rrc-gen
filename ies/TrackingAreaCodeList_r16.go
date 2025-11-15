package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TrackingAreaCodeList_r16 struct {
	Value []TrackingAreaCode `lb:1,ub:8,madatory`
}

func (ie *TrackingAreaCodeList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*TrackingAreaCode]([]*TrackingAreaCode{}, uper.Constraint{Lb: 1, Ub: 8}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode TrackingAreaCodeList_r16", err)
	}
	return nil
}

func (ie *TrackingAreaCodeList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*TrackingAreaCode]([]*TrackingAreaCode{}, uper.Constraint{Lb: 1, Ub: 8}, false)
	fn := func() *TrackingAreaCode {
		return new(TrackingAreaCode)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode TrackingAreaCodeList_r16", err)
	}
	ie.Value = []TrackingAreaCode{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
