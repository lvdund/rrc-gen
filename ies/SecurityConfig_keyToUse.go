package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityConfig_keyToUse_Enum_master    uper.Enumerated = 0
	SecurityConfig_keyToUse_Enum_secondary uper.Enumerated = 1
)

type SecurityConfig_keyToUse struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SecurityConfig_keyToUse) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SecurityConfig_keyToUse", err)
	}
	return nil
}

func (ie *SecurityConfig_keyToUse) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SecurityConfig_keyToUse", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
