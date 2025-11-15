package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddModListExt_v1710 struct {
	Value []CellsToAddModExt_v1710 `lb:1,ub:maxNrofCellMeas,madatory`
}

func (ie *CellsToAddModListExt_v1710) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CellsToAddModExt_v1710]([]*CellsToAddModExt_v1710{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CellsToAddModListExt_v1710", err)
	}
	return nil
}

func (ie *CellsToAddModListExt_v1710) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CellsToAddModExt_v1710]([]*CellsToAddModExt_v1710{}, uper.Constraint{Lb: 1, Ub: maxNrofCellMeas}, false)
	fn := func() *CellsToAddModExt_v1710 {
		return new(CellsToAddModExt_v1710)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CellsToAddModListExt_v1710", err)
	}
	ie.Value = []CellsToAddModExt_v1710{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
