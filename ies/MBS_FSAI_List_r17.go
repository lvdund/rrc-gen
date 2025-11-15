package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MBS_FSAI_List_r17 struct {
	Value []MBS_FSAI_r17 `lb:1,ub:maxFSAI_MBS_r17,madatory`
}

func (ie *MBS_FSAI_List_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MBS_FSAI_r17]([]*MBS_FSAI_r17{}, uper.Constraint{Lb: 1, Ub: maxFSAI_MBS_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MBS_FSAI_List_r17", err)
	}
	return nil
}

func (ie *MBS_FSAI_List_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MBS_FSAI_r17]([]*MBS_FSAI_r17{}, uper.Constraint{Lb: 1, Ub: maxFSAI_MBS_r17}, false)
	fn := func() *MBS_FSAI_r17 {
		return new(MBS_FSAI_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MBS_FSAI_List_r17", err)
	}
	ie.Value = []MBS_FSAI_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
