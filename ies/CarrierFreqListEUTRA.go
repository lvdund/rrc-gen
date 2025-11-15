package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierFreqListEUTRA struct {
	Value []CarrierFreqEUTRA `lb:1,ub:maxEUTRA_Carrier,madatory`
}

func (ie *CarrierFreqListEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CarrierFreqEUTRA]([]*CarrierFreqEUTRA{}, uper.Constraint{Lb: 1, Ub: maxEUTRA_Carrier}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreqListEUTRA", err)
	}
	return nil
}

func (ie *CarrierFreqListEUTRA) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CarrierFreqEUTRA]([]*CarrierFreqEUTRA{}, uper.Constraint{Lb: 1, Ub: maxEUTRA_Carrier}, false)
	fn := func() *CarrierFreqEUTRA {
		return new(CarrierFreqEUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CarrierFreqListEUTRA", err)
	}
	ie.Value = []CarrierFreqEUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
