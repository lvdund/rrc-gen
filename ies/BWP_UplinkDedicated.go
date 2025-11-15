package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_UplinkDedicated struct {
	pucch_Config                                        *PUCCH_Config                                        `optional,setuprelease`
	pusch_Config                                        *PUSCH_Config                                        `optional,setuprelease`
	configuredGrantConfig                               *ConfiguredGrantConfig                               `optional,setuprelease`
	srs_Config                                          *SRS_Config                                          `optional,setuprelease`
	beamFailureRecoveryConfig                           *BeamFailureRecoveryConfig                           `optional,setuprelease`
	sl_PUCCH_Config_r16                                 *PUCCH_Config                                        `optional,ext-1,setuprelease`
	cp_ExtensionC2_r16                                  *int64                                               `lb:1,ub:28,optional,ext-1`
	cp_ExtensionC3_r16                                  *int64                                               `lb:1,ub:28,optional,ext-1`
	useInterlacePUCCH_PUSCH_r16                         *BWP_UplinkDedicated_useInterlacePUCCH_PUSCH_r16     `optional,ext-1`
	pucch_ConfigurationList_r16                         *PUCCH_ConfigurationList_r16                         `optional,ext-1,setuprelease`
	lbt_FailureRecoveryConfig_r16                       *LBT_FailureRecoveryConfig_r16                       `optional,ext-1,setuprelease`
	configuredGrantConfigToAddModList_r16               *ConfiguredGrantConfigToAddModList_r16               `optional,ext-1`
	configuredGrantConfigToReleaseList_r16              *ConfiguredGrantConfigToReleaseList_r16              `optional,ext-1`
	configuredGrantConfigType2DeactivationStateList_r16 *ConfiguredGrantConfigType2DeactivationStateList_r16 `optional,ext-1`
	ul_TCI_StateList_r17                                *BWP_UplinkDedicated_ul_TCI_StateList_r17            `lb:1,ub:maxUL_TCI_r17,optional,ext-2`
	ul_powerControl_r17                                 *Uplink_powerControlId_r17                           `optional,ext-2`
	pucch_ConfigurationListMulticast1_r17               *PUCCH_ConfigurationList_r16                         `optional,ext-2,setuprelease`
	pucch_ConfigurationListMulticast2_r17               *PUCCH_ConfigurationList_r16                         `optional,ext-2,setuprelease`
	pucch_ConfigMulticast1_r17                          *PUCCH_Config                                        `optional,ext-3,setuprelease`
	pucch_ConfigMulticast2_r17                          *PUCCH_Config                                        `optional,ext-3,setuprelease`
	pathlossReferenceRSToAddModList_r17                 []PathlossReferenceRS_r17                            `lb:1,ub:maxNrofPathlossReferenceRSs_r17,optional,ext-4`
	pathlossReferenceRSToReleaseList_r17                []PathlossReferenceRS_Id_r17                         `lb:1,ub:maxNrofPathlossReferenceRSs_r17,optional,ext-4`
}

