package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n16  uper.Enumerated = 0
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n32  uper.Enumerated = 1
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n48  uper.Enumerated = 2
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n64  uper.Enumerated = 3
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n80  uper.Enumerated = 4
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n96  uper.Enumerated = 5
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n112 uper.Enumerated = 6
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n128 uper.Enumerated = 7
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n144 uper.Enumerated = 8
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n160 uper.Enumerated = 9
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n176 uper.Enumerated = 10
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n192 uper.Enumerated = 11
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n208 uper.Enumerated = 12
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n224 uper.Enumerated = 13
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n240 uper.Enumerated = 14
	Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot_Enum_n256 uper.Enumerated = 15
)

type Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot", err)
	}
	return nil
}

func (ie *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
