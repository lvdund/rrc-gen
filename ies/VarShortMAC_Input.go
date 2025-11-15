package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarShortMAC_Input struct {
	sourcePhysCellId   PhysCellId   `madatory`
	targetCellIdentity CellIdentity `madatory`
	source_c_RNTI      RNTI_Value   `madatory`
}

func (ie *VarShortMAC_Input) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sourcePhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode sourcePhysCellId", err)
	}
	if err = ie.targetCellIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode targetCellIdentity", err)
	}
	if err = ie.source_c_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode source_c_RNTI", err)
	}
	return nil
}

func (ie *VarShortMAC_Input) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sourcePhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode sourcePhysCellId", err)
	}
	if err = ie.targetCellIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode targetCellIdentity", err)
	}
	if err = ie.source_c_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode source_c_RNTI", err)
	}
	return nil
}
