package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_Reassembly_Enum_ms0    uper.Enumerated = 0
	T_Reassembly_Enum_ms5    uper.Enumerated = 1
	T_Reassembly_Enum_ms10   uper.Enumerated = 2
	T_Reassembly_Enum_ms15   uper.Enumerated = 3
	T_Reassembly_Enum_ms20   uper.Enumerated = 4
	T_Reassembly_Enum_ms25   uper.Enumerated = 5
	T_Reassembly_Enum_ms30   uper.Enumerated = 6
	T_Reassembly_Enum_ms35   uper.Enumerated = 7
	T_Reassembly_Enum_ms40   uper.Enumerated = 8
	T_Reassembly_Enum_ms45   uper.Enumerated = 9
	T_Reassembly_Enum_ms50   uper.Enumerated = 10
	T_Reassembly_Enum_ms55   uper.Enumerated = 11
	T_Reassembly_Enum_ms60   uper.Enumerated = 12
	T_Reassembly_Enum_ms65   uper.Enumerated = 13
	T_Reassembly_Enum_ms70   uper.Enumerated = 14
	T_Reassembly_Enum_ms75   uper.Enumerated = 15
	T_Reassembly_Enum_ms80   uper.Enumerated = 16
	T_Reassembly_Enum_ms85   uper.Enumerated = 17
	T_Reassembly_Enum_ms90   uper.Enumerated = 18
	T_Reassembly_Enum_ms95   uper.Enumerated = 19
	T_Reassembly_Enum_ms100  uper.Enumerated = 20
	T_Reassembly_Enum_ms110  uper.Enumerated = 21
	T_Reassembly_Enum_ms120  uper.Enumerated = 22
	T_Reassembly_Enum_ms130  uper.Enumerated = 23
	T_Reassembly_Enum_ms140  uper.Enumerated = 24
	T_Reassembly_Enum_ms150  uper.Enumerated = 25
	T_Reassembly_Enum_ms160  uper.Enumerated = 26
	T_Reassembly_Enum_ms170  uper.Enumerated = 27
	T_Reassembly_Enum_ms180  uper.Enumerated = 28
	T_Reassembly_Enum_ms190  uper.Enumerated = 29
	T_Reassembly_Enum_ms200  uper.Enumerated = 30
	T_Reassembly_Enum_spare1 uper.Enumerated = 31
)

type T_Reassembly struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *T_Reassembly) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode T_Reassembly", err)
	}
	return nil
}

func (ie *T_Reassembly) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode T_Reassembly", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
