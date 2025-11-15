package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol_Enum_n10 uper.Enumerated = 0
	Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol_Enum_n20 uper.Enumerated = 1
)

type Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol", err)
	}
	return nil
}

func (ie *Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFR1_pdsch_RE_MappingFR1_PerSymbol", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
