package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1730 struct {
	srs_SwitchingAffectedBandsListNR_r17 []SRS_SwitchingAffectedBandsNR_r17 `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandParameters_v1730) Encode(w *uper.UperWriter) error {
	var err error
	tmp_srs_SwitchingAffectedBandsListNR_r17 := utils.NewSequence[*SRS_SwitchingAffectedBandsNR_r17]([]*SRS_SwitchingAffectedBandsNR_r17{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.srs_SwitchingAffectedBandsListNR_r17 {
		tmp_srs_SwitchingAffectedBandsListNR_r17.Value = append(tmp_srs_SwitchingAffectedBandsListNR_r17.Value, &i)
	}
	if err = tmp_srs_SwitchingAffectedBandsListNR_r17.Encode(w); err != nil {
		return utils.WrapError("Encode srs_SwitchingAffectedBandsListNR_r17", err)
	}
	return nil
}

func (ie *BandParameters_v1730) Decode(r *uper.UperReader) error {
	var err error
	tmp_srs_SwitchingAffectedBandsListNR_r17 := utils.NewSequence[*SRS_SwitchingAffectedBandsNR_r17]([]*SRS_SwitchingAffectedBandsNR_r17{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_srs_SwitchingAffectedBandsListNR_r17 := func() *SRS_SwitchingAffectedBandsNR_r17 {
		return new(SRS_SwitchingAffectedBandsNR_r17)
	}
	if err = tmp_srs_SwitchingAffectedBandsListNR_r17.Decode(r, fn_srs_SwitchingAffectedBandsListNR_r17); err != nil {
		return utils.WrapError("Decode srs_SwitchingAffectedBandsListNR_r17", err)
	}
	ie.srs_SwitchingAffectedBandsListNR_r17 = []SRS_SwitchingAffectedBandsNR_r17{}
	for _, i := range tmp_srs_SwitchingAffectedBandsListNR_r17.Value {
		ie.srs_SwitchingAffectedBandsListNR_r17 = append(ie.srs_SwitchingAffectedBandsListNR_r17, *i)
	}
	return nil
}
