package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellGlobalIdList_r16 struct {
	Value []CGI_Info_Logging_r16 `lb:1,ub:32,madatory`
}

func (ie *CellGlobalIdList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*CGI_Info_Logging_r16]([]*CGI_Info_Logging_r16{}, uper.Constraint{Lb: 1, Ub: 32}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CellGlobalIdList_r16", err)
	}
	return nil
}

func (ie *CellGlobalIdList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*CGI_Info_Logging_r16]([]*CGI_Info_Logging_r16{}, uper.Constraint{Lb: 1, Ub: 32}, false)
	fn := func() *CGI_Info_Logging_r16 {
		return new(CGI_Info_Logging_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CellGlobalIdList_r16", err)
	}
	ie.Value = []CGI_Info_Logging_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
