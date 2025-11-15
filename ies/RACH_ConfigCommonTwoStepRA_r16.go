package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommonTwoStepRA_r16 struct {
	rach_ConfigGenericTwoStepRA_r16                    RACH_ConfigGenericTwoStepRA_r16                                                    `madatory`
	msgA_TotalNumberOfRA_Preambles_r16                 *int64                                                                             `lb:1,ub:63,optional`
	msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 `lb:1,ub:16,optional`
	msgA_CB_PreamblesPerSSB_PerSharedRO_r16            *int64                                                                             `lb:1,ub:60,optional`
	msgA_SSB_SharedRO_MaskIndex_r16                    *int64                                                                             `lb:1,ub:15,optional`
	groupB_ConfiguredTwoStepRA_r16                     *GroupB_ConfiguredTwoStepRA_r16                                                    `optional`
	msgA_PRACH_RootSequenceIndex_r16                   *RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16                   `lb:0,ub:837,optional`
	msgA_TransMax_r16                                  *RACH_ConfigCommonTwoStepRA_r16_msgA_TransMax_r16                                  `optional`
	msgA_RSRP_Threshold_r16                            *RSRP_Range                                                                        `optional`
	msgA_RSRP_ThresholdSSB_r16                         *RSRP_Range                                                                        `optional`
	msgA_SubcarrierSpacing_r16                         *SubcarrierSpacing                                                                 `optional`
	msgA_RestrictedSetConfig_r16                       *RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16                       `optional`
	ra_PrioritizationForAccessIdentityTwoStep_r16      *RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16      `lb:2,ub:2,optional`
	ra_ContentionResolutionTimer_r16                   *RACH_ConfigCommonTwoStepRA_r16_ra_ContentionResolutionTimer_r16                   `optional`
	ra_PrioritizationForSlicingTwoStep_r17             *RA_PrioritizationForSlicing_r17                                                   `optional,ext-1`
	featureCombinationPreamblesList_r17                []FeatureCombinationPreambles_r17                                                  `lb:1,ub:maxFeatureCombPreamblesPerRACHResource_r17,optional,ext-1`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ra_PrioritizationForSlicingTwoStep_r17 != nil || len(ie.featureCombinationPreamblesList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.msgA_TotalNumberOfRA_Preambles_r16 != nil, ie.msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 != nil, ie.msgA_CB_PreamblesPerSSB_PerSharedRO_r16 != nil, ie.msgA_SSB_SharedRO_MaskIndex_r16 != nil, ie.groupB_ConfiguredTwoStepRA_r16 != nil, ie.msgA_PRACH_RootSequenceIndex_r16 != nil, ie.msgA_TransMax_r16 != nil, ie.msgA_RSRP_Threshold_r16 != nil, ie.msgA_RSRP_ThresholdSSB_r16 != nil, ie.msgA_SubcarrierSpacing_r16 != nil, ie.msgA_RestrictedSetConfig_r16 != nil, ie.ra_PrioritizationForAccessIdentityTwoStep_r16 != nil, ie.ra_ContentionResolutionTimer_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rach_ConfigGenericTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rach_ConfigGenericTwoStepRA_r16", err)
	}
	if ie.msgA_TotalNumberOfRA_Preambles_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_TotalNumberOfRA_Preambles_r16, &uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode msgA_TotalNumberOfRA_Preambles_r16", err)
		}
	}
	if ie.msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 != nil {
		if err = ie.msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16", err)
		}
	}
	if ie.msgA_CB_PreamblesPerSSB_PerSharedRO_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_CB_PreamblesPerSSB_PerSharedRO_r16, &uper.Constraint{Lb: 1, Ub: 60}, false); err != nil {
			return utils.WrapError("Encode msgA_CB_PreamblesPerSSB_PerSharedRO_r16", err)
		}
	}
	if ie.msgA_SSB_SharedRO_MaskIndex_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_SSB_SharedRO_MaskIndex_r16, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode msgA_SSB_SharedRO_MaskIndex_r16", err)
		}
	}
	if ie.groupB_ConfiguredTwoStepRA_r16 != nil {
		if err = ie.groupB_ConfiguredTwoStepRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode groupB_ConfiguredTwoStepRA_r16", err)
		}
	}
	if ie.msgA_PRACH_RootSequenceIndex_r16 != nil {
		if err = ie.msgA_PRACH_RootSequenceIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_PRACH_RootSequenceIndex_r16", err)
		}
	}
	if ie.msgA_TransMax_r16 != nil {
		if err = ie.msgA_TransMax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_TransMax_r16", err)
		}
	}
	if ie.msgA_RSRP_Threshold_r16 != nil {
		if err = ie.msgA_RSRP_Threshold_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_RSRP_Threshold_r16", err)
		}
	}
	if ie.msgA_RSRP_ThresholdSSB_r16 != nil {
		if err = ie.msgA_RSRP_ThresholdSSB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_RSRP_ThresholdSSB_r16", err)
		}
	}
	if ie.msgA_SubcarrierSpacing_r16 != nil {
		if err = ie.msgA_SubcarrierSpacing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_SubcarrierSpacing_r16", err)
		}
	}
	if ie.msgA_RestrictedSetConfig_r16 != nil {
		if err = ie.msgA_RestrictedSetConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_RestrictedSetConfig_r16", err)
		}
	}
	if ie.ra_PrioritizationForAccessIdentityTwoStep_r16 != nil {
		if err = ie.ra_PrioritizationForAccessIdentityTwoStep_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ra_PrioritizationForAccessIdentityTwoStep_r16", err)
		}
	}
	if ie.ra_ContentionResolutionTimer_r16 != nil {
		if err = ie.ra_ContentionResolutionTimer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ra_ContentionResolutionTimer_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.ra_PrioritizationForSlicingTwoStep_r17 != nil || len(ie.featureCombinationPreamblesList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigCommonTwoStepRA_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ra_PrioritizationForSlicingTwoStep_r17 != nil, len(ie.featureCombinationPreamblesList_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ra_PrioritizationForSlicingTwoStep_r17 optional
			if ie.ra_PrioritizationForSlicingTwoStep_r17 != nil {
				if err = ie.ra_PrioritizationForSlicingTwoStep_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ra_PrioritizationForSlicingTwoStep_r17", err)
				}
			}
			// encode featureCombinationPreamblesList_r17 optional
			if len(ie.featureCombinationPreamblesList_r17) > 0 {
				tmp_featureCombinationPreamblesList_r17 := utils.NewSequence[*FeatureCombinationPreambles_r17]([]*FeatureCombinationPreambles_r17{}, uper.Constraint{Lb: 1, Ub: maxFeatureCombPreamblesPerRACHResource_r17}, false)
				for _, i := range ie.featureCombinationPreamblesList_r17 {
					tmp_featureCombinationPreamblesList_r17.Value = append(tmp_featureCombinationPreamblesList_r17.Value, &i)
				}
				if err = tmp_featureCombinationPreamblesList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featureCombinationPreamblesList_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *RACH_ConfigCommonTwoStepRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_TotalNumberOfRA_Preambles_r16Present bool
	if msgA_TotalNumberOfRA_Preambles_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16Present bool
	if msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_CB_PreamblesPerSSB_PerSharedRO_r16Present bool
	if msgA_CB_PreamblesPerSSB_PerSharedRO_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_SSB_SharedRO_MaskIndex_r16Present bool
	if msgA_SSB_SharedRO_MaskIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var groupB_ConfiguredTwoStepRA_r16Present bool
	if groupB_ConfiguredTwoStepRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_PRACH_RootSequenceIndex_r16Present bool
	if msgA_PRACH_RootSequenceIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_TransMax_r16Present bool
	if msgA_TransMax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_RSRP_Threshold_r16Present bool
	if msgA_RSRP_Threshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_RSRP_ThresholdSSB_r16Present bool
	if msgA_RSRP_ThresholdSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_SubcarrierSpacing_r16Present bool
	if msgA_SubcarrierSpacing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_RestrictedSetConfig_r16Present bool
	if msgA_RestrictedSetConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_PrioritizationForAccessIdentityTwoStep_r16Present bool
	if ra_PrioritizationForAccessIdentityTwoStep_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_ContentionResolutionTimer_r16Present bool
	if ra_ContentionResolutionTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rach_ConfigGenericTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rach_ConfigGenericTwoStepRA_r16", err)
	}
	if msgA_TotalNumberOfRA_Preambles_r16Present {
		var tmp_int_msgA_TotalNumberOfRA_Preambles_r16 int64
		if tmp_int_msgA_TotalNumberOfRA_Preambles_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode msgA_TotalNumberOfRA_Preambles_r16", err)
		}
		ie.msgA_TotalNumberOfRA_Preambles_r16 = &tmp_int_msgA_TotalNumberOfRA_Preambles_r16
	}
	if msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16Present {
		ie.msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16)
		if err = ie.msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16", err)
		}
	}
	if msgA_CB_PreamblesPerSSB_PerSharedRO_r16Present {
		var tmp_int_msgA_CB_PreamblesPerSSB_PerSharedRO_r16 int64
		if tmp_int_msgA_CB_PreamblesPerSSB_PerSharedRO_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 60}, false); err != nil {
			return utils.WrapError("Decode msgA_CB_PreamblesPerSSB_PerSharedRO_r16", err)
		}
		ie.msgA_CB_PreamblesPerSSB_PerSharedRO_r16 = &tmp_int_msgA_CB_PreamblesPerSSB_PerSharedRO_r16
	}
	if msgA_SSB_SharedRO_MaskIndex_r16Present {
		var tmp_int_msgA_SSB_SharedRO_MaskIndex_r16 int64
		if tmp_int_msgA_SSB_SharedRO_MaskIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode msgA_SSB_SharedRO_MaskIndex_r16", err)
		}
		ie.msgA_SSB_SharedRO_MaskIndex_r16 = &tmp_int_msgA_SSB_SharedRO_MaskIndex_r16
	}
	if groupB_ConfiguredTwoStepRA_r16Present {
		ie.groupB_ConfiguredTwoStepRA_r16 = new(GroupB_ConfiguredTwoStepRA_r16)
		if err = ie.groupB_ConfiguredTwoStepRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode groupB_ConfiguredTwoStepRA_r16", err)
		}
	}
	if msgA_PRACH_RootSequenceIndex_r16Present {
		ie.msgA_PRACH_RootSequenceIndex_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_PRACH_RootSequenceIndex_r16)
		if err = ie.msgA_PRACH_RootSequenceIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_PRACH_RootSequenceIndex_r16", err)
		}
	}
	if msgA_TransMax_r16Present {
		ie.msgA_TransMax_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_TransMax_r16)
		if err = ie.msgA_TransMax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_TransMax_r16", err)
		}
	}
	if msgA_RSRP_Threshold_r16Present {
		ie.msgA_RSRP_Threshold_r16 = new(RSRP_Range)
		if err = ie.msgA_RSRP_Threshold_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_RSRP_Threshold_r16", err)
		}
	}
	if msgA_RSRP_ThresholdSSB_r16Present {
		ie.msgA_RSRP_ThresholdSSB_r16 = new(RSRP_Range)
		if err = ie.msgA_RSRP_ThresholdSSB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_RSRP_ThresholdSSB_r16", err)
		}
	}
	if msgA_SubcarrierSpacing_r16Present {
		ie.msgA_SubcarrierSpacing_r16 = new(SubcarrierSpacing)
		if err = ie.msgA_SubcarrierSpacing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_SubcarrierSpacing_r16", err)
		}
	}
	if msgA_RestrictedSetConfig_r16Present {
		ie.msgA_RestrictedSetConfig_r16 = new(RACH_ConfigCommonTwoStepRA_r16_msgA_RestrictedSetConfig_r16)
		if err = ie.msgA_RestrictedSetConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_RestrictedSetConfig_r16", err)
		}
	}
	if ra_PrioritizationForAccessIdentityTwoStep_r16Present {
		ie.ra_PrioritizationForAccessIdentityTwoStep_r16 = new(RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16)
		if err = ie.ra_PrioritizationForAccessIdentityTwoStep_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ra_PrioritizationForAccessIdentityTwoStep_r16", err)
		}
	}
	if ra_ContentionResolutionTimer_r16Present {
		ie.ra_ContentionResolutionTimer_r16 = new(RACH_ConfigCommonTwoStepRA_r16_ra_ContentionResolutionTimer_r16)
		if err = ie.ra_ContentionResolutionTimer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ra_ContentionResolutionTimer_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ra_PrioritizationForSlicingTwoStep_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureCombinationPreamblesList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ra_PrioritizationForSlicingTwoStep_r17 optional
			if ra_PrioritizationForSlicingTwoStep_r17Present {
				ie.ra_PrioritizationForSlicingTwoStep_r17 = new(RA_PrioritizationForSlicing_r17)
				if err = ie.ra_PrioritizationForSlicingTwoStep_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ra_PrioritizationForSlicingTwoStep_r17", err)
				}
			}
			// decode featureCombinationPreamblesList_r17 optional
			if featureCombinationPreamblesList_r17Present {
				tmp_featureCombinationPreamblesList_r17 := utils.NewSequence[*FeatureCombinationPreambles_r17]([]*FeatureCombinationPreambles_r17{}, uper.Constraint{Lb: 1, Ub: maxFeatureCombPreamblesPerRACHResource_r17}, false)
				fn_featureCombinationPreamblesList_r17 := func() *FeatureCombinationPreambles_r17 {
					return new(FeatureCombinationPreambles_r17)
				}
				if err = tmp_featureCombinationPreamblesList_r17.Decode(extReader, fn_featureCombinationPreamblesList_r17); err != nil {
					return utils.WrapError("Decode featureCombinationPreamblesList_r17", err)
				}
				ie.featureCombinationPreamblesList_r17 = []FeatureCombinationPreambles_r17{}
				for _, i := range tmp_featureCombinationPreamblesList_r17.Value {
					ie.featureCombinationPreamblesList_r17 = append(ie.featureCombinationPreamblesList_r17, *i)
				}
			}
		}
	}
	return nil
}
