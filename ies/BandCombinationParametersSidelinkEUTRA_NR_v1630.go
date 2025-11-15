package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationParametersSidelinkEUTRA_NR_v1630 struct {
	Value []BandParametersSidelinkEUTRA_NR_v1630 `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1630) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelinkEUTRA_NR_v1630]([]*BandParametersSidelinkEUTRA_NR_v1630{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationParametersSidelinkEUTRA_NR_v1630", err)
	}
	return nil
}

func (ie *BandCombinationParametersSidelinkEUTRA_NR_v1630) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelinkEUTRA_NR_v1630]([]*BandParametersSidelinkEUTRA_NR_v1630{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *BandParametersSidelinkEUTRA_NR_v1630 {
		return new(BandParametersSidelinkEUTRA_NR_v1630)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationParametersSidelinkEUTRA_NR_v1630", err)
	}
	ie.Value = []BandParametersSidelinkEUTRA_NR_v1630{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
