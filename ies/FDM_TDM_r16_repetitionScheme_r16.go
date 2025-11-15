package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FDM_TDM_r16_repetitionScheme_r16_Enum_fdmSchemeA uper.Enumerated = 0
	FDM_TDM_r16_repetitionScheme_r16_Enum_fdmSchemeB uper.Enumerated = 1
	FDM_TDM_r16_repetitionScheme_r16_Enum_tdmSchemeA uper.Enumerated = 2
)

type FDM_TDM_r16_repetitionScheme_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FDM_TDM_r16_repetitionScheme_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FDM_TDM_r16_repetitionScheme_r16", err)
	}
	return nil
}

func (ie *FDM_TDM_r16_repetitionScheme_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FDM_TDM_r16_repetitionScheme_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
