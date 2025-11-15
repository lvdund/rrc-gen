package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PowSav_ParametersFRX_Diff_r16 struct {
	maxBW_Preference_r16        *PowSav_ParametersFRX_Diff_r16_maxBW_Preference_r16        `optional`
	maxMIMO_LayerPreference_r16 *PowSav_ParametersFRX_Diff_r16_maxMIMO_LayerPreference_r16 `optional`
}

func (ie *PowSav_ParametersFRX_Diff_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.maxBW_Preference_r16 != nil, ie.maxMIMO_LayerPreference_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.maxBW_Preference_r16 != nil {
		if err = ie.maxBW_Preference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxBW_Preference_r16", err)
		}
	}
	if ie.maxMIMO_LayerPreference_r16 != nil {
		if err = ie.maxMIMO_LayerPreference_r16.Encode(w); err != nil {
			return utils.WrapError("Encode maxMIMO_LayerPreference_r16", err)
		}
	}
	return nil
}

func (ie *PowSav_ParametersFRX_Diff_r16) Decode(r *uper.UperReader) error {
	var err error
	var maxBW_Preference_r16Present bool
	if maxBW_Preference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxMIMO_LayerPreference_r16Present bool
	if maxMIMO_LayerPreference_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if maxBW_Preference_r16Present {
		ie.maxBW_Preference_r16 = new(PowSav_ParametersFRX_Diff_r16_maxBW_Preference_r16)
		if err = ie.maxBW_Preference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxBW_Preference_r16", err)
		}
	}
	if maxMIMO_LayerPreference_r16Present {
		ie.maxMIMO_LayerPreference_r16 = new(PowSav_ParametersFRX_Diff_r16_maxMIMO_LayerPreference_r16)
		if err = ie.maxMIMO_LayerPreference_r16.Decode(r); err != nil {
			return utils.WrapError("Decode maxMIMO_LayerPreference_r16", err)
		}
	}
	return nil
}
