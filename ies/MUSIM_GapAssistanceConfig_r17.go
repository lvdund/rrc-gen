package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_GapAssistanceConfig_r17 struct {
	musim_GapProhibitTimer_r17 MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17 `madatory`
}

func (ie *MUSIM_GapAssistanceConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.musim_GapProhibitTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode musim_GapProhibitTimer_r17", err)
	}
	return nil
}

func (ie *MUSIM_GapAssistanceConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.musim_GapProhibitTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode musim_GapProhibitTimer_r17", err)
	}
	return nil
}
