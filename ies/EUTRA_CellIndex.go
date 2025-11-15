package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_CellIndex struct {
	Value uint64 `lb:1,ub:maxCellMeasEUTRA,madatory`
}

func (ie *EUTRA_CellIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false); err != nil {
		return utils.WrapError("Encode EUTRA_CellIndex", err)
	}
	return nil
}

func (ie *EUTRA_CellIndex) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxCellMeasEUTRA}, false); err != nil {
		return utils.WrapError("Decode EUTRA_CellIndex", err)
	}
	ie.Value = uint64(v)
	return nil
}
