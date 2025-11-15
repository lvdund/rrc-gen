package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreferenceConfig_r16 struct {
	minSchedulingOffsetPreferenceProhibitTimer_r16 MinSchedulingOffsetPreferenceConfig_r16_minSchedulingOffsetPreferenceProhibitTimer_r16 `madatory`
}

func (ie *MinSchedulingOffsetPreferenceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.minSchedulingOffsetPreferenceProhibitTimer_r16.Encode(w); err != nil {
		return utils.WrapError("Encode minSchedulingOffsetPreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MinSchedulingOffsetPreferenceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.minSchedulingOffsetPreferenceProhibitTimer_r16.Decode(r); err != nil {
		return utils.WrapError("Decode minSchedulingOffsetPreferenceProhibitTimer_r16", err)
	}
	return nil
}
