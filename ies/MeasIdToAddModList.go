package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdToAddModList struct {
	Value []MeasIdToAddMod `lb:1,ub:maxNrofMeasId,madatory`
}

func (ie *MeasIdToAddModList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasIdToAddMod]([]*MeasIdToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasIdToAddModList", err)
	}
	return nil
}

func (ie *MeasIdToAddModList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasIdToAddMod]([]*MeasIdToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false)
	fn := func() *MeasIdToAddMod {
		return new(MeasIdToAddMod)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasIdToAddModList", err)
	}
	ie.Value = []MeasIdToAddMod{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
