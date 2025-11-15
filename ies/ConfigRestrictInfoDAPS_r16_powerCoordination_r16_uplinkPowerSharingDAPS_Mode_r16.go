package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16_Enum_semi_static_mode1 uper.Enumerated = 0
	ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16_Enum_semi_static_mode2 uper.Enumerated = 1
	ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16_Enum_dynamic           uper.Enumerated = 2
)

type ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16", err)
	}
	return nil
}

func (ie *ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ConfigRestrictInfoDAPS_r16_powerCoordination_r16_uplinkPowerSharingDAPS_Mode_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
