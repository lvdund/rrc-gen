package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_TimeDomainResourceAllocation_mappingType_Enum_typeA uper.Enumerated = 0
	PDSCH_TimeDomainResourceAllocation_mappingType_Enum_typeB uper.Enumerated = 1
)

type PDSCH_TimeDomainResourceAllocation_mappingType struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocation_mappingType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PDSCH_TimeDomainResourceAllocation_mappingType", err)
	}
	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocation_mappingType) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PDSCH_TimeDomainResourceAllocation_mappingType", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
