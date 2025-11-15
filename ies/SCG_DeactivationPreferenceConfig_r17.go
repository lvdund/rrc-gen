package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCG_DeactivationPreferenceConfig_r17 struct {
	scg_DeactivationPreferenceProhibitTimer_r17 SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17 `madatory`
}

func (ie *SCG_DeactivationPreferenceConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.scg_DeactivationPreferenceProhibitTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	return nil
}

func (ie *SCG_DeactivationPreferenceConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.scg_DeactivationPreferenceProhibitTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	return nil
}
