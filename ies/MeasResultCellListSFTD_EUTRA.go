package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCellListSFTD_EUTRA struct {
	Value []MeasResultSFTD_EUTRA `lb:1,ub:maxCellSFTD,madatory`
}

func (ie *MeasResultCellListSFTD_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultSFTD_EUTRA]([]*MeasResultSFTD_EUTRA{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultCellListSFTD_EUTRA", err)
	}
	return nil
}

func (ie *MeasResultCellListSFTD_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultSFTD_EUTRA]([]*MeasResultSFTD_EUTRA{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	fn := func() *MeasResultSFTD_EUTRA {
		return new(MeasResultSFTD_EUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultCellListSFTD_EUTRA", err)
	}
	ie.Value = []MeasResultSFTD_EUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
