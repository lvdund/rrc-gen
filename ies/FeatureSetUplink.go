package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink struct {
	featureSetListPerUplinkCC                 []FeatureSetUplinkPerCC_Id                                  `lb:1,ub:maxNrofServingCells,madatory`
	scalingFactor                             *FeatureSetUplink_scalingFactor                             `optional`
	dummy3                                    *FeatureSetUplink_dummy3                                    `optional`
	intraBandFreqSeparationUL                 *FreqSeparationClass                                        `optional`
	searchSpaceSharingCA_UL                   *FeatureSetUplink_searchSpaceSharingCA_UL                   `optional`
	dummy1                                    *DummyI                                                     `optional`
	supportedSRS_Resources                    *SRS_Resources                                              `optional`
	twoPUCCH_Group                            *FeatureSetUplink_twoPUCCH_Group                            `optional`
	dynamicSwitchSUL                          *FeatureSetUplink_dynamicSwitchSUL                          `optional`
	simultaneousTxSUL_NonSUL                  *FeatureSetUplink_simultaneousTxSUL_NonSUL                  `optional`
	pusch_ProcessingType1_DifferentTB_PerSlot *FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot `optional`
	dummy2                                    *DummyF                                                     `optional`
}

func (ie *FeatureSetUplink) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scalingFactor != nil, ie.dummy3 != nil, ie.intraBandFreqSeparationUL != nil, ie.searchSpaceSharingCA_UL != nil, ie.dummy1 != nil, ie.supportedSRS_Resources != nil, ie.twoPUCCH_Group != nil, ie.dynamicSwitchSUL != nil, ie.simultaneousTxSUL_NonSUL != nil, ie.pusch_ProcessingType1_DifferentTB_PerSlot != nil, ie.dummy2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_featureSetListPerUplinkCC := utils.NewSequence[*FeatureSetUplinkPerCC_Id]([]*FeatureSetUplinkPerCC_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.featureSetListPerUplinkCC {
		tmp_featureSetListPerUplinkCC.Value = append(tmp_featureSetListPerUplinkCC.Value, &i)
	}
	if err = tmp_featureSetListPerUplinkCC.Encode(w); err != nil {
		return utils.WrapError("Encode featureSetListPerUplinkCC", err)
	}
	if ie.scalingFactor != nil {
		if err = ie.scalingFactor.Encode(w); err != nil {
			return utils.WrapError("Encode scalingFactor", err)
		}
	}
	if ie.dummy3 != nil {
		if err = ie.dummy3.Encode(w); err != nil {
			return utils.WrapError("Encode dummy3", err)
		}
	}
	if ie.intraBandFreqSeparationUL != nil {
		if err = ie.intraBandFreqSeparationUL.Encode(w); err != nil {
			return utils.WrapError("Encode intraBandFreqSeparationUL", err)
		}
	}
	if ie.searchSpaceSharingCA_UL != nil {
		if err = ie.searchSpaceSharingCA_UL.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceSharingCA_UL", err)
		}
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.supportedSRS_Resources != nil {
		if err = ie.supportedSRS_Resources.Encode(w); err != nil {
			return utils.WrapError("Encode supportedSRS_Resources", err)
		}
	}
	if ie.twoPUCCH_Group != nil {
		if err = ie.twoPUCCH_Group.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_Group", err)
		}
	}
	if ie.dynamicSwitchSUL != nil {
		if err = ie.dynamicSwitchSUL.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicSwitchSUL", err)
		}
	}
	if ie.simultaneousTxSUL_NonSUL != nil {
		if err = ie.simultaneousTxSUL_NonSUL.Encode(w); err != nil {
			return utils.WrapError("Encode simultaneousTxSUL_NonSUL", err)
		}
	}
	if ie.pusch_ProcessingType1_DifferentTB_PerSlot != nil {
		if err = ie.pusch_ProcessingType1_DifferentTB_PerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if ie.dummy2 != nil {
		if err = ie.dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode dummy2", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink) Decode(r *uper.UperReader) error {
	var err error
	var scalingFactorPresent bool
	if scalingFactorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy3Present bool
	if dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraBandFreqSeparationULPresent bool
	if intraBandFreqSeparationULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSharingCA_ULPresent bool
	if searchSpaceSharingCA_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedSRS_ResourcesPresent bool
	if supportedSRS_ResourcesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_GroupPresent bool
	if twoPUCCH_GroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicSwitchSULPresent bool
	if dynamicSwitchSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var simultaneousTxSUL_NonSULPresent bool
	if simultaneousTxSUL_NonSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_ProcessingType1_DifferentTB_PerSlotPresent bool
	if pusch_ProcessingType1_DifferentTB_PerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy2Present bool
	if dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_featureSetListPerUplinkCC := utils.NewSequence[*FeatureSetUplinkPerCC_Id]([]*FeatureSetUplinkPerCC_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn_featureSetListPerUplinkCC := func() *FeatureSetUplinkPerCC_Id {
		return new(FeatureSetUplinkPerCC_Id)
	}
	if err = tmp_featureSetListPerUplinkCC.Decode(r, fn_featureSetListPerUplinkCC); err != nil {
		return utils.WrapError("Decode featureSetListPerUplinkCC", err)
	}
	ie.featureSetListPerUplinkCC = []FeatureSetUplinkPerCC_Id{}
	for _, i := range tmp_featureSetListPerUplinkCC.Value {
		ie.featureSetListPerUplinkCC = append(ie.featureSetListPerUplinkCC, *i)
	}
	if scalingFactorPresent {
		ie.scalingFactor = new(FeatureSetUplink_scalingFactor)
		if err = ie.scalingFactor.Decode(r); err != nil {
			return utils.WrapError("Decode scalingFactor", err)
		}
	}
	if dummy3Present {
		ie.dummy3 = new(FeatureSetUplink_dummy3)
		if err = ie.dummy3.Decode(r); err != nil {
			return utils.WrapError("Decode dummy3", err)
		}
	}
	if intraBandFreqSeparationULPresent {
		ie.intraBandFreqSeparationUL = new(FreqSeparationClass)
		if err = ie.intraBandFreqSeparationUL.Decode(r); err != nil {
			return utils.WrapError("Decode intraBandFreqSeparationUL", err)
		}
	}
	if searchSpaceSharingCA_ULPresent {
		ie.searchSpaceSharingCA_UL = new(FeatureSetUplink_searchSpaceSharingCA_UL)
		if err = ie.searchSpaceSharingCA_UL.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceSharingCA_UL", err)
		}
	}
	if dummy1Present {
		ie.dummy1 = new(DummyI)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if supportedSRS_ResourcesPresent {
		ie.supportedSRS_Resources = new(SRS_Resources)
		if err = ie.supportedSRS_Resources.Decode(r); err != nil {
			return utils.WrapError("Decode supportedSRS_Resources", err)
		}
	}
	if twoPUCCH_GroupPresent {
		ie.twoPUCCH_Group = new(FeatureSetUplink_twoPUCCH_Group)
		if err = ie.twoPUCCH_Group.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_Group", err)
		}
	}
	if dynamicSwitchSULPresent {
		ie.dynamicSwitchSUL = new(FeatureSetUplink_dynamicSwitchSUL)
		if err = ie.dynamicSwitchSUL.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicSwitchSUL", err)
		}
	}
	if simultaneousTxSUL_NonSULPresent {
		ie.simultaneousTxSUL_NonSUL = new(FeatureSetUplink_simultaneousTxSUL_NonSUL)
		if err = ie.simultaneousTxSUL_NonSUL.Decode(r); err != nil {
			return utils.WrapError("Decode simultaneousTxSUL_NonSUL", err)
		}
	}
	if pusch_ProcessingType1_DifferentTB_PerSlotPresent {
		ie.pusch_ProcessingType1_DifferentTB_PerSlot = new(FeatureSetUplink_pusch_ProcessingType1_DifferentTB_PerSlot)
		if err = ie.pusch_ProcessingType1_DifferentTB_PerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_ProcessingType1_DifferentTB_PerSlot", err)
		}
	}
	if dummy2Present {
		ie.dummy2 = new(DummyF)
		if err = ie.dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode dummy2", err)
		}
	}
	return nil
}
