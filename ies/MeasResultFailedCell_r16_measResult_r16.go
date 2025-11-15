package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFailedCell_r16_measResult_r16 struct {
	cellResults_r16    MeasResultFailedCell_r16_measResult_r16_cellResults_r16    `madatory`
	rsIndexResults_r16 MeasResultFailedCell_r16_measResult_r16_rsIndexResults_r16 `madatory`
}

func (ie *MeasResultFailedCell_r16_measResult_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellResults_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cellResults_r16", err)
	}
	if err = ie.rsIndexResults_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rsIndexResults_r16", err)
	}
	return nil
}

func (ie *MeasResultFailedCell_r16_measResult_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellResults_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cellResults_r16", err)
	}
	if err = ie.rsIndexResults_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rsIndexResults_r16", err)
	}
	return nil
}
