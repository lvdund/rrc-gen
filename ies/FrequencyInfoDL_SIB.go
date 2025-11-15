package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FrequencyInfoDL_SIB struct {
	frequencyBandList       MultiFrequencyBandListNR_SIB `madatory`
	offsetToPointA          int64                        `lb:0,ub:2199,madatory`
	scs_SpecificCarrierList []SCS_SpecificCarrier        `lb:1,ub:maxSCSs,madatory`
}

func (ie *FrequencyInfoDL_SIB) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.frequencyBandList.Encode(w); err != nil {
		return utils.WrapError("Encode frequencyBandList", err)
	}
	if err = w.WriteInteger(ie.offsetToPointA, &uper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("WriteInteger offsetToPointA", err)
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

func (ie *FrequencyInfoDL_SIB) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.frequencyBandList.Decode(r); err != nil {
		return utils.WrapError("Decode frequencyBandList", err)
	}
	var tmp_int_offsetToPointA int64
	if tmp_int_offsetToPointA, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2199}, false); err != nil {
		return utils.WrapError("ReadInteger offsetToPointA", err)
	}
	ie.offsetToPointA = tmp_int_offsetToPointA
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
