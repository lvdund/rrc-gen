package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultListEUTRA struct {
	Value []MeasResultEUTRA `lb:1,ub:maxCellReport,madatory`
}

func (ie *MeasResultListEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultEUTRA]([]*MeasResultEUTRA{}, uper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListEUTRA", err)
	}
	return nil
}

func (ie *MeasResultListEUTRA) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultEUTRA]([]*MeasResultEUTRA{}, uper.Constraint{Lb: 1, Ub: maxCellReport}, false)
	fn := func() *MeasResultEUTRA {
		return new(MeasResultEUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultListEUTRA", err)
	}
	ie.Value = []MeasResultEUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
