package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForNCSG_EUTRA_r17 struct {
	bandEUTRA_r17     FreqBandIndicatorEUTRA                  `madatory`
	gapIndication_r17 NeedForNCSG_EUTRA_r17_gapIndication_r17 `madatory`
}

func (ie *NeedForNCSG_EUTRA_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.bandEUTRA_r17.Encode(w); err != nil {
		return utils.WrapError("Encode bandEUTRA_r17", err)
	}
	if err = ie.gapIndication_r17.Encode(w); err != nil {
		return utils.WrapError("Encode gapIndication_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_EUTRA_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.bandEUTRA_r17.Decode(r); err != nil {
		return utils.WrapError("Decode bandEUTRA_r17", err)
	}
	if err = ie.gapIndication_r17.Decode(r); err != nil {
		return utils.WrapError("Decode gapIndication_r17", err)
	}
	return nil
}
