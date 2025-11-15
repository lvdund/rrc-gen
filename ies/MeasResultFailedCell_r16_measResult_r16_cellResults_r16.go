package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFailedCell_r16_measResult_r16_cellResults_r16 struct {
	resultsSSB_Cell_r16 MeasQuantityResults `madatory`
}

func (ie *MeasResultFailedCell_r16_measResult_r16_cellResults_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.resultsSSB_Cell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resultsSSB_Cell_r16", err)
	}
	return nil
}

func (ie *MeasResultFailedCell_r16_measResult_r16_cellResults_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.resultsSSB_Cell_r16.Decode(r); err != nil {
		return utils.WrapError("Decode resultsSSB_Cell_r16", err)
	}
	return nil
}
