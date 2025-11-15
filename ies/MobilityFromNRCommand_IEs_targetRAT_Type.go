package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_eutra          uper.Enumerated = 0
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_utra_fdd_v1610 uper.Enumerated = 1
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_spare2         uper.Enumerated = 2
	MobilityFromNRCommand_IEs_targetRAT_Type_Enum_spare1         uper.Enumerated = 3
)

type MobilityFromNRCommand_IEs_targetRAT_Type struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MobilityFromNRCommand_IEs_targetRAT_Type) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MobilityFromNRCommand_IEs_targetRAT_Type", err)
	}
	return nil
}

func (ie *MobilityFromNRCommand_IEs_targetRAT_Type) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MobilityFromNRCommand_IEs_targetRAT_Type", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
