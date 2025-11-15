package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n2  uper.Enumerated = 0
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n3  uper.Enumerated = 1
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n4  uper.Enumerated = 2
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n5  uper.Enumerated = 3
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n6  uper.Enumerated = 4
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n7  uper.Enumerated = 5
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n8  uper.Enumerated = 6
	PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730_Enum_n16 uper.Enumerated = 7
)

type PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730", err)
	}
	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
