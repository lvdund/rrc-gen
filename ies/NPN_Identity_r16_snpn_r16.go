package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NPN_Identity_r16_snpn_r16 struct {
	plmn_Identity_r16 PLMN_Identity `madatory`
	nid_List_r16      []NID_r16     `lb:1,ub:maxNPN_r16,madatory`
}

func (ie *NPN_Identity_r16_snpn_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_Identity_r16", err)
	}
	tmp_nid_List_r16 := utils.NewSequence[*NID_r16]([]*NID_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.nid_List_r16 {
		tmp_nid_List_r16.Value = append(tmp_nid_List_r16.Value, &i)
	}
	if err = tmp_nid_List_r16.Encode(w); err != nil {
		return utils.WrapError("Encode nid_List_r16", err)
	}
	return nil
}

func (ie *NPN_Identity_r16_snpn_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_Identity_r16", err)
	}
	tmp_nid_List_r16 := utils.NewSequence[*NID_r16]([]*NID_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn_nid_List_r16 := func() *NID_r16 {
		return new(NID_r16)
	}
	if err = tmp_nid_List_r16.Decode(r, fn_nid_List_r16); err != nil {
		return utils.WrapError("Decode nid_List_r16", err)
	}
	ie.nid_List_r16 = []NID_r16{}
	for _, i := range tmp_nid_List_r16.Value {
		ie.nid_List_r16 = append(ie.nid_List_r16, *i)
	}
	return nil
}
