package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultList2UTRA struct {
	Value []MeasResult2UTRA_FDD_r16 `lb:1,ub:maxFreq,madatory`
}

func (ie *MeasResultList2UTRA) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2UTRA_FDD_r16]([]*MeasResult2UTRA_FDD_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultList2UTRA", err)
	}
	return nil
}

func (ie *MeasResultList2UTRA) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResult2UTRA_FDD_r16]([]*MeasResult2UTRA_FDD_r16{}, uper.Constraint{Lb: 1, Ub: maxFreq}, false)
	fn := func() *MeasResult2UTRA_FDD_r16 {
		return new(MeasResult2UTRA_FDD_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultList2UTRA", err)
	}
	ie.Value = []MeasResult2UTRA_FDD_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
