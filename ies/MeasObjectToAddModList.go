package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectToAddModList struct {
	Value []MeasObjectToAddMod `lb:1,ub:maxNrofObjectId,madatory`
}

func (ie *MeasObjectToAddModList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasObjectToAddMod]([]*MeasObjectToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofObjectId}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasObjectToAddModList", err)
	}
	return nil
}

func (ie *MeasObjectToAddModList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasObjectToAddMod]([]*MeasObjectToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofObjectId}, false)
	fn := func() *MeasObjectToAddMod {
		return new(MeasObjectToAddMod)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasObjectToAddModList", err)
	}
	ie.Value = []MeasObjectToAddMod{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
