package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_resourceAllocation_Enum_resourceAllocationType0 uper.Enumerated = 0
	PDSCH_Config_resourceAllocation_Enum_resourceAllocationType1 uper.Enumerated = 1
	PDSCH_Config_resourceAllocation_Enum_dynamicSwitch           uper.Enumerated = 2
)

type PDSCH_Config_resourceAllocation struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PDSCH_Config_resourceAllocation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_resourceAllocation", err)
	}
	return nil
}

func (ie *PDSCH_Config_resourceAllocation) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_resourceAllocation", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
