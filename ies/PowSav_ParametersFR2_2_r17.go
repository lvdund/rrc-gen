package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_ParametersFR2_2_r17 struct {
	maxBW_Preference_r17        *PowSav_ParametersFR2_2_r17_maxBW_Preference_r17        `optional`
	maxMIMO_LayerPreference_r17 *PowSav_ParametersFR2_2_r17_maxMIMO_LayerPreference_r17 `optional`
}

func (ie *PowSav_ParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxBW_Preference_r17 != nil, ie.maxMIMO_LayerPreference_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxBW_Preference_r17 != nil {
		if err = ie.maxBW_Preference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxBW_Preference_r17", err)
		}
	}
	if ie.maxMIMO_LayerPreference_r17 != nil {
		if err = ie.maxMIMO_LayerPreference_r17.Encode(w); err != nil {
			return utils.WrapError("Encode maxMIMO_LayerPreference_r17", err)
		}
	}
	return nil
}

func (ie *PowSav_ParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var maxBW_Preference_r17Present bool
	if maxBW_Preference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxMIMO_LayerPreference_r17Present bool
	if maxMIMO_LayerPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxBW_Preference_r17Present {
		ie.maxBW_Preference_r17 = new(PowSav_ParametersFR2_2_r17_maxBW_Preference_r17)
		if err = ie.maxBW_Preference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxBW_Preference_r17", err)
		}
	}
	if maxMIMO_LayerPreference_r17Present {
		ie.maxMIMO_LayerPreference_r17 = new(PowSav_ParametersFR2_2_r17_maxMIMO_LayerPreference_r17)
		if err = ie.maxMIMO_LayerPreference_r17.Decode(r); err != nil {
			return utils.WrapError("Decode maxMIMO_LayerPreference_r17", err)
		}
	}
	return nil
}
