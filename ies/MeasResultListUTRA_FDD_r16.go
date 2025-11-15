package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultListUTRA_FDD_r16 struct {
	Value []MeasResultUTRA_FDD_r16 `lb:1,ub:maxCellReport,madatory`
}

func (ie *MeasResultListUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultUTRA_FDD_r16]([]*MeasResultUTRA_FDD_r16{}, uper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListUTRA_FDD_r16", err)
	}
	return nil
}

func (ie *MeasResultListUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultUTRA_FDD_r16]([]*MeasResultUTRA_FDD_r16{}, uper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	fn := func() *MeasResultUTRA_FDD_r16 {
		return new(MeasResultUTRA_FDD_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultListUTRA_FDD_r16", err)
	}
	ie.Value = []MeasResultUTRA_FDD_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