func (ie *BWP_UplinkDedicated) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_PUCCH_Config_r16 != nil || ie.cp_ExtensionC2_r16 != nil || ie.cp_ExtensionC3_r16 != nil || ie.useInterlacePUCCH_PUSCH_r16 != nil || ie.pucch_ConfigurationList_r16 != nil || ie.lbt_FailureRecoveryConfig_r16 != nil || ie.configuredGrantConfigToAddModList_r16 != nil || ie.configuredGrantConfigToReleaseList_r16 != nil || ie.configuredGrantConfigType2DeactivationStateList_r16 != nil || ie.ul_TCI_StateList_r17 != nil || ie.ul_powerControl_r17 != nil || ie.pucch_ConfigurationListMulticast1_r17 != nil || ie.pucch_ConfigurationListMulticast2_r17 != nil || ie.pucch_ConfigMulticast1_r17 != nil || ie.pucch_ConfigMulticast2_r17 != nil || len(ie.pathlossReferenceRSToAddModList_r17) > 0 || len(ie.pathlossReferenceRSToReleaseList_r17) > 0
	preambleBits := []bool{hasExtensions, ie.pucch_Config != nil, ie.pusch_Config != nil, ie.configuredGrantConfig != nil, ie.srs_Config != nil, ie.beamFailureRecoveryConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pucch_Config != nil {
		tmp_pucch_Config := utils.SetupRelease[*PUCCH_Config]{
			Setup: ie.pucch_Config,
		}
		if err = tmp_pucch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_Config", err)
		}
	}
	if ie.pusch_Config != nil {
		tmp_pusch_Config := utils.SetupRelease[*PUSCH_Config]{
			Setup: ie.pusch_Config,
		}
		if err = tmp_pusch_Config.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_Config", err)
		}
	}
	if ie.configuredGrantConfig != nil {
		tmp_configuredGrantConfig := utils.SetupRelease[*ConfiguredGrantConfig]{
			Setup: ie.configuredGrantConfig,
		}
		if err = tmp_configuredGrantConfig.Encode(w); err != nil {
			return utils.WrapError("Encode configuredGrantConfig", err)
		}
	}
	if ie.srs_Config != nil {
		tmp_srs_Config := utils.SetupRelease[*SRS_Config]{
			Setup: ie.srs_Config,
		}
		if err = tmp_srs_Config.Encode(w); err != nil {
			return utils.WrapError("Encode srs_Config", err)
		}
	}
	if ie.beamFailureRecoveryConfig != nil {
		tmp_beamFailureRecoveryConfig := utils.SetupRelease[*BeamFailureRecoveryConfig]{
			Setup: ie.beamFailureRecoveryConfig,
		}
		if err = tmp_beamFailureRecoveryConfig.Encode(w); err != nil {
			return utils.WrapError("Encode beamFailureRecoveryConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.sl_PUCCH_Config_r16 != nil || ie.cp_ExtensionC2_r16 != nil || ie.cp_ExtensionC3_r16 != nil || ie.useInterlacePUCCH_PUSCH_r16 != nil || ie.pucch_ConfigurationList_r16 != nil || ie.lbt_FailureRecoveryConfig_r16 != nil || ie.configuredGrantConfigToAddModList_r16 != nil || ie.configuredGrantConfigToReleaseList_r16 != nil || ie.configuredGrantConfigType2DeactivationStateList_r16 != nil, ie.ul_TCI_StateList_r17 != nil || ie.ul_powerControl_r17 != nil || ie.pucch_ConfigurationListMulticast1_r17 != nil || ie.pucch_ConfigurationListMulticast2_r17 != nil, ie.pucch_ConfigMulticast1_r17 != nil || ie.pucch_ConfigMulticast2_r17 != nil, len(ie.pathlossReferenceRSToAddModList_r17) > 0 || len(ie.pathlossReferenceRSToReleaseList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BWP_UplinkDedicated", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_PUCCH_Config_r16 != nil, ie.cp_ExtensionC2_r16 != nil, ie.cp_ExtensionC3_r16 != nil, ie.useInterlacePUCCH_PUSCH_r16 != nil, ie.pucch_ConfigurationList_r16 != nil, ie.lbt_FailureRecoveryConfig_r16 != nil, ie.configuredGrantConfigToAddModList_r16 != nil, ie.configuredGrantConfigToReleaseList_r16 != nil, ie.configuredGrantConfigType2DeactivationStateList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_PUCCH_Config_r16 optional
			if ie.sl_PUCCH_Config_r16 != nil {
				tmp_sl_PUCCH_Config_r16 := utils.SetupRelease[*PUCCH_Config]{
					Setup: ie.sl_PUCCH_Config_r16,
				}
				if err = tmp_sl_PUCCH_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_PUCCH_Config_r16", err)
				}
			}
			// encode cp_ExtensionC2_r16 optional
			if ie.cp_ExtensionC2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.cp_ExtensionC2_r16, &uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Encode cp_ExtensionC2_r16", err)
				}
			}
			// encode cp_ExtensionC3_r16 optional
			if ie.cp_ExtensionC3_r16 != nil {
				if err = extWriter.WriteInteger(*ie.cp_ExtensionC3_r16, &uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Encode cp_ExtensionC3_r16", err)
				}
			}
			// encode useInterlacePUCCH_PUSCH_r16 optional
			if ie.useInterlacePUCCH_PUSCH_r16 != nil {
				if err = ie.useInterlacePUCCH_PUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode useInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// encode pucch_ConfigurationList_r16 optional
			if ie.pucch_ConfigurationList_r16 != nil {
				tmp_pucch_ConfigurationList_r16 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{
					Setup: ie.pucch_ConfigurationList_r16,
				}
				if err = tmp_pucch_ConfigurationList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_ConfigurationList_r16", err)
				}
			}
			// encode lbt_FailureRecoveryConfig_r16 optional
			if ie.lbt_FailureRecoveryConfig_r16 != nil {
				tmp_lbt_FailureRecoveryConfig_r16 := utils.SetupRelease[*LBT_FailureRecoveryConfig_r16]{
					Setup: ie.lbt_FailureRecoveryConfig_r16,
				}
				if err = tmp_lbt_FailureRecoveryConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lbt_FailureRecoveryConfig_r16", err)
				}
			}
			// encode configuredGrantConfigToAddModList_r16 optional
			if ie.configuredGrantConfigToAddModList_r16 != nil {
				if err = ie.configuredGrantConfigToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configuredGrantConfigToAddModList_r16", err)
				}
			}
			// encode configuredGrantConfigToReleaseList_r16 optional
			if ie.configuredGrantConfigToReleaseList_r16 != nil {
				if err = ie.configuredGrantConfigToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configuredGrantConfigToReleaseList_r16", err)
				}
			}
			// encode configuredGrantConfigType2DeactivationStateList_r16 optional
			if ie.configuredGrantConfigType2DeactivationStateList_r16 != nil {
				if err = ie.configuredGrantConfigType2DeactivationStateList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configuredGrantConfigType2DeactivationStateList_r16", err)
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
			optionals_ext_2 := []bool{ie.ul_TCI_StateList_r17 != nil, ie.ul_powerControl_r17 != nil, ie.pucch_ConfigurationListMulticast1_r17 != nil, ie.pucch_ConfigurationListMulticast2_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ul_TCI_StateList_r17 optional
			if ie.ul_TCI_StateList_r17 != nil {
				if err = ie.ul_TCI_StateList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_TCI_StateList_r17", err)
				}
			}
			// encode ul_powerControl_r17 optional
			if ie.ul_powerControl_r17 != nil {
				if err = ie.ul_powerControl_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_powerControl_r17", err)
				}
			}
			// encode pucch_ConfigurationListMulticast1_r17 optional
			if ie.pucch_ConfigurationListMulticast1_r17 != nil {
				tmp_pucch_ConfigurationListMulticast1_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{
					Setup: ie.pucch_ConfigurationListMulticast1_r17,
				}
				if err = tmp_pucch_ConfigurationListMulticast1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_ConfigurationListMulticast1_r17", err)
				}
			}
			// encode pucch_ConfigurationListMulticast2_r17 optional
			if ie.pucch_ConfigurationListMulticast2_r17 != nil {
				tmp_pucch_ConfigurationListMulticast2_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{
					Setup: ie.pucch_ConfigurationListMulticast2_r17,
				}
				if err = tmp_pucch_ConfigurationListMulticast2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_ConfigurationListMulticast2_r17", err)
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
			optionals_ext_3 := []bool{ie.pucch_ConfigMulticast1_r17 != nil, ie.pucch_ConfigMulticast2_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pucch_ConfigMulticast1_r17 optional
			if ie.pucch_ConfigMulticast1_r17 != nil {
				tmp_pucch_ConfigMulticast1_r17 := utils.SetupRelease[*PUCCH_Config]{
					Setup: ie.pucch_ConfigMulticast1_r17,
				}
				if err = tmp_pucch_ConfigMulticast1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_ConfigMulticast1_r17", err)
				}
			}
			// encode pucch_ConfigMulticast2_r17 optional
			if ie.pucch_ConfigMulticast2_r17 != nil {
				tmp_pucch_ConfigMulticast2_r17 := utils.SetupRelease[*PUCCH_Config]{
					Setup: ie.pucch_ConfigMulticast2_r17,
				}
				if err = tmp_pucch_ConfigMulticast2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_ConfigMulticast2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{len(ie.pathlossReferenceRSToAddModList_r17) > 0, len(ie.pathlossReferenceRSToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pathlossReferenceRSToAddModList_r17 optional
			if len(ie.pathlossReferenceRSToAddModList_r17) > 0 {
				tmp_pathlossReferenceRSToAddModList_r17 := utils.NewSequence[*PathlossReferenceRS_r17]([]*PathlossReferenceRS_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				for _, i := range ie.pathlossReferenceRSToAddModList_r17 {
					tmp_pathlossReferenceRSToAddModList_r17.Value = append(tmp_pathlossReferenceRSToAddModList_r17.Value, &i)
				}
				if err = tmp_pathlossReferenceRSToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pathlossReferenceRSToAddModList_r17", err)
				}
			}
			// encode pathlossReferenceRSToReleaseList_r17 optional
			if len(ie.pathlossReferenceRSToReleaseList_r17) > 0 {
				tmp_pathlossReferenceRSToReleaseList_r17 := utils.NewSequence[*PathlossReferenceRS_Id_r17]([]*PathlossReferenceRS_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				for _, i := range ie.pathlossReferenceRSToReleaseList_r17 {
					tmp_pathlossReferenceRSToReleaseList_r17.Value = append(tmp_pathlossReferenceRSToReleaseList_r17.Value, &i)
				}
				if err = tmp_pathlossReferenceRSToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pathlossReferenceRSToReleaseList_r17", err)
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

func (ie *BWP_UplinkDedicated) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_ConfigPresent bool
	if pucch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_ConfigPresent bool
	if pusch_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredGrantConfigPresent bool
	if configuredGrantConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_ConfigPresent bool
	if srs_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var beamFailureRecoveryConfigPresent bool
	if beamFailureRecoveryConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pucch_ConfigPresent {
		tmp_pucch_Config := utils.SetupRelease[*PUCCH_Config]{}
		if err = tmp_pucch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_Config", err)
		}
		ie.pucch_Config = tmp_pucch_Config.Setup
	}
	if pusch_ConfigPresent {
		tmp_pusch_Config := utils.SetupRelease[*PUSCH_Config]{}
		if err = tmp_pusch_Config.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_Config", err)
		}
		ie.pusch_Config = tmp_pusch_Config.Setup
	}
	if configuredGrantConfigPresent {
		tmp_configuredGrantConfig := utils.SetupRelease[*ConfiguredGrantConfig]{}
		if err = tmp_configuredGrantConfig.Decode(r); err != nil {
			return utils.WrapError("Decode configuredGrantConfig", err)
		}
		ie.configuredGrantConfig = tmp_configuredGrantConfig.Setup
	}
	if srs_ConfigPresent {
		tmp_srs_Config := utils.SetupRelease[*SRS_Config]{}
		if err = tmp_srs_Config.Decode(r); err != nil {
			return utils.WrapError("Decode srs_Config", err)
		}
		ie.srs_Config = tmp_srs_Config.Setup
	}
	if beamFailureRecoveryConfigPresent {
		tmp_beamFailureRecoveryConfig := utils.SetupRelease[*BeamFailureRecoveryConfig]{}
		if err = tmp_beamFailureRecoveryConfig.Decode(r); err != nil {
			return utils.WrapError("Decode beamFailureRecoveryConfig", err)
		}
		ie.beamFailureRecoveryConfig = tmp_beamFailureRecoveryConfig.Setup
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			sl_PUCCH_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cp_ExtensionC2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cp_ExtensionC3_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			useInterlacePUCCH_PUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_ConfigurationList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lbt_FailureRecoveryConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredGrantConfigToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredGrantConfigToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredGrantConfigType2DeactivationStateList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_PUCCH_Config_r16 optional
			if sl_PUCCH_Config_r16Present {
				tmp_sl_PUCCH_Config_r16 := utils.SetupRelease[*PUCCH_Config]{}
				if err = tmp_sl_PUCCH_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_PUCCH_Config_r16", err)
				}
				ie.sl_PUCCH_Config_r16 = tmp_sl_PUCCH_Config_r16.Setup
			}
			// decode cp_ExtensionC2_r16 optional
			if cp_ExtensionC2_r16Present {
				var tmp_int_cp_ExtensionC2_r16 int64
				if tmp_int_cp_ExtensionC2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Decode cp_ExtensionC2_r16", err)
				}
				ie.cp_ExtensionC2_r16 = &tmp_int_cp_ExtensionC2_r16
			}
			// decode cp_ExtensionC3_r16 optional
			if cp_ExtensionC3_r16Present {
				var tmp_int_cp_ExtensionC3_r16 int64
				if tmp_int_cp_ExtensionC3_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 28}, false); err != nil {
					return utils.WrapError("Decode cp_ExtensionC3_r16", err)
				}
				ie.cp_ExtensionC3_r16 = &tmp_int_cp_ExtensionC3_r16
			}
			// decode useInterlacePUCCH_PUSCH_r16 optional
			if useInterlacePUCCH_PUSCH_r16Present {
				ie.useInterlacePUCCH_PUSCH_r16 = new(BWP_UplinkDedicated_useInterlacePUCCH_PUSCH_r16)
				if err = ie.useInterlacePUCCH_PUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode useInterlacePUCCH_PUSCH_r16", err)
				}
			}
			// decode pucch_ConfigurationList_r16 optional
			if pucch_ConfigurationList_r16Present {
				tmp_pucch_ConfigurationList_r16 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{}
				if err = tmp_pucch_ConfigurationList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_ConfigurationList_r16", err)
				}
				ie.pucch_ConfigurationList_r16 = tmp_pucch_ConfigurationList_r16.Setup
			}
			// decode lbt_FailureRecoveryConfig_r16 optional
			if lbt_FailureRecoveryConfig_r16Present {
				tmp_lbt_FailureRecoveryConfig_r16 := utils.SetupRelease[*LBT_FailureRecoveryConfig_r16]{}
				if err = tmp_lbt_FailureRecoveryConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lbt_FailureRecoveryConfig_r16", err)
				}
				ie.lbt_FailureRecoveryConfig_r16 = tmp_lbt_FailureRecoveryConfig_r16.Setup
			}
			// decode configuredGrantConfigToAddModList_r16 optional
			if configuredGrantConfigToAddModList_r16Present {
				ie.configuredGrantConfigToAddModList_r16 = new(ConfiguredGrantConfigToAddModList_r16)
				if err = ie.configuredGrantConfigToAddModList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode configuredGrantConfigToAddModList_r16", err)
				}
			}
			// decode configuredGrantConfigToReleaseList_r16 optional
			if configuredGrantConfigToReleaseList_r16Present {
				ie.configuredGrantConfigToReleaseList_r16 = new(ConfiguredGrantConfigToReleaseList_r16)
				if err = ie.configuredGrantConfigToReleaseList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode configuredGrantConfigToReleaseList_r16", err)
				}
			}
			// decode configuredGrantConfigType2DeactivationStateList_r16 optional
			if configuredGrantConfigType2DeactivationStateList_r16Present {
				ie.configuredGrantConfigType2DeactivationStateList_r16 = new(ConfiguredGrantConfigType2DeactivationStateList_r16)
				if err = ie.configuredGrantConfigType2DeactivationStateList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode configuredGrantConfigType2DeactivationStateList_r16", err)
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

			ul_TCI_StateList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_powerControl_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_ConfigurationListMulticast1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_ConfigurationListMulticast2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ul_TCI_StateList_r17 optional
			if ul_TCI_StateList_r17Present {
				ie.ul_TCI_StateList_r17 = new(BWP_UplinkDedicated_ul_TCI_StateList_r17)
				if err = ie.ul_TCI_StateList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_TCI_StateList_r17", err)
				}
			}
			// decode ul_powerControl_r17 optional
			if ul_powerControl_r17Present {
				ie.ul_powerControl_r17 = new(Uplink_powerControlId_r17)
				if err = ie.ul_powerControl_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_powerControl_r17", err)
				}
			}
			// decode pucch_ConfigurationListMulticast1_r17 optional
			if pucch_ConfigurationListMulticast1_r17Present {
				tmp_pucch_ConfigurationListMulticast1_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{}
				if err = tmp_pucch_ConfigurationListMulticast1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_ConfigurationListMulticast1_r17", err)
				}
				ie.pucch_ConfigurationListMulticast1_r17 = tmp_pucch_ConfigurationListMulticast1_r17.Setup
			}
			// decode pucch_ConfigurationListMulticast2_r17 optional
			if pucch_ConfigurationListMulticast2_r17Present {
				tmp_pucch_ConfigurationListMulticast2_r17 := utils.SetupRelease[*PUCCH_ConfigurationList_r16]{}
				if err = tmp_pucch_ConfigurationListMulticast2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_ConfigurationListMulticast2_r17", err)
				}
				ie.pucch_ConfigurationListMulticast2_r17 = tmp_pucch_ConfigurationListMulticast2_r17.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pucch_ConfigMulticast1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_ConfigMulticast2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pucch_ConfigMulticast1_r17 optional
			if pucch_ConfigMulticast1_r17Present {
				tmp_pucch_ConfigMulticast1_r17 := utils.SetupRelease[*PUCCH_Config]{}
				if err = tmp_pucch_ConfigMulticast1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_ConfigMulticast1_r17", err)
				}
				ie.pucch_ConfigMulticast1_r17 = tmp_pucch_ConfigMulticast1_r17.Setup
			}
			// decode pucch_ConfigMulticast2_r17 optional
			if pucch_ConfigMulticast2_r17Present {
				tmp_pucch_ConfigMulticast2_r17 := utils.SetupRelease[*PUCCH_Config]{}
				if err = tmp_pucch_ConfigMulticast2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_ConfigMulticast2_r17", err)
				}
				ie.pucch_ConfigMulticast2_r17 = tmp_pucch_ConfigMulticast2_r17.Setup
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pathlossReferenceRSToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pathlossReferenceRSToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pathlossReferenceRSToAddModList_r17 optional
			if pathlossReferenceRSToAddModList_r17Present {
				tmp_pathlossReferenceRSToAddModList_r17 := utils.NewSequence[*PathlossReferenceRS_r17]([]*PathlossReferenceRS_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				fn_pathlossReferenceRSToAddModList_r17 := func() *PathlossReferenceRS_r17 {
					return new(PathlossReferenceRS_r17)
				}
				if err = tmp_pathlossReferenceRSToAddModList_r17.Decode(extReader, fn_pathlossReferenceRSToAddModList_r17); err != nil {
					return utils.WrapError("Decode pathlossReferenceRSToAddModList_r17", err)
				}
				ie.pathlossReferenceRSToAddModList_r17 = []PathlossReferenceRS_r17{}
				for _, i := range tmp_pathlossReferenceRSToAddModList_r17.Value {
					ie.pathlossReferenceRSToAddModList_r17 = append(ie.pathlossReferenceRSToAddModList_r17, *i)
				}
			}
			// decode pathlossReferenceRSToReleaseList_r17 optional
			if pathlossReferenceRSToReleaseList_r17Present {
				tmp_pathlossReferenceRSToReleaseList_r17 := utils.NewSequence[*PathlossReferenceRS_Id_r17]([]*PathlossReferenceRS_Id_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPathlossReferenceRSs_r17}, false)
				fn_pathlossReferenceRSToReleaseList_r17 := func() *PathlossReferenceRS_Id_r17 {
					return new(PathlossReferenceRS_Id_r17)
				}
				if err = tmp_pathlossReferenceRSToReleaseList_r17.Decode(extReader, fn_pathlossReferenceRSToReleaseList_r17); err != nil {
					return utils.WrapError("Decode pathlossReferenceRSToReleaseList_r17", err)
				}
				ie.pathlossReferenceRSToReleaseList_r17 = []PathlossReferenceRS_Id_r17{}
				for _, i := range tmp_pathlossReferenceRSToReleaseList_r17.Value {
					ie.pathlossReferenceRSToReleaseList_r17 = append(ie.pathlossReferenceRSToReleaseList_r17, *i)
				}
			}
		}
	}
	return nil
}
