package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarConnEstFailReport_r16 struct {
	connEstFailReport_r16 ConnEstFailReport_r16 `madatory`
	plmn_Identity_r16     PLMN_Identity         `madatory`
}

func (ie *VarConnEstFailReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.connEstFailReport_r16.Encode(w); err != nil {
		return utils.WrapError("Encode connEstFailReport_r16", err)
	}
	if err = ie.plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_Identity_r16", err)
	}
	return nil
}

func (ie *VarConnEstFailReport_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.connEstFailReport_r16.Decode(r); err != nil {
		return utils.WrapError("Decode connEstFailReport_r16", err)
	}
	if err = ie.plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_Identity_r16", err)
	}
	return nil
}
