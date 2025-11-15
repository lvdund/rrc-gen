package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_supportedDMRS_TypeDL_Enum_type1     uper.Enumerated = 0
	Phy_ParametersFRX_Diff_supportedDMRS_TypeDL_Enum_type1And2 uper.Enumerated = 1
)

type Phy_ParametersFRX_Diff_supportedDMRS_TypeDL struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *Phy_ParametersFRX_Diff_supportedDMRS_TypeDL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_supportedDMRS_TypeDL", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_supportedDMRS_TypeDL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_supportedDMRS_TypeDL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
