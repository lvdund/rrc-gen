package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_Parameters struct {
	supportedBandListEUTRA   []FreqBandIndicatorEUTRA  `lb:1,ub:maxBandsEUTRA,madatory`
	eutra_ParametersCommon   *EUTRA_ParametersCommon   `optional`
	eutra_ParametersXDD_Diff *EUTRA_ParametersXDD_Diff `optional`
}

func (ie *EUTRA_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.eutra_ParametersCommon != nil, ie.eutra_ParametersXDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_supportedBandListEUTRA := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	for _, i := range ie.supportedBandListEUTRA {
		tmp_supportedBandListEUTRA.Value = append(tmp_supportedBandListEUTRA.Value, &i)
	}
	if err = tmp_supportedBandListEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode supportedBandListEUTRA", err)
	}
	if ie.eutra_ParametersCommon != nil {
		if err = ie.eutra_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_ParametersCommon", err)
		}
	}
	if ie.eutra_ParametersXDD_Diff != nil {
		if err = ie.eutra_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode eutra_ParametersXDD_Diff", err)
		}
	}
	return nil
}

func (ie *EUTRA_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var eutra_ParametersCommonPresent bool
	if eutra_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var eutra_ParametersXDD_DiffPresent bool
	if eutra_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_supportedBandListEUTRA := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, uper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	fn_supportedBandListEUTRA := func() *FreqBandIndicatorEUTRA {
		return new(FreqBandIndicatorEUTRA)
	}
	if err = tmp_supportedBandListEUTRA.Decode(r, fn_supportedBandListEUTRA); err != nil {
		return utils.WrapError("Decode supportedBandListEUTRA", err)
	}
	ie.supportedBandListEUTRA = []FreqBandIndicatorEUTRA{}
	for _, i := range tmp_supportedBandListEUTRA.Value {
		ie.supportedBandListEUTRA = append(ie.supportedBandListEUTRA, *i)
	}
	if eutra_ParametersCommonPresent {
		ie.eutra_ParametersCommon = new(EUTRA_ParametersCommon)
		if err = ie.eutra_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_ParametersCommon", err)
		}
	}
	if eutra_ParametersXDD_DiffPresent {
		ie.eutra_ParametersXDD_Diff = new(EUTRA_ParametersXDD_Diff)
		if err = ie.eutra_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode eutra_ParametersXDD_Diff", err)
		}
	}
	return nil
}
