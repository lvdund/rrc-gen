package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_LeaveAssistanceConfig_r17 struct {
	musim_LeaveWithoutResponseTimer_r17 MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17 `madatory`
}

func (ie *MUSIM_LeaveAssistanceConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.musim_LeaveWithoutResponseTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode musim_LeaveWithoutResponseTimer_r17", err)
	}
	return nil
}

func (ie *MUSIM_LeaveAssistanceConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.musim_LeaveWithoutResponseTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode musim_LeaveWithoutResponseTimer_r17", err)
	}
	return nil
}
