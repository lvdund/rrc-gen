package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_PositionQCL_Cell_r17 struct {
	physCellId_r17      PhysCellId                   `madatory`
	ssb_PositionQCL_r17 SSB_PositionQCL_Relation_r17 `madatory`
}

func (ie *SSB_PositionQCL_Cell_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.physCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r17", err)
	}
	if err = ie.ssb_PositionQCL_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_PositionQCL_r17", err)
	}
	return nil
}

func (ie *SSB_PositionQCL_Cell_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.physCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r17", err)
	}
	if err = ie.ssb_PositionQCL_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_PositionQCL_r17", err)
	}
	return nil
}
