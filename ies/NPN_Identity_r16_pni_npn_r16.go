package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NPN_Identity_r16_pni_npn_r16 struct {
	plmn_Identity_r16    PLMN_Identity          `madatory`
	cag_IdentityList_r16 []CAG_IdentityInfo_r16 `lb:1,ub:maxNPN_r16,madatory`
}

func (ie *NPN_Identity_r16_pni_npn_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_Identity_r16", err)
	}
	tmp_cag_IdentityList_r16 := utils.NewSequence[*CAG_IdentityInfo_r16]([]*CAG_IdentityInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.cag_IdentityList_r16 {
		tmp_cag_IdentityList_r16.Value = append(tmp_cag_IdentityList_r16.Value, &i)
	}
	if err = tmp_cag_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cag_IdentityList_r16", err)
	}
	return nil
}

func (ie *NPN_Identity_r16_pni_npn_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_Identity_r16", err)
	}
	tmp_cag_IdentityList_r16 := utils.NewSequence[*CAG_IdentityInfo_r16]([]*CAG_IdentityInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn_cag_IdentityList_r16 := func() *CAG_IdentityInfo_r16 {
		return new(CAG_IdentityInfo_r16)
	}
	if err = tmp_cag_IdentityList_r16.Decode(r, fn_cag_IdentityList_r16); err != nil {
		return utils.WrapError("Decode cag_IdentityList_r16", err)
	}
	ie.cag_IdentityList_r16 = []CAG_IdentityInfo_r16{}
	for _, i := range tmp_cag_IdentityList_r16.Value {
		ie.cag_IdentityList_r16 = append(ie.cag_IdentityList_r16, *i)
	}
	return nil
}
