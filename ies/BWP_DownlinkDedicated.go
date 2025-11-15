package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_DownlinkDedicated struct {
	pdcch_Config                          *PDCCH_Config                        `optional,setuprelease`
	pdsch_Config                          *PDSCH_Config                        `optional,setuprelease`
	sps_Config                            *SPS_Config                          `optional,setuprelease`
	radioLinkMonitoringConfig             *RadioLinkMonitoringConfig           `optional,setuprelease`
	sps_ConfigToAddModList_r16            *SPS_ConfigToAddModList_r16          `optional,ext-1`
	sps_ConfigToReleaseList_r16           *SPS_ConfigToReleaseList_r16         `optional,ext-1`
	sps_ConfigDeactivationStateList_r16   *SPS_ConfigDeactivationStateList_r16 `optional,ext-1`
	beamFailureRecoverySCellConfig_r16    *BeamFailureRecoveryRSConfig_r16     `optional,ext-1,setuprelease`
	sl_PDCCH_Config_r16                   *PDCCH_Config                        `optional,ext-1,setuprelease`
	sl_V2X_PDCCH_Config_r16               *PDCCH_Config                        `optional,ext-1,setuprelease`
	preConfGapStatus_r17                  *uper.BitString                      `lb:maxNrofGapId_r17,ub:maxNrofGapId_r17,optional,ext-2`
	beamFailureRecoverySpCellConfig_r17   *BeamFailureRecoveryRSConfig_r16     `optional,ext-2,setuprelease`
	harq_FeedbackEnablingforSPSactive_r17 *bool                                `optional,ext-2`
	cfr_ConfigMulticast_r17               *CFR_ConfigMulticast_r17             `optional,ext-2,setuprelease`
	dl_PPW_PreConfigToAddModList_r17      *DL_PPW_PreConfigToAddModList_r17    `optional,ext-2`
	dl_PPW_PreConfigToReleaseList_r17     *DL_PPW_PreConfigToReleaseList_r17   `optional,ext-2`
	nonCellDefiningSSB_r17                *NonCellDefiningSSB_r17              `optional,ext-2`
	servingCellMO_r17                     *MeasObjectId                        `optional,ext-2`
}

