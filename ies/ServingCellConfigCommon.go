package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfigCommon struct {
	physCellId                     *PhysCellId                                             `optional`
	downlinkConfigCommon           *DownlinkConfigCommon                                   `optional`
	uplinkConfigCommon             *UplinkConfigCommon                                     `optional`
	supplementaryUplinkConfig      *UplinkConfigCommon                                     `optional`
	n_TimingAdvanceOffset          *ServingCellConfigCommon_n_TimingAdvanceOffset          `optional`
	ssb_PositionsInBurst           *ServingCellConfigCommon_ssb_PositionsInBurst           `lb:4,ub:4,optional`
	ssb_periodicityServingCell     *ServingCellConfigCommon_ssb_periodicityServingCell     `optional`
	dmrs_TypeA_Position            ServingCellConfigCommon_dmrs_TypeA_Position             `madatory`
	lte_CRS_ToMatchAround          *RateMatchPatternLTE_CRS                                `optional,setuprelease`
	rateMatchPatternToAddModList   []RateMatchPattern                                      `lb:1,ub:maxNrofRateMatchPatterns,optional`
	rateMatchPatternToReleaseList  []RateMatchPatternId                                    `lb:1,ub:maxNrofRateMatchPatterns,optional`
	ssbSubcarrierSpacing           *SubcarrierSpacing                                      `optional`
	tdd_UL_DL_ConfigurationCommon  *TDD_UL_DL_ConfigCommon                                 `optional`
	ss_PBCH_BlockPower             int64                                                   `lb:-60,ub:50,madatory`
	channelAccessMode_r16          *ServingCellConfigCommon_channelAccessMode_r16          `optional,ext-1`
	discoveryBurstWindowLength_r16 *ServingCellConfigCommon_discoveryBurstWindowLength_r16 `optional,ext-1`
	ssb_PositionQCL_r16            *SSB_PositionQCL_Relation_r16                           `optional,ext-1`
	highSpeedConfig_r16            *HighSpeedConfig_r16                                    `optional,ext-1`
	highSpeedConfig_v1700          *HighSpeedConfig_v1700                                  `optional,ext-2`
	channelAccessMode2_r17         *ServingCellConfigCommon_channelAccessMode2_r17         `optional,ext-2`
	discoveryBurstWindowLength_r17 *ServingCellConfigCommon_discoveryBurstWindowLength_r17 `optional,ext-2`
	ssb_PositionQCL_r17            *SSB_PositionQCL_Relation_r17                           `optional,ext-2`
	highSpeedConfigFR2_r17         *HighSpeedConfigFR2_r17                                 `optional,ext-2`
	uplinkConfigCommon_v1700       *UplinkConfigCommon_v1700                               `optional,ext-2`
	ntn_Config_r17                 *NTN_Config_r17                                         `optional,ext-2`
	featurePriorities_r17          *FeaturePriorities_r17                                  `optional,ext-3`
}

