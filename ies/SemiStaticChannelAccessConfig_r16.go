package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SemiStaticChannelAccessConfig_r16 struct {
	period SemiStaticChannelAccessConfig_r16_period `madatory`
}

func (ie *SemiStaticChannelAccessConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.period.Encode(w); err != nil {
		return utils.WrapError("Encode period", err)
	}
	return nil
}

func (ie *SemiStaticChannelAccessConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.period.Decode(r); err != nil {
		return utils.WrapError("Decode period", err)
	}
	return nil
}
