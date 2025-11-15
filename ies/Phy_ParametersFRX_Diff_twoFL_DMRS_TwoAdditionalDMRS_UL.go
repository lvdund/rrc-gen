package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL_Enum_supported uper.Enumerated = 0
)

type Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL", err)
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
