package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfig struct {
	initialUplinkBWP                    *BWP_UplinkDedicated                              `optional`
	uplinkBWP_ToReleaseList             []BWP_Id                                          `lb:1,ub:maxNrofBWPs,optional`
	uplinkBWP_ToAddModList              []BWP_Uplink                                      `lb:1,ub:maxNrofBWPs,optional`
	firstActiveUplinkBWP_Id             *BWP_Id                                           `optional`
	pusch_ServingCellConfig             *PUSCH_ServingCellConfig                          `optional,setuprelease`
	carrierSwitching                    *SRS_CarrierSwitching                             `optional,setuprelease`
	powerBoostPi2BPSK                   *bool                                             `optional,ext-1`
	uplinkChannelBW_PerSCS_List         []SCS_SpecificCarrier                             `lb:1,ub:maxSCSs,optional,ext-1`
	enablePL_RS_UpdateForPUSCH_SRS_r16  *UplinkConfig_enablePL_RS_UpdateForPUSCH_SRS_r16  `optional,ext-2`
	enableDefaultBeamPL_ForPUSCH0_0_r16 *UplinkConfig_enableDefaultBeamPL_ForPUSCH0_0_r16 `optional,ext-2`
	enableDefaultBeamPL_ForPUCCH_r16    *UplinkConfig_enableDefaultBeamPL_ForPUCCH_r16    `optional,ext-2`
	enableDefaultBeamPL_ForSRS_r16      *UplinkConfig_enableDefaultBeamPL_ForSRS_r16      `optional,ext-2`
	uplinkTxSwitching_r16               *UplinkTxSwitching_r16                            `optional,ext-2,setuprelease`
	mpr_PowerBoost_FR2_r16              *UplinkConfig_mpr_PowerBoost_FR2_r16              `optional,ext-2`
}

