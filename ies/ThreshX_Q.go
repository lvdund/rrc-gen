package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ThreshX_Q struct {
	threshX_HighQ ReselectionThresholdQ `madatory`
	threshX_LowQ  ReselectionThresholdQ `madatory`
}

func (ie *ThreshX_Q) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.threshX_HighQ.Encode(w); err != nil {
		return utils.WrapError("Encode threshX_HighQ", err)
	}
	if err = ie.threshX_LowQ.Encode(w); err != nil {
		return utils.WrapError("Encode threshX_LowQ", err)
	}
	return nil
}

func (ie *ThreshX_Q) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.threshX_HighQ.Decode(r); err != nil {
		return utils.WrapError("Decode threshX_HighQ", err)
	}
	if err = ie.threshX_LowQ.Decode(r); err != nil {
		return utils.WrapError("Decode threshX_LowQ", err)
	}
	return nil
}
