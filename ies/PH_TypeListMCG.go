package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PH_TypeListMCG struct {
	Value []PH_InfoMCG `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *PH_TypeListMCG) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PH_InfoMCG]([]*PH_InfoMCG{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PH_TypeListMCG", err)
	}
	return nil
}

func (ie *PH_TypeListMCG) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PH_InfoMCG]([]*PH_InfoMCG{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *PH_InfoMCG {
		return new(PH_InfoMCG)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PH_TypeListMCG", err)
	}
	ie.Value = []PH_InfoMCG{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
