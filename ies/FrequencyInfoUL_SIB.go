package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyInfoUL_SIB struct {
	frequencyBandList       *MultiFrequencyBandListNR_SIB             `optional`
	absoluteFrequencyPointA *ARFCN_ValueNR                            `optional`
	scs_SpecificCarrierList []SCS_SpecificCarrier                     `lb:1,ub:maxSCSs,madatory`
	p_Max                   *P_Max                                    `optional`
	frequencyShift7p5khz    *FrequencyInfoUL_SIB_frequencyShift7p5khz `optional`
}

func (ie *FrequencyInfoUL_SIB) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.frequencyBandList != nil, ie.absoluteFrequencyPointA != nil, ie.p_Max != nil, ie.frequencyShift7p5khz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.frequencyBandList != nil {
		if err = ie.frequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyBandList", err)
		}
	}
	if ie.absoluteFrequencyPointA != nil {
		if err = ie.absoluteFrequencyPointA.Encode(w); err != nil {
			return utils.WrapError("Encode absoluteFrequencyPointA", err)
		}
	}
	tmp_scs_SpecificCarrierList := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	for _, i := range ie.scs_SpecificCarrierList {
		tmp_scs_SpecificCarrierList.Value = append(tmp_scs_SpecificCarrierList.Value, &i)
	}
	if err = tmp_scs_SpecificCarrierList.Encode(w); err != nil {
		return utils.WrapError("Encode scs_SpecificCarrierList", err)
	}
	if ie.p_Max != nil {
		if err = ie.p_Max.Encode(w); err != nil {
			return utils.WrapError("Encode p_Max", err)
		}
	}
	if ie.frequencyShift7p5khz != nil {
		if err = ie.frequencyShift7p5khz.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyShift7p5khz", err)
		}
	}
	return nil
}

func (ie *FrequencyInfoUL_SIB) Decode(r *uper.UperReader) error {
	var err error
	var frequencyBandListPresent bool
	if frequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var absoluteFrequencyPointAPresent bool
	if absoluteFrequencyPointAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p_MaxPresent bool
	if p_MaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var frequencyShift7p5khzPresent bool
	if frequencyShift7p5khzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyBandListPresent {
		ie.frequencyBandList = new(MultiFrequencyBandListNR_SIB)
		if err = ie.frequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyBandList", err)
		}
	}
	if absoluteFrequencyPointAPresent {
		ie.absoluteFrequencyPointA = new(ARFCN_ValueNR)
		if err = ie.absoluteFrequencyPointA.Decode(r); err != nil {
			return utils.WrapError("Decode absoluteFrequencyPointA", err)
		}
	}
	tmp_scs_SpecificCarrierList := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	fn_scs_SpecificCarrierList := func() *SCS_SpecificCarrier {
		return new(SCS_SpecificCarrier)
	}
	if err = tmp_scs_SpecificCarrierList.Decode(r, fn_scs_SpecificCarrierList); err != nil {
		return utils.WrapError("Decode scs_SpecificCarrierList", err)
	}
	ie.scs_SpecificCarrierList = []SCS_SpecificCarrier{}
	for _, i := range tmp_scs_SpecificCarrierList.Value {
		ie.scs_SpecificCarrierList = append(ie.scs_SpecificCarrierList, *i)
	}
	if p_MaxPresent {
		ie.p_Max = new(P_Max)
		if err = ie.p_Max.Decode(r); err != nil {
			return utils.WrapError("Decode p_Max", err)
		}
	}
	if frequencyShift7p5khzPresent {
		ie.frequencyShift7p5khz = new(FrequencyInfoUL_SIB_frequencyShift7p5khz)
		if err = ie.frequencyShift7p5khz.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyShift7p5khz", err)
		}
	}
	return nil
}
