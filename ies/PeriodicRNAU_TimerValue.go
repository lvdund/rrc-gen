package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PeriodicRNAU_TimerValue_Enum_min5   uper.Enumerated = 0
	PeriodicRNAU_TimerValue_Enum_min10  uper.Enumerated = 1
	PeriodicRNAU_TimerValue_Enum_min20  uper.Enumerated = 2
	PeriodicRNAU_TimerValue_Enum_min30  uper.Enumerated = 3
	PeriodicRNAU_TimerValue_Enum_min60  uper.Enumerated = 4
	PeriodicRNAU_TimerValue_Enum_min120 uper.Enumerated = 5
	PeriodicRNAU_TimerValue_Enum_min360 uper.Enumerated = 6
	PeriodicRNAU_TimerValue_Enum_min720 uper.Enumerated = 7
)

type PeriodicRNAU_TimerValue struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PeriodicRNAU_TimerValue) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PeriodicRNAU_TimerValue", err)
	}
	return nil
}

func (ie *PeriodicRNAU_TimerValue) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PeriodicRNAU_TimerValue", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
