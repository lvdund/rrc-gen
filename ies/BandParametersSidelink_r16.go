package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParametersSidelink_r16 struct {
	freqBandSidelink_r16 FreqBandIndicatorNR `madatory`
}

func (ie *BandParametersSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.freqBandSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode freqBandSidelink_r16", err)
	}
	return nil
}

func (ie *BandParametersSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.freqBandSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode freqBandSidelink_r16", err)
	}
	return nil
}
