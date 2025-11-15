package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BCCH_Config_modificationPeriodCoeff_Enum_n2  uper.Enumerated = 0
	BCCH_Config_modificationPeriodCoeff_Enum_n4  uper.Enumerated = 1
	BCCH_Config_modificationPeriodCoeff_Enum_n8  uper.Enumerated = 2
	BCCH_Config_modificationPeriodCoeff_Enum_n16 uper.Enumerated = 3
)

type BCCH_Config_modificationPeriodCoeff struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *BCCH_Config_modificationPeriodCoeff) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode BCCH_Config_modificationPeriodCoeff", err)
	}
	return nil
}

func (ie *BCCH_Config_modificationPeriodCoeff) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode BCCH_Config_modificationPeriodCoeff", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
