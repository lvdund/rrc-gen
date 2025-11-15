package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n1     uper.Enumerated = 0
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n2     uper.Enumerated = 1
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n3     uper.Enumerated = 2
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n4     uper.Enumerated = 3
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n7     uper.Enumerated = 4
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n8     uper.Enumerated = 5
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n12    uper.Enumerated = 6
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n16    uper.Enumerated = 7
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n20    uper.Enumerated = 8
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n24    uper.Enumerated = 9
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n28    uper.Enumerated = 10
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_n32    uper.Enumerated = 11
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare4 uper.Enumerated = 12
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare3 uper.Enumerated = 13
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare2 uper.Enumerated = 14
	PUSCH_Allocation_r16_numberOfRepetitionsExt_r17_Enum_spare1 uper.Enumerated = 15
)

type PUSCH_Allocation_r16_numberOfRepetitionsExt_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PUSCH_Allocation_r16_numberOfRepetitionsExt_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Allocation_r16_numberOfRepetitionsExt_r17", err)
	}
	return nil
}

func (ie *PUSCH_Allocation_r16_numberOfRepetitionsExt_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Allocation_r16_numberOfRepetitionsExt_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
