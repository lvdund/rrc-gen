package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_Cell struct {
	cellIndexEUTRA       EUTRA_CellIndex     `madatory`
	physCellId           EUTRA_PhysCellId    `madatory`
	cellIndividualOffset EUTRA_Q_OffsetRange `madatory`
}

func (ie *EUTRA_Cell) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellIndexEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode cellIndexEUTRA", err)
	}
	if err = ie.physCellId.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId", err)
	}
	if err = ie.cellIndividualOffset.Encode(w); err != nil {
		return utils.WrapError("Encode cellIndividualOffset", err)
	}
	return nil
}

func (ie *EUTRA_Cell) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellIndexEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode cellIndexEUTRA", err)
	}
	if err = ie.physCellId.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId", err)
	}
	if err = ie.cellIndividualOffset.Decode(r); err != nil {
		return utils.WrapError("Decode cellIndividualOffset", err)
	}
	return nil
}
