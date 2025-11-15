package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16_Enum_interRepetition uper.Enumerated = 0
	ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16_Enum_interSlot       uper.Enumerated = 1
)

type ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16", err)
	}
	return nil
}

func (ie *ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode ConfiguredGrantConfig_rrc_ConfiguredUplinkGrant_frequencyHoppingPUSCH_RepTypeB_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
