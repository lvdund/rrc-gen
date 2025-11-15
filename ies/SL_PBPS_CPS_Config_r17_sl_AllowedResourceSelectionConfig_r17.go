package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c1 uper.Enumerated = 0
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c2 uper.Enumerated = 1
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c3 uper.Enumerated = 2
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c4 uper.Enumerated = 3
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c5 uper.Enumerated = 4
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c6 uper.Enumerated = 5
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c7 uper.Enumerated = 6
)

type SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17", err)
	}
	return nil
}

func (ie *SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
