package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellGroupForSwitch_r16 struct {
	Value []ServCellIndex `lb:1,ub:16,madatory`
}

func (ie *CellGroupForSwitch_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CellGroupForSwitch_r16", err)
	}
	return nil
}

func (ie *CellGroupForSwitch_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: 16}, false)
	fn := func() *ServCellIndex {
		return new(ServCellIndex)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CellGroupForSwitch_r16", err)
	}
	ie.Value = []ServCellIndex{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
