package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MultiFrequencyBandListNR struct {
	Value []FreqBandIndicatorNR `lb:1,ub:maxNrofMultiBands,madatory`
}

func (ie *MultiFrequencyBandListNR) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxNrofMultiBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MultiFrequencyBandListNR", err)
	}
	return nil
}

func (ie *MultiFrequencyBandListNR) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FreqBandIndicatorNR]([]*FreqBandIndicatorNR{}, uper.Constraint{Lb: 1, Ub: maxNrofMultiBands}, false)
	fn := func() *FreqBandIndicatorNR {
		return new(FreqBandIndicatorNR)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MultiFrequencyBandListNR", err)
	}
	ie.Value = []FreqBandIndicatorNR{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
