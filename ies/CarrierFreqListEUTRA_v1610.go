package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierFreqListEUTRA_v1610 struct {
	Value []CarrierFreqEUTRA_v1610 `lb:1,ub:maxEUTRA_Carrier,madatory`
}

func (ie *CarrierFreqListEUTRA_v1610) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CarrierFreqEUTRA_v1610]([]*CarrierFreqEUTRA_v1610{}, uper.Constraint{Lb: 1, Ub: maxEUTRA_Carrier}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreqListEUTRA_v1610", err)
	}
	return nil
}

func (ie *CarrierFreqListEUTRA_v1610) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CarrierFreqEUTRA_v1610]([]*CarrierFreqEUTRA_v1610{}, uper.Constraint{Lb: 1, Ub: maxEUTRA_Carrier}, false)
	fn := func() *CarrierFreqEUTRA_v1610 {
		return new(CarrierFreqEUTRA_v1610)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CarrierFreqListEUTRA_v1610", err)
	}
	ie.Value = []CarrierFreqEUTRA_v1610{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
