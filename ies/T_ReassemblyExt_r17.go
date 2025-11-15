package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_ReassemblyExt_r17_Enum_ms210  uper.Enumerated = 0
	T_ReassemblyExt_r17_Enum_ms220  uper.Enumerated = 1
	T_ReassemblyExt_r17_Enum_ms340  uper.Enumerated = 2
	T_ReassemblyExt_r17_Enum_ms350  uper.Enumerated = 3
	T_ReassemblyExt_r17_Enum_ms550  uper.Enumerated = 4
	T_ReassemblyExt_r17_Enum_ms1100 uper.Enumerated = 5
	T_ReassemblyExt_r17_Enum_ms1650 uper.Enumerated = 6
	T_ReassemblyExt_r17_Enum_ms2200 uper.Enumerated = 7
)

type T_ReassemblyExt_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *T_ReassemblyExt_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode T_ReassemblyExt_r17", err)
	}
	return nil
}

func (ie *T_ReassemblyExt_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode T_ReassemblyExt_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
