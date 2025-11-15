package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqBandList struct {
	Value []FreqBandInformation `lb:1,ub:maxBandsMRDC,madatory`
}

func (ie *FreqBandList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FreqBandInformation]([]*FreqBandInformation{}, uper.Constraint{Lb: 1, Ub: maxBandsMRDC}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FreqBandList", err)
	}
	return nil
}

func (ie *FreqBandList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FreqBandInformation]([]*FreqBandInformation{}, uper.Constraint{Lb: 1, Ub: maxBandsMRDC}, false)
	fn := func() *FreqBandInformation {
		return new(FreqBandInformation)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FreqBandList", err)
	}
	ie.Value = []FreqBandInformation{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
