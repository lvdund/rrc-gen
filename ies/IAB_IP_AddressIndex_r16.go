package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressIndex_r16 struct {
	Value uint64 `lb:1,ub:maxIAB_IP_Address_r16,madatory`
}

func (ie *IAB_IP_AddressIndex_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false); err != nil {
		return utils.WrapError("Encode IAB_IP_AddressIndex_r16", err)
	}
	return nil
}

func (ie *IAB_IP_AddressIndex_r16) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false); err != nil {
		return utils.WrapError("Decode IAB_IP_AddressIndex_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
