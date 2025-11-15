package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MRDC_Capability struct {
	measAndMobParametersMRDC     *MeasAndMobParametersMRDC       `optional`
	phy_ParametersMRDC_v1530     *Phy_ParametersMRDC             `optional`
	rf_ParametersMRDC            RF_ParametersMRDC               `madatory`
	generalParametersMRDC        *GeneralParametersMRDC_XDD_Diff `optional`
	fdd_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	tdd_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddXDD_Mode  `optional`
	fr1_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	fr2_Add_UE_MRDC_Capabilities *UE_MRDC_CapabilityAddFRX_Mode  `optional`
	featureSetCombinations       []FeatureSetCombination         `lb:1,ub:maxFeatureSetCombinations,optional`
	pdcp_ParametersMRDC_v1530    *PDCP_ParametersMRDC            `optional`
	lateNonCriticalExtension     *[]byte                         `optional`
	nonCriticalExtension         *UE_MRDC_Capability_v1560       `optional`
}

func (ie *UE_MRDC_Capability) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measAndMobParametersMRDC != nil, ie.phy_ParametersMRDC_v1530 != nil, ie.generalParametersMRDC != nil, ie.fdd_Add_UE_MRDC_Capabilities != nil, ie.tdd_Add_UE_MRDC_Capabilities != nil, ie.fr1_Add_UE_MRDC_Capabilities != nil, ie.fr2_Add_UE_MRDC_Capabilities != nil, len(ie.featureSetCombinations) > 0, ie.pdcp_ParametersMRDC_v1530 != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measAndMobParametersMRDC != nil {
		if err = ie.measAndMobParametersMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode measAndMobParametersMRDC", err)
		}
	}
	if ie.phy_ParametersMRDC_v1530 != nil {
		if err = ie.phy_ParametersMRDC_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode phy_ParametersMRDC_v1530", err)
		}
	}
	if err = ie.rf_ParametersMRDC.Encode(w); err != nil {
		return utils.WrapError("Encode rf_ParametersMRDC", err)
	}
	if ie.generalParametersMRDC != nil {
		if err = ie.generalParametersMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode generalParametersMRDC", err)
		}
	}
	if ie.fdd_Add_UE_MRDC_Capabilities != nil {
		if err = ie.fdd_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if ie.tdd_Add_UE_MRDC_Capabilities != nil {
		if err = ie.tdd_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if ie.fr1_Add_UE_MRDC_Capabilities != nil {
		if err = ie.fr1_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fr1_Add_UE_MRDC_Capabilities", err)
		}
	}
	if ie.fr2_Add_UE_MRDC_Capabilities != nil {
		if err = ie.fr2_Add_UE_MRDC_Capabilities.Encode(w); err != nil {
			return utils.WrapError("Encode fr2_Add_UE_MRDC_Capabilities", err)
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
	if ie.pdcp_ParametersMRDC_v1530 != nil {
		if err = ie.pdcp_ParametersMRDC_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_ParametersMRDC_v1530", err)
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

func (ie *UE_MRDC_Capability) Decode(r *uper.UperReader) error {
	var err error
	var measAndMobParametersMRDCPresent bool
	if measAndMobParametersMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var phy_ParametersMRDC_v1530Present bool
	if phy_ParametersMRDC_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	var generalParametersMRDCPresent bool
	if generalParametersMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fdd_Add_UE_MRDC_CapabilitiesPresent bool
	if fdd_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_Add_UE_MRDC_CapabilitiesPresent bool
	if tdd_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr1_Add_UE_MRDC_CapabilitiesPresent bool
	if fr1_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr2_Add_UE_MRDC_CapabilitiesPresent bool
	if fr2_Add_UE_MRDC_CapabilitiesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var featureSetCombinationsPresent bool
	if featureSetCombinationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_ParametersMRDC_v1530Present bool
	if pdcp_ParametersMRDC_v1530Present, err = r.ReadBool(); err != nil {
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
	if measAndMobParametersMRDCPresent {
		ie.measAndMobParametersMRDC = new(MeasAndMobParametersMRDC)
		if err = ie.measAndMobParametersMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode measAndMobParametersMRDC", err)
		}
	}
	if phy_ParametersMRDC_v1530Present {
		ie.phy_ParametersMRDC_v1530 = new(Phy_ParametersMRDC)
		if err = ie.phy_ParametersMRDC_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode phy_ParametersMRDC_v1530", err)
		}
	}
	if err = ie.rf_ParametersMRDC.Decode(r); err != nil {
		return utils.WrapError("Decode rf_ParametersMRDC", err)
	}
	if generalParametersMRDCPresent {
		ie.generalParametersMRDC = new(GeneralParametersMRDC_XDD_Diff)
		if err = ie.generalParametersMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode generalParametersMRDC", err)
		}
	}
	if fdd_Add_UE_MRDC_CapabilitiesPresent {
		ie.fdd_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.fdd_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if tdd_Add_UE_MRDC_CapabilitiesPresent {
		ie.tdd_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddXDD_Mode)
		if err = ie.tdd_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_Add_UE_MRDC_Capabilities", err)
		}
	}
	if fr1_Add_UE_MRDC_CapabilitiesPresent {
		ie.fr1_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.fr1_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fr1_Add_UE_MRDC_Capabilities", err)
		}
	}
	if fr2_Add_UE_MRDC_CapabilitiesPresent {
		ie.fr2_Add_UE_MRDC_Capabilities = new(UE_MRDC_CapabilityAddFRX_Mode)
		if err = ie.fr2_Add_UE_MRDC_Capabilities.Decode(r); err != nil {
			return utils.WrapError("Decode fr2_Add_UE_MRDC_Capabilities", err)
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
	if pdcp_ParametersMRDC_v1530Present {
		ie.pdcp_ParametersMRDC_v1530 = new(PDCP_ParametersMRDC)
		if err = ie.pdcp_ParametersMRDC_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_ParametersMRDC_v1530", err)
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
		ie.nonCriticalExtension = new(UE_MRDC_Capability_v1560)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
