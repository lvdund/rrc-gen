package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_ExcludedCell struct {
	cellIndexEUTRA  EUTRA_CellIndex       `madatory`
	physCellIdRange EUTRA_PhysCellIdRange `madatory`
}

func (ie *EUTRA_ExcludedCell) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellIndexEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode cellIndexEUTRA", err)
	}
	if err = ie.physCellIdRange.Encode(w); err != nil {
		return utils.WrapError("Encode physCellIdRange", err)
	}
	return nil
}

func (ie *EUTRA_ExcludedCell) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellIndexEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode cellIndexEUTRA", err)
	}
	if err = ie.physCellIdRange.Decode(r); err != nil {
		return utils.WrapError("Decode physCellIdRange", err)
	}
	return nil
}
