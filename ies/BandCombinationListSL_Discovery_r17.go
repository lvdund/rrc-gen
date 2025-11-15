package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationListSL_Discovery_r17 struct {
	Value []BandParametersSidelinkDiscovery_r17 `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandCombinationListSL_Discovery_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelinkDiscovery_r17]([]*BandParametersSidelinkDiscovery_r17{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationListSL_Discovery_r17", err)
	}
	return nil
}

func (ie *BandCombinationListSL_Discovery_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelinkDiscovery_r17]([]*BandParametersSidelinkDiscovery_r17{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *BandParametersSidelinkDiscovery_r17 {
		return new(BandParametersSidelinkDiscovery_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationListSL_Discovery_r17", err)
	}
	ie.Value = []BandParametersSidelinkDiscovery_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
