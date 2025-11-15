package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellListEUTRA_r16 struct {
	Value []EUTRA_PhysCellIdRange `lb:1,ub:maxCellMeasIdle_r16,madatory`
}

func (ie *CellListEUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_PhysCellIdRange]([]*EUTRA_PhysCellIdRange{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CellListEUTRA_r16", err)
	}
	return nil
}

func (ie *CellListEUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_PhysCellIdRange]([]*EUTRA_PhysCellIdRange{}, uper.Constraint{Lb: 1, Ub: maxCellMeasIdle_r16}, false)
	fn := func() *EUTRA_PhysCellIdRange {
		return new(EUTRA_PhysCellIdRange)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CellListEUTRA_r16", err)
	}
	ie.Value = []EUTRA_PhysCellIdRange{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
