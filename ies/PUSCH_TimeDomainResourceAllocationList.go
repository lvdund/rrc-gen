package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_TimeDomainResourceAllocationList struct {
	Value []PUSCH_TimeDomainResourceAllocation `lb:1,ub:maxNrofUL_Allocations,madatory`
}

func (ie *PUSCH_TimeDomainResourceAllocationList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PUSCH_TimeDomainResourceAllocation]([]*PUSCH_TimeDomainResourceAllocation{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PUSCH_TimeDomainResourceAllocationList", err)
	}
	return nil
}

func (ie *PUSCH_TimeDomainResourceAllocationList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PUSCH_TimeDomainResourceAllocation]([]*PUSCH_TimeDomainResourceAllocation{}, uper.Constraint{Lb: 1, Ub: maxNrofUL_Allocations}, false)
	fn := func() *PUSCH_TimeDomainResourceAllocation {
		return new(PUSCH_TimeDomainResourceAllocation)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PUSCH_TimeDomainResourceAllocationList", err)
	}
	ie.Value = []PUSCH_TimeDomainResourceAllocation{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
