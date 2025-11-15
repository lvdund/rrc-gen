package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type WLAN_Name_r16 struct {
	Value []byte `lb:1,ub:32,madatory`
}

func (ie *WLAN_Name_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteOctetString(ie.Value, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("Encode WLAN_Name_r16", err)
	}
	return nil
}

func (ie *WLAN_Name_r16) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	if v, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("Decode WLAN_Name_r16", err)
	}
	ie.Value = v
	return nil
}
