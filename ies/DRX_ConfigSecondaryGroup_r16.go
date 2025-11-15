package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigSecondaryGroup_r16 struct {
	drx_onDurationTimer_r16 DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16 `lb:1,ub:31,madatory`
	drx_InactivityTimer_r16 DRX_ConfigSecondaryGroup_r16_drx_InactivityTimer_r16 `madatory`
}

func (ie *DRX_ConfigSecondaryGroup_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.drx_onDurationTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode drx_onDurationTimer_r16", err)
	}
	if err = ie.drx_InactivityTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode drx_InactivityTimer_r16", err)
	}
	return nil
}

func (ie *DRX_ConfigSecondaryGroup_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.drx_onDurationTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode drx_onDurationTimer_r16", err)
	}
	if err = ie.drx_InactivityTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode drx_InactivityTimer_r16", err)
	}
	return nil
}
