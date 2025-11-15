package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ConfigurationList_r16 struct {
	Value []PUCCH_Config `lb:1,ub:2,madatory`
}

func (ie *PUCCH_ConfigurationList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*PUCCH_Config]([]*PUCCH_Config{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PUCCH_ConfigurationList_r16", err)
	}
	return nil
}

func (ie *PUCCH_ConfigurationList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*PUCCH_Config]([]*PUCCH_Config{}, uper.Constraint{Lb: 1, Ub: 2}, false)
	fn := func() *PUCCH_Config {
		return new(PUCCH_Config)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PUCCH_ConfigurationList_r16", err)
	}
	ie.Value = []PUCCH_Config{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