func (ie *UplinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.powerBoostPi2BPSK != nil || len(ie.uplinkChannelBW_PerSCS_List) > 0 || ie.enablePL_RS_UpdateForPUSCH_SRS_r16 != nil || ie.enableDefaultBeamPL_ForPUSCH0_0_r16 != nil || ie.enableDefaultBeamPL_ForPUCCH_r16 != nil || ie.enableDefaultBeamPL_ForSRS_r16 != nil || ie.uplinkTxSwitching_r16 != nil || ie.mpr_PowerBoost_FR2_r16 != nil
	preambleBits := []bool{hasExtensions, ie.initialUplinkBWP != nil, len(ie.uplinkBWP_ToReleaseList) > 0, len(ie.uplinkBWP_ToAddModList) > 0, ie.firstActiveUplinkBWP_Id != nil, ie.pusch_ServingCellConfig != nil, ie.carrierSwitching != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.initialUplinkBWP != nil {
		if err = ie.initialUplinkBWP.Encode(w); err != nil {
			return utils.WrapError("Encode initialUplinkBWP", err)
		}
	}
	if len(ie.uplinkBWP_ToReleaseList) > 0 {
		tmp_uplinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.uplinkBWP_ToReleaseList {
			tmp_uplinkBWP_ToReleaseList.Value = append(tmp_uplinkBWP_ToReleaseList.Value, &i)
		}
		if err = tmp_uplinkBWP_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkBWP_ToReleaseList", err)
		}
	}
	if len(ie.uplinkBWP_ToAddModList) > 0 {
		tmp_uplinkBWP_ToAddModList := utils.NewSequence[*BWP_Uplink]([]*BWP_Uplink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		for _, i := range ie.uplinkBWP_ToAddModList {
			tmp_uplinkBWP_ToAddModList.Value = append(tmp_uplinkBWP_ToAddModList.Value, &i)
		}
		if err = tmp_uplinkBWP_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkBWP_ToAddModList", err)
		}
	}
	if ie.firstActiveUplinkBWP_Id != nil {
		if err = ie.firstActiveUplinkBWP_Id.Encode(w); err != nil {
			return utils.WrapError("Encode firstActiveUplinkBWP_Id", err)
		}
	}
	if ie.pusch_ServingCellConfig != nil {
		tmp_pusch_ServingCellConfig := utils.SetupRelease[*PUSCH_ServingCellConfig]{
			Setup: ie.pusch_ServingCellConfig,
		}
		if err = tmp_pusch_ServingCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_ServingCellConfig", err)
		}
	}
	if ie.carrierSwitching != nil {
		tmp_carrierSwitching := utils.SetupRelease[*SRS_CarrierSwitching]{
			Setup: ie.carrierSwitching,
		}
		if err = tmp_carrierSwitching.Encode(w); err != nil {
			return utils.WrapError("Encode carrierSwitching", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.powerBoostPi2BPSK != nil || len(ie.uplinkChannelBW_PerSCS_List) > 0, ie.enablePL_RS_UpdateForPUSCH_SRS_r16 != nil || ie.enableDefaultBeamPL_ForPUSCH0_0_r16 != nil || ie.enableDefaultBeamPL_ForPUCCH_r16 != nil || ie.enableDefaultBeamPL_ForSRS_r16 != nil || ie.uplinkTxSwitching_r16 != nil || ie.mpr_PowerBoost_FR2_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UplinkConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.powerBoostPi2BPSK != nil, len(ie.uplinkChannelBW_PerSCS_List) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode powerBoostPi2BPSK optional
			if ie.powerBoostPi2BPSK != nil {
				if err = extWriter.WriteBoolean(*ie.powerBoostPi2BPSK); err != nil {
					return utils.WrapError("Encode powerBoostPi2BPSK", err)
				}
			}
			// encode uplinkChannelBW_PerSCS_List optional
			if len(ie.uplinkChannelBW_PerSCS_List) > 0 {
				tmp_uplinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				for _, i := range ie.uplinkChannelBW_PerSCS_List {
					tmp_uplinkChannelBW_PerSCS_List.Value = append(tmp_uplinkChannelBW_PerSCS_List.Value, &i)
				}
				if err = tmp_uplinkChannelBW_PerSCS_List.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkChannelBW_PerSCS_List", err)
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
			optionals_ext_2 := []bool{ie.enablePL_RS_UpdateForPUSCH_SRS_r16 != nil, ie.enableDefaultBeamPL_ForPUSCH0_0_r16 != nil, ie.enableDefaultBeamPL_ForPUCCH_r16 != nil, ie.enableDefaultBeamPL_ForSRS_r16 != nil, ie.uplinkTxSwitching_r16 != nil, ie.mpr_PowerBoost_FR2_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enablePL_RS_UpdateForPUSCH_SRS_r16 optional
			if ie.enablePL_RS_UpdateForPUSCH_SRS_r16 != nil {
				if err = ie.enablePL_RS_UpdateForPUSCH_SRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enablePL_RS_UpdateForPUSCH_SRS_r16", err)
				}
			}
			// encode enableDefaultBeamPL_ForPUSCH0_0_r16 optional
			if ie.enableDefaultBeamPL_ForPUSCH0_0_r16 != nil {
				if err = ie.enableDefaultBeamPL_ForPUSCH0_0_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableDefaultBeamPL_ForPUSCH0_0_r16", err)
				}
			}
			// encode enableDefaultBeamPL_ForPUCCH_r16 optional
			if ie.enableDefaultBeamPL_ForPUCCH_r16 != nil {
				if err = ie.enableDefaultBeamPL_ForPUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableDefaultBeamPL_ForPUCCH_r16", err)
				}
			}
			// encode enableDefaultBeamPL_ForSRS_r16 optional
			if ie.enableDefaultBeamPL_ForSRS_r16 != nil {
				if err = ie.enableDefaultBeamPL_ForSRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableDefaultBeamPL_ForSRS_r16", err)
				}
			}
			// encode uplinkTxSwitching_r16 optional
			if ie.uplinkTxSwitching_r16 != nil {
				tmp_uplinkTxSwitching_r16 := utils.SetupRelease[*UplinkTxSwitching_r16]{
					Setup: ie.uplinkTxSwitching_r16,
				}
				if err = tmp_uplinkTxSwitching_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkTxSwitching_r16", err)
				}
			}
			// encode mpr_PowerBoost_FR2_r16 optional
			if ie.mpr_PowerBoost_FR2_r16 != nil {
				if err = ie.mpr_PowerBoost_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mpr_PowerBoost_FR2_r16", err)
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

func (ie *UplinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var initialUplinkBWPPresent bool
	if initialUplinkBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkBWP_ToReleaseListPresent bool
	if uplinkBWP_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkBWP_ToAddModListPresent bool
	if uplinkBWP_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var firstActiveUplinkBWP_IdPresent bool
	if firstActiveUplinkBWP_IdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_ServingCellConfigPresent bool
	if pusch_ServingCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var carrierSwitchingPresent bool
	if carrierSwitchingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if initialUplinkBWPPresent {
		ie.initialUplinkBWP = new(BWP_UplinkDedicated)
		if err = ie.initialUplinkBWP.Decode(r); err != nil {
			return utils.WrapError("Decode initialUplinkBWP", err)
		}
	}
	if uplinkBWP_ToReleaseListPresent {
		tmp_uplinkBWP_ToReleaseList := utils.NewSequence[*BWP_Id]([]*BWP_Id{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_uplinkBWP_ToReleaseList := func() *BWP_Id {
			return new(BWP_Id)
		}
		if err = tmp_uplinkBWP_ToReleaseList.Decode(r, fn_uplinkBWP_ToReleaseList); err != nil {
			return utils.WrapError("Decode uplinkBWP_ToReleaseList", err)
		}
		ie.uplinkBWP_ToReleaseList = []BWP_Id{}
		for _, i := range tmp_uplinkBWP_ToReleaseList.Value {
			ie.uplinkBWP_ToReleaseList = append(ie.uplinkBWP_ToReleaseList, *i)
		}
	}
	if uplinkBWP_ToAddModListPresent {
		tmp_uplinkBWP_ToAddModList := utils.NewSequence[*BWP_Uplink]([]*BWP_Uplink{}, uper.Constraint{Lb: 1, Ub: maxNrofBWPs}, false)
		fn_uplinkBWP_ToAddModList := func() *BWP_Uplink {
			return new(BWP_Uplink)
		}
		if err = tmp_uplinkBWP_ToAddModList.Decode(r, fn_uplinkBWP_ToAddModList); err != nil {
			return utils.WrapError("Decode uplinkBWP_ToAddModList", err)
		}
		ie.uplinkBWP_ToAddModList = []BWP_Uplink{}
		for _, i := range tmp_uplinkBWP_ToAddModList.Value {
			ie.uplinkBWP_ToAddModList = append(ie.uplinkBWP_ToAddModList, *i)
		}
	}
	if firstActiveUplinkBWP_IdPresent {
		ie.firstActiveUplinkBWP_Id = new(BWP_Id)
		if err = ie.firstActiveUplinkBWP_Id.Decode(r); err != nil {
			return utils.WrapError("Decode firstActiveUplinkBWP_Id", err)
		}
	}
	if pusch_ServingCellConfigPresent {
		tmp_pusch_ServingCellConfig := utils.SetupRelease[*PUSCH_ServingCellConfig]{}
		if err = tmp_pusch_ServingCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_ServingCellConfig", err)
		}
		ie.pusch_ServingCellConfig = tmp_pusch_ServingCellConfig.Setup
	}
	if carrierSwitchingPresent {
		tmp_carrierSwitching := utils.SetupRelease[*SRS_CarrierSwitching]{}
		if err = tmp_carrierSwitching.Decode(r); err != nil {
			return utils.WrapError("Decode carrierSwitching", err)
		}
		ie.carrierSwitching = tmp_carrierSwitching.Setup
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

			powerBoostPi2BPSKPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkChannelBW_PerSCS_ListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode powerBoostPi2BPSK optional
			if powerBoostPi2BPSKPresent {
				var tmp_bool_powerBoostPi2BPSK bool
				if tmp_bool_powerBoostPi2BPSK, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode powerBoostPi2BPSK", err)
				}
				ie.powerBoostPi2BPSK = &tmp_bool_powerBoostPi2BPSK
			}
			// decode uplinkChannelBW_PerSCS_List optional
			if uplinkChannelBW_PerSCS_ListPresent {
				tmp_uplinkChannelBW_PerSCS_List := utils.NewSequence[*SCS_SpecificCarrier]([]*SCS_SpecificCarrier{}, uper.Constraint{Lb: 1, Ub: maxSCSs}, false)
				fn_uplinkChannelBW_PerSCS_List := func() *SCS_SpecificCarrier {
					return new(SCS_SpecificCarrier)
				}
				if err = tmp_uplinkChannelBW_PerSCS_List.Decode(extReader, fn_uplinkChannelBW_PerSCS_List); err != nil {
					return utils.WrapError("Decode uplinkChannelBW_PerSCS_List", err)
				}
				ie.uplinkChannelBW_PerSCS_List = []SCS_SpecificCarrier{}
				for _, i := range tmp_uplinkChannelBW_PerSCS_List.Value {
					ie.uplinkChannelBW_PerSCS_List = append(ie.uplinkChannelBW_PerSCS_List, *i)
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

			enablePL_RS_UpdateForPUSCH_SRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enableDefaultBeamPL_ForPUSCH0_0_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enableDefaultBeamPL_ForPUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enableDefaultBeamPL_ForSRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkTxSwitching_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mpr_PowerBoost_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enablePL_RS_UpdateForPUSCH_SRS_r16 optional
			if enablePL_RS_UpdateForPUSCH_SRS_r16Present {
				ie.enablePL_RS_UpdateForPUSCH_SRS_r16 = new(UplinkConfig_enablePL_RS_UpdateForPUSCH_SRS_r16)
				if err = ie.enablePL_RS_UpdateForPUSCH_SRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enablePL_RS_UpdateForPUSCH_SRS_r16", err)
				}
			}
			// decode enableDefaultBeamPL_ForPUSCH0_0_r16 optional
			if enableDefaultBeamPL_ForPUSCH0_0_r16Present {
				ie.enableDefaultBeamPL_ForPUSCH0_0_r16 = new(UplinkConfig_enableDefaultBeamPL_ForPUSCH0_0_r16)
				if err = ie.enableDefaultBeamPL_ForPUSCH0_0_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableDefaultBeamPL_ForPUSCH0_0_r16", err)
				}
			}
			// decode enableDefaultBeamPL_ForPUCCH_r16 optional
			if enableDefaultBeamPL_ForPUCCH_r16Present {
				ie.enableDefaultBeamPL_ForPUCCH_r16 = new(UplinkConfig_enableDefaultBeamPL_ForPUCCH_r16)
				if err = ie.enableDefaultBeamPL_ForPUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableDefaultBeamPL_ForPUCCH_r16", err)
				}
			}
			// decode enableDefaultBeamPL_ForSRS_r16 optional
			if enableDefaultBeamPL_ForSRS_r16Present {
				ie.enableDefaultBeamPL_ForSRS_r16 = new(UplinkConfig_enableDefaultBeamPL_ForSRS_r16)
				if err = ie.enableDefaultBeamPL_ForSRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableDefaultBeamPL_ForSRS_r16", err)
				}
			}
			// decode uplinkTxSwitching_r16 optional
			if uplinkTxSwitching_r16Present {
				tmp_uplinkTxSwitching_r16 := utils.SetupRelease[*UplinkTxSwitching_r16]{}
				if err = tmp_uplinkTxSwitching_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkTxSwitching_r16", err)
				}
				ie.uplinkTxSwitching_r16 = tmp_uplinkTxSwitching_r16.Setup
			}
			// decode mpr_PowerBoost_FR2_r16 optional
			if mpr_PowerBoost_FR2_r16Present {
				ie.mpr_PowerBoost_FR2_r16 = new(UplinkConfig_mpr_PowerBoost_FR2_r16)
				if err = ie.mpr_PowerBoost_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mpr_PowerBoost_FR2_r16", err)
				}
			}
		}
	}
	return nil
}
