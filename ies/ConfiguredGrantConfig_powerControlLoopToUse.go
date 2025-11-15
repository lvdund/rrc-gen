package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_powerControlLoopToUse_Enum_n0 uper.Enumerated = 0
	ConfiguredGrantConfig_powerControlLoopToUse_Enum_n1 uper.Enumerated = 1
)

type ConfiguredGrantConfig_powerControlLoopToUse struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ConfiguredGrantConfig_powerControlLoopToUse) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_powerControlLoopToUse", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_powerControlLoopToUse) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_powerControlLoopToUse", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
