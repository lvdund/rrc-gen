package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DAPS_UplinkPowerConfig_r16 struct {
	p_DAPS_Source_r16               P_Max                                                      `madatory`
	p_DAPS_Target_r16               P_Max                                                      `madatory`
	uplinkPowerSharingDAPS_Mode_r16 DAPS_UplinkPowerConfig_r16_uplinkPowerSharingDAPS_Mode_r16 `madatory`
}

func (ie *DAPS_UplinkPowerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.p_DAPS_Source_r16.Encode(w); err != nil {
		return utils.WrapError("Encode p_DAPS_Source_r16", err)
	}
	if err = ie.p_DAPS_Target_r16.Encode(w); err != nil {
		return utils.WrapError("Encode p_DAPS_Target_r16", err)
	}
	if err = ie.uplinkPowerSharingDAPS_Mode_r16.Encode(w); err != nil {
		return utils.WrapError("Encode uplinkPowerSharingDAPS_Mode_r16", err)
	}
	return nil
}

func (ie *DAPS_UplinkPowerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.p_DAPS_Source_r16.Decode(r); err != nil {
		return utils.WrapError("Decode p_DAPS_Source_r16", err)
	}
	if err = ie.p_DAPS_Target_r16.Decode(r); err != nil {
		return utils.WrapError("Decode p_DAPS_Target_r16", err)
	}
	if err = ie.uplinkPowerSharingDAPS_Mode_r16.Decode(r); err != nil {
		return utils.WrapError("Decode uplinkPowerSharingDAPS_Mode_r16", err)
	}
	return nil
}
