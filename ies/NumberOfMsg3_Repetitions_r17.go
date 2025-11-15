package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NumberOfMsg3_Repetitions_r17_Enum_n1  uper.Enumerated = 0
	NumberOfMsg3_Repetitions_r17_Enum_n2  uper.Enumerated = 1
	NumberOfMsg3_Repetitions_r17_Enum_n3  uper.Enumerated = 2
	NumberOfMsg3_Repetitions_r17_Enum_n4  uper.Enumerated = 3
	NumberOfMsg3_Repetitions_r17_Enum_n7  uper.Enumerated = 4
	NumberOfMsg3_Repetitions_r17_Enum_n8  uper.Enumerated = 5
	NumberOfMsg3_Repetitions_r17_Enum_n12 uper.Enumerated = 6
	NumberOfMsg3_Repetitions_r17_Enum_n16 uper.Enumerated = 7
)

type NumberOfMsg3_Repetitions_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *NumberOfMsg3_Repetitions_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode NumberOfMsg3_Repetitions_r17", err)
	}
	return nil
}

func (ie *NumberOfMsg3_Repetitions_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode NumberOfMsg3_Repetitions_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
