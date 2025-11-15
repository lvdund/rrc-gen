package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_bwp_SwitchingDelay_Enum_type1 uper.Enumerated = 0
	Phy_ParametersCommon_bwp_SwitchingDelay_Enum_type2 uper.Enumerated = 1
)

type Phy_ParametersCommon_bwp_SwitchingDelay struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *Phy_ParametersCommon_bwp_SwitchingDelay) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersCommon_bwp_SwitchingDelay", err)
	}
	return nil
}

func (ie *Phy_ParametersCommon_bwp_SwitchingDelay) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersCommon_bwp_SwitchingDelay", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
