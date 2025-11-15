package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRB_ToReleaseList struct {
	Value []DRB_Identity `lb:1,ub:maxDRB,madatory`
}

func (ie *DRB_ToReleaseList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode DRB_ToReleaseList", err)
	}
	return nil
}

func (ie *DRB_ToReleaseList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn := func() *DRB_Identity {
		return new(DRB_Identity)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode DRB_ToReleaseList", err)
	}
	ie.Value = []DRB_Identity{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
