package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreferenceConfig_r16 struct {
	maxMIMO_LayerPreferenceProhibitTimer_r16 MaxMIMO_LayerPreferenceConfig_r16_maxMIMO_LayerPreferenceProhibitTimer_r16 `madatory`
}

func (ie *MaxMIMO_LayerPreferenceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxMIMO_LayerPreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxMIMO_LayerPreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxMIMO_LayerPreferenceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxMIMO_LayerPreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxMIMO_LayerPreferenceProhibitTimer_r16", err)
	}
	return nil
}
