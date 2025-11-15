package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdToRemoveList struct {
	Value []MeasId `lb:1,ub:maxNrofMeasId,madatory`
}

func (ie *MeasIdToRemoveList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasId]([]*MeasId{}, uper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasIdToRemoveList", err)
	}
	return nil
}

func (ie *MeasIdToRemoveList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasId]([]*MeasId{}, uper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false)
	fn := func() *MeasId {
		return new(MeasId)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasIdToRemoveList", err)
	}
	ie.Value = []MeasId{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
