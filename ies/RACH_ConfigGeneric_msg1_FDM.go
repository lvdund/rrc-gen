package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_msg1_FDM_Enum_one   uper.Enumerated = 0
	RACH_ConfigGeneric_msg1_FDM_Enum_two   uper.Enumerated = 1
	RACH_ConfigGeneric_msg1_FDM_Enum_four  uper.Enumerated = 2
	RACH_ConfigGeneric_msg1_FDM_Enum_eight uper.Enumerated = 3
)

type RACH_ConfigGeneric_msg1_FDM struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RACH_ConfigGeneric_msg1_FDM) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_msg1_FDM", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_msg1_FDM) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_msg1_FDM", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
