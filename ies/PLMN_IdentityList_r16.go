package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_IdentityList_r16 struct {
	Value []PLMN_Identity `lb:1,ub:maxPLMN,madatory`
}

func (ie *PLMN_IdentityList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PLMN_IdentityList_r16", err)
	}
	return nil
}

func (ie *PLMN_IdentityList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	fn := func() *PLMN_Identity {
		return new(PLMN_Identity)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PLMN_IdentityList_r16", err)
	}
	ie.Value = []PLMN_Identity{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
