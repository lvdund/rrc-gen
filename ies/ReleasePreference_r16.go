package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReleasePreference_r16 struct {
	preferredRRC_State_r16 ReleasePreference_r16_preferredRRC_State_r16 `madatory`
}

func (ie *ReleasePreference_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.preferredRRC_State_r16.Encode(w); err != nil {
		return utils.WrapError("Encode preferredRRC_State_r16", err)
	}
	return nil
}

func (ie *ReleasePreference_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.preferredRRC_State_r16.Decode(r); err != nil {
		return utils.WrapError("Decode preferredRRC_State_r16", err)
	}
	return nil
}
