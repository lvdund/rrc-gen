package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommon struct {
	rach_ConfigGeneric                        RACH_ConfigGeneric                                           `madatory`
	totalNumberOfRA_Preambles                 *int64                                                       `lb:1,ub:63,optional`
	ssb_perRACH_OccasionAndCB_PreamblesPerSSB *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB `lb:1,ub:16,optional`
	groupBconfigured                          *RACH_ConfigCommon_groupBconfigured                          `lb:1,ub:64,optional`
	ra_ContentionResolutionTimer              RACH_ConfigCommon_ra_ContentionResolutionTimer               `madatory`
	rsrp_ThresholdSSB                         *RSRP_Range                                                  `optional`
	rsrp_ThresholdSSB_SUL                     *RSRP_Range                                                  `optional`
	prach_RootSequenceIndex                   RACH_ConfigCommon_prach_RootSequenceIndex                    `lb:0,ub:837,madatory`
	msg1_SubcarrierSpacing                    *SubcarrierSpacing                                           `optional`
	restrictedSetConfig                       RACH_ConfigCommon_restrictedSetConfig                        `madatory`
	msg3_transformPrecoder                    *RACH_ConfigCommon_msg3_transformPrecoder                    `optional`
	ra_PrioritizationForAccessIdentity_r16    *RACH_ConfigCommon_ra_PrioritizationForAccessIdentity_r16    `lb:2,ub:2,optional,ext-1`
	prach_RootSequenceIndex_r16               *RACH_ConfigCommon_prach_RootSequenceIndex_r16               `lb:0,ub:569,optional,ext-1`
	ra_PrioritizationForSlicing_r17           *RA_PrioritizationForSlicing_r17                             `optional,ext-2`
	featureCombinationPreamblesList_r17       []FeatureCombinationPreambles_r17                            `lb:1,ub:maxFeatureCombPreamblesPerRACHResource_r17,optional,ext-2`
}

