package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqSeparationClass_Enum_mhz800       uper.Enumerated = 0
	FreqSeparationClass_Enum_mhz1200      uper.Enumerated = 1
	FreqSeparationClass_Enum_mhz1400      uper.Enumerated = 2
	FreqSeparationClass_Enum_mhz400_v1650 uper.Enumerated = 3
	FreqSeparationClass_Enum_mhz600_v1650 uper.Enumerated = 4
)

type FreqSeparationClass struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FreqSeparationClass) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FreqSeparationClass", err)
	}
	return nil
}

func (ie *FreqSeparationClass) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FreqSeparationClass", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
