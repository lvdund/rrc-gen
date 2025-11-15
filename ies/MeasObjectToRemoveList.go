package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectToRemoveList struct {
	Value []MeasObjectId `lb:1,ub:maxNrofObjectId,madatory`
}

func (ie *MeasObjectToRemoveList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasObjectId]([]*MeasObjectId{}, uper.Constraint{Lb: 1, Ub: maxNrofObjectId}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasObjectToRemoveList", err)
	}
	return nil
}

func (ie *MeasObjectToRemoveList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasObjectId]([]*MeasObjectId{}, uper.Constraint{Lb: 1, Ub: maxNrofObjectId}, false)
	fn := func() *MeasObjectId {
		return new(MeasObjectId)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasObjectToRemoveList", err)
	}
	ie.Value = []MeasObjectId{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
