package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_Failure_r16 struct {
	sl_DestinationIdentity_r16 SL_DestinationIdentity_r16    `madatory`
	sl_Failure_r16             SL_Failure_r16_sl_Failure_r16 `madatory`
}

func (ie *SL_Failure_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_DestinationIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_DestinationIdentity_r16", err)
	}
	if err = ie.sl_Failure_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_Failure_r16", err)
	}
	return nil
}

func (ie *SL_Failure_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_DestinationIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_DestinationIdentity_r16", err)
	}
	if err = ie.sl_Failure_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_Failure_r16", err)
	}
	return nil
}
