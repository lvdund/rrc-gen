package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_PositionQCL_CellsToAddMod_r16 struct {
	physCellId_r16      PhysCellId                   `madatory`
	ssb_PositionQCL_r16 SSB_PositionQCL_Relation_r16 `madatory`
}

func (ie *SSB_PositionQCL_CellsToAddMod_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.physCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r16", err)
	}
	if err = ie.ssb_PositionQCL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_PositionQCL_r16", err)
	}
	return nil
}

func (ie *SSB_PositionQCL_CellsToAddMod_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.physCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r16", err)
	}
	if err = ie.ssb_PositionQCL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_PositionQCL_r16", err)
	}
	return nil
}
