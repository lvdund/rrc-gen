package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_StatusProhibit_v1610_Enum_ms1    uper.Enumerated = 0
	T_StatusProhibit_v1610_Enum_ms2    uper.Enumerated = 1
	T_StatusProhibit_v1610_Enum_ms3    uper.Enumerated = 2
	T_StatusProhibit_v1610_Enum_ms4    uper.Enumerated = 3
	T_StatusProhibit_v1610_Enum_spare4 uper.Enumerated = 4
	T_StatusProhibit_v1610_Enum_spare3 uper.Enumerated = 5
	T_StatusProhibit_v1610_Enum_spare2 uper.Enumerated = 6
	T_StatusProhibit_v1610_Enum_spare1 uper.Enumerated = 7
)

type T_StatusProhibit_v1610 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *T_StatusProhibit_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode T_StatusProhibit_v1610", err)
	}
	return nil
}

func (ie *T_StatusProhibit_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode T_StatusProhibit_v1610", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
