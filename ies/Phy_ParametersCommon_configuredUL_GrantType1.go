package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_configuredUL_GrantType1_Enum_supported uper.Enumerated = 0
)

type Phy_ParametersCommon_configuredUL_GrantType1 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersCommon_configuredUL_GrantType1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersCommon_configuredUL_GrantType1", err)
	}
	return nil
}

func (ie *Phy_ParametersCommon_configuredUL_GrantType1) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersCommon_configuredUL_GrantType1", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
