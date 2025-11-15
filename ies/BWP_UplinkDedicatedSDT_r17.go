package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkDedicatedSDT_r17 struct {
	pusch_Config_r17                       *PUSCH_Config                           `optional,setuprelease`
	configuredGrantConfigToAddModList_r17  *ConfiguredGrantConfigToAddModList_r16  `optional`
	configuredGrantConfigToReleaseList_r17 *ConfiguredGrantConfigToReleaseList_r16 `optional`
}

func (ie *BWP_UplinkDedicatedSDT_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pusch_Config_r17 != nil, ie.configuredGrantConfigToAddModList_r17 != nil, ie.configuredGrantConfigToReleaseList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pusch_Config_r17 != nil {
		tmp_pusch_Config_r17 := utils.SetupRelease[*PUSCH_Config]{
			Setup: ie.pusch_Config_r17,
		}
		if err = tmp_pusch_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_Config_r17", err)
		}
	}
	if ie.configuredGrantConfigToAddModList_r17 != nil {
		if err = ie.configuredGrantConfigToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode configuredGrantConfigToAddModList_r17", err)
		}
	}
	if ie.configuredGrantConfigToReleaseList_r17 != nil {
		if err = ie.configuredGrantConfigToReleaseList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode configuredGrantConfigToReleaseList_r17", err)
		}
	}
	return nil
}

func (ie *BWP_UplinkDedicatedSDT_r17) Decode(r *uper.UperReader) error {
	var err error
	var pusch_Config_r17Present bool
	if pusch_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredGrantConfigToAddModList_r17Present bool
	if configuredGrantConfigToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredGrantConfigToReleaseList_r17Present bool
	if configuredGrantConfigToReleaseList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if pusch_Config_r17Present {
		tmp_pusch_Config_r17 := utils.SetupRelease[*PUSCH_Config]{}
		if err = tmp_pusch_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_Config_r17", err)
		}
		ie.pusch_Config_r17 = tmp_pusch_Config_r17.Setup
	}
	if configuredGrantConfigToAddModList_r17Present {
		ie.configuredGrantConfigToAddModList_r17 = new(ConfiguredGrantConfigToAddModList_r16)
		if err = ie.configuredGrantConfigToAddModList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode configuredGrantConfigToAddModList_r17", err)
		}
	}
	if configuredGrantConfigToReleaseList_r17Present {
		ie.configuredGrantConfigToReleaseList_r17 = new(ConfiguredGrantConfigToReleaseList_r16)
		if err = ie.configuredGrantConfigToReleaseList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode configuredGrantConfigToReleaseList_r17", err)
		}
	}
	return nil
}
