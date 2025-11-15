package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellGroupConfig struct {
	cellGroupId                                CellGroupId                                                `madatory`
	rlc_BearerToAddModList                     []RLC_BearerConfig                                         `lb:1,ub:maxLC_ID,optional`
	rlc_BearerToReleaseList                    []LogicalChannelIdentity                                   `lb:1,ub:maxLC_ID,optional`
	mac_CellGroupConfig                        *MAC_CellGroupConfig                                       `optional`
	physicalCellGroupConfig                    *PhysicalCellGroupConfig                                   `optional`
	spCellConfig                               *SpCellConfig                                              `optional`
	sCellToAddModList                          []SCellConfig                                              `lb:1,ub:maxNrofSCells,optional`
	sCellToReleaseList                         []SCellIndex                                               `lb:1,ub:maxNrofSCells,optional`
	reportUplinkTxDirectCurrent                *CellGroupConfig_reportUplinkTxDirectCurrent               `optional,ext-1`
	bap_Address_r16                            *uper.BitString                                            `lb:10,ub:10,optional,ext-2`
	bh_RLC_ChannelToAddModList_r16             []BH_RLC_ChannelConfig_r16                                 `lb:1,ub:maxBH_RLC_ChannelID_r16,optional,ext-2`
	bh_RLC_ChannelToReleaseList_r16            []BH_RLC_ChannelID_r16                                     `lb:1,ub:maxBH_RLC_ChannelID_r16,optional,ext-2`
	f1c_TransferPath_r16                       *CellGroupConfig_f1c_TransferPath_r16                      `optional,ext-2`
	simultaneousTCI_UpdateList1_r16            []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	simultaneousTCI_UpdateList2_r16            []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	simultaneousSpatial_UpdatedList1_r16       []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	simultaneousSpatial_UpdatedList2_r16       []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	uplinkTxSwitchingOption_r16                *CellGroupConfig_uplinkTxSwitchingOption_r16               `optional,ext-2`
	uplinkTxSwitchingPowerBoosting_r16         *CellGroupConfig_uplinkTxSwitchingPowerBoosting_r16        `optional,ext-2`
	reportUplinkTxDirectCurrentTwoCarrier_r16  *CellGroupConfig_reportUplinkTxDirectCurrentTwoCarrier_r16 `optional,ext-3`
	f1c_TransferPathNRDC_r17                   *CellGroupConfig_f1c_TransferPathNRDC_r17                  `optional,ext-4`
	uplinkTxSwitching_2T_Mode_r17              *CellGroupConfig_uplinkTxSwitching_2T_Mode_r17             `optional,ext-4`
	uplinkTxSwitching_DualUL_TxState_r17       *CellGroupConfig_uplinkTxSwitching_DualUL_TxState_r17      `optional,ext-4`
	uu_RelayRLC_ChannelToAddModList_r17        []Uu_RelayRLC_ChannelConfig_r17                            `lb:1,ub:maxUu_RelayRLC_ChannelID_r17,optional,ext-4`
	uu_RelayRLC_ChannelToReleaseList_r17       []Uu_RelayRLC_ChannelID_r17                                `lb:1,ub:maxUu_RelayRLC_ChannelID_r17,optional,ext-4`
	simultaneousU_TCI_UpdateList1_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	simultaneousU_TCI_UpdateList2_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	simultaneousU_TCI_UpdateList3_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	simultaneousU_TCI_UpdateList4_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	rlc_BearerToReleaseListExt_r17             []LogicalChannelIdentityExt_r17                            `lb:1,ub:maxLC_ID,optional,ext-4`
	iab_ResourceConfigToAddModList_r17         []IAB_ResourceConfig_r17                                   `lb:1,ub:maxNrofIABResourceConfig_r17,optional,ext-4`
	iab_ResourceConfigToReleaseList_r17        []IAB_ResourceConfigID_r17                                 `lb:1,ub:maxNrofIABResourceConfig_r17,optional,ext-4`
	reportUplinkTxDirectCurrentMoreCarrier_r17 *ReportUplinkTxDirectCurrentMoreCarrier_r17                `optional,ext-5`
}

