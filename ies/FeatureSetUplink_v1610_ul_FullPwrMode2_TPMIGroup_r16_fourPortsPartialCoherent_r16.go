package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g0 uper.Enumerated = 0
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g1 uper.Enumerated = 1
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g2 uper.Enumerated = 2
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g3 uper.Enumerated = 3
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g4 uper.Enumerated = 4
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g5 uper.Enumerated = 5
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g6 uper.Enumerated = 6
)

type FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
