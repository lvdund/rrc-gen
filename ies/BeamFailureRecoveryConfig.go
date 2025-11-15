package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BeamFailureRecoveryConfig struct {
	rootSequenceIndex_BFR        *int64                                              `lb:0,ub:137,optional`
	rach_ConfigBFR               *RACH_ConfigGeneric                                 `optional`
	rsrp_ThresholdSSB            *RSRP_Range                                         `optional`
	candidateBeamRSList          []PRACH_ResourceDedicatedBFR                        `lb:1,ub:maxNrofCandidateBeams,optional`
	ssb_perRACH_Occasion         *BeamFailureRecoveryConfig_ssb_perRACH_Occasion     `optional`
	ra_ssb_OccasionMaskIndex     *int64                                              `lb:0,ub:15,optional`
	recoverySearchSpaceId        *SearchSpaceId                                      `optional`
	ra_Prioritization            *RA_Prioritization                                  `optional`
	beamFailureRecoveryTimer     *BeamFailureRecoveryConfig_beamFailureRecoveryTimer `optional`
	msg1_SubcarrierSpacing       *SubcarrierSpacing                                  `optional,ext-1`
	ra_PrioritizationTwoStep_r16 *RA_Prioritization                                  `optional,ext-2`
	candidateBeamRSListExt_v1610 *CandidateBeamRSListExt_r16                         `optional,ext-2,setuprelease`
	spCell_BFR_CBRA_r16          *BeamFailureRecoveryConfig_spCell_BFR_CBRA_r16      `optional,ext-3`
}

