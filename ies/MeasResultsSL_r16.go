package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsSL_r16 struct {
	measResultsListSL_r16 MeasResultsSL_r16_measResultsListSL_r16 `madatory`
}

func (ie *MeasResultsSL_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.measResultsListSL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultsListSL_r16", err)
	}
	return nil
}

func (ie *MeasResultsSL_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.measResultsListSL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResultsListSL_r16", err)
	}
	return nil
}
