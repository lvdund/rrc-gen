package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCellListSFTD_NR struct {
	Value []MeasResultCellSFTD_NR `lb:1,ub:maxCellSFTD,madatory`
}

func (ie *MeasResultCellListSFTD_NR) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultCellSFTD_NR]([]*MeasResultCellSFTD_NR{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultCellListSFTD_NR", err)
	}
	return nil
}

func (ie *MeasResultCellListSFTD_NR) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultCellSFTD_NR]([]*MeasResultCellSFTD_NR{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
	fn := func() *MeasResultCellSFTD_NR {
		return new(MeasResultCellSFTD_NR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultCellListSFTD_NR", err)
	}
	ie.Value = []MeasResultCellSFTD_NR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
