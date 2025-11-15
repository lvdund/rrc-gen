package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_CellIndexList struct {
	Value []EUTRA_CellIndex `lb:1,ub:maxCellMeasEUTRA,madatory`
}

func (ie *EUTRA_CellIndexList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_CellIndex]([]*EUTRA_CellIndex{}, uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode EUTRA_CellIndexList", err)
	}
	return nil
}

func (ie *EUTRA_CellIndexList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_CellIndex]([]*EUTRA_CellIndex{}, uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false)
	fn := func() *EUTRA_CellIndex {
		return new(EUTRA_CellIndex)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode EUTRA_CellIndexList", err)
	}
	ie.Value = []EUTRA_CellIndex{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
