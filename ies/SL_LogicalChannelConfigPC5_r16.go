package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_LogicalChannelConfigPC5_r16 struct {
	sl_LogicalChannelIdentity_r16 LogicalChannelIdentity `madatory`
}

func (ie *SL_LogicalChannelConfigPC5_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_LogicalChannelIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_LogicalChannelIdentity_r16", err)
	}
	return nil
}

func (ie *SL_LogicalChannelConfigPC5_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_LogicalChannelIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_LogicalChannelIdentity_r16", err)
	}
	return nil
}