func (ie *RACH_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ra_PrioritizationForAccessIdentity_r16 != nil || ie.prach_RootSequenceIndex_r16 != nil || ie.ra_PrioritizationForSlicing_r17 != nil || len(ie.featureCombinationPreamblesList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.totalNumberOfRA_Preambles != nil, ie.ssb_perRACH_OccasionAndCB_PreamblesPerSSB != nil, ie.groupBconfigured != nil, ie.rsrp_ThresholdSSB != nil, ie.rsrp_ThresholdSSB_SUL != nil, ie.msg1_SubcarrierSpacing != nil, ie.msg3_transformPrecoder != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rach_ConfigGeneric.Encode(w); err != nil {
		return utils.WrapError("Encode rach_ConfigGeneric", err)
	}
	if ie.totalNumberOfRA_Preambles != nil {
		if err = w.WriteInteger(*ie.totalNumberOfRA_Preambles, &uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Encode totalNumberOfRA_Preambles", err)
		}
	}
	if ie.ssb_perRACH_OccasionAndCB_PreamblesPerSSB != nil {
		if err = ie.ssb_perRACH_OccasionAndCB_PreamblesPerSSB.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_perRACH_OccasionAndCB_PreamblesPerSSB", err)
		}
	}
	if ie.groupBconfigured != nil {
		if err = ie.groupBconfigured.Encode(w); err != nil {
			return utils.WrapError("Encode groupBconfigured", err)
		}
	}
	if err = ie.ra_ContentionResolutionTimer.Encode(w); err != nil {
		return utils.WrapError("Encode ra_ContentionResolutionTimer", err)
	}
	if ie.rsrp_ThresholdSSB != nil {
		if err = ie.rsrp_ThresholdSSB.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_ThresholdSSB", err)
		}
	}
	if ie.rsrp_ThresholdSSB_SUL != nil {
		if err = ie.rsrp_ThresholdSSB_SUL.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_ThresholdSSB_SUL", err)
		}
	}
	if err = ie.prach_RootSequenceIndex.Encode(w); err != nil {
		return utils.WrapError("Encode prach_RootSequenceIndex", err)
	}
	if ie.msg1_SubcarrierSpacing != nil {
		if err = ie.msg1_SubcarrierSpacing.Encode(w); err != nil {
			return utils.WrapError("Encode msg1_SubcarrierSpacing", err)
		}
	}
	if err = ie.restrictedSetConfig.Encode(w); err != nil {
		return utils.WrapError("Encode restrictedSetConfig", err)
	}
	if ie.msg3_transformPrecoder != nil {
		if err = ie.msg3_transformPrecoder.Encode(w); err != nil {
			return utils.WrapError("Encode msg3_transformPrecoder", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.ra_PrioritizationForAccessIdentity_r16 != nil || ie.prach_RootSequenceIndex_r16 != nil, ie.ra_PrioritizationForSlicing_r17 != nil || len(ie.featureCombinationPreamblesList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ra_PrioritizationForAccessIdentity_r16 != nil, ie.prach_RootSequenceIndex_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ra_PrioritizationForAccessIdentity_r16 optional
			if ie.ra_PrioritizationForAccessIdentity_r16 != nil {
				if err = ie.ra_PrioritizationForAccessIdentity_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ra_PrioritizationForAccessIdentity_r16", err)
				}
			}
			// encode prach_RootSequenceIndex_r16 optional
			if ie.prach_RootSequenceIndex_r16 != nil {
				if err = ie.prach_RootSequenceIndex_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prach_RootSequenceIndex_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.ra_PrioritizationForSlicing_r17 != nil, len(ie.featureCombinationPreamblesList_r17) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ra_PrioritizationForSlicing_r17 optional
			if ie.ra_PrioritizationForSlicing_r17 != nil {
				if err = ie.ra_PrioritizationForSlicing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ra_PrioritizationForSlicing_r17", err)
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

func (ie *RACH_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var totalNumberOfRA_PreamblesPresent bool
	if totalNumberOfRA_PreamblesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_perRACH_OccasionAndCB_PreamblesPerSSBPresent bool
	if ssb_perRACH_OccasionAndCB_PreamblesPerSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var groupBconfiguredPresent bool
	if groupBconfiguredPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrp_ThresholdSSBPresent bool
	if rsrp_ThresholdSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrp_ThresholdSSB_SULPresent bool
	if rsrp_ThresholdSSB_SULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var msg1_SubcarrierSpacingPresent bool
	if msg1_SubcarrierSpacingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var msg3_transformPrecoderPresent bool
	if msg3_transformPrecoderPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rach_ConfigGeneric.Decode(r); err != nil {
		return utils.WrapError("Decode rach_ConfigGeneric", err)
	}
	if totalNumberOfRA_PreamblesPresent {
		var tmp_int_totalNumberOfRA_Preambles int64
		if tmp_int_totalNumberOfRA_Preambles, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode totalNumberOfRA_Preambles", err)
		}
		ie.totalNumberOfRA_Preambles = &tmp_int_totalNumberOfRA_Preambles
	}
	if ssb_perRACH_OccasionAndCB_PreamblesPerSSBPresent {
		ie.ssb_perRACH_OccasionAndCB_PreamblesPerSSB = new(RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB)
		if err = ie.ssb_perRACH_OccasionAndCB_PreamblesPerSSB.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_perRACH_OccasionAndCB_PreamblesPerSSB", err)
		}
	}
	if groupBconfiguredPresent {
		ie.groupBconfigured = new(RACH_ConfigCommon_groupBconfigured)
		if err = ie.groupBconfigured.Decode(r); err != nil {
			return utils.WrapError("Decode groupBconfigured", err)
		}
	}
	if err = ie.ra_ContentionResolutionTimer.Decode(r); err != nil {
		return utils.WrapError("Decode ra_ContentionResolutionTimer", err)
	}
	if rsrp_ThresholdSSBPresent {
		ie.rsrp_ThresholdSSB = new(RSRP_Range)
		if err = ie.rsrp_ThresholdSSB.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_ThresholdSSB", err)
		}
	}
	if rsrp_ThresholdSSB_SULPresent {
		ie.rsrp_ThresholdSSB_SUL = new(RSRP_Range)
		if err = ie.rsrp_ThresholdSSB_SUL.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_ThresholdSSB_SUL", err)
		}
	}
	if err = ie.prach_RootSequenceIndex.Decode(r); err != nil {
		return utils.WrapError("Decode prach_RootSequenceIndex", err)
	}
	if msg1_SubcarrierSpacingPresent {
		ie.msg1_SubcarrierSpacing = new(SubcarrierSpacing)
		if err = ie.msg1_SubcarrierSpacing.Decode(r); err != nil {
			return utils.WrapError("Decode msg1_SubcarrierSpacing", err)
		}
	}
	if err = ie.restrictedSetConfig.Decode(r); err != nil {
		return utils.WrapError("Decode restrictedSetConfig", err)
	}
	if msg3_transformPrecoderPresent {
		ie.msg3_transformPrecoder = new(RACH_ConfigCommon_msg3_transformPrecoder)
		if err = ie.msg3_transformPrecoder.Decode(r); err != nil {
			return utils.WrapError("Decode msg3_transformPrecoder", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			ra_PrioritizationForAccessIdentity_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prach_RootSequenceIndex_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ra_PrioritizationForAccessIdentity_r16 optional
			if ra_PrioritizationForAccessIdentity_r16Present {
				ie.ra_PrioritizationForAccessIdentity_r16 = new(RACH_ConfigCommon_ra_PrioritizationForAccessIdentity_r16)
				if err = ie.ra_PrioritizationForAccessIdentity_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ra_PrioritizationForAccessIdentity_r16", err)
				}
			}
			// decode prach_RootSequenceIndex_r16 optional
			if prach_RootSequenceIndex_r16Present {
				ie.prach_RootSequenceIndex_r16 = new(RACH_ConfigCommon_prach_RootSequenceIndex_r16)
				if err = ie.prach_RootSequenceIndex_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode prach_RootSequenceIndex_r16", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ra_PrioritizationForSlicing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			featureCombinationPreamblesList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ra_PrioritizationForSlicing_r17 optional
			if ra_PrioritizationForSlicing_r17Present {
				ie.ra_PrioritizationForSlicing_r17 = new(RA_PrioritizationForSlicing_r17)
				if err = ie.ra_PrioritizationForSlicing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ra_PrioritizationForSlicing_r17", err)
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
