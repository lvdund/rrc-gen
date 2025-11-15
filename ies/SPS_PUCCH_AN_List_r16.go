package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SPS_PUCCH_AN_List_r16 struct {
	Value []SPS_PUCCH_AN_r16 `lb:1,ub:4,madatory`
}

func (ie *SPS_PUCCH_AN_List_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SPS_PUCCH_AN_r16]([]*SPS_PUCCH_AN_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SPS_PUCCH_AN_List_r16", err)
	}
	return nil
}

func (ie *SPS_PUCCH_AN_List_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SPS_PUCCH_AN_r16]([]*SPS_PUCCH_AN_r16{}, uper.Constraint{Lb: 1, Ub: 4}, false)
	fn := func() *SPS_PUCCH_AN_r16 {
		return new(SPS_PUCCH_AN_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SPS_PUCCH_AN_List_r16", err)
	}
	ie.Value = []SPS_PUCCH_AN_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
