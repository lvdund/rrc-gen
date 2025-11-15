package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationParametersSidelinkNR_v1710 struct {
	Value []BandParametersSidelink_v1710 `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandCombinationParametersSidelinkNR_v1710) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelink_v1710]([]*BandParametersSidelink_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationParametersSidelinkNR_v1710", err)
	}
	return nil
}

func (ie *BandCombinationParametersSidelinkNR_v1710) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelink_v1710]([]*BandParametersSidelink_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *BandParametersSidelink_v1710 {
		return new(BandParametersSidelink_v1710)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationParametersSidelinkNR_v1710", err)
	}
	ie.Value = []BandParametersSidelink_v1710{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
