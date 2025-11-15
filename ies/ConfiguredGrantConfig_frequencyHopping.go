package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_frequencyHopping_Enum_intraSlot uper.Enumerated = 0
	ConfiguredGrantConfig_frequencyHopping_Enum_interSlot uper.Enumerated = 1
)

type ConfiguredGrantConfig_frequencyHopping struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ConfiguredGrantConfig_frequencyHopping) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_frequencyHopping", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_frequencyHopping) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_frequencyHopping", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
