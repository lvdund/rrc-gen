package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_FreqNeighHSDN_CellList_r17 struct {
	Value []EUTRA_PhysCellIdRange `lb:1,ub:maxCellEUTRA,madatory`
}

func (ie *EUTRA_FreqNeighHSDN_CellList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_PhysCellIdRange]([]*EUTRA_PhysCellIdRange{}, uper.Constraint{Lb: 1, Ub: maxCellEUTRA}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode EUTRA_FreqNeighHSDN_CellList_r17", err)
	}
	return nil
}

func (ie *EUTRA_FreqNeighHSDN_CellList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_PhysCellIdRange]([]*EUTRA_PhysCellIdRange{}, uper.Constraint{Lb: 1, Ub: maxCellEUTRA}, false)
	fn := func() *EUTRA_PhysCellIdRange {
		return new(EUTRA_PhysCellIdRange)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode EUTRA_FreqNeighHSDN_CellList_r17", err)
	}
	ie.Value = []EUTRA_PhysCellIdRange{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
