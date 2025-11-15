package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB0 uper.Enumerated = 0
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB2 uper.Enumerated = 1
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB4 uper.Enumerated = 2
	RA_Prioritization_powerRampingStepHighPriority_Enum_dB6 uper.Enumerated = 3
)

type RA_Prioritization_powerRampingStepHighPriority struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RA_Prioritization_powerRampingStepHighPriority) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RA_Prioritization_powerRampingStepHighPriority", err)
	}
	return nil
}

func (ie *RA_Prioritization_powerRampingStepHighPriority) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RA_Prioritization_powerRampingStepHighPriority", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
