package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarRLF_Report_r16 struct {
	rlf_Report_r16        RLF_Report_r16         `madatory`
	plmn_IdentityList_r16 PLMN_IdentityList2_r16 `madatory`
}

func (ie *VarRLF_Report_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rlf_Report_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rlf_Report_r16", err)
	}
	if err = ie.plmn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityList_r16", err)
	}
	return nil
}

func (ie *VarRLF_Report_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rlf_Report_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rlf_Report_r16", err)
	}
	if err = ie.plmn_IdentityList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_IdentityList_r16", err)
	}
	return nil
}