func (ie *ServingCellConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.channelAccessMode_r16 != nil || ie.discoveryBurstWindowLength_r16 != nil || ie.ssb_PositionQCL_r16 != nil || ie.highSpeedConfig_r16 != nil || ie.highSpeedConfig_v1700 != nil || ie.channelAccessMode2_r17 != nil || ie.discoveryBurstWindowLength_r17 != nil || ie.ssb_PositionQCL_r17 != nil || ie.highSpeedConfigFR2_r17 != nil || ie.uplinkConfigCommon_v1700 != nil || ie.ntn_Config_r17 != nil || ie.featurePriorities_r17 != nil
	preambleBits := []bool{hasExtensions, ie.physCellId != nil, ie.downlinkConfigCommon != nil, ie.uplinkConfigCommon != nil, ie.supplementaryUplinkConfig != nil, ie.n_TimingAdvanceOffset != nil, ie.ssb_PositionsInBurst != nil, ie.ssb_periodicityServingCell != nil, ie.lte_CRS_ToMatchAround != nil, len(ie.rateMatchPatternToAddModList) > 0, len(ie.rateMatchPatternToReleaseList) > 0, ie.ssbSubcarrierSpacing != nil, ie.tdd_UL_DL_ConfigurationCommon != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.physCellId != nil {
		if err = ie.physCellId.Encode(w); err != nil {
			return utils.WrapError("Encode physCellId", err)
		}
	}
	if ie.downlinkConfigCommon != nil {
		if err = ie.downlinkConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode downlinkConfigCommon", err)
		}
	}
	if ie.uplinkConfigCommon != nil {
		if err = ie.uplinkConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkConfigCommon", err)
		}
	}
	if ie.supplementaryUplinkConfig != nil {
		if err = ie.supplementaryUplinkConfig.Encode(w); err != nil {
			return utils.WrapError("Encode supplementaryUplinkConfig", err)
		}
	}
	if ie.n_TimingAdvanceOffset != nil {
		if err = ie.n_TimingAdvanceOffset.Encode(w); err != nil {
			return utils.WrapError("Encode n_TimingAdvanceOffset", err)
		}
	}
	if ie.ssb_PositionsInBurst != nil {
		if err = ie.ssb_PositionsInBurst.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_PositionsInBurst", err)
		}
	}
	if ie.ssb_periodicityServingCell != nil {
		if err = ie.ssb_periodicityServingCell.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_periodicityServingCell", err)
		}
	}
	if err = ie.dmrs_TypeA_Position.Encode(w); err != nil {
		return utils.WrapError("Encode dmrs_TypeA_Position", err)
	}
	if ie.lte_CRS_ToMatchAround != nil {
		tmp_lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{
			Setup: ie.lte_CRS_ToMatchAround,
		}
		if err = tmp_lte_CRS_ToMatchAround.Encode(w); err != nil {
			return utils.WrapError("Encode lte_CRS_ToMatchAround", err)
		}
	}
	if len(ie.rateMatchPatternToAddModList) > 0 {
		tmp_rateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.rateMatchPatternToAddModList {
			tmp_rateMatchPatternToAddModList.Value = append(tmp_rateMatchPatternToAddModList.Value, &i)
		}
		if err = tmp_rateMatchPatternToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchPatternToAddModList", err)
		}
	}
	if len(ie.rateMatchPatternToReleaseList) > 0 {
		tmp_rateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.rateMatchPatternToReleaseList {
			tmp_rateMatchPatternToReleaseList.Value = append(tmp_rateMatchPatternToReleaseList.Value, &i)
		}
		if err = tmp_rateMatchPatternToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchPatternToReleaseList", err)
		}
	}
	if ie.ssbSubcarrierSpacing != nil {
		if err = ie.ssbSubcarrierSpacing.Encode(w); err != nil {
			return utils.WrapError("Encode ssbSubcarrierSpacing", err)
		}
	}
	if ie.tdd_UL_DL_ConfigurationCommon != nil {
		if err = ie.tdd_UL_DL_ConfigurationCommon.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	if err = w.WriteInteger(ie.ss_PBCH_BlockPower, &uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("WriteInteger ss_PBCH_BlockPower", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.channelAccessMode_r16 != nil || ie.discoveryBurstWindowLength_r16 != nil || ie.ssb_PositionQCL_r16 != nil || ie.highSpeedConfig_r16 != nil, ie.highSpeedConfig_v1700 != nil || ie.channelAccessMode2_r17 != nil || ie.discoveryBurstWindowLength_r17 != nil || ie.ssb_PositionQCL_r17 != nil || ie.highSpeedConfigFR2_r17 != nil || ie.uplinkConfigCommon_v1700 != nil || ie.ntn_Config_r17 != nil, ie.featurePriorities_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ServingCellConfigCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.channelAccessMode_r16 != nil, ie.discoveryBurstWindowLength_r16 != nil, ie.ssb_PositionQCL_r16 != nil, ie.highSpeedConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode channelAccessMode_r16 optional
			if ie.channelAccessMode_r16 != nil {
				if err = ie.channelAccessMode_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelAccessMode_r16", err)
				}
			}
			// encode discoveryBurstWindowLength_r16 optional
			if ie.discoveryBurstWindowLength_r16 != nil {
				if err = ie.discoveryBurstWindowLength_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode discoveryBurstWindowLength_r16", err)
				}
			}
			// encode ssb_PositionQCL_r16 optional
			if ie.ssb_PositionQCL_r16 != nil {
				if err = ie.ssb_PositionQCL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_PositionQCL_r16", err)
				}
			}
			// encode highSpeedConfig_r16 optional
			if ie.highSpeedConfig_r16 != nil {
				if err = ie.highSpeedConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode highSpeedConfig_r16", err)
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
			optionals_ext_2 := []bool{ie.highSpeedConfig_v1700 != nil, ie.channelAccessMode2_r17 != nil, ie.discoveryBurstWindowLength_r17 != nil, ie.ssb_PositionQCL_r17 != nil, ie.highSpeedConfigFR2_r17 != nil, ie.uplinkConfigCommon_v1700 != nil, ie.ntn_Config_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode highSpeedConfig_v1700 optional
			if ie.highSpeedConfig_v1700 != nil {
				if err = ie.highSpeedConfig_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode highSpeedConfig_v1700", err)
				}
			}
			// encode channelAccessMode2_r17 optional
			if ie.channelAccessMode2_r17 != nil {
				if err = ie.channelAccessMode2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelAccessMode2_r17", err)
				}
			}
			// encode discoveryBurstWindowLength_r17 optional
			if ie.discoveryBurstWindowLength_r17 != nil {
				if err = ie.discoveryBurstWindowLength_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode discoveryBurstWindowLength_r17", err)
				}
			}
			// encode ssb_PositionQCL_r17 optional
			if ie.ssb_PositionQCL_r17 != nil {
				if err = ie.ssb_PositionQCL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_PositionQCL_r17", err)
				}
			}
			// encode highSpeedConfigFR2_r17 optional
			if ie.highSpeedConfigFR2_r17 != nil {
				if err = ie.highSpeedConfigFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode highSpeedConfigFR2_r17", err)
				}
			}
			// encode uplinkConfigCommon_v1700 optional
			if ie.uplinkConfigCommon_v1700 != nil {
				if err = ie.uplinkConfigCommon_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkConfigCommon_v1700", err)
				}
			}
			// encode ntn_Config_r17 optional
			if ie.ntn_Config_r17 != nil {
				if err = ie.ntn_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ntn_Config_r17", err)
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
			optionals_ext_3 := []bool{ie.featurePriorities_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode featurePriorities_r17 optional
			if ie.featurePriorities_r17 != nil {
				if err = ie.featurePriorities_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode featurePriorities_r17", err)
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

func (ie *ServingCellConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var physCellIdPresent bool
	if physCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var downlinkConfigCommonPresent bool
	if downlinkConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkConfigCommonPresent bool
	if uplinkConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supplementaryUplinkConfigPresent bool
	if supplementaryUplinkConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var n_TimingAdvanceOffsetPresent bool
	if n_TimingAdvanceOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_PositionsInBurstPresent bool
	if ssb_PositionsInBurstPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_periodicityServingCellPresent bool
	if ssb_periodicityServingCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lte_CRS_ToMatchAroundPresent bool
	if lte_CRS_ToMatchAroundPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchPatternToAddModListPresent bool
	if rateMatchPatternToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchPatternToReleaseListPresent bool
	if rateMatchPatternToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssbSubcarrierSpacingPresent bool
	if ssbSubcarrierSpacingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_UL_DL_ConfigurationCommonPresent bool
	if tdd_UL_DL_ConfigurationCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if physCellIdPresent {
		ie.physCellId = new(PhysCellId)
		if err = ie.physCellId.Decode(r); err != nil {
			return utils.WrapError("Decode physCellId", err)
		}
	}
	if downlinkConfigCommonPresent {
		ie.downlinkConfigCommon = new(DownlinkConfigCommon)
		if err = ie.downlinkConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode downlinkConfigCommon", err)
		}
	}
	if uplinkConfigCommonPresent {
		ie.uplinkConfigCommon = new(UplinkConfigCommon)
		if err = ie.uplinkConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkConfigCommon", err)
		}
	}
	if supplementaryUplinkConfigPresent {
		ie.supplementaryUplinkConfig = new(UplinkConfigCommon)
		if err = ie.supplementaryUplinkConfig.Decode(r); err != nil {
			return utils.WrapError("Decode supplementaryUplinkConfig", err)
		}
	}
	if n_TimingAdvanceOffsetPresent {
		ie.n_TimingAdvanceOffset = new(ServingCellConfigCommon_n_TimingAdvanceOffset)
		if err = ie.n_TimingAdvanceOffset.Decode(r); err != nil {
			return utils.WrapError("Decode n_TimingAdvanceOffset", err)
		}
	}
	if ssb_PositionsInBurstPresent {
		ie.ssb_PositionsInBurst = new(ServingCellConfigCommon_ssb_PositionsInBurst)
		if err = ie.ssb_PositionsInBurst.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_PositionsInBurst", err)
		}
	}
	if ssb_periodicityServingCellPresent {
		ie.ssb_periodicityServingCell = new(ServingCellConfigCommon_ssb_periodicityServingCell)
		if err = ie.ssb_periodicityServingCell.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_periodicityServingCell", err)
		}
	}
	if err = ie.dmrs_TypeA_Position.Decode(r); err != nil {
		return utils.WrapError("Decode dmrs_TypeA_Position", err)
	}
	if lte_CRS_ToMatchAroundPresent {
		tmp_lte_CRS_ToMatchAround := utils.SetupRelease[*RateMatchPatternLTE_CRS]{}
		if err = tmp_lte_CRS_ToMatchAround.Decode(r); err != nil {
			return utils.WrapError("Decode lte_CRS_ToMatchAround", err)
		}
		ie.lte_CRS_ToMatchAround = tmp_lte_CRS_ToMatchAround.Setup
	}
	if rateMatchPatternToAddModListPresent {
		tmp_rateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_rateMatchPatternToAddModList := func() *RateMatchPattern {
			return new(RateMatchPattern)
		}
		if err = tmp_rateMatchPatternToAddModList.Decode(r, fn_rateMatchPatternToAddModList); err != nil {
			return utils.WrapError("Decode rateMatchPatternToAddModList", err)
		}
		ie.rateMatchPatternToAddModList = []RateMatchPattern{}
		for _, i := range tmp_rateMatchPatternToAddModList.Value {
			ie.rateMatchPatternToAddModList = append(ie.rateMatchPatternToAddModList, *i)
		}
	}
	if rateMatchPatternToReleaseListPresent {
		tmp_rateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_rateMatchPatternToReleaseList := func() *RateMatchPatternId {
			return new(RateMatchPatternId)
		}
		if err = tmp_rateMatchPatternToReleaseList.Decode(r, fn_rateMatchPatternToReleaseList); err != nil {
			return utils.WrapError("Decode rateMatchPatternToReleaseList", err)
		}
		ie.rateMatchPatternToReleaseList = []RateMatchPatternId{}
		for _, i := range tmp_rateMatchPatternToReleaseList.Value {
			ie.rateMatchPatternToReleaseList = append(ie.rateMatchPatternToReleaseList, *i)
		}
	}
	if ssbSubcarrierSpacingPresent {
		ie.ssbSubcarrierSpacing = new(SubcarrierSpacing)
		if err = ie.ssbSubcarrierSpacing.Decode(r); err != nil {
			return utils.WrapError("Decode ssbSubcarrierSpacing", err)
		}
	}
	if tdd_UL_DL_ConfigurationCommonPresent {
		ie.tdd_UL_DL_ConfigurationCommon = new(TDD_UL_DL_ConfigCommon)
		if err = ie.tdd_UL_DL_ConfigurationCommon.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_UL_DL_ConfigurationCommon", err)
		}
	}
	var tmp_int_ss_PBCH_BlockPower int64
	if tmp_int_ss_PBCH_BlockPower, err = r.ReadInteger(&uper.Constraint{Lb: -60, Ub: 50}, false); err != nil {
		return utils.WrapError("ReadInteger ss_PBCH_BlockPower", err)
	}
	ie.ss_PBCH_BlockPower = tmp_int_ss_PBCH_BlockPower

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

			channelAccessMode_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			discoveryBurstWindowLength_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssb_PositionQCL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			highSpeedConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode channelAccessMode_r16 optional
			if channelAccessMode_r16Present {
				ie.channelAccessMode_r16 = new(ServingCellConfigCommon_channelAccessMode_r16)
				if err = ie.channelAccessMode_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelAccessMode_r16", err)
				}
			}
			// decode discoveryBurstWindowLength_r16 optional
			if discoveryBurstWindowLength_r16Present {
				ie.discoveryBurstWindowLength_r16 = new(ServingCellConfigCommon_discoveryBurstWindowLength_r16)
				if err = ie.discoveryBurstWindowLength_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode discoveryBurstWindowLength_r16", err)
				}
			}
			// decode ssb_PositionQCL_r16 optional
			if ssb_PositionQCL_r16Present {
				ie.ssb_PositionQCL_r16 = new(SSB_PositionQCL_Relation_r16)
				if err = ie.ssb_PositionQCL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_PositionQCL_r16", err)
				}
			}
			// decode highSpeedConfig_r16 optional
			if highSpeedConfig_r16Present {
				ie.highSpeedConfig_r16 = new(HighSpeedConfig_r16)
				if err = ie.highSpeedConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode highSpeedConfig_r16", err)
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

			highSpeedConfig_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelAccessMode2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			discoveryBurstWindowLength_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssb_PositionQCL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			highSpeedConfigFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkConfigCommon_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ntn_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode highSpeedConfig_v1700 optional
			if highSpeedConfig_v1700Present {
				ie.highSpeedConfig_v1700 = new(HighSpeedConfig_v1700)
				if err = ie.highSpeedConfig_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode highSpeedConfig_v1700", err)
				}
			}
			// decode channelAccessMode2_r17 optional
			if channelAccessMode2_r17Present {
				ie.channelAccessMode2_r17 = new(ServingCellConfigCommon_channelAccessMode2_r17)
				if err = ie.channelAccessMode2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelAccessMode2_r17", err)
				}
			}
			// decode discoveryBurstWindowLength_r17 optional
			if discoveryBurstWindowLength_r17Present {
				ie.discoveryBurstWindowLength_r17 = new(ServingCellConfigCommon_discoveryBurstWindowLength_r17)
				if err = ie.discoveryBurstWindowLength_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode discoveryBurstWindowLength_r17", err)
				}
			}
			// decode ssb_PositionQCL_r17 optional
			if ssb_PositionQCL_r17Present {
				ie.ssb_PositionQCL_r17 = new(SSB_PositionQCL_Relation_r17)
				if err = ie.ssb_PositionQCL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_PositionQCL_r17", err)
				}
			}
			// decode highSpeedConfigFR2_r17 optional
			if highSpeedConfigFR2_r17Present {
				ie.highSpeedConfigFR2_r17 = new(HighSpeedConfigFR2_r17)
				if err = ie.highSpeedConfigFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode highSpeedConfigFR2_r17", err)
				}
			}
			// decode uplinkConfigCommon_v1700 optional
			if uplinkConfigCommon_v1700Present {
				ie.uplinkConfigCommon_v1700 = new(UplinkConfigCommon_v1700)
				if err = ie.uplinkConfigCommon_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkConfigCommon_v1700", err)
				}
			}
			// decode ntn_Config_r17 optional
			if ntn_Config_r17Present {
				ie.ntn_Config_r17 = new(NTN_Config_r17)
				if err = ie.ntn_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ntn_Config_r17", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			featurePriorities_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode featurePriorities_r17 optional
			if featurePriorities_r17Present {
				ie.featurePriorities_r17 = new(FeaturePriorities_r17)
				if err = ie.featurePriorities_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode featurePriorities_r17", err)
				}
			}
		}
	}
	return nil
}
