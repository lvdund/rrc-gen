package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TwoPUCCH_Grp_Configurations_r17 struct {
	primaryPUCCH_GroupConfig_r17   PUCCH_Group_Config_r17 `madatory`
	secondaryPUCCH_GroupConfig_r17 PUCCH_Group_Config_r17 `madatory`
}

func (ie *TwoPUCCH_Grp_Configurations_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.primaryPUCCH_GroupConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode primaryPUCCH_GroupConfig_r17", err)
	}
	if err = ie.secondaryPUCCH_GroupConfig_r17.Encode(w); err != nil {
		return utils.WrapError("Encode secondaryPUCCH_GroupConfig_r17", err)
	}
	return nil
}

func (ie *TwoPUCCH_Grp_Configurations_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.primaryPUCCH_GroupConfig_r17.Decode(r); err != nil {
		return utils.WrapError("Decode primaryPUCCH_GroupConfig_r17", err)
	}
	if err = ie.secondaryPUCCH_GroupConfig_r17.Decode(r); err != nil {
		return utils.WrapError("Decode secondaryPUCCH_GroupConfig_r17", err)
	}
	return nil
}
