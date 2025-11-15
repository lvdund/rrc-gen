package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540_srs_CarrierSwitch_eutra struct {
	srs_SwitchingTimesListEUTRA []SRS_SwitchingTimeEUTRA `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_eutra) Encode(w *uper.UperWriter) error {
	var err error
	tmp_srs_SwitchingTimesListEUTRA := utils.NewSequence[*SRS_SwitchingTimeEUTRA]([]*SRS_SwitchingTimeEUTRA{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.srs_SwitchingTimesListEUTRA {
		tmp_srs_SwitchingTimesListEUTRA.Value = append(tmp_srs_SwitchingTimesListEUTRA.Value, &i)
	}
	if err = tmp_srs_SwitchingTimesListEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode srs_SwitchingTimesListEUTRA", err)
	}
	return nil
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_eutra) Decode(r *uper.UperReader) error {
	var err error
	tmp_srs_SwitchingTimesListEUTRA := utils.NewSequence[*SRS_SwitchingTimeEUTRA]([]*SRS_SwitchingTimeEUTRA{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_srs_SwitchingTimesListEUTRA := func() *SRS_SwitchingTimeEUTRA {
		return new(SRS_SwitchingTimeEUTRA)
	}
	if err = tmp_srs_SwitchingTimesListEUTRA.Decode(r, fn_srs_SwitchingTimesListEUTRA); err != nil {
		return utils.WrapError("Decode srs_SwitchingTimesListEUTRA", err)
	}
	ie.srs_SwitchingTimesListEUTRA = []SRS_SwitchingTimeEUTRA{}
	for _, i := range tmp_srs_SwitchingTimesListEUTRA.Value {
		ie.srs_SwitchingTimesListEUTRA = append(ie.srs_SwitchingTimesListEUTRA, *i)
	}
	return nil
}
