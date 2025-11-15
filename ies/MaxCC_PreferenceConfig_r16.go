package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxCC_PreferenceConfig_r16 struct {
	maxCC_PreferenceProhibitTimer_r16 MaxCC_PreferenceConfig_r16_maxCC_PreferenceProhibitTimer_r16 `madatory`
}

func (ie *MaxCC_PreferenceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxCC_PreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxCC_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxCC_PreferenceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxCC_PreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxCC_PreferenceProhibitTimer_r16", err)
	}
	return nil
}
