package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxBW_PreferenceConfig_r16 struct {
	maxBW_PreferenceProhibitTimer_r16 MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16 `madatory`
}

func (ie *MaxBW_PreferenceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxBW_PreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxBW_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxBW_PreferenceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxBW_PreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxBW_PreferenceProhibitTimer_r16", err)
	}
	return nil
}
