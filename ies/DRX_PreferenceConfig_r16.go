package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_PreferenceConfig_r16 struct {
	drx_PreferenceProhibitTimer_r16 DRX_PreferenceConfig_r16_drx_PreferenceProhibitTimer_r16 `madatory`
}

func (ie *DRX_PreferenceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.drx_PreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode drx_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *DRX_PreferenceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.drx_PreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode drx_PreferenceProhibitTimer_r16", err)
	}
	return nil
}
