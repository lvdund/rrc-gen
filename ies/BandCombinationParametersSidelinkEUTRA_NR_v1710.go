package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationParametersSidelinkEUTRA_NR_v1710 struct {
	Value []BandParametersSidelinkEUTRA_NR_v1710 `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1710) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelinkEUTRA_NR_v1710]([]*BandParametersSidelinkEUTRA_NR_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationParametersSidelinkEUTRA_NR_v1710", err)
	}
	return nil
}

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1710) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelinkEUTRA_NR_v1710]([]*BandParametersSidelinkEUTRA_NR_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *BandParametersSidelinkEUTRA_NR_v1710 {
		return new(BandParametersSidelinkEUTRA_NR_v1710)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationParametersSidelinkEUTRA_NR_v1710", err)
	}
	ie.Value = []BandParametersSidelinkEUTRA_NR_v1710{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
