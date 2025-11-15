package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForNCSG_IntraFreq_r17 struct {
	servCellId_r17         ServCellIndex                                    `madatory`
	gapIndicationIntra_r17 NeedForNCSG_IntraFreq_r17_gapIndicationIntra_r17 `madatory`
}

func (ie *NeedForNCSG_IntraFreq_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.servCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode servCellId_r17", err)
	}
	if err = ie.gapIndicationIntra_r17.Encode(w); err != nil {
		return utils.WrapError("Encode gapIndicationIntra_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_IntraFreq_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.servCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode servCellId_r17", err)
	}
	if err = ie.gapIndicationIntra_r17.Decode(r); err != nil {
		return utils.WrapError("Decode gapIndicationIntra_r17", err)
	}
	return nil
}
