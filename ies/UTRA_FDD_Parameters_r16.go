package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UTRA_FDD_Parameters_r16 struct {
	supportedBandListUTRA_FDD_r16 []SupportedBandUTRA_FDD_r16 `lb:1,ub:maxBandsUTRA_FDD_r16,madatory`
}

func (ie *UTRA_FDD_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp_supportedBandListUTRA_FDD_r16 := utils.NewSequence[*SupportedBandUTRA_FDD_r16]([]*SupportedBandUTRA_FDD_r16{}, uper.Constraint{Lb: 1, Ub: maxBandsUTRA_FDD_r16}, false)
	for _, i := range ie.supportedBandListUTRA_FDD_r16 {
		tmp_supportedBandListUTRA_FDD_r16.Value = append(tmp_supportedBandListUTRA_FDD_r16.Value, &i)
	}
	if err = tmp_supportedBandListUTRA_FDD_r16.Encode(w); err != nil {
		return utils.WrapError("Encode supportedBandListUTRA_FDD_r16", err)
	}
	return nil
}

func (ie *UTRA_FDD_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp_supportedBandListUTRA_FDD_r16 := utils.NewSequence[*SupportedBandUTRA_FDD_r16]([]*SupportedBandUTRA_FDD_r16{}, uper.Constraint{Lb: 1, Ub: maxBandsUTRA_FDD_r16}, false)
	fn_supportedBandListUTRA_FDD_r16 := func() *SupportedBandUTRA_FDD_r16 {
		return new(SupportedBandUTRA_FDD_r16)
	}
	if err = tmp_supportedBandListUTRA_FDD_r16.Decode(r, fn_supportedBandListUTRA_FDD_r16); err != nil {
		return utils.WrapError("Decode supportedBandListUTRA_FDD_r16", err)
	}
	ie.supportedBandListUTRA_FDD_r16 = []SupportedBandUTRA_FDD_r16{}
	for _, i := range tmp_supportedBandListUTRA_FDD_r16.Value {
		ie.supportedBandListUTRA_FDD_r16 = append(ie.supportedBandListUTRA_FDD_r16, *i)
	}
	return nil
}
