package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasResults_r16 struct {
	sl_MeasId_r16     SL_MeasId_r16     `madatory`
	sl_MeasResult_r16 SL_MeasResult_r16 `madatory`
}

func (ie *SL_MeasResults_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_MeasId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_MeasId_r16", err)
	}
	if err = ie.sl_MeasResult_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_MeasResult_r16", err)
	}
	return nil
}

func (ie *SL_MeasResults_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_MeasId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_MeasId_r16", err)
	}
	if err = ie.sl_MeasResult_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_MeasResult_r16", err)
	}
	return nil
}
