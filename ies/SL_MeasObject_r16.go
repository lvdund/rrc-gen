package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasObject_r16 struct {
	frequencyInfoSL_r16 ARFCN_ValueNR `madatory`
}

func (ie *SL_MeasObject_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.frequencyInfoSL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode frequencyInfoSL_r16", err)
	}
	return nil
}

func (ie *SL_MeasObject_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.frequencyInfoSL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode frequencyInfoSL_r16", err)
	}
	return nil
}
