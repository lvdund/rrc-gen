package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarRA_Report_r16 struct {
	ra_ReportList_r16     RA_ReportList_r16     `madatory`
	plmn_IdentityList_r16 PLMN_IdentityList_r16 `madatory`
}

func (ie *VarRA_Report_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ra_ReportList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ra_ReportList_r16", err)
	}
	if err = ie.plmn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityList_r16", err)
	}
	return nil
}

func (ie *VarRA_Report_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ra_ReportList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ra_ReportList_r16", err)
	}
	if err = ie.plmn_IdentityList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_IdentityList_r16", err)
	}
	return nil
}
