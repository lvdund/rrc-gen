package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_FormatConfig_nrofSlots_Enum_n2 uper.Enumerated = 0
	PUCCH_FormatConfig_nrofSlots_Enum_n4 uper.Enumerated = 1
	PUCCH_FormatConfig_nrofSlots_Enum_n8 uper.Enumerated = 2
)

type PUCCH_FormatConfig_nrofSlots struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUCCH_FormatConfig_nrofSlots) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUCCH_FormatConfig_nrofSlots", err)
	}
	return nil
}

func (ie *PUCCH_FormatConfig_nrofSlots) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUCCH_FormatConfig_nrofSlots", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
