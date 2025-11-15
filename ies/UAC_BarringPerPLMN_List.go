package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringPerPLMN_List struct {
	Value []UAC_BarringPerPLMN `lb:1,ub:maxPLMN,madatory`
}

func (ie *UAC_BarringPerPLMN_List) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringPerPLMN]([]*UAC_BarringPerPLMN{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode UAC_BarringPerPLMN_List", err)
	}
	return nil
}

func (ie *UAC_BarringPerPLMN_List) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*UAC_BarringPerPLMN]([]*UAC_BarringPerPLMN{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	fn := func() *UAC_BarringPerPLMN {
		return new(UAC_BarringPerPLMN)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode UAC_BarringPerPLMN_List", err)
	}
	ie.Value = []UAC_BarringPerPLMN{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
