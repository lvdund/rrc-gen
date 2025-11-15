package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PathlossReferenceRSList_r16 struct {
	Value []PathlossReferenceRS_r16 `lb:1,ub:maxNrofSRS_PathlossReferenceRS_r16,madatory`
}

func (ie *PathlossReferenceRSList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PathlossReferenceRS_r16]([]*PathlossReferenceRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PathlossReferenceRS_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PathlossReferenceRSList_r16", err)
	}
	return nil
}

func (ie *PathlossReferenceRSList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PathlossReferenceRS_r16]([]*PathlossReferenceRS_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSRS_PathlossReferenceRS_r16}, false)
	fn := func() *PathlossReferenceRS_r16 {
		return new(PathlossReferenceRS_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PathlossReferenceRSList_r16", err)
	}
	ie.Value = []PathlossReferenceRS_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
