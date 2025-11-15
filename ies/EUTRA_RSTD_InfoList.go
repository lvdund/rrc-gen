package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_RSTD_InfoList struct {
	Value []EUTRA_RSTD_Info `lb:1,ub:maxInterRAT_RSTD_Freq,madatory`
}

func (ie *EUTRA_RSTD_InfoList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_RSTD_Info]([]*EUTRA_RSTD_Info{}, uper.Constraint{Lb: 1, Ub: maxInterRAT_RSTD_Freq}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode EUTRA_RSTD_InfoList", err)
	}
	return nil
}

func (ie *EUTRA_RSTD_InfoList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*EUTRA_RSTD_Info]([]*EUTRA_RSTD_Info{}, uper.Constraint{Lb: 1, Ub: maxInterRAT_RSTD_Freq}, false)
	fn := func() *EUTRA_RSTD_Info {
		return new(EUTRA_RSTD_Info)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode EUTRA_RSTD_InfoList", err)
	}
	ie.Value = []EUTRA_RSTD_Info{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
