package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_TimeDomainResourceAllocation_r16 struct {
	k2_r16                  *int64                 `lb:0,ub:32,optional`
	puschAllocationList_r16 []PUSCH_Allocation_r16 `lb:1,ub:maxNrofMultiplePUSCHs_r16,madatory`
}

func (ie *PUSCH_TimeDomainResourceAllocation_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.k2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.k2_r16 != nil {
		if err = w.WriteInteger(*ie.k2_r16, &uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode k2_r16", err)
		}
	}
	tmp_puschAllocationList_r16 := utils.NewSequence[*PUSCH_Allocation_r16]([]*PUSCH_Allocation_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofMultiplePUSCHs_r16}, false)
	for _, i := range ie.puschAllocationList_r16 {
		tmp_puschAllocationList_r16.Value = append(tmp_puschAllocationList_r16.Value, &i)
	}
	if err = tmp_puschAllocationList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode puschAllocationList_r16", err)
	}
	return nil
}

func (ie *PUSCH_TimeDomainResourceAllocation_r16) Decode(r *uper.UperReader) error {
	var err error
	var k2_r16Present bool
	if k2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if k2_r16Present {
		var tmp_int_k2_r16 int64
		if tmp_int_k2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode k2_r16", err)
		}
		ie.k2_r16 = &tmp_int_k2_r16
	}
	tmp_puschAllocationList_r16 := utils.NewSequence[*PUSCH_Allocation_r16]([]*PUSCH_Allocation_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofMultiplePUSCHs_r16}, false)
	fn_puschAllocationList_r16 := func() *PUSCH_Allocation_r16 {
		return new(PUSCH_Allocation_r16)
	}
	if err = tmp_puschAllocationList_r16.Decode(r, fn_puschAllocationList_r16); err != nil {
		return utils.WrapError("Decode puschAllocationList_r16", err)
	}
	ie.puschAllocationList_r16 = []PUSCH_Allocation_r16{}
	for _, i := range tmp_puschAllocationList_r16.Value {
		ie.puschAllocationList_r16 = append(ie.puschAllocationList_r16, *i)
	}
	return nil
}
