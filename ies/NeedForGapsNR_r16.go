package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsNR_r16 struct {
	bandNR_r16        FreqBandIndicatorNR                 `madatory`
	gapIndication_r16 NeedForGapsNR_r16_gapIndication_r16 `madatory`
}

func (ie *NeedForGapsNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.bandNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode bandNR_r16", err)
	}
	if err = ie.gapIndication_r16.Encode(w); err != nil {
		return utils.WrapError("Encode gapIndication_r16", err)
	}
	return nil
}

func (ie *NeedForGapsNR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.bandNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode bandNR_r16", err)
	}
	if err = ie.gapIndication_r16.Decode(r); err != nil {
		return utils.WrapError("Decode gapIndication_r16", err)
	}
	return nil
}
