package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddMod struct {
	physCellId           PhysCellId        `madatory`
	cellIndividualOffset Q_OffsetRangeList `madatory`
}

func (ie *CellsToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.physCellId.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId", err)
	}
	if err = ie.cellIndividualOffset.Encode(w); err != nil {
		return utils.WrapError("Encode cellIndividualOffset", err)
	}
	return nil
}

func (ie *CellsToAddMod) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.physCellId.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId", err)
	}
	if err = ie.cellIndividualOffset.Decode(r); err != nil {
		return utils.WrapError("Decode cellIndividualOffset", err)
	}
	return nil
}
