package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarSuccessHO_Report_r17_IEs struct {
	successHO_Report_r17  SuccessHO_Report_r17   `madatory`
	plmn_IdentityList_r17 PLMN_IdentityList2_r16 `madatory`
}

func (ie *VarSuccessHO_Report_r17_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.successHO_Report_r17.Encode(w); err != nil {
		return utils.WrapError("Encode successHO_Report_r17", err)
	}
	if err = ie.plmn_IdentityList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityList_r17", err)
	}
	return nil
}

func (ie *VarSuccessHO_Report_r17_IEs) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.successHO_Report_r17.Decode(r); err != nil {
		return utils.WrapError("Decode successHO_Report_r17", err)
	}
	if err = ie.plmn_IdentityList_r17.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_IdentityList_r17", err)
	}
	return nil
}
