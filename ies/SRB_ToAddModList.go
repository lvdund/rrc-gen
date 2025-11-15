package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRB_ToAddModList struct {
	Value []SRB_ToAddMod `lb:1,ub:2,madatory`
}

func (ie *SRB_ToAddModList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*SRB_ToAddMod]([]*SRB_ToAddMod{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode SRB_ToAddModList", err)
	}
	return nil
}

func (ie *SRB_ToAddModList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*SRB_ToAddMod]([]*SRB_ToAddMod{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	fn := func() *SRB_ToAddMod {
		return new(SRB_ToAddMod)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode SRB_ToAddModList", err)
	}
	ie.Value = []SRB_ToAddMod{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
