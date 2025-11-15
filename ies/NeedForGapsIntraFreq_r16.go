package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsIntraFreq_r16 struct {
	servCellId_r16         ServCellIndex                                   `madatory`
	gapIndicationIntra_r16 NeedForGapsIntraFreq_r16_gapIndicationIntra_r16 `madatory`
}

func (ie *NeedForGapsIntraFreq_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.servCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode servCellId_r16", err)
	}
	if err = ie.gapIndicationIntra_r16.Encode(w); err != nil {
		return utils.WrapError("Encode gapIndicationIntra_r16", err)
	}
	return nil
}

func (ie *NeedForGapsIntraFreq_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.servCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode servCellId_r16", err)
	}
	if err = ie.gapIndicationIntra_r16.Decode(r); err != nil {
		return utils.WrapError("Decode gapIndicationIntra_r16", err)
	}
	return nil
}
