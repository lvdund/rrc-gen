package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServMOList struct {
	Value []MeasResultServMO `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *MeasResultServMOList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultServMO]([]*MeasResultServMO{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultServMOList", err)
	}
	return nil
}

func (ie *MeasResultServMOList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultServMO]([]*MeasResultServMO{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *MeasResultServMO {
		return new(MeasResultServMO)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultServMOList", err)
	}
	ie.Value = []MeasResultServMO{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
