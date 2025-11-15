package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_Resource_intraSlotFrequencyHopping_Enum_enabled uper.Enumerated = 0
)

type PUCCH_Resource_intraSlotFrequencyHopping struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PUCCH_Resource_intraSlotFrequencyHopping) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PUCCH_Resource_intraSlotFrequencyHopping", err)
	}
	return nil
}

func (ie *PUCCH_Resource_intraSlotFrequencyHopping) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PUCCH_Resource_intraSlotFrequencyHopping", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