func (ie *CellGroupConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.reportUplinkTxDirectCurrent != nil || ie.bap_Address_r16 != nil || len(ie.bh_RLC_ChannelToAddModList_r16) > 0 || len(ie.bh_RLC_ChannelToReleaseList_r16) > 0 || ie.f1c_TransferPath_r16 != nil || len(ie.simultaneousTCI_UpdateList1_r16) > 0 || len(ie.simultaneousTCI_UpdateList2_r16) > 0 || len(ie.simultaneousSpatial_UpdatedList1_r16) > 0 || len(ie.simultaneousSpatial_UpdatedList2_r16) > 0 || ie.uplinkTxSwitchingOption_r16 != nil || ie.uplinkTxSwitchingPowerBoosting_r16 != nil || ie.reportUplinkTxDirectCurrentTwoCarrier_r16 != nil || ie.f1c_TransferPathNRDC_r17 != nil || ie.uplinkTxSwitching_2T_Mode_r17 != nil || ie.uplinkTxSwitching_DualUL_TxState_r17 != nil || len(ie.uu_RelayRLC_ChannelToAddModList_r17) > 0 || len(ie.uu_RelayRLC_ChannelToReleaseList_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList1_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList2_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList3_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList4_r17) > 0 || len(ie.rlc_BearerToReleaseListExt_r17) > 0 || len(ie.iab_ResourceConfigToAddModList_r17) > 0 || len(ie.iab_ResourceConfigToReleaseList_r17) > 0 || ie.reportUplinkTxDirectCurrentMoreCarrier_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.rlc_BearerToAddModList) > 0, len(ie.rlc_BearerToReleaseList) > 0, ie.mac_CellGroupConfig != nil, ie.physicalCellGroupConfig != nil, ie.spCellConfig != nil, len(ie.sCellToAddModList) > 0, len(ie.sCellToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.cellGroupId.Encode(w); err != nil {
		return utils.WrapError("Encode cellGroupId", err)
	}
	if len(ie.rlc_BearerToAddModList) > 0 {
		tmp_rlc_BearerToAddModList := utils.NewSequence[*RLC_BearerConfig]([]*RLC_BearerConfig{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.rlc_BearerToAddModList {
			tmp_rlc_BearerToAddModList.Value = append(tmp_rlc_BearerToAddModList.Value, &i)
		}
		if err = tmp_rlc_BearerToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_BearerToAddModList", err)
		}
	}
	if len(ie.rlc_BearerToReleaseList) > 0 {
		tmp_rlc_BearerToReleaseList := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.rlc_BearerToReleaseList {
			tmp_rlc_BearerToReleaseList.Value = append(tmp_rlc_BearerToReleaseList.Value, &i)
		}
		if err = tmp_rlc_BearerToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_BearerToReleaseList", err)
		}
	}
	if ie.mac_CellGroupConfig != nil {
		if err = ie.mac_CellGroupConfig.Encode(w); err != nil {
			return utils.WrapError("Encode mac_CellGroupConfig", err)
		}
	}
	if ie.physicalCellGroupConfig != nil {
		if err = ie.physicalCellGroupConfig.Encode(w); err != nil {
			return utils.WrapError("Encode physicalCellGroupConfig", err)
		}
	}
	if ie.spCellConfig != nil {
		if err = ie.spCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode spCellConfig", err)
		}
	}
	if len(ie.sCellToAddModList) > 0 {
		tmp_sCellToAddModList := utils.NewSequence[*SCellConfig]([]*SCellConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		for _, i := range ie.sCellToAddModList {
			tmp_sCellToAddModList.Value = append(tmp_sCellToAddModList.Value, &i)
		}
		if err = tmp_sCellToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode sCellToAddModList", err)
		}
	}
	if len(ie.sCellToReleaseList) > 0 {
		tmp_sCellToReleaseList := utils.NewSequence[*SCellIndex]([]*SCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		for _, i := range ie.sCellToReleaseList {
			tmp_sCellToReleaseList.Value = append(tmp_sCellToReleaseList.Value, &i)
		}
		if err = tmp_sCellToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode sCellToReleaseList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.reportUplinkTxDirectCurrent != nil, ie.bap_Address_r16 != nil || len(ie.bh_RLC_ChannelToAddModList_r16) > 0 || len(ie.bh_RLC_ChannelToReleaseList_r16) > 0 || ie.f1c_TransferPath_r16 != nil || len(ie.simultaneousTCI_UpdateList1_r16) > 0 || len(ie.simultaneousTCI_UpdateList2_r16) > 0 || len(ie.simultaneousSpatial_UpdatedList1_r16) > 0 || len(ie.simultaneousSpatial_UpdatedList2_r16) > 0 || ie.uplinkTxSwitchingOption_r16 != nil || ie.uplinkTxSwitchingPowerBoosting_r16 != nil, ie.reportUplinkTxDirectCurrentTwoCarrier_r16 != nil, ie.f1c_TransferPathNRDC_r17 != nil || ie.uplinkTxSwitching_2T_Mode_r17 != nil || ie.uplinkTxSwitching_DualUL_TxState_r17 != nil || len(ie.uu_RelayRLC_ChannelToAddModList_r17) > 0 || len(ie.uu_RelayRLC_ChannelToReleaseList_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList1_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList2_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList3_r17) > 0 || len(ie.simultaneousU_TCI_UpdateList4_r17) > 0 || len(ie.rlc_BearerToReleaseListExt_r17) > 0 || len(ie.iab_ResourceConfigToAddModList_r17) > 0 || len(ie.iab_ResourceConfigToReleaseList_r17) > 0, ie.reportUplinkTxDirectCurrentMoreCarrier_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CellGroupConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.reportUplinkTxDirectCurrent != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode reportUplinkTxDirectCurrent optional
			if ie.reportUplinkTxDirectCurrent != nil {
				if err = ie.reportUplinkTxDirectCurrent.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportUplinkTxDirectCurrent", err)
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
			optionals_ext_2 := []bool{ie.bap_Address_r16 != nil, len(ie.bh_RLC_ChannelToAddModList_r16) > 0, len(ie.bh_RLC_ChannelToReleaseList_r16) > 0, ie.f1c_TransferPath_r16 != nil, len(ie.simultaneousTCI_UpdateList1_r16) > 0, len(ie.simultaneousTCI_UpdateList2_r16) > 0, len(ie.simultaneousSpatial_UpdatedList1_r16) > 0, len(ie.simultaneousSpatial_UpdatedList2_r16) > 0, ie.uplinkTxSwitchingOption_r16 != nil, ie.uplinkTxSwitchingPowerBoosting_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode bap_Address_r16 optional
			if ie.bap_Address_r16 != nil {
				if err = extWriter.WriteBitString(ie.bap_Address_r16.Bytes, uint(ie.bap_Address_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Encode bap_Address_r16", err)
				}
			}
			// encode bh_RLC_ChannelToAddModList_r16 optional
			if len(ie.bh_RLC_ChannelToAddModList_r16) > 0 {
				tmp_bh_RLC_ChannelToAddModList_r16 := utils.NewSequence[*BH_RLC_ChannelConfig_r16]([]*BH_RLC_ChannelConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				for _, i := range ie.bh_RLC_ChannelToAddModList_r16 {
					tmp_bh_RLC_ChannelToAddModList_r16.Value = append(tmp_bh_RLC_ChannelToAddModList_r16.Value, &i)
				}
				if err = tmp_bh_RLC_ChannelToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bh_RLC_ChannelToAddModList_r16", err)
				}
			}
			// encode bh_RLC_ChannelToReleaseList_r16 optional
			if len(ie.bh_RLC_ChannelToReleaseList_r16) > 0 {
				tmp_bh_RLC_ChannelToReleaseList_r16 := utils.NewSequence[*BH_RLC_ChannelID_r16]([]*BH_RLC_ChannelID_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				for _, i := range ie.bh_RLC_ChannelToReleaseList_r16 {
					tmp_bh_RLC_ChannelToReleaseList_r16.Value = append(tmp_bh_RLC_ChannelToReleaseList_r16.Value, &i)
				}
				if err = tmp_bh_RLC_ChannelToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bh_RLC_ChannelToReleaseList_r16", err)
				}
			}
			// encode f1c_TransferPath_r16 optional
			if ie.f1c_TransferPath_r16 != nil {
				if err = ie.f1c_TransferPath_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode f1c_TransferPath_r16", err)
				}
			}
			// encode simultaneousTCI_UpdateList1_r16 optional
			if len(ie.simultaneousTCI_UpdateList1_r16) > 0 {
				tmp_simultaneousTCI_UpdateList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousTCI_UpdateList1_r16 {
					tmp_simultaneousTCI_UpdateList1_r16.Value = append(tmp_simultaneousTCI_UpdateList1_r16.Value, &i)
				}
				if err = tmp_simultaneousTCI_UpdateList1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousTCI_UpdateList1_r16", err)
				}
			}
			// encode simultaneousTCI_UpdateList2_r16 optional
			if len(ie.simultaneousTCI_UpdateList2_r16) > 0 {
				tmp_simultaneousTCI_UpdateList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousTCI_UpdateList2_r16 {
					tmp_simultaneousTCI_UpdateList2_r16.Value = append(tmp_simultaneousTCI_UpdateList2_r16.Value, &i)
				}
				if err = tmp_simultaneousTCI_UpdateList2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousTCI_UpdateList2_r16", err)
				}
			}
			// encode simultaneousSpatial_UpdatedList1_r16 optional
			if len(ie.simultaneousSpatial_UpdatedList1_r16) > 0 {
				tmp_simultaneousSpatial_UpdatedList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousSpatial_UpdatedList1_r16 {
					tmp_simultaneousSpatial_UpdatedList1_r16.Value = append(tmp_simultaneousSpatial_UpdatedList1_r16.Value, &i)
				}
				if err = tmp_simultaneousSpatial_UpdatedList1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousSpatial_UpdatedList1_r16", err)
				}
			}
			// encode simultaneousSpatial_UpdatedList2_r16 optional
			if len(ie.simultaneousSpatial_UpdatedList2_r16) > 0 {
				tmp_simultaneousSpatial_UpdatedList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousSpatial_UpdatedList2_r16 {
					tmp_simultaneousSpatial_UpdatedList2_r16.Value = append(tmp_simultaneousSpatial_UpdatedList2_r16.Value, &i)
				}
				if err = tmp_simultaneousSpatial_UpdatedList2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousSpatial_UpdatedList2_r16", err)
				}
			}
			// encode uplinkTxSwitchingOption_r16 optional
			if ie.uplinkTxSwitchingOption_r16 != nil {
				if err = ie.uplinkTxSwitchingOption_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkTxSwitchingOption_r16", err)
				}
			}
			// encode uplinkTxSwitchingPowerBoosting_r16 optional
			if ie.uplinkTxSwitchingPowerBoosting_r16 != nil {
				if err = ie.uplinkTxSwitchingPowerBoosting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkTxSwitchingPowerBoosting_r16", err)
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
			optionals_ext_3 := []bool{ie.reportUplinkTxDirectCurrentTwoCarrier_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode reportUplinkTxDirectCurrentTwoCarrier_r16 optional
			if ie.reportUplinkTxDirectCurrentTwoCarrier_r16 != nil {
				if err = ie.reportUplinkTxDirectCurrentTwoCarrier_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportUplinkTxDirectCurrentTwoCarrier_r16", err)
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
			optionals_ext_4 := []bool{ie.f1c_TransferPathNRDC_r17 != nil, ie.uplinkTxSwitching_2T_Mode_r17 != nil, ie.uplinkTxSwitching_DualUL_TxState_r17 != nil, len(ie.uu_RelayRLC_ChannelToAddModList_r17) > 0, len(ie.uu_RelayRLC_ChannelToReleaseList_r17) > 0, len(ie.simultaneousU_TCI_UpdateList1_r17) > 0, len(ie.simultaneousU_TCI_UpdateList2_r17) > 0, len(ie.simultaneousU_TCI_UpdateList3_r17) > 0, len(ie.simultaneousU_TCI_UpdateList4_r17) > 0, len(ie.rlc_BearerToReleaseListExt_r17) > 0, len(ie.iab_ResourceConfigToAddModList_r17) > 0, len(ie.iab_ResourceConfigToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode f1c_TransferPathNRDC_r17 optional
			if ie.f1c_TransferPathNRDC_r17 != nil {
				if err = ie.f1c_TransferPathNRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode f1c_TransferPathNRDC_r17", err)
				}
			}
			// encode uplinkTxSwitching_2T_Mode_r17 optional
			if ie.uplinkTxSwitching_2T_Mode_r17 != nil {
				if err = ie.uplinkTxSwitching_2T_Mode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkTxSwitching_2T_Mode_r17", err)
				}
			}
			// encode uplinkTxSwitching_DualUL_TxState_r17 optional
			if ie.uplinkTxSwitching_DualUL_TxState_r17 != nil {
				if err = ie.uplinkTxSwitching_DualUL_TxState_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkTxSwitching_DualUL_TxState_r17", err)
				}
			}
			// encode uu_RelayRLC_ChannelToAddModList_r17 optional
			if len(ie.uu_RelayRLC_ChannelToAddModList_r17) > 0 {
				tmp_uu_RelayRLC_ChannelToAddModList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelConfig_r17]([]*Uu_RelayRLC_ChannelConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				for _, i := range ie.uu_RelayRLC_ChannelToAddModList_r17 {
					tmp_uu_RelayRLC_ChannelToAddModList_r17.Value = append(tmp_uu_RelayRLC_ChannelToAddModList_r17.Value, &i)
				}
				if err = tmp_uu_RelayRLC_ChannelToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uu_RelayRLC_ChannelToAddModList_r17", err)
				}
			}
			// encode uu_RelayRLC_ChannelToReleaseList_r17 optional
			if len(ie.uu_RelayRLC_ChannelToReleaseList_r17) > 0 {
				tmp_uu_RelayRLC_ChannelToReleaseList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelID_r17]([]*Uu_RelayRLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				for _, i := range ie.uu_RelayRLC_ChannelToReleaseList_r17 {
					tmp_uu_RelayRLC_ChannelToReleaseList_r17.Value = append(tmp_uu_RelayRLC_ChannelToReleaseList_r17.Value, &i)
				}
				if err = tmp_uu_RelayRLC_ChannelToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uu_RelayRLC_ChannelToReleaseList_r17", err)
				}
			}
			// encode simultaneousU_TCI_UpdateList1_r17 optional
			if len(ie.simultaneousU_TCI_UpdateList1_r17) > 0 {
				tmp_simultaneousU_TCI_UpdateList1_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousU_TCI_UpdateList1_r17 {
					tmp_simultaneousU_TCI_UpdateList1_r17.Value = append(tmp_simultaneousU_TCI_UpdateList1_r17.Value, &i)
				}
				if err = tmp_simultaneousU_TCI_UpdateList1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousU_TCI_UpdateList1_r17", err)
				}
			}
			// encode simultaneousU_TCI_UpdateList2_r17 optional
			if len(ie.simultaneousU_TCI_UpdateList2_r17) > 0 {
				tmp_simultaneousU_TCI_UpdateList2_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousU_TCI_UpdateList2_r17 {
					tmp_simultaneousU_TCI_UpdateList2_r17.Value = append(tmp_simultaneousU_TCI_UpdateList2_r17.Value, &i)
				}
				if err = tmp_simultaneousU_TCI_UpdateList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousU_TCI_UpdateList2_r17", err)
				}
			}
			// encode simultaneousU_TCI_UpdateList3_r17 optional
			if len(ie.simultaneousU_TCI_UpdateList3_r17) > 0 {
				tmp_simultaneousU_TCI_UpdateList3_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousU_TCI_UpdateList3_r17 {
					tmp_simultaneousU_TCI_UpdateList3_r17.Value = append(tmp_simultaneousU_TCI_UpdateList3_r17.Value, &i)
				}
				if err = tmp_simultaneousU_TCI_UpdateList3_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousU_TCI_UpdateList3_r17", err)
				}
			}
			// encode simultaneousU_TCI_UpdateList4_r17 optional
			if len(ie.simultaneousU_TCI_UpdateList4_r17) > 0 {
				tmp_simultaneousU_TCI_UpdateList4_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.simultaneousU_TCI_UpdateList4_r17 {
					tmp_simultaneousU_TCI_UpdateList4_r17.Value = append(tmp_simultaneousU_TCI_UpdateList4_r17.Value, &i)
				}
				if err = tmp_simultaneousU_TCI_UpdateList4_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousU_TCI_UpdateList4_r17", err)
				}
			}
			// encode rlc_BearerToReleaseListExt_r17 optional
			if len(ie.rlc_BearerToReleaseListExt_r17) > 0 {
				tmp_rlc_BearerToReleaseListExt_r17 := utils.NewSequence[*LogicalChannelIdentityExt_r17]([]*LogicalChannelIdentityExt_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
				for _, i := range ie.rlc_BearerToReleaseListExt_r17 {
					tmp_rlc_BearerToReleaseListExt_r17.Value = append(tmp_rlc_BearerToReleaseListExt_r17.Value, &i)
				}
				if err = tmp_rlc_BearerToReleaseListExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rlc_BearerToReleaseListExt_r17", err)
				}
			}
			// encode iab_ResourceConfigToAddModList_r17 optional
			if len(ie.iab_ResourceConfigToAddModList_r17) > 0 {
				tmp_iab_ResourceConfigToAddModList_r17 := utils.NewSequence[*IAB_ResourceConfig_r17]([]*IAB_ResourceConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				for _, i := range ie.iab_ResourceConfigToAddModList_r17 {
					tmp_iab_ResourceConfigToAddModList_r17.Value = append(tmp_iab_ResourceConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_iab_ResourceConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode iab_ResourceConfigToAddModList_r17", err)
				}
			}
			// encode iab_ResourceConfigToReleaseList_r17 optional
			if len(ie.iab_ResourceConfigToReleaseList_r17) > 0 {
				tmp_iab_ResourceConfigToReleaseList_r17 := utils.NewSequence[*IAB_ResourceConfigID_r17]([]*IAB_ResourceConfigID_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				for _, i := range ie.iab_ResourceConfigToReleaseList_r17 {
					tmp_iab_ResourceConfigToReleaseList_r17.Value = append(tmp_iab_ResourceConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_iab_ResourceConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode iab_ResourceConfigToReleaseList_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.reportUplinkTxDirectCurrentMoreCarrier_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode reportUplinkTxDirectCurrentMoreCarrier_r17 optional
			if ie.reportUplinkTxDirectCurrentMoreCarrier_r17 != nil {
				if err = ie.reportUplinkTxDirectCurrentMoreCarrier_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportUplinkTxDirectCurrentMoreCarrier_r17", err)
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

func (ie *CellGroupConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rlc_BearerToAddModListPresent bool
	if rlc_BearerToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rlc_BearerToReleaseListPresent bool
	if rlc_BearerToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_CellGroupConfigPresent bool
	if mac_CellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var physicalCellGroupConfigPresent bool
	if physicalCellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var spCellConfigPresent bool
	if spCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sCellToAddModListPresent bool
	if sCellToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sCellToReleaseListPresent bool
	if sCellToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.cellGroupId.Decode(r); err != nil {
		return utils.WrapError("Decode cellGroupId", err)
	}
	if rlc_BearerToAddModListPresent {
		tmp_rlc_BearerToAddModList := utils.NewSequence[*RLC_BearerConfig]([]*RLC_BearerConfig{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_rlc_BearerToAddModList := func() *RLC_BearerConfig {
			return new(RLC_BearerConfig)
		}
		if err = tmp_rlc_BearerToAddModList.Decode(r, fn_rlc_BearerToAddModList); err != nil {
			return utils.WrapError("Decode rlc_BearerToAddModList", err)
		}
		ie.rlc_BearerToAddModList = []RLC_BearerConfig{}
		for _, i := range tmp_rlc_BearerToAddModList.Value {
			ie.rlc_BearerToAddModList = append(ie.rlc_BearerToAddModList, *i)
		}
	}
	if rlc_BearerToReleaseListPresent {
		tmp_rlc_BearerToReleaseList := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_rlc_BearerToReleaseList := func() *LogicalChannelIdentity {
			return new(LogicalChannelIdentity)
		}
		if err = tmp_rlc_BearerToReleaseList.Decode(r, fn_rlc_BearerToReleaseList); err != nil {
			return utils.WrapError("Decode rlc_BearerToReleaseList", err)
		}
		ie.rlc_BearerToReleaseList = []LogicalChannelIdentity{}
		for _, i := range tmp_rlc_BearerToReleaseList.Value {
			ie.rlc_BearerToReleaseList = append(ie.rlc_BearerToReleaseList, *i)
		}
	}
	if mac_CellGroupConfigPresent {
		ie.mac_CellGroupConfig = new(MAC_CellGroupConfig)
		if err = ie.mac_CellGroupConfig.Decode(r); err != nil {
			return utils.WrapError("Decode mac_CellGroupConfig", err)
		}
	}
	if physicalCellGroupConfigPresent {
		ie.physicalCellGroupConfig = new(PhysicalCellGroupConfig)
		if err = ie.physicalCellGroupConfig.Decode(r); err != nil {
			return utils.WrapError("Decode physicalCellGroupConfig", err)
		}
	}
	if spCellConfigPresent {
		ie.spCellConfig = new(SpCellConfig)
		if err = ie.spCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode spCellConfig", err)
		}
	}
	if sCellToAddModListPresent {
		tmp_sCellToAddModList := utils.NewSequence[*SCellConfig]([]*SCellConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		fn_sCellToAddModList := func() *SCellConfig {
			return new(SCellConfig)
		}
		if err = tmp_sCellToAddModList.Decode(r, fn_sCellToAddModList); err != nil {
			return utils.WrapError("Decode sCellToAddModList", err)
		}
		ie.sCellToAddModList = []SCellConfig{}
		for _, i := range tmp_sCellToAddModList.Value {
			ie.sCellToAddModList = append(ie.sCellToAddModList, *i)
		}
	}
	if sCellToReleaseListPresent {
		tmp_sCellToReleaseList := utils.NewSequence[*SCellIndex]([]*SCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		fn_sCellToReleaseList := func() *SCellIndex {
			return new(SCellIndex)
		}
		if err = tmp_sCellToReleaseList.Decode(r, fn_sCellToReleaseList); err != nil {
			return utils.WrapError("Decode sCellToReleaseList", err)
		}
		ie.sCellToReleaseList = []SCellIndex{}
		for _, i := range tmp_sCellToReleaseList.Value {
			ie.sCellToReleaseList = append(ie.sCellToReleaseList, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
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

			reportUplinkTxDirectCurrentPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode reportUplinkTxDirectCurrent optional
			if reportUplinkTxDirectCurrentPresent {
				ie.reportUplinkTxDirectCurrent = new(CellGroupConfig_reportUplinkTxDirectCurrent)
				if err = ie.reportUplinkTxDirectCurrent.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportUplinkTxDirectCurrent", err)
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

			bap_Address_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			bh_RLC_ChannelToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			bh_RLC_ChannelToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			f1c_TransferPath_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousTCI_UpdateList1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousTCI_UpdateList2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousSpatial_UpdatedList1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousSpatial_UpdatedList2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkTxSwitchingOption_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkTxSwitchingPowerBoosting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode bap_Address_r16 optional
			if bap_Address_r16Present {
				var tmp_bs_bap_Address_r16 []byte
				var n_bap_Address_r16 uint
				if tmp_bs_bap_Address_r16, n_bap_Address_r16, err = extReader.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Decode bap_Address_r16", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_bap_Address_r16,
					NumBits: uint64(n_bap_Address_r16),
				}
				ie.bap_Address_r16 = &tmp_bitstring
			}
			// decode bh_RLC_ChannelToAddModList_r16 optional
			if bh_RLC_ChannelToAddModList_r16Present {
				tmp_bh_RLC_ChannelToAddModList_r16 := utils.NewSequence[*BH_RLC_ChannelConfig_r16]([]*BH_RLC_ChannelConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				fn_bh_RLC_ChannelToAddModList_r16 := func() *BH_RLC_ChannelConfig_r16 {
					return new(BH_RLC_ChannelConfig_r16)
				}
				if err = tmp_bh_RLC_ChannelToAddModList_r16.Decode(extReader, fn_bh_RLC_ChannelToAddModList_r16); err != nil {
					return utils.WrapError("Decode bh_RLC_ChannelToAddModList_r16", err)
				}
				ie.bh_RLC_ChannelToAddModList_r16 = []BH_RLC_ChannelConfig_r16{}
				for _, i := range tmp_bh_RLC_ChannelToAddModList_r16.Value {
					ie.bh_RLC_ChannelToAddModList_r16 = append(ie.bh_RLC_ChannelToAddModList_r16, *i)
				}
			}
			// decode bh_RLC_ChannelToReleaseList_r16 optional
			if bh_RLC_ChannelToReleaseList_r16Present {
				tmp_bh_RLC_ChannelToReleaseList_r16 := utils.NewSequence[*BH_RLC_ChannelID_r16]([]*BH_RLC_ChannelID_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				fn_bh_RLC_ChannelToReleaseList_r16 := func() *BH_RLC_ChannelID_r16 {
					return new(BH_RLC_ChannelID_r16)
				}
				if err = tmp_bh_RLC_ChannelToReleaseList_r16.Decode(extReader, fn_bh_RLC_ChannelToReleaseList_r16); err != nil {
					return utils.WrapError("Decode bh_RLC_ChannelToReleaseList_r16", err)
				}
				ie.bh_RLC_ChannelToReleaseList_r16 = []BH_RLC_ChannelID_r16{}
				for _, i := range tmp_bh_RLC_ChannelToReleaseList_r16.Value {
					ie.bh_RLC_ChannelToReleaseList_r16 = append(ie.bh_RLC_ChannelToReleaseList_r16, *i)
				}
			}
			// decode f1c_TransferPath_r16 optional
			if f1c_TransferPath_r16Present {
				ie.f1c_TransferPath_r16 = new(CellGroupConfig_f1c_TransferPath_r16)
				if err = ie.f1c_TransferPath_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode f1c_TransferPath_r16", err)
				}
			}
			// decode simultaneousTCI_UpdateList1_r16 optional
			if simultaneousTCI_UpdateList1_r16Present {
				tmp_simultaneousTCI_UpdateList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousTCI_UpdateList1_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousTCI_UpdateList1_r16.Decode(extReader, fn_simultaneousTCI_UpdateList1_r16); err != nil {
					return utils.WrapError("Decode simultaneousTCI_UpdateList1_r16", err)
				}
				ie.simultaneousTCI_UpdateList1_r16 = []ServCellIndex{}
				for _, i := range tmp_simultaneousTCI_UpdateList1_r16.Value {
					ie.simultaneousTCI_UpdateList1_r16 = append(ie.simultaneousTCI_UpdateList1_r16, *i)
				}
			}
			// decode simultaneousTCI_UpdateList2_r16 optional
			if simultaneousTCI_UpdateList2_r16Present {
				tmp_simultaneousTCI_UpdateList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousTCI_UpdateList2_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousTCI_UpdateList2_r16.Decode(extReader, fn_simultaneousTCI_UpdateList2_r16); err != nil {
					return utils.WrapError("Decode simultaneousTCI_UpdateList2_r16", err)
				}
				ie.simultaneousTCI_UpdateList2_r16 = []ServCellIndex{}
				for _, i := range tmp_simultaneousTCI_UpdateList2_r16.Value {
					ie.simultaneousTCI_UpdateList2_r16 = append(ie.simultaneousTCI_UpdateList2_r16, *i)
				}
			}
			// decode simultaneousSpatial_UpdatedList1_r16 optional
			if simultaneousSpatial_UpdatedList1_r16Present {
				tmp_simultaneousSpatial_UpdatedList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousSpatial_UpdatedList1_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousSpatial_UpdatedList1_r16.Decode(extReader, fn_simultaneousSpatial_UpdatedList1_r16); err != nil {
					return utils.WrapError("Decode simultaneousSpatial_UpdatedList1_r16", err)
				}
				ie.simultaneousSpatial_UpdatedList1_r16 = []ServCellIndex{}
				for _, i := range tmp_simultaneousSpatial_UpdatedList1_r16.Value {
					ie.simultaneousSpatial_UpdatedList1_r16 = append(ie.simultaneousSpatial_UpdatedList1_r16, *i)
				}
			}
			// decode simultaneousSpatial_UpdatedList2_r16 optional
			if simultaneousSpatial_UpdatedList2_r16Present {
				tmp_simultaneousSpatial_UpdatedList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousSpatial_UpdatedList2_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousSpatial_UpdatedList2_r16.Decode(extReader, fn_simultaneousSpatial_UpdatedList2_r16); err != nil {
					return utils.WrapError("Decode simultaneousSpatial_UpdatedList2_r16", err)
				}
				ie.simultaneousSpatial_UpdatedList2_r16 = []ServCellIndex{}
				for _, i := range tmp_simultaneousSpatial_UpdatedList2_r16.Value {
					ie.simultaneousSpatial_UpdatedList2_r16 = append(ie.simultaneousSpatial_UpdatedList2_r16, *i)
				}
			}
			// decode uplinkTxSwitchingOption_r16 optional
			if uplinkTxSwitchingOption_r16Present {
				ie.uplinkTxSwitchingOption_r16 = new(CellGroupConfig_uplinkTxSwitchingOption_r16)
				if err = ie.uplinkTxSwitchingOption_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkTxSwitchingOption_r16", err)
				}
			}
			// decode uplinkTxSwitchingPowerBoosting_r16 optional
			if uplinkTxSwitchingPowerBoosting_r16Present {
				ie.uplinkTxSwitchingPowerBoosting_r16 = new(CellGroupConfig_uplinkTxSwitchingPowerBoosting_r16)
				if err = ie.uplinkTxSwitchingPowerBoosting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkTxSwitchingPowerBoosting_r16", err)
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

			reportUplinkTxDirectCurrentTwoCarrier_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode reportUplinkTxDirectCurrentTwoCarrier_r16 optional
			if reportUplinkTxDirectCurrentTwoCarrier_r16Present {
				ie.reportUplinkTxDirectCurrentTwoCarrier_r16 = new(CellGroupConfig_reportUplinkTxDirectCurrentTwoCarrier_r16)
				if err = ie.reportUplinkTxDirectCurrentTwoCarrier_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportUplinkTxDirectCurrentTwoCarrier_r16", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			f1c_TransferPathNRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkTxSwitching_2T_Mode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkTxSwitching_DualUL_TxState_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uu_RelayRLC_ChannelToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uu_RelayRLC_ChannelToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousU_TCI_UpdateList1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousU_TCI_UpdateList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousU_TCI_UpdateList3_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousU_TCI_UpdateList4_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rlc_BearerToReleaseListExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			iab_ResourceConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			iab_ResourceConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode f1c_TransferPathNRDC_r17 optional
			if f1c_TransferPathNRDC_r17Present {
				ie.f1c_TransferPathNRDC_r17 = new(CellGroupConfig_f1c_TransferPathNRDC_r17)
				if err = ie.f1c_TransferPathNRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode f1c_TransferPathNRDC_r17", err)
				}
			}
			// decode uplinkTxSwitching_2T_Mode_r17 optional
			if uplinkTxSwitching_2T_Mode_r17Present {
				ie.uplinkTxSwitching_2T_Mode_r17 = new(CellGroupConfig_uplinkTxSwitching_2T_Mode_r17)
				if err = ie.uplinkTxSwitching_2T_Mode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkTxSwitching_2T_Mode_r17", err)
				}
			}
			// decode uplinkTxSwitching_DualUL_TxState_r17 optional
			if uplinkTxSwitching_DualUL_TxState_r17Present {
				ie.uplinkTxSwitching_DualUL_TxState_r17 = new(CellGroupConfig_uplinkTxSwitching_DualUL_TxState_r17)
				if err = ie.uplinkTxSwitching_DualUL_TxState_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkTxSwitching_DualUL_TxState_r17", err)
				}
			}
			// decode uu_RelayRLC_ChannelToAddModList_r17 optional
			if uu_RelayRLC_ChannelToAddModList_r17Present {
				tmp_uu_RelayRLC_ChannelToAddModList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelConfig_r17]([]*Uu_RelayRLC_ChannelConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				fn_uu_RelayRLC_ChannelToAddModList_r17 := func() *Uu_RelayRLC_ChannelConfig_r17 {
					return new(Uu_RelayRLC_ChannelConfig_r17)
				}
				if err = tmp_uu_RelayRLC_ChannelToAddModList_r17.Decode(extReader, fn_uu_RelayRLC_ChannelToAddModList_r17); err != nil {
					return utils.WrapError("Decode uu_RelayRLC_ChannelToAddModList_r17", err)
				}
				ie.uu_RelayRLC_ChannelToAddModList_r17 = []Uu_RelayRLC_ChannelConfig_r17{}
				for _, i := range tmp_uu_RelayRLC_ChannelToAddModList_r17.Value {
					ie.uu_RelayRLC_ChannelToAddModList_r17 = append(ie.uu_RelayRLC_ChannelToAddModList_r17, *i)
				}
			}
			// decode uu_RelayRLC_ChannelToReleaseList_r17 optional
			if uu_RelayRLC_ChannelToReleaseList_r17Present {
				tmp_uu_RelayRLC_ChannelToReleaseList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelID_r17]([]*Uu_RelayRLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				fn_uu_RelayRLC_ChannelToReleaseList_r17 := func() *Uu_RelayRLC_ChannelID_r17 {
					return new(Uu_RelayRLC_ChannelID_r17)
				}
				if err = tmp_uu_RelayRLC_ChannelToReleaseList_r17.Decode(extReader, fn_uu_RelayRLC_ChannelToReleaseList_r17); err != nil {
					return utils.WrapError("Decode uu_RelayRLC_ChannelToReleaseList_r17", err)
				}
				ie.uu_RelayRLC_ChannelToReleaseList_r17 = []Uu_RelayRLC_ChannelID_r17{}
				for _, i := range tmp_uu_RelayRLC_ChannelToReleaseList_r17.Value {
					ie.uu_RelayRLC_ChannelToReleaseList_r17 = append(ie.uu_RelayRLC_ChannelToReleaseList_r17, *i)
				}
			}
			// decode simultaneousU_TCI_UpdateList1_r17 optional
			if simultaneousU_TCI_UpdateList1_r17Present {
				tmp_simultaneousU_TCI_UpdateList1_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousU_TCI_UpdateList1_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousU_TCI_UpdateList1_r17.Decode(extReader, fn_simultaneousU_TCI_UpdateList1_r17); err != nil {
					return utils.WrapError("Decode simultaneousU_TCI_UpdateList1_r17", err)
				}
				ie.simultaneousU_TCI_UpdateList1_r17 = []ServCellIndex{}
				for _, i := range tmp_simultaneousU_TCI_UpdateList1_r17.Value {
					ie.simultaneousU_TCI_UpdateList1_r17 = append(ie.simultaneousU_TCI_UpdateList1_r17, *i)
				}
			}
			// decode simultaneousU_TCI_UpdateList2_r17 optional
			if simultaneousU_TCI_UpdateList2_r17Present {
				tmp_simultaneousU_TCI_UpdateList2_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousU_TCI_UpdateList2_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousU_TCI_UpdateList2_r17.Decode(extReader, fn_simultaneousU_TCI_UpdateList2_r17); err != nil {
					return utils.WrapError("Decode simultaneousU_TCI_UpdateList2_r17", err)
				}
				ie.simultaneousU_TCI_UpdateList2_r17 = []ServCellIndex{}
				for _, i := range tmp_simultaneousU_TCI_UpdateList2_r17.Value {
					ie.simultaneousU_TCI_UpdateList2_r17 = append(ie.simultaneousU_TCI_UpdateList2_r17, *i)
				}
			}
			// decode simultaneousU_TCI_UpdateList3_r17 optional
			if simultaneousU_TCI_UpdateList3_r17Present {
				tmp_simultaneousU_TCI_UpdateList3_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousU_TCI_UpdateList3_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousU_TCI_UpdateList3_r17.Decode(extReader, fn_simultaneousU_TCI_UpdateList3_r17); err != nil {
					return utils.WrapError("Decode simultaneousU_TCI_UpdateList3_r17", err)
				}
				ie.simultaneousU_TCI_UpdateList3_r17 = []ServCellIndex{}
				for _, i := range tmp_simultaneousU_TCI_UpdateList3_r17.Value {
					ie.simultaneousU_TCI_UpdateList3_r17 = append(ie.simultaneousU_TCI_UpdateList3_r17, *i)
				}
			}
			// decode simultaneousU_TCI_UpdateList4_r17 optional
			if simultaneousU_TCI_UpdateList4_r17Present {
				tmp_simultaneousU_TCI_UpdateList4_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_simultaneousU_TCI_UpdateList4_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_simultaneousU_TCI_UpdateList4_r17.Decode(extReader, fn_simultaneousU_TCI_UpdateList4_r17); err != nil {
					return utils.WrapError("Decode simultaneousU_TCI_UpdateList4_r17", err)
				}
				ie.simultaneousU_TCI_UpdateList4_r17 = []ServCellIndex{}
				for _, i := range tmp_simultaneousU_TCI_UpdateList4_r17.Value {
					ie.simultaneousU_TCI_UpdateList4_r17 = append(ie.simultaneousU_TCI_UpdateList4_r17, *i)
				}
			}
			// decode rlc_BearerToReleaseListExt_r17 optional
			if rlc_BearerToReleaseListExt_r17Present {
				tmp_rlc_BearerToReleaseListExt_r17 := utils.NewSequence[*LogicalChannelIdentityExt_r17]([]*LogicalChannelIdentityExt_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
				fn_rlc_BearerToReleaseListExt_r17 := func() *LogicalChannelIdentityExt_r17 {
					return new(LogicalChannelIdentityExt_r17)
				}
				if err = tmp_rlc_BearerToReleaseListExt_r17.Decode(extReader, fn_rlc_BearerToReleaseListExt_r17); err != nil {
					return utils.WrapError("Decode rlc_BearerToReleaseListExt_r17", err)
				}
				ie.rlc_BearerToReleaseListExt_r17 = []LogicalChannelIdentityExt_r17{}
				for _, i := range tmp_rlc_BearerToReleaseListExt_r17.Value {
					ie.rlc_BearerToReleaseListExt_r17 = append(ie.rlc_BearerToReleaseListExt_r17, *i)
				}
			}
			// decode iab_ResourceConfigToAddModList_r17 optional
			if iab_ResourceConfigToAddModList_r17Present {
				tmp_iab_ResourceConfigToAddModList_r17 := utils.NewSequence[*IAB_ResourceConfig_r17]([]*IAB_ResourceConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				fn_iab_ResourceConfigToAddModList_r17 := func() *IAB_ResourceConfig_r17 {
					return new(IAB_ResourceConfig_r17)
				}
				if err = tmp_iab_ResourceConfigToAddModList_r17.Decode(extReader, fn_iab_ResourceConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode iab_ResourceConfigToAddModList_r17", err)
				}
				ie.iab_ResourceConfigToAddModList_r17 = []IAB_ResourceConfig_r17{}
				for _, i := range tmp_iab_ResourceConfigToAddModList_r17.Value {
					ie.iab_ResourceConfigToAddModList_r17 = append(ie.iab_ResourceConfigToAddModList_r17, *i)
				}
			}
			// decode iab_ResourceConfigToReleaseList_r17 optional
			if iab_ResourceConfigToReleaseList_r17Present {
				tmp_iab_ResourceConfigToReleaseList_r17 := utils.NewSequence[*IAB_ResourceConfigID_r17]([]*IAB_ResourceConfigID_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				fn_iab_ResourceConfigToReleaseList_r17 := func() *IAB_ResourceConfigID_r17 {
					return new(IAB_ResourceConfigID_r17)
				}
				if err = tmp_iab_ResourceConfigToReleaseList_r17.Decode(extReader, fn_iab_ResourceConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode iab_ResourceConfigToReleaseList_r17", err)
				}
				ie.iab_ResourceConfigToReleaseList_r17 = []IAB_ResourceConfigID_r17{}
				for _, i := range tmp_iab_ResourceConfigToReleaseList_r17.Value {
					ie.iab_ResourceConfigToReleaseList_r17 = append(ie.iab_ResourceConfigToReleaseList_r17, *i)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			reportUplinkTxDirectCurrentMoreCarrier_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode reportUplinkTxDirectCurrentMoreCarrier_r17 optional
			if reportUplinkTxDirectCurrentMoreCarrier_r17Present {
				ie.reportUplinkTxDirectCurrentMoreCarrier_r17 = new(ReportUplinkTxDirectCurrentMoreCarrier_r17)
				if err = ie.reportUplinkTxDirectCurrentMoreCarrier_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportUplinkTxDirectCurrentMoreCarrier_r17", err)
				}
			}
		}
	}
	return nil
}
