package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyInfoDL struct {
	absoluteFrequencySSB    *ARFCN_ValueNR           `optional`
	frequencyBandList       MultiFrequencyBandListNR `madatory`
	absoluteFrequencyPointA ARFCN_ValueNR            `madatory`
	scs_SpecificCarrierList []SCS_SpecificCarrier    `lb:1,ub:maxSCSs,madatory`
}

func (ie *FrequencyInfoDL) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.absoluteFrequencySSB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.absoluteFrequencySSB != nil {
		if err = ie.absoluteFrequencySSB.Encode(w); err != nil {
			return utils.WrapError("Encode absoluteFrequencySSB", err)
		}
	}
	if err = ie.frequencyBandList.Encode(w); err != nil {
		return utils.WrapError("Encode frequencyBandList", err)
	}
	if err = ie.absoluteFrequencyPointA.Encode(w); err != nil {
		return utils.WrapError("Encode absoluteFrequencyPointA", err)
	}
	tmp_scs_SpecificCarrierList := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
	for _, i := range ie.scs_SpecificCarrierList {
		tmp_scs_SpecificCarrierList.Value = append(tmp_scs_SpecificCarrierList.Value, &i)
	}
	if err = tmp_scs_SpecificCarrierList.Encode(w); err != nil {
		return utils.WrapError("Encode scs_SpecificCarrierList", err)
	}
	return nil
}

func (ie *FrequencyInfoDL) Decode(r *uper.UperReader) error {
	var err error
	var absoluteFrequencySSBPresent bool
	if absoluteFrequencySSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if absoluteFrequencySSBPresent {
		ie.absoluteFrequencySSB = new(ARFCN_ValueNR)
		if err = ie.absoluteFrequencySSB.Decode(r); err != nil {
			return utils.WrapError("Decode absoluteFrequencySSB", err)
		}
	}
	if err = ie.frequencyBandList.Decode(r); err != nil {
		return utils.WrapError("Decode frequencyBandList", err)
	}
	if err = ie.absoluteFrequencyPointA.Decode(r); err != nil {
		return utils.WrapError("Decode absoluteFrequencyPointA", err)
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
	return nil
}
