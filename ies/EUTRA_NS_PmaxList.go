package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_NS_PmaxList struct {
	Value []EUTRA_NS_PmaxValue `lb:1,ub:maxEUTRA_NS_Pmax,madatory`
}

func (ie *EUTRA_NS_PmaxList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_NS_PmaxValue]([]*EUTRA_NS_PmaxValue{}, uper.Constraint{Lb: 1, Ub: maxEUTRA_NS_Pmax}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode EUTRA_NS_PmaxList", err)
	}
	return nil
}

func (ie *EUTRA_NS_PmaxList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_NS_PmaxValue]([]*EUTRA_NS_PmaxValue{}, uper.Constraint{Lb: 1, Ub: maxEUTRA_NS_Pmax}, false)
	fn := func() *EUTRA_NS_PmaxValue {
		return new(EUTRA_NS_PmaxValue)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode EUTRA_NS_PmaxList", err)
	}
	ie.Value = []EUTRA_NS_PmaxValue{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
