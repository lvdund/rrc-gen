package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_ToReleaseList_r17 struct {
	Value []MRB_Identity_r17 `lb:1,ub:maxMRB_r17,madatory`
}

func (ie *MRB_ToReleaseList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MRB_Identity_r17]([]*MRB_Identity_r17{}, uper.Constraint{Lb: 1, Ub: maxMRB_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MRB_ToReleaseList_r17", err)
	}
	return nil
}

func (ie *MRB_ToReleaseList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MRB_Identity_r17]([]*MRB_Identity_r17{}, uper.Constraint{Lb: 1, Ub: maxMRB_r17}, false)
	fn := func() *MRB_Identity_r17 {
		return new(MRB_Identity_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MRB_ToReleaseList_r17", err)
	}
	ie.Value = []MRB_Identity_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
