package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_Offset_r16_Enum_ms0dot5  uper.Enumerated = 0
	T_Offset_r16_Enum_ms0dot75 uper.Enumerated = 1
	T_Offset_r16_Enum_ms1      uper.Enumerated = 2
	T_Offset_r16_Enum_ms1dot5  uper.Enumerated = 3
	T_Offset_r16_Enum_ms2      uper.Enumerated = 4
	T_Offset_r16_Enum_ms2dot5  uper.Enumerated = 5
	T_Offset_r16_Enum_ms3      uper.Enumerated = 6
	T_Offset_r16_Enum_spare1   uper.Enumerated = 7
)

type T_Offset_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *T_Offset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode T_Offset_r16", err)
	}
	return nil
}

func (ie *T_Offset_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode T_Offset_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
