package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_TimeDomainResourceAllocationList struct {
	Value []PDSCH_TimeDomainResourceAllocation `lb:1,ub:maxNrofDL_Allocations,madatory`
}

func (ie *PDSCH_TimeDomainResourceAllocationList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PDSCH_TimeDomainResourceAllocation]([]*PDSCH_TimeDomainResourceAllocation{}, uper.Constraint{Lb: 1, Ub: maxNrofDL_Allocations}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PDSCH_TimeDomainResourceAllocationList", err)
	}
	return nil
}

func (ie *PDSCH_TimeDomainResourceAllocationList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PDSCH_TimeDomainResourceAllocation]([]*PDSCH_TimeDomainResourceAllocation{}, uper.Constraint{Lb: 1, Ub: maxNrofDL_Allocations}, false)
	fn := func() *PDSCH_TimeDomainResourceAllocation {
		return new(PDSCH_TimeDomainResourceAllocation)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PDSCH_TimeDomainResourceAllocationList", err)
	}
	ie.Value = []PDSCH_TimeDomainResourceAllocation{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
