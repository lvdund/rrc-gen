package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Dummy_TDRA_List struct {
	Value []MultiPDSCH_TDRA_r17 `lb:1,ub:maxNrofDL_Allocations,madatory`
}

func (ie *Dummy_TDRA_List) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MultiPDSCH_TDRA_r17]([]*MultiPDSCH_TDRA_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDL_Allocations}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy_TDRA_List", err)
	}
	return nil
}

func (ie *Dummy_TDRA_List) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MultiPDSCH_TDRA_r17]([]*MultiPDSCH_TDRA_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDL_Allocations}, false)
	fn := func() *MultiPDSCH_TDRA_r17 {
		return new(MultiPDSCH_TDRA_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode Dummy_TDRA_List", err)
	}
	ie.Value = []MultiPDSCH_TDRA_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
