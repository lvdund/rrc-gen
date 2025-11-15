package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UCI_OnPUSCH_ListDCI_0_2_r16 struct {
	Value []UCI_OnPUSCH_DCI_0_2_r16 `lb:1,ub:2,madatory`
}

func (ie *UCI_OnPUSCH_ListDCI_0_2_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UCI_OnPUSCH_DCI_0_2_r16]([]*UCI_OnPUSCH_DCI_0_2_r16{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UCI_OnPUSCH_ListDCI_0_2_r16", err)
	}
	return nil
}

func (ie *UCI_OnPUSCH_ListDCI_0_2_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UCI_OnPUSCH_DCI_0_2_r16]([]*UCI_OnPUSCH_DCI_0_2_r16{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	fn := func() *UCI_OnPUSCH_DCI_0_2_r16 {
		return new(UCI_OnPUSCH_DCI_0_2_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UCI_OnPUSCH_ListDCI_0_2_r16", err)
	}
	ie.Value = []UCI_OnPUSCH_DCI_0_2_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
