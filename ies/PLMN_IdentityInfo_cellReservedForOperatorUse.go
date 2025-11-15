package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PLMN_IdentityInfo_cellReservedForOperatorUse_Enum_reserved    uper.Enumerated = 0
	PLMN_IdentityInfo_cellReservedForOperatorUse_Enum_notReserved uper.Enumerated = 1
)

type PLMN_IdentityInfo_cellReservedForOperatorUse struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PLMN_IdentityInfo_cellReservedForOperatorUse) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PLMN_IdentityInfo_cellReservedForOperatorUse", err)
	}
	return nil
}

func (ie *PLMN_IdentityInfo_cellReservedForOperatorUse) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PLMN_IdentityInfo_cellReservedForOperatorUse", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
