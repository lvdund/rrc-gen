package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UTRA_FDD_CellIndexList_r16 struct {
	Value []UTRA_FDD_CellIndex_r16 `lb:1,ub:maxCellMeasUTRA_FDD_r16,madatory`
}

func (ie *UTRA_FDD_CellIndexList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UTRA_FDD_CellIndex_r16]([]*UTRA_FDD_CellIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasUTRA_FDD_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UTRA_FDD_CellIndexList_r16", err)
	}
	return nil
}

func (ie *UTRA_FDD_CellIndexList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UTRA_FDD_CellIndex_r16]([]*UTRA_FDD_CellIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxCellMeasUTRA_FDD_r16}, false)
	fn := func() *UTRA_FDD_CellIndex_r16 {
		return new(UTRA_FDD_CellIndex_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UTRA_FDD_CellIndexList_r16", err)
	}
	ie.Value = []UTRA_FDD_CellIndex_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
