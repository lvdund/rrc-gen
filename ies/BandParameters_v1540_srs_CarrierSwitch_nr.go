package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540_srs_CarrierSwitch_nr struct {
	srs_SwitchingTimesListNR []SRS_SwitchingTimeNR `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_nr) Encode(w *uper.UperWriter) error {
	var err error
	tmp_srs_SwitchingTimesListNR := utils.NewSequence[*SRS_SwitchingTimeNR]([]*SRS_SwitchingTimeNR{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.srs_SwitchingTimesListNR {
		tmp_srs_SwitchingTimesListNR.Value = append(tmp_srs_SwitchingTimesListNR.Value, &i)
	}
	if err = tmp_srs_SwitchingTimesListNR.Encode(w); err != nil {
		return utils.WrapError("Encode srs_SwitchingTimesListNR", err)
	}
	return nil
}

func (ie *BandParameters_v1540_srs_CarrierSwitch_nr) Decode(r *uper.UperReader) error {
	var err error
	tmp_srs_SwitchingTimesListNR := utils.NewSequence[*SRS_SwitchingTimeNR]([]*SRS_SwitchingTimeNR{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_srs_SwitchingTimesListNR := func() *SRS_SwitchingTimeNR {
		return new(SRS_SwitchingTimeNR)
	}
	if err = tmp_srs_SwitchingTimesListNR.Decode(r, fn_srs_SwitchingTimesListNR); err != nil {
		return utils.WrapError("Decode srs_SwitchingTimesListNR", err)
	}
	ie.srs_SwitchingTimesListNR = []SRS_SwitchingTimeNR{}
	for _, i := range tmp_srs_SwitchingTimesListNR.Value {
		ie.srs_SwitchingTimesListNR = append(ie.srs_SwitchingTimesListNR, *i)
	}
	return nil
}
