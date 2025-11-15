package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols_Enum_supported uper.Enumerated = 0
)

type Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
