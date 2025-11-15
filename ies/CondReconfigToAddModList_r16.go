package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondReconfigToAddModList_r16 struct {
	Value []CondReconfigToAddMod_r16 `lb:1,ub:maxNrofCondCells_r16,madatory`
}

func (ie *CondReconfigToAddModList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CondReconfigToAddMod_r16]([]*CondReconfigToAddMod_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CondReconfigToAddModList_r16", err)
	}
	return nil
}

func (ie *CondReconfigToAddModList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CondReconfigToAddMod_r16]([]*CondReconfigToAddMod_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofCondCells_r16}, false)
	fn := func() *CondReconfigToAddMod_r16 {
		return new(CondReconfigToAddMod_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CondReconfigToAddModList_r16", err)
	}
	ie.Value = []CondReconfigToAddMod_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
