package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUSCH_Config_resourceAllocationDCI_0_2_r16_Enum_resourceAllocationType0 uper.Enumerated = 0
	PUSCH_Config_resourceAllocationDCI_0_2_r16_Enum_resourceAllocationType1 uper.Enumerated = 1
	PUSCH_Config_resourceAllocationDCI_0_2_r16_Enum_dynamicSwitch           uper.Enumerated = 2
)

type PUSCH_Config_resourceAllocationDCI_0_2_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUSCH_Config_resourceAllocationDCI_0_2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUSCH_Config_resourceAllocationDCI_0_2_r16", err)
	}
	return nil
}

func (ie *PUSCH_Config_resourceAllocationDCI_0_2_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUSCH_Config_resourceAllocationDCI_0_2_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
