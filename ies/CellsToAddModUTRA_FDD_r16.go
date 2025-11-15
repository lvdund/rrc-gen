package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddModUTRA_FDD_r16 struct {
	cellIndexUTRA_FDD_r16 UTRA_FDD_CellIndex_r16 `madatory`
	physCellId_r16        PhysCellIdUTRA_FDD_r16 `madatory`
}

func (ie *CellsToAddModUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellIndexUTRA_FDD_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cellIndexUTRA_FDD_r16", err)
	}
	if err = ie.physCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r16", err)
	}
	return nil
}

func (ie *CellsToAddModUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellIndexUTRA_FDD_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cellIndexUTRA_FDD_r16", err)
	}
	if err = ie.physCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r16", err)
	}
	return nil
}