func (ie *BWP_DownlinkDedicated) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sps_ConfigToAddModList_r16 != nil || ie.sps_ConfigToReleaseList_r16 != nil || ie.sps_ConfigDeactivationStateList_r16 != nil || ie.beamFailureRecoverySCellConfig_r16 != nil || ie.sl_PDCCH_Config_r16 != nil || ie.sl_V2X_PDCCH_Config_r16 != nil || ie.preConfGapStatus_r17 != nil || ie.beamFailureRecoverySpCellConfig_r17 != nil || ie.harq_FeedbackEnablingforSPSactive_r17 != nil || ie.cfr_ConfigMulticast_r17 != nil || ie.dl_PPW_PreConfigToAddModList_r17 != nil || ie.dl_PPW_PreConfigToReleaseList_r17 != nil || ie.nonCellDefiningSSB_r17 != nil || ie.servingCellMO_r17 != nil
	preambleBits := []bool{hasExtensions, ie.pdcch_Config != nil, ie.pdsch_Config != nil, ie.sps_Config != nil, ie.radioLinkMonitoringConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcch_Config != nil {
		tmp_pdcch_Config := utils.SetupRelease[*PDCCH_Config]{
			Setup: ie.pdcch_Config,
		}
		if err = tmp_pdcch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode pdcch_Config", err)
		}
	}
	if ie.pdsch_Config != nil {
		tmp_pdsch_Config := utils.SetupRelease[*PDSCH_Config]{
			Setup: ie.pdsch_Config,
		}
		if err = tmp_pdsch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_Config", err)
		}
	}
	if ie.sps_Config != nil {
		tmp_sps_Config := utils.SetupRelease[*SPS_Config]{
			Setup: ie.sps_Config,
		}
		if err = tmp_sps_Config.Encode(w); err != nil {
			return utils.WrapError("Encode sps_Config", err)
		}
	}
	if ie.radioLinkMonitoringConfig != nil {
		tmp_radioLinkMonitoringConfig := utils.SetupRelease[*RadioLinkMonitoringConfig]{
			Setup: ie.radioLinkMonitoringConfig,
		}
		if err = tmp_radioLinkMonitoringConfig.Encode(w); err != nil {
			return utils.WrapError("Encode radioLinkMonitoringConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.sps_ConfigToAddModList_r16 != nil || ie.sps_ConfigToReleaseList_r16 != nil || ie.sps_ConfigDeactivationStateList_r16 != nil || ie.beamFailureRecoverySCellConfig_r16 != nil || ie.sl_PDCCH_Config_r16 != nil || ie.sl_V2X_PDCCH_Config_r16 != nil, ie.preConfGapStatus_r17 != nil || ie.beamFailureRecoverySpCellConfig_r17 != nil || ie.harq_FeedbackEnablingforSPSactive_r17 != nil || ie.cfr_ConfigMulticast_r17 != nil || ie.dl_PPW_PreConfigToAddModList_r17 != nil || ie.dl_PPW_PreConfigToReleaseList_r17 != nil || ie.nonCellDefiningSSB_r17 != nil || ie.servingCellMO_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BWP_DownlinkDedicated", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sps_ConfigToAddModList_r16 != nil, ie.sps_ConfigToReleaseList_r16 != nil, ie.sps_ConfigDeactivationStateList_r16 != nil, ie.beamFailureRecoverySCellConfig_r16 != nil, ie.sl_PDCCH_Config_r16 != nil, ie.sl_V2X_PDCCH_Config_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sps_ConfigToAddModList_r16 optional
			if ie.sps_ConfigToAddModList_r16 != nil {
				if err = ie.sps_ConfigToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_ConfigToAddModList_r16", err)
				}
			}
			// encode sps_ConfigToReleaseList_r16 optional
			if ie.sps_ConfigToReleaseList_r16 != nil {
				if err = ie.sps_ConfigToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_ConfigToReleaseList_r16", err)
				}
			}
			// encode sps_ConfigDeactivationStateList_r16 optional
			if ie.sps_ConfigDeactivationStateList_r16 != nil {
				if err = ie.sps_ConfigDeactivationStateList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_ConfigDeactivationStateList_r16", err)
				}
			}
			// encode beamFailureRecoverySCellConfig_r16 optional
			if ie.beamFailureRecoverySCellConfig_r16 != nil {
				tmp_beamFailureRecoverySCellConfig_r16 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{
					Setup: ie.beamFailureRecoverySCellConfig_r16,
				}
				if err = tmp_beamFailureRecoverySCellConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamFailureRecoverySCellConfig_r16", err)
				}
			}
			// encode sl_PDCCH_Config_r16 optional
			if ie.sl_PDCCH_Config_r16 != nil {
				tmp_sl_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{
					Setup: ie.sl_PDCCH_Config_r16,
				}
				if err = tmp_sl_PDCCH_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_PDCCH_Config_r16", err)
				}
			}
			// encode sl_V2X_PDCCH_Config_r16 optional
			if ie.sl_V2X_PDCCH_Config_r16 != nil {
				tmp_sl_V2X_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{
					Setup: ie.sl_V2X_PDCCH_Config_r16,
				}
				if err = tmp_sl_V2X_PDCCH_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_V2X_PDCCH_Config_r16", err)
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
			optionals_ext_2 := []bool{ie.preConfGapStatus_r17 != nil, ie.beamFailureRecoverySpCellConfig_r17 != nil, ie.harq_FeedbackEnablingforSPSactive_r17 != nil, ie.cfr_ConfigMulticast_r17 != nil, ie.dl_PPW_PreConfigToAddModList_r17 != nil, ie.dl_PPW_PreConfigToReleaseList_r17 != nil, ie.nonCellDefiningSSB_r17 != nil, ie.servingCellMO_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode preConfGapStatus_r17 optional
			if ie.preConfGapStatus_r17 != nil {
				if err = extWriter.WriteBitString(ie.preConfGapStatus_r17.Bytes, uint(ie.preConfGapStatus_r17.NumBits), &uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Encode preConfGapStatus_r17", err)
				}
			}
			// encode beamFailureRecoverySpCellConfig_r17 optional
			if ie.beamFailureRecoverySpCellConfig_r17 != nil {
				tmp_beamFailureRecoverySpCellConfig_r17 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{
					Setup: ie.beamFailureRecoverySpCellConfig_r17,
				}
				if err = tmp_beamFailureRecoverySpCellConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamFailureRecoverySpCellConfig_r17", err)
				}
			}
			// encode harq_FeedbackEnablingforSPSactive_r17 optional
			if ie.harq_FeedbackEnablingforSPSactive_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.harq_FeedbackEnablingforSPSactive_r17); err != nil {
					return utils.WrapError("Encode harq_FeedbackEnablingforSPSactive_r17", err)
				}
			}
			// encode cfr_ConfigMulticast_r17 optional
			if ie.cfr_ConfigMulticast_r17 != nil {
				tmp_cfr_ConfigMulticast_r17 := utils.SetupRelease[*CFR_ConfigMulticast_r17]{
					Setup: ie.cfr_ConfigMulticast_r17,
				}
				if err = tmp_cfr_ConfigMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cfr_ConfigMulticast_r17", err)
				}
			}
			// encode dl_PPW_PreConfigToAddModList_r17 optional
			if ie.dl_PPW_PreConfigToAddModList_r17 != nil {
				if err = ie.dl_PPW_PreConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_PPW_PreConfigToAddModList_r17", err)
				}
			}
			// encode dl_PPW_PreConfigToReleaseList_r17 optional
			if ie.dl_PPW_PreConfigToReleaseList_r17 != nil {
				if err = ie.dl_PPW_PreConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_PPW_PreConfigToReleaseList_r17", err)
				}
			}
			// encode nonCellDefiningSSB_r17 optional
			if ie.nonCellDefiningSSB_r17 != nil {
				if err = ie.nonCellDefiningSSB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nonCellDefiningSSB_r17", err)
				}
			}
			// encode servingCellMO_r17 optional
			if ie.servingCellMO_r17 != nil {
				if err = ie.servingCellMO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode servingCellMO_r17", err)
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

func (ie *BWP_DownlinkDedicated) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_ConfigPresent bool
	if pdcch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_ConfigPresent bool
	if pdsch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sps_ConfigPresent bool
	if sps_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var radioLinkMonitoringConfigPresent bool
	if radioLinkMonitoringConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcch_ConfigPresent {
		tmp_pdcch_Config := utils.SetupRelease[*PDCCH_Config]{}
		if err = tmp_pdcch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode pdcch_Config", err)
		}
		ie.pdcch_Config = tmp_pdcch_Config.Setup
	}
	if pdsch_ConfigPresent {
		tmp_pdsch_Config := utils.SetupRelease[*PDSCH_Config]{}
		if err = tmp_pdsch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_Config", err)
		}
		ie.pdsch_Config = tmp_pdsch_Config.Setup
	}
	if sps_ConfigPresent {
		tmp_sps_Config := utils.SetupRelease[*SPS_Config]{}
		if err = tmp_sps_Config.Decode(r); err != nil {
			return utils.WrapError("Decode sps_Config", err)
		}
		ie.sps_Config = tmp_sps_Config.Setup
	}
	if radioLinkMonitoringConfigPresent {
		tmp_radioLinkMonitoringConfig := utils.SetupRelease[*RadioLinkMonitoringConfig]{}
		if err = tmp_radioLinkMonitoringConfig.Decode(r); err != nil {
			return utils.WrapError("Decode radioLinkMonitoringConfig", err)
		}
		ie.radioLinkMonitoringConfig = tmp_radioLinkMonitoringConfig.Setup
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

			sps_ConfigToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_ConfigToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_ConfigDeactivationStateList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamFailureRecoverySCellConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_PDCCH_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_V2X_PDCCH_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sps_ConfigToAddModList_r16 optional
			if sps_ConfigToAddModList_r16Present {
				ie.sps_ConfigToAddModList_r16 = new(SPS_ConfigToAddModList_r16)
				if err = ie.sps_ConfigToAddModList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_ConfigToAddModList_r16", err)
				}
			}
			// decode sps_ConfigToReleaseList_r16 optional
			if sps_ConfigToReleaseList_r16Present {
				ie.sps_ConfigToReleaseList_r16 = new(SPS_ConfigToReleaseList_r16)
				if err = ie.sps_ConfigToReleaseList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_ConfigToReleaseList_r16", err)
				}
			}
			// decode sps_ConfigDeactivationStateList_r16 optional
			if sps_ConfigDeactivationStateList_r16Present {
				ie.sps_ConfigDeactivationStateList_r16 = new(SPS_ConfigDeactivationStateList_r16)
				if err = ie.sps_ConfigDeactivationStateList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_ConfigDeactivationStateList_r16", err)
				}
			}
			// decode beamFailureRecoverySCellConfig_r16 optional
			if beamFailureRecoverySCellConfig_r16Present {
				tmp_beamFailureRecoverySCellConfig_r16 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{}
				if err = tmp_beamFailureRecoverySCellConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamFailureRecoverySCellConfig_r16", err)
				}
				ie.beamFailureRecoverySCellConfig_r16 = tmp_beamFailureRecoverySCellConfig_r16.Setup
			}
			// decode sl_PDCCH_Config_r16 optional
			if sl_PDCCH_Config_r16Present {
				tmp_sl_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{}
				if err = tmp_sl_PDCCH_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_PDCCH_Config_r16", err)
				}
				ie.sl_PDCCH_Config_r16 = tmp_sl_PDCCH_Config_r16.Setup
			}
			// decode sl_V2X_PDCCH_Config_r16 optional
			if sl_V2X_PDCCH_Config_r16Present {
				tmp_sl_V2X_PDCCH_Config_r16 := utils.SetupRelease[*PDCCH_Config]{}
				if err = tmp_sl_V2X_PDCCH_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_V2X_PDCCH_Config_r16", err)
				}
				ie.sl_V2X_PDCCH_Config_r16 = tmp_sl_V2X_PDCCH_Config_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			preConfGapStatus_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamFailureRecoverySpCellConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_FeedbackEnablingforSPSactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cfr_ConfigMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_PPW_PreConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_PPW_PreConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nonCellDefiningSSB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			servingCellMO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode preConfGapStatus_r17 optional
			if preConfGapStatus_r17Present {
				var tmp_bs_preConfGapStatus_r17 []byte
				var n_preConfGapStatus_r17 uint
				if tmp_bs_preConfGapStatus_r17, n_preConfGapStatus_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Decode preConfGapStatus_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_preConfGapStatus_r17,
					NumBits: uint64(n_preConfGapStatus_r17),
				}
				ie.preConfGapStatus_r17 = &tmp_bitstring
			}
			// decode beamFailureRecoverySpCellConfig_r17 optional
			if beamFailureRecoverySpCellConfig_r17Present {
				tmp_beamFailureRecoverySpCellConfig_r17 := utils.SetupRelease[*BeamFailureRecoveryRSConfig_r16]{}
				if err = tmp_beamFailureRecoverySpCellConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamFailureRecoverySpCellConfig_r17", err)
				}
				ie.beamFailureRecoverySpCellConfig_r17 = tmp_beamFailureRecoverySpCellConfig_r17.Setup
			}
			// decode harq_FeedbackEnablingforSPSactive_r17 optional
			if harq_FeedbackEnablingforSPSactive_r17Present {
				var tmp_bool_harq_FeedbackEnablingforSPSactive_r17 bool
				if tmp_bool_harq_FeedbackEnablingforSPSactive_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode harq_FeedbackEnablingforSPSactive_r17", err)
				}
				ie.harq_FeedbackEnablingforSPSactive_r17 = &tmp_bool_harq_FeedbackEnablingforSPSactive_r17
			}
			// decode cfr_ConfigMulticast_r17 optional
			if cfr_ConfigMulticast_r17Present {
				tmp_cfr_ConfigMulticast_r17 := utils.SetupRelease[*CFR_ConfigMulticast_r17]{}
				if err = tmp_cfr_ConfigMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cfr_ConfigMulticast_r17", err)
				}
				ie.cfr_ConfigMulticast_r17 = tmp_cfr_ConfigMulticast_r17.Setup
			}
			// decode dl_PPW_PreConfigToAddModList_r17 optional
			if dl_PPW_PreConfigToAddModList_r17Present {
				ie.dl_PPW_PreConfigToAddModList_r17 = new(DL_PPW_PreConfigToAddModList_r17)
				if err = ie.dl_PPW_PreConfigToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_PPW_PreConfigToAddModList_r17", err)
				}
			}
			// decode dl_PPW_PreConfigToReleaseList_r17 optional
			if dl_PPW_PreConfigToReleaseList_r17Present {
				ie.dl_PPW_PreConfigToReleaseList_r17 = new(DL_PPW_PreConfigToReleaseList_r17)
				if err = ie.dl_PPW_PreConfigToReleaseList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_PPW_PreConfigToReleaseList_r17", err)
				}
			}
			// decode nonCellDefiningSSB_r17 optional
			if nonCellDefiningSSB_r17Present {
				ie.nonCellDefiningSSB_r17 = new(NonCellDefiningSSB_r17)
				if err = ie.nonCellDefiningSSB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nonCellDefiningSSB_r17", err)
				}
			}
			// decode servingCellMO_r17 optional
			if servingCellMO_r17Present {
				ie.servingCellMO_r17 = new(MeasObjectId)
				if err = ie.servingCellMO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode servingCellMO_r17", err)
				}
			}
		}
	}
	return nil
}
