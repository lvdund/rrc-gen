package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddModList struct {
	Value []CellsToAddMod `lb:1,ub:maxNrofCellMeas,madatory`
}

func (ie *CellsToAddModList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CellsToAddMod]([]*CellsToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CellsToAddModList", err)
	}
	return nil
}

func (ie *CellsToAddModList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CellsToAddMod]([]*CellsToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	fn := func() *CellsToAddMod {
		return new(CellsToAddMod)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CellsToAddModList", err)
	}
	ie.Value = []CellsToAddMod{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
