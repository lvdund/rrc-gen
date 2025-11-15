package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MultiPDSCH_TDRA_List_r17 struct {
	Value []MultiPDSCH_TDRA_r17 `lb:1,ub:maxNrofDL_AllocationsExt_r17,madatory`
}

func (ie *MultiPDSCH_TDRA_List_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MultiPDSCH_TDRA_r17]([]*MultiPDSCH_TDRA_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDL_AllocationsExt_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MultiPDSCH_TDRA_List_r17", err)
	}
	return nil
}

func (ie *MultiPDSCH_TDRA_List_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MultiPDSCH_TDRA_r17]([]*MultiPDSCH_TDRA_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofDL_AllocationsExt_r17}, false)
	fn := func() *MultiPDSCH_TDRA_r17 {
		return new(MultiPDSCH_TDRA_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MultiPDSCH_TDRA_List_r17", err)
	}
	ie.Value = []MultiPDSCH_TDRA_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
