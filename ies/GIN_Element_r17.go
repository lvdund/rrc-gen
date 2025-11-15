package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GIN_Element_r17 struct {
	plmn_Identity_r17 PLMN_Identity `madatory`
	nid_List_r17      []NID_r16     `lb:1,ub:maxGIN_r17,madatory`
}

func (ie *GIN_Element_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.plmn_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_Identity_r17", err)
	}
	tmp_nid_List_r17 := utils.NewSequence[*NID_r16]([]*NID_r16{}, uper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
	for _, i := range ie.nid_List_r17 {
		tmp_nid_List_r17.Value = append(tmp_nid_List_r17.Value, &i)
	}
	if err = tmp_nid_List_r17.Encode(w); err != nil {
		return utils.WrapError("Encode nid_List_r17", err)
	}
	return nil
}

func (ie *GIN_Element_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.plmn_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_Identity_r17", err)
	}
	tmp_nid_List_r17 := utils.NewSequence[*NID_r16]([]*NID_r16{}, uper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
	fn_nid_List_r17 := func() *NID_r16 {
		return new(NID_r16)
	}
	if err = tmp_nid_List_r17.Decode(r, fn_nid_List_r17); err != nil {
		return utils.WrapError("Decode nid_List_r17", err)
	}
	ie.nid_List_r17 = []NID_r16{}
	for _, i := range tmp_nid_List_r17.Value {
		ie.nid_List_r17 = append(ie.nid_List_r17, *i)
	}
	return nil
}
