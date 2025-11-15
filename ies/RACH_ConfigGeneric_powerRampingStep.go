package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigGeneric_powerRampingStep_Enum_dB0 uper.Enumerated = 0
	RACH_ConfigGeneric_powerRampingStep_Enum_dB2 uper.Enumerated = 1
	RACH_ConfigGeneric_powerRampingStep_Enum_dB4 uper.Enumerated = 2
	RACH_ConfigGeneric_powerRampingStep_Enum_dB6 uper.Enumerated = 3
)

type RACH_ConfigGeneric_powerRampingStep struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RACH_ConfigGeneric_powerRampingStep) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RACH_ConfigGeneric_powerRampingStep", err)
	}
	return nil
}

func (ie *RACH_ConfigGeneric_powerRampingStep) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RACH_ConfigGeneric_powerRampingStep", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
