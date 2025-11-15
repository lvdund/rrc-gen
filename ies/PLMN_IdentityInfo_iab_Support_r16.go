package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PLMN_IdentityInfo_iab_Support_r16_Enum_true uper.Enumerated = 0
)

type PLMN_IdentityInfo_iab_Support_r16 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PLMN_IdentityInfo_iab_Support_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PLMN_IdentityInfo_iab_Support_r16", err)
	}
	return nil
}

func (ie *PLMN_IdentityInfo_iab_Support_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PLMN_IdentityInfo_iab_Support_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
