package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRAInfoList_r16 struct {
	Value []PerRAInfo_r16 `lb:1,ub:200,madatory`
}

func (ie *PerRAInfoList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PerRAInfo_r16]([]*PerRAInfo_r16{}, uper.Constraint{Lb: 1, Ub: 200}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PerRAInfoList_r16", err)
	}
	return nil
}

func (ie *PerRAInfoList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PerRAInfo_r16]([]*PerRAInfo_r16{}, uper.Constraint{Lb: 1, Ub: 200}, false)
	fn := func() *PerRAInfo_r16 {
		return new(PerRAInfo_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PerRAInfoList_r16", err)
	}
	ie.Value = []PerRAInfo_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
