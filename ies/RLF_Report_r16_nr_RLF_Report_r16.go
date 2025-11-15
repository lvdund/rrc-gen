package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLF_Report_r16_nr_RLF_Report_r16 struct {
	measResultLastServCell_r16     MeasResultRLFNR_r16                                        `madatory`
	measResultNeighCells_r16       *RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16 `optional`
	c_RNTI_r16                     RNTI_Value                                                 `madatory`
	previousPCellId_r16            *RLF_Report_r16_nr_RLF_Report_r16_previousPCellId_r16      `optional`
	failedPCellId_r16              RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16         `madatory`
	reconnectCellId_r16            *RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16      `optional`
	timeUntilReconnection_r16      *TimeUntilReconnection_r16                                 `optional`
	reestablishmentCellId_r16      *CGI_Info_Logging_r16                                      `optional`
	timeConnFailure_r16            *int64                                                     `lb:0,ub:1023,optional`
	timeSinceFailure_r16           TimeSinceFailure_r16                                       `madatory`
	connectionFailureType_r16      RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16 `madatory`
	rlf_Cause_r16                  RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16             `madatory`
	locationInfo_r16               *LocationInfo_r16                                          `optional`
	noSuitableCellFound_r16        *RLF_Report_r16_nr_RLF_Report_r16_noSuitableCellFound_r16  `optional`
	ra_InformationCommon_r16       *RA_InformationCommon_r16                                  `optional`
	csi_rsRLMConfigBitmap_v1650    *uper.BitString                                            `lb:96,ub:96,optional`
	lastHO_Type_r17                *RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17          `optional`
	timeConnSourceDAPS_Failure_r17 *TimeConnSourceDAPS_Failure_r17                            `optional`
	timeSinceCHO_Reconfig_r17      *TimeSinceCHO_Reconfig_r17                                 `optional`
	choCellId_r17                  *RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17            `optional`
	choCandidateCellList_r17       *ChoCandidateCellList_r17                                  `optional`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultNeighCells_r16 != nil, ie.previousPCellId_r16 != nil, ie.reconnectCellId_r16 != nil, ie.timeUntilReconnection_r16 != nil, ie.reestablishmentCellId_r16 != nil, ie.timeConnFailure_r16 != nil, ie.locationInfo_r16 != nil, ie.noSuitableCellFound_r16 != nil, ie.ra_InformationCommon_r16 != nil, ie.csi_rsRLMConfigBitmap_v1650 != nil, ie.lastHO_Type_r17 != nil, ie.timeConnSourceDAPS_Failure_r17 != nil, ie.timeSinceCHO_Reconfig_r17 != nil, ie.choCellId_r17 != nil, ie.choCandidateCellList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measResultLastServCell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResultLastServCell_r16", err)
	}
	if ie.measResultNeighCells_r16 != nil {
		if err = ie.measResultNeighCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultNeighCells_r16", err)
		}
	}
	if err = ie.c_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode c_RNTI_r16", err)
	}
	if ie.previousPCellId_r16 != nil {
		if err = ie.previousPCellId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode previousPCellId_r16", err)
		}
	}
	if err = ie.failedPCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode failedPCellId_r16", err)
	}
	if ie.reconnectCellId_r16 != nil {
		if err = ie.reconnectCellId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reconnectCellId_r16", err)
		}
	}
	if ie.timeUntilReconnection_r16 != nil {
		if err = ie.timeUntilReconnection_r16.Encode(w); err != nil {
			return utils.WrapError("Encode timeUntilReconnection_r16", err)
		}
	}
	if ie.reestablishmentCellId_r16 != nil {
		if err = ie.reestablishmentCellId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishmentCellId_r16", err)
		}
	}
	if ie.timeConnFailure_r16 != nil {
		if err = w.WriteInteger(*ie.timeConnFailure_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode timeConnFailure_r16", err)
		}
	}
	if err = ie.timeSinceFailure_r16.Encode(w); err != nil {
		return utils.WrapError("Encode timeSinceFailure_r16", err)
	}
	if err = ie.connectionFailureType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode connectionFailureType_r16", err)
	}
	if err = ie.rlf_Cause_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rlf_Cause_r16", err)
	}
	if ie.locationInfo_r16 != nil {
		if err = ie.locationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode locationInfo_r16", err)
		}
	}
	if ie.noSuitableCellFound_r16 != nil {
		if err = ie.noSuitableCellFound_r16.Encode(w); err != nil {
			return utils.WrapError("Encode noSuitableCellFound_r16", err)
		}
	}
	if ie.ra_InformationCommon_r16 != nil {
		if err = ie.ra_InformationCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ra_InformationCommon_r16", err)
		}
	}
	if ie.csi_rsRLMConfigBitmap_v1650 != nil {
		if err = w.WriteBitString(ie.csi_rsRLMConfigBitmap_v1650.Bytes, uint(ie.csi_rsRLMConfigBitmap_v1650.NumBits), &uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Encode csi_rsRLMConfigBitmap_v1650", err)
		}
	}
	if ie.lastHO_Type_r17 != nil {
		if err = ie.lastHO_Type_r17.Encode(w); err != nil {
			return utils.WrapError("Encode lastHO_Type_r17", err)
		}
	}
	if ie.timeConnSourceDAPS_Failure_r17 != nil {
		if err = ie.timeConnSourceDAPS_Failure_r17.Encode(w); err != nil {
			return utils.WrapError("Encode timeConnSourceDAPS_Failure_r17", err)
		}
	}
	if ie.timeSinceCHO_Reconfig_r17 != nil {
		if err = ie.timeSinceCHO_Reconfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode timeSinceCHO_Reconfig_r17", err)
		}
	}
	if ie.choCellId_r17 != nil {
		if err = ie.choCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode choCellId_r17", err)
		}
	}
	if ie.choCandidateCellList_r17 != nil {
		if err = ie.choCandidateCellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode choCandidateCellList_r17", err)
		}
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16) Decode(r *uper.UperReader) error {
	var err error
	var measResultNeighCells_r16Present bool
	if measResultNeighCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var previousPCellId_r16Present bool
	if previousPCellId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reconnectCellId_r16Present bool
	if reconnectCellId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeUntilReconnection_r16Present bool
	if timeUntilReconnection_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishmentCellId_r16Present bool
	if reestablishmentCellId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeConnFailure_r16Present bool
	if timeConnFailure_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var locationInfo_r16Present bool
	if locationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var noSuitableCellFound_r16Present bool
	if noSuitableCellFound_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_InformationCommon_r16Present bool
	if ra_InformationCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_rsRLMConfigBitmap_v1650Present bool
	if csi_rsRLMConfigBitmap_v1650Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lastHO_Type_r17Present bool
	if lastHO_Type_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeConnSourceDAPS_Failure_r17Present bool
	if timeConnSourceDAPS_Failure_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timeSinceCHO_Reconfig_r17Present bool
	if timeSinceCHO_Reconfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var choCellId_r17Present bool
	if choCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var choCandidateCellList_r17Present bool
	if choCandidateCellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measResultLastServCell_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResultLastServCell_r16", err)
	}
	if measResultNeighCells_r16Present {
		ie.measResultNeighCells_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16)
		if err = ie.measResultNeighCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultNeighCells_r16", err)
		}
	}
	if err = ie.c_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode c_RNTI_r16", err)
	}
	if previousPCellId_r16Present {
		ie.previousPCellId_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_previousPCellId_r16)
		if err = ie.previousPCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode previousPCellId_r16", err)
		}
	}
	if err = ie.failedPCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode failedPCellId_r16", err)
	}
	if reconnectCellId_r16Present {
		ie.reconnectCellId_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16)
		if err = ie.reconnectCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reconnectCellId_r16", err)
		}
	}
	if timeUntilReconnection_r16Present {
		ie.timeUntilReconnection_r16 = new(TimeUntilReconnection_r16)
		if err = ie.timeUntilReconnection_r16.Decode(r); err != nil {
			return utils.WrapError("Decode timeUntilReconnection_r16", err)
		}
	}
	if reestablishmentCellId_r16Present {
		ie.reestablishmentCellId_r16 = new(CGI_Info_Logging_r16)
		if err = ie.reestablishmentCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishmentCellId_r16", err)
		}
	}
	if timeConnFailure_r16Present {
		var tmp_int_timeConnFailure_r16 int64
		if tmp_int_timeConnFailure_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode timeConnFailure_r16", err)
		}
		ie.timeConnFailure_r16 = &tmp_int_timeConnFailure_r16
	}
	if err = ie.timeSinceFailure_r16.Decode(r); err != nil {
		return utils.WrapError("Decode timeSinceFailure_r16", err)
	}
	if err = ie.connectionFailureType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode connectionFailureType_r16", err)
	}
	if err = ie.rlf_Cause_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rlf_Cause_r16", err)
	}
	if locationInfo_r16Present {
		ie.locationInfo_r16 = new(LocationInfo_r16)
		if err = ie.locationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode locationInfo_r16", err)
		}
	}
	if noSuitableCellFound_r16Present {
		ie.noSuitableCellFound_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_noSuitableCellFound_r16)
		if err = ie.noSuitableCellFound_r16.Decode(r); err != nil {
			return utils.WrapError("Decode noSuitableCellFound_r16", err)
		}
	}
	if ra_InformationCommon_r16Present {
		ie.ra_InformationCommon_r16 = new(RA_InformationCommon_r16)
		if err = ie.ra_InformationCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ra_InformationCommon_r16", err)
		}
	}
	if csi_rsRLMConfigBitmap_v1650Present {
		var tmp_bs_csi_rsRLMConfigBitmap_v1650 []byte
		var n_csi_rsRLMConfigBitmap_v1650 uint
		if tmp_bs_csi_rsRLMConfigBitmap_v1650, n_csi_rsRLMConfigBitmap_v1650, err = r.ReadBitString(&uper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode csi_rsRLMConfigBitmap_v1650", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_csi_rsRLMConfigBitmap_v1650,
			NumBits: uint64(n_csi_rsRLMConfigBitmap_v1650),
		}
		ie.csi_rsRLMConfigBitmap_v1650 = &tmp_bitstring
	}
	if lastHO_Type_r17Present {
		ie.lastHO_Type_r17 = new(RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17)
		if err = ie.lastHO_Type_r17.Decode(r); err != nil {
			return utils.WrapError("Decode lastHO_Type_r17", err)
		}
	}
	if timeConnSourceDAPS_Failure_r17Present {
		ie.timeConnSourceDAPS_Failure_r17 = new(TimeConnSourceDAPS_Failure_r17)
		if err = ie.timeConnSourceDAPS_Failure_r17.Decode(r); err != nil {
			return utils.WrapError("Decode timeConnSourceDAPS_Failure_r17", err)
		}
	}
	if timeSinceCHO_Reconfig_r17Present {
		ie.timeSinceCHO_Reconfig_r17 = new(TimeSinceCHO_Reconfig_r17)
		if err = ie.timeSinceCHO_Reconfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode timeSinceCHO_Reconfig_r17", err)
		}
	}
	if choCellId_r17Present {
		ie.choCellId_r17 = new(RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17)
		if err = ie.choCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode choCellId_r17", err)
		}
	}
	if choCandidateCellList_r17Present {
		ie.choCandidateCellList_r17 = new(ChoCandidateCellList_r17)
		if err = ie.choCandidateCellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode choCandidateCellList_r17", err)
		}
	}
	return nil
}
