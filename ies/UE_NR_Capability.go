package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability struct {
	accessStratumRelease       AccessStratumRelease         `madatory`
	pdcp_Parameters            PDCP_Parameters              `madatory`
	rlc_Parameters             *RLC_Parameters              `optional`
	mac_Parameters             *MAC_Parameters              `optional`
	phy_Parameters             Phy_Parameters               `madatory`
	rf_Parameters              RF_Parameters                `madatory`
	measAndMobParameters       *MeasAndMobParameters        `optional`
	fdd_Add_UE_NR_Capabilities *UE_NR_CapabilityAddXDD_Mode `optional`
	tdd_Add_UE_NR_Capabilities *UE_NR_CapabilityAddXDD_Mode `optional`
	fr1_Add_UE_NR_Capabilities *UE_NR_CapabilityAddFRX_Mode `optional`
	fr2_Add_UE_NR_Capabilities *UE_NR_CapabilityAddFRX_Mode `optional`
	featureSets                *FeatureSets                 `optional`
	featureSetCombinations     []FeatureSetCombination      `lb:1,ub:maxFeatureSetCombinations,optional`
	lateNonCriticalExtension   *[]byte                      `optional`
	nonCriticalExtension       *UE_NR_Capability_v1530      `optional`
}

func (ie *UE_NR_Capability) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rlc_Parameters != nil, ie.mac_Parameters != nil, ie.measAndMobParameters != nil, ie.fdd_Add_UE_NR_Capabilities != nil, ie.tdd_Add_UE_NR_Capabilities != nil, ie.fr1_Add_UE_NR_Capabilities != nil, ie.fr2_Add_UE_NR_Capabilities != nil, ie.featureSets != nil, len(ie.featureSetCombinations) > 0, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.accessStratumRelease.Encode(w); err != nil {
		return utils.WrapError("Encode accessStratumRelease", err)
	}
	if err = ie.pdcp_Parameters.Encode(w); err != nil {
		return utils.WrapError("Encode pdcp_Parameters", err)
	}
	if ie.rlc_Parameters != nil {
		if err = ie.rlc_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_Parameters", err)
		}
	}
	if ie.mac_Parameters != nil {
		if err = ie.mac_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode mac_Parameters", err)
		}
	}
	if err = ie.phy_Parameters.Encode(w); err != nil {
		return utils.WrapError("Encode phy_Parameters", err)
	}
	if err = ie.rf_Parameters.Encode(w); err != nil {
		return utils.WrapError("Encode rf_Parameters", err)
	}
	if ie.measAndMobParameters != nil {
		if err = ie.measAndMobParameters.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParameters", err)
		}
	}
	if ie.fdd_Add_UE_NR_Capabilities != nil {
		if err = ie.fdd_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.tdd_Add_UE_NR_Capabilities != nil {
		if err = ie.tdd_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.fr1_Add_UE_NR_Capabilities != nil {
		if err = ie.fr1_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.fr2_Add_UE_NR_Capabilities != nil {
		if err = ie.fr2_Add_UE_NR_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if ie.featureSets != nil {
		if err = ie.featureSets.Encode(w); err != nil {
			return utils.WrapError("Encode featureSets", err)
		}
	}
	if len(ie.featureSetCombinations) > 0 {
		tmp_featureSetCombinations := utils.NewSequence[*FeatureSetCombination]([]*FeatureSetCombination{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetCombinations}, false)
		for _, i := range ie.featureSetCombinations {
			tmp_featureSetCombinations.Value = append(tmp_featureSetCombinations.Value, &i)
		}
		if err = tmp_featureSetCombinations.Encode(w); err != nil {
			return utils.WrapError("Encode featureSetCombinations", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability) Decode(r *uper.UperReader) error {
	var err error
	var rlc_ParametersPresent bool
	if rlc_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_ParametersPresent bool
	if mac_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measAndMobParametersPresent bool
	if measAndMobParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fdd_Add_UE_NR_CapabilitiesPresent bool
	if fdd_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_Add_UE_NR_CapabilitiesPresent bool
	if tdd_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_Add_UE_NR_CapabilitiesPresent bool
	if fr1_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_Add_UE_NR_CapabilitiesPresent bool
	if fr2_Add_UE_NR_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetsPresent bool
	if featureSetsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetCombinationsPresent bool
	if featureSetCombinationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.accessStratumRelease.Decode(r); err != nil {
		return utils.WrapError("Decode accessStratumRelease", err)
	}
	if err = ie.pdcp_Parameters.Decode(r); err != nil {
		return utils.WrapError("Decode pdcp_Parameters", err)
	}
	if rlc_ParametersPresent {
		ie.rlc_Parameters = new(RLC_Parameters)
		if err = ie.rlc_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode rlc_Parameters", err)
		}
	}
	if mac_ParametersPresent {
		ie.mac_Parameters = new(MAC_Parameters)
		if err = ie.mac_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode mac_Parameters", err)
		}
	}
	if err = ie.phy_Parameters.Decode(r); err != nil {
		return utils.WrapError("Decode phy_Parameters", err)
	}
	if err = ie.rf_Parameters.Decode(r); err != nil {
		return utils.WrapError("Decode rf_Parameters", err)
	}
	if measAndMobParametersPresent {
		ie.measAndMobParameters = new(MeasAndMobParameters)
		if err = ie.measAndMobParameters.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParameters", err)
		}
	}
	if fdd_Add_UE_NR_CapabilitiesPresent {
		ie.fdd_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddXDD_Mode)
		if err = ie.fdd_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_Add_UE_NR_Capabilities", err)
		}
	}
	if tdd_Add_UE_NR_CapabilitiesPresent {
		ie.tdd_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddXDD_Mode)
		if err = ie.tdd_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_Add_UE_NR_Capabilities", err)
		}
	}
	if fr1_Add_UE_NR_CapabilitiesPresent {
		ie.fr1_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.fr1_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_Add_UE_NR_Capabilities", err)
		}
	}
	if fr2_Add_UE_NR_CapabilitiesPresent {
		ie.fr2_Add_UE_NR_Capabilities = new(UE_NR_CapabilityAddFRX_Mode)
		if err = ie.fr2_Add_UE_NR_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_Add_UE_NR_Capabilities", err)
		}
	}
	if featureSetsPresent {
		ie.featureSets = new(FeatureSets)
		if err = ie.featureSets.Decode(r); err != nil {
			return utils.WrapError("Decode featureSets", err)
		}
	}
	if featureSetCombinationsPresent {
		tmp_featureSetCombinations := utils.NewSequence[*FeatureSetCombination]([]*FeatureSetCombination{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetCombinations}, false)
		fn_featureSetCombinations := func() *FeatureSetCombination {
			return new(FeatureSetCombination)
		}
		if err = tmp_featureSetCombinations.Decode(r, fn_featureSetCombinations); err != nil {
			return utils.WrapError("Decode featureSetCombinations", err)
		}
		ie.featureSetCombinations = []FeatureSetCombination{}
		for _, i := range tmp_featureSetCombinations.Value {
			ie.featureSetCombinations = append(ie.featureSetCombinations, *i)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1530)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
