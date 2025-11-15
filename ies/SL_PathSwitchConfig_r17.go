package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PathSwitchConfig_r17 struct {
	targetRelayUE_Identity_r17 SL_SourceIdentity_r17            `madatory`
	t420_r17                   SL_PathSwitchConfig_r17_t420_r17 `madatory`
}

func (ie *SL_PathSwitchConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.targetRelayUE_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode targetRelayUE_Identity_r17", err)
	}
	if err = ie.t420_r17.Encode(w); err != nil {
		return utils.WrapError("Encode t420_r17", err)
	}
	return nil
}

func (ie *SL_PathSwitchConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.targetRelayUE_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode targetRelayUE_Identity_r17", err)
	}
	if err = ie.t420_r17.Decode(r); err != nil {
		return utils.WrapError("Decode t420_r17", err)
	}
	return nil
}