func (ie *BeamFailureRecoveryConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.msg1_SubcarrierSpacing != nil || ie.ra_PrioritizationTwoStep_r16 != nil || ie.candidateBeamRSListExt_v1610 != nil || ie.spCell_BFR_CBRA_r16 != nil
	preambleBits := []bool{hasExtensions, ie.rootSequenceIndex_BFR != nil, ie.rach_ConfigBFR != nil, ie.rsrp_ThresholdSSB != nil, len(ie.candidateBeamRSList) > 0, ie.ssb_perRACH_Occasion != nil, ie.ra_ssb_OccasionMaskIndex != nil, ie.recoverySearchSpaceId != nil, ie.ra_Prioritization != nil, ie.beamFailureRecoveryTimer != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rootSequenceIndex_BFR != nil {
		if err = w.WriteInteger(*ie.rootSequenceIndex_BFR, &uper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			return utils.WrapError("Encode rootSequenceIndex_BFR", err)
		}
	}
	if ie.rach_ConfigBFR != nil {
		if err = ie.rach_ConfigBFR.Encode(w); err != nil {
			return utils.WrapError("Encode rach_ConfigBFR", err)
		}
	}
	if ie.rsrp_ThresholdSSB != nil {
		if err = ie.rsrp_ThresholdSSB.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_ThresholdSSB", err)
		}
	}
	if len(ie.candidateBeamRSList) > 0 {
		tmp_candidateBeamRSList := utils.NewSequence[*PRACH_ResourceDedicatedBFR]([]*PRACH_ResourceDedicatedBFR{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams}, false)
		for _, i := range ie.candidateBeamRSList {
			tmp_candidateBeamRSList.Value = append(tmp_candidateBeamRSList.Value, &i)
		}
		if err = tmp_candidateBeamRSList.Encode(w); err != nil {
			return utils.WrapError("Encode candidateBeamRSList", err)
		}
	}
	if ie.ssb_perRACH_Occasion != nil {
		if err = ie.ssb_perRACH_Occasion.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_perRACH_Occasion", err)
		}
	}
	if ie.ra_ssb_OccasionMaskIndex != nil {
		if err = w.WriteInteger(*ie.ra_ssb_OccasionMaskIndex, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode ra_ssb_OccasionMaskIndex", err)
		}
	}
	if ie.recoverySearchSpaceId != nil {
		if err = ie.recoverySearchSpaceId.Encode(w); err != nil {
			return utils.WrapError("Encode recoverySearchSpaceId", err)
		}
	}
	if ie.ra_Prioritization != nil {
		if err = ie.ra_Prioritization.Encode(w); err != nil {
			return utils.WrapError("Encode ra_Prioritization", err)
		}
	}
	if ie.beamFailureRecoveryTimer != nil {
		if err = ie.beamFailureRecoveryTimer.Encode(w); err != nil {
			return utils.WrapError("Encode beamFailureRecoveryTimer", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.msg1_SubcarrierSpacing != nil, ie.ra_PrioritizationTwoStep_r16 != nil || ie.candidateBeamRSListExt_v1610 != nil, ie.spCell_BFR_CBRA_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BeamFailureRecoveryConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.msg1_SubcarrierSpacing != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode msg1_SubcarrierSpacing optional
			if ie.msg1_SubcarrierSpacing != nil {
				if err = ie.msg1_SubcarrierSpacing.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msg1_SubcarrierSpacing", err)
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
			optionals_ext_2 := []bool{ie.ra_PrioritizationTwoStep_r16 != nil, ie.candidateBeamRSListExt_v1610 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ra_PrioritizationTwoStep_r16 optional
			if ie.ra_PrioritizationTwoStep_r16 != nil {
				if err = ie.ra_PrioritizationTwoStep_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ra_PrioritizationTwoStep_r16", err)
				}
			}
			// encode candidateBeamRSListExt_v1610 optional
			if ie.candidateBeamRSListExt_v1610 != nil {
				tmp_candidateBeamRSListExt_v1610 := utils.SetupRelease[*CandidateBeamRSListExt_r16]{
					Setup: ie.candidateBeamRSListExt_v1610,
				}
				if err = tmp_candidateBeamRSListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode candidateBeamRSListExt_v1610", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.spCell_BFR_CBRA_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode spCell_BFR_CBRA_r16 optional
			if ie.spCell_BFR_CBRA_r16 != nil {
				if err = ie.spCell_BFR_CBRA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spCell_BFR_CBRA_r16", err)
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

func (ie *BeamFailureRecoveryConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rootSequenceIndex_BFRPresent bool
	if rootSequenceIndex_BFRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rach_ConfigBFRPresent bool
	if rach_ConfigBFRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrp_ThresholdSSBPresent bool
	if rsrp_ThresholdSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateBeamRSListPresent bool
	if candidateBeamRSListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_perRACH_OccasionPresent bool
	if ssb_perRACH_OccasionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_ssb_OccasionMaskIndexPresent bool
	if ra_ssb_OccasionMaskIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var recoverySearchSpaceIdPresent bool
	if recoverySearchSpaceIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_PrioritizationPresent bool
	if ra_PrioritizationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var beamFailureRecoveryTimerPresent bool
	if beamFailureRecoveryTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if rootSequenceIndex_BFRPresent {
		var tmp_int_rootSequenceIndex_BFR int64
		if tmp_int_rootSequenceIndex_BFR, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			return utils.WrapError("Decode rootSequenceIndex_BFR", err)
		}
		ie.rootSequenceIndex_BFR = &tmp_int_rootSequenceIndex_BFR
	}
	if rach_ConfigBFRPresent {
		ie.rach_ConfigBFR = new(RACH_ConfigGeneric)
		if err = ie.rach_ConfigBFR.Decode(r); err != nil {
			return utils.WrapError("Decode rach_ConfigBFR", err)
		}
	}
	if rsrp_ThresholdSSBPresent {
		ie.rsrp_ThresholdSSB = new(RSRP_Range)
		if err = ie.rsrp_ThresholdSSB.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_ThresholdSSB", err)
		}
	}
	if candidateBeamRSListPresent {
		tmp_candidateBeamRSList := utils.NewSequence[*PRACH_ResourceDedicatedBFR]([]*PRACH_ResourceDedicatedBFR{}, uper.Constraint{Lb: 1, Ub: maxNrofCandidateBeams}, false)
		fn_candidateBeamRSList := func() *PRACH_ResourceDedicatedBFR {
			return new(PRACH_ResourceDedicatedBFR)
		}
		if err = tmp_candidateBeamRSList.Decode(r, fn_candidateBeamRSList); err != nil {
			return utils.WrapError("Decode candidateBeamRSList", err)
		}
		ie.candidateBeamRSList = []PRACH_ResourceDedicatedBFR{}
		for _, i := range tmp_candidateBeamRSList.Value {
			ie.candidateBeamRSList = append(ie.candidateBeamRSList, *i)
		}
	}
	if ssb_perRACH_OccasionPresent {
		ie.ssb_perRACH_Occasion = new(BeamFailureRecoveryConfig_ssb_perRACH_Occasion)
		if err = ie.ssb_perRACH_Occasion.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_perRACH_Occasion", err)
		}
	}
	if ra_ssb_OccasionMaskIndexPresent {
		var tmp_int_ra_ssb_OccasionMaskIndex int64
		if tmp_int_ra_ssb_OccasionMaskIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode ra_ssb_OccasionMaskIndex", err)
		}
		ie.ra_ssb_OccasionMaskIndex = &tmp_int_ra_ssb_OccasionMaskIndex
	}
	if recoverySearchSpaceIdPresent {
		ie.recoverySearchSpaceId = new(SearchSpaceId)
		if err = ie.recoverySearchSpaceId.Decode(r); err != nil {
			return utils.WrapError("Decode recoverySearchSpaceId", err)
		}
	}
	if ra_PrioritizationPresent {
		ie.ra_Prioritization = new(RA_Prioritization)
		if err = ie.ra_Prioritization.Decode(r); err != nil {
			return utils.WrapError("Decode ra_Prioritization", err)
		}
	}
	if beamFailureRecoveryTimerPresent {
		ie.beamFailureRecoveryTimer = new(BeamFailureRecoveryConfig_beamFailureRecoveryTimer)
		if err = ie.beamFailureRecoveryTimer.Decode(r); err != nil {
			return utils.WrapError("Decode beamFailureRecoveryTimer", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			msg1_SubcarrierSpacingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode msg1_SubcarrierSpacing optional
			if msg1_SubcarrierSpacingPresent {
				ie.msg1_SubcarrierSpacing = new(SubcarrierSpacing)
				if err = ie.msg1_SubcarrierSpacing.Decode(extReader); err != nil {
					return utils.WrapError("Decode msg1_SubcarrierSpacing", err)
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

			ra_PrioritizationTwoStep_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			candidateBeamRSListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ra_PrioritizationTwoStep_r16 optional
			if ra_PrioritizationTwoStep_r16Present {
				ie.ra_PrioritizationTwoStep_r16 = new(RA_Prioritization)
				if err = ie.ra_PrioritizationTwoStep_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ra_PrioritizationTwoStep_r16", err)
				}
			}
			// decode candidateBeamRSListExt_v1610 optional
			if candidateBeamRSListExt_v1610Present {
				tmp_candidateBeamRSListExt_v1610 := utils.SetupRelease[*CandidateBeamRSListExt_r16]{}
				if err = tmp_candidateBeamRSListExt_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode candidateBeamRSListExt_v1610", err)
				}
				ie.candidateBeamRSListExt_v1610 = tmp_candidateBeamRSListExt_v1610.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			spCell_BFR_CBRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode spCell_BFR_CBRA_r16 optional
			if spCell_BFR_CBRA_r16Present {
				ie.spCell_BFR_CBRA_r16 = new(BeamFailureRecoveryConfig_spCell_BFR_CBRA_r16)
				if err = ie.spCell_BFR_CBRA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode spCell_BFR_CBRA_r16", err)
				}
			}
		}
	}
	return nil
}
