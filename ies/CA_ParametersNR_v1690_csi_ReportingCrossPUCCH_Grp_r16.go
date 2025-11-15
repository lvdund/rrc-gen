package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16 struct {
	computationTimeForA_CSI_r16 CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_computationTimeForA_CSI_r16  `madatory`
	additionalSymbols_r16       *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_additionalSymbols_r16       `optional`
	sp_CSI_ReportingOnPUCCH_r16 *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUCCH_r16 `optional`
	sp_CSI_ReportingOnPUSCH_r16 *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUSCH_r16 `optional`
	carrierTypePairList_r16     []CarrierTypePair_r16                                                              `lb:1,ub:maxCarrierTypePairList_r16,madatory`
}

func (ie *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.additionalSymbols_r16 != nil, ie.sp_CSI_ReportingOnPUCCH_r16 != nil, ie.sp_CSI_ReportingOnPUSCH_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.computationTimeForA_CSI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode computationTimeForA_CSI_r16", err)
	}
	if ie.additionalSymbols_r16 != nil {
		if err = ie.additionalSymbols_r16.Encode(w); err != nil {
			return utils.WrapError("Encode additionalSymbols_r16", err)
		}
	}
	if ie.sp_CSI_ReportingOnPUCCH_r16 != nil {
		if err = ie.sp_CSI_ReportingOnPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_ReportingOnPUCCH_r16", err)
		}
	}
	if ie.sp_CSI_ReportingOnPUSCH_r16 != nil {
		if err = ie.sp_CSI_ReportingOnPUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_ReportingOnPUSCH_r16", err)
		}
	}
	tmp_carrierTypePairList_r16 := utils.NewSequence[*CarrierTypePair_r16]([]*CarrierTypePair_r16{}, uper.Constraint{Lb: 1, Ub: maxCarrierTypePairList_r16}, false)
	for _, i := range ie.carrierTypePairList_r16 {
		tmp_carrierTypePairList_r16.Value = append(tmp_carrierTypePairList_r16.Value, &i)
	}
	if err = tmp_carrierTypePairList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierTypePairList_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16) Decode(r *uper.UperReader) error {
	var err error
	var additionalSymbols_r16Present bool
	if additionalSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_ReportingOnPUCCH_r16Present bool
	if sp_CSI_ReportingOnPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_ReportingOnPUSCH_r16Present bool
	if sp_CSI_ReportingOnPUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.computationTimeForA_CSI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode computationTimeForA_CSI_r16", err)
	}
	if additionalSymbols_r16Present {
		ie.additionalSymbols_r16 = new(CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_additionalSymbols_r16)
		if err = ie.additionalSymbols_r16.Decode(r); err != nil {
			return utils.WrapError("Decode additionalSymbols_r16", err)
		}
	}
	if sp_CSI_ReportingOnPUCCH_r16Present {
		ie.sp_CSI_ReportingOnPUCCH_r16 = new(CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUCCH_r16)
		if err = ie.sp_CSI_ReportingOnPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_ReportingOnPUCCH_r16", err)
		}
	}
	if sp_CSI_ReportingOnPUSCH_r16Present {
		ie.sp_CSI_ReportingOnPUSCH_r16 = new(CA_ParametersNR_v1690_csi_ReportingCrossPUCCH_Grp_r16_sp_CSI_ReportingOnPUSCH_r16)
		if err = ie.sp_CSI_ReportingOnPUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_ReportingOnPUSCH_r16", err)
		}
	}
	tmp_carrierTypePairList_r16 := utils.NewSequence[*CarrierTypePair_r16]([]*CarrierTypePair_r16{}, uper.Constraint{Lb: 1, Ub: maxCarrierTypePairList_r16}, false)
	fn_carrierTypePairList_r16 := func() *CarrierTypePair_r16 {
		return new(CarrierTypePair_r16)
	}
	if err = tmp_carrierTypePairList_r16.Decode(r, fn_carrierTypePairList_r16); err != nil {
		return utils.WrapError("Decode carrierTypePairList_r16", err)
	}
	ie.carrierTypePairList_r16 = []CarrierTypePair_r16{}
	for _, i := range tmp_carrierTypePairList_r16.Value {
		ie.carrierTypePairList_r16 = append(ie.carrierTypePairList_r16, *i)
	}
	return nil
}
