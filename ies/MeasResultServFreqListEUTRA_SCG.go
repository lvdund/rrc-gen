package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServFreqListEUTRA_SCG struct {
	Value []MeasResult2EUTRA `lb:1,ub:maxNrofServingCellsEUTRA,madatory`
}

func (ie *MeasResultServFreqListEUTRA_SCG) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2EUTRA]([]*MeasResult2EUTRA{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsEUTRA}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultServFreqListEUTRA_SCG", err)
	}
	return nil
}

func (ie *MeasResultServFreqListEUTRA_SCG) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2EUTRA]([]*MeasResult2EUTRA{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsEUTRA}, false)
	fn := func() *MeasResult2EUTRA {
		return new(MeasResult2EUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultServFreqListEUTRA_SCG", err)
	}
	ie.Value = []MeasResult2EUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
