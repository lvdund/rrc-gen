package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PeriodicalReportConfig struct {
	rsType                        NR_RS_Type                                            `madatory`
	reportInterval                ReportInterval                                        `madatory`
	reportAmount                  PeriodicalReportConfig_reportAmount                   `madatory`
	reportQuantityCell            MeasReportQuantity                                    `madatory`
	maxReportCells                int64                                                 `lb:1,ub:maxCellReport,madatory`
	reportQuantityRS_Indexes      *MeasReportQuantity                                   `optional`
	maxNrofRS_IndexesToReport     *int64                                                `lb:1,ub:maxNrofIndexesToReport,optional`
	includeBeamMeasurements       bool                                                  `madatory`
	useAllowedCellList            bool                                                  `madatory`
	measRSSI_ReportConfig_r16     *MeasRSSI_ReportConfig_r16                            `optional,ext-1`
	includeCommonLocationInfo_r16 *PeriodicalReportConfig_includeCommonLocationInfo_r16 `optional,ext-1`
	includeBT_Meas_r16            *BT_NameList_r16                                      `optional,ext-1,setuprelease`
	includeWLAN_Meas_r16          *WLAN_NameList_r16                                    `optional,ext-1,setuprelease`
	includeSensor_Meas_r16        *Sensor_NameList_r16                                  `optional,ext-1,setuprelease`
	ul_DelayValueConfig_r16       *UL_DelayValueConfig_r16                              `optional,ext-1,setuprelease`
	reportAddNeighMeas_r16        *PeriodicalReportConfig_reportAddNeighMeas_r16        `optional,ext-1`
	ul_ExcessDelayConfig_r17      *UL_ExcessDelayConfig_r17                             `optional,ext-2,setuprelease`
	coarseLocationRequest_r17     *PeriodicalReportConfig_coarseLocationRequest_r17     `optional,ext-2`
	reportQuantityRelay_r17       *SL_MeasReportQuantity_r16                            `optional,ext-2`
}

func (ie *PeriodicalReportConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.measRSSI_ReportConfig_r16 != nil || ie.includeCommonLocationInfo_r16 != nil || ie.includeBT_Meas_r16 != nil || ie.includeWLAN_Meas_r16 != nil || ie.includeSensor_Meas_r16 != nil || ie.ul_DelayValueConfig_r16 != nil || ie.reportAddNeighMeas_r16 != nil || ie.ul_ExcessDelayConfig_r17 != nil || ie.coarseLocationRequest_r17 != nil || ie.reportQuantityRelay_r17 != nil
	preambleBits := []bool{hasExtensions, ie.reportQuantityRS_Indexes != nil, ie.maxNrofRS_IndexesToReport != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rsType.Encode(w); err != nil {
		return utils.WrapError("Encode rsType", err)
	}
	if err = ie.reportInterval.Encode(w); err != nil {
		return utils.WrapError("Encode reportInterval", err)
	}
	if err = ie.reportAmount.Encode(w); err != nil {
		return utils.WrapError("Encode reportAmount", err)
	}
	if err = ie.reportQuantityCell.Encode(w); err != nil {
		return utils.WrapError("Encode reportQuantityCell", err)
	}
	if err = w.WriteInteger(ie.maxReportCells, &uper.Constraint{Lb: 1, Ub: maxCellReport}, false); err != nil {
		return utils.WrapError("WriteInteger maxReportCells", err)
	}
	if ie.reportQuantityRS_Indexes != nil {
		if err = ie.reportQuantityRS_Indexes.Encode(w); err != nil {
			return utils.WrapError("Encode reportQuantityRS_Indexes", err)
		}
	}
	if ie.maxNrofRS_IndexesToReport != nil {
		if err = w.WriteInteger(*ie.maxNrofRS_IndexesToReport, &uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
			return utils.WrapError("Encode maxNrofRS_IndexesToReport", err)
		}
	}
	if err = w.WriteBoolean(ie.includeBeamMeasurements); err != nil {
		return utils.WrapError("WriteBoolean includeBeamMeasurements", err)
	}
	if err = w.WriteBoolean(ie.useAllowedCellList); err != nil {
		return utils.WrapError("WriteBoolean useAllowedCellList", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.measRSSI_ReportConfig_r16 != nil || ie.includeCommonLocationInfo_r16 != nil || ie.includeBT_Meas_r16 != nil || ie.includeWLAN_Meas_r16 != nil || ie.includeSensor_Meas_r16 != nil || ie.ul_DelayValueConfig_r16 != nil || ie.reportAddNeighMeas_r16 != nil, ie.ul_ExcessDelayConfig_r17 != nil || ie.coarseLocationRequest_r17 != nil || ie.reportQuantityRelay_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PeriodicalReportConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.measRSSI_ReportConfig_r16 != nil, ie.includeCommonLocationInfo_r16 != nil, ie.includeBT_Meas_r16 != nil, ie.includeWLAN_Meas_r16 != nil, ie.includeSensor_Meas_r16 != nil, ie.ul_DelayValueConfig_r16 != nil, ie.reportAddNeighMeas_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode measRSSI_ReportConfig_r16 optional
			if ie.measRSSI_ReportConfig_r16 != nil {
				if err = ie.measRSSI_ReportConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measRSSI_ReportConfig_r16", err)
				}
			}
			// encode includeCommonLocationInfo_r16 optional
			if ie.includeCommonLocationInfo_r16 != nil {
				if err = ie.includeCommonLocationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode includeCommonLocationInfo_r16", err)
				}
			}
			// encode includeBT_Meas_r16 optional
			if ie.includeBT_Meas_r16 != nil {
				tmp_includeBT_Meas_r16 := utils.SetupRelease[*BT_NameList_r16]{
					Setup: ie.includeBT_Meas_r16,
				}
				if err = tmp_includeBT_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode includeBT_Meas_r16", err)
				}
			}
			// encode includeWLAN_Meas_r16 optional
			if ie.includeWLAN_Meas_r16 != nil {
				tmp_includeWLAN_Meas_r16 := utils.SetupRelease[*WLAN_NameList_r16]{
					Setup: ie.includeWLAN_Meas_r16,
				}
				if err = tmp_includeWLAN_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode includeWLAN_Meas_r16", err)
				}
			}
			// encode includeSensor_Meas_r16 optional
			if ie.includeSensor_Meas_r16 != nil {
				tmp_includeSensor_Meas_r16 := utils.SetupRelease[*Sensor_NameList_r16]{
					Setup: ie.includeSensor_Meas_r16,
				}
				if err = tmp_includeSensor_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode includeSensor_Meas_r16", err)
				}
			}
			// encode ul_DelayValueConfig_r16 optional
			if ie.ul_DelayValueConfig_r16 != nil {
				tmp_ul_DelayValueConfig_r16 := utils.SetupRelease[*UL_DelayValueConfig_r16]{
					Setup: ie.ul_DelayValueConfig_r16,
				}
				if err = tmp_ul_DelayValueConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_DelayValueConfig_r16", err)
				}
			}
			// encode reportAddNeighMeas_r16 optional
			if ie.reportAddNeighMeas_r16 != nil {
				if err = ie.reportAddNeighMeas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportAddNeighMeas_r16", err)
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
			optionals_ext_2 := []bool{ie.ul_ExcessDelayConfig_r17 != nil, ie.coarseLocationRequest_r17 != nil, ie.reportQuantityRelay_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ul_ExcessDelayConfig_r17 optional
			if ie.ul_ExcessDelayConfig_r17 != nil {
				tmp_ul_ExcessDelayConfig_r17 := utils.SetupRelease[*UL_ExcessDelayConfig_r17]{
					Setup: ie.ul_ExcessDelayConfig_r17,
				}
				if err = tmp_ul_ExcessDelayConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_ExcessDelayConfig_r17", err)
				}
			}
			// encode coarseLocationRequest_r17 optional
			if ie.coarseLocationRequest_r17 != nil {
				if err = ie.coarseLocationRequest_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode coarseLocationRequest_r17", err)
				}
			}
			// encode reportQuantityRelay_r17 optional
			if ie.reportQuantityRelay_r17 != nil {
				if err = ie.reportQuantityRelay_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportQuantityRelay_r17", err)
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

func (ie *PeriodicalReportConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var reportQuantityRS_IndexesPresent bool
	if reportQuantityRS_IndexesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNrofRS_IndexesToReportPresent bool
	if maxNrofRS_IndexesToReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rsType.Decode(r); err != nil {
		return utils.WrapError("Decode rsType", err)
	}
	if err = ie.reportInterval.Decode(r); err != nil {
		return utils.WrapError("Decode reportInterval", err)
	}
	if err = ie.reportAmount.Decode(r); err != nil {
		return utils.WrapError("Decode reportAmount", err)
	}
	if err = ie.reportQuantityCell.Decode(r); err != nil {
		return utils.WrapError("Decode reportQuantityCell", err)
	}
	var tmp_int_maxReportCells int64
	if tmp_int_maxReportCells, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxCellReport}, false); err != nil {
		return utils.WrapError("ReadInteger maxReportCells", err)
	}
	ie.maxReportCells = tmp_int_maxReportCells
	if reportQuantityRS_IndexesPresent {
		ie.reportQuantityRS_Indexes = new(MeasReportQuantity)
		if err = ie.reportQuantityRS_Indexes.Decode(r); err != nil {
			return utils.WrapError("Decode reportQuantityRS_Indexes", err)
		}
	}
	if maxNrofRS_IndexesToReportPresent {
		var tmp_int_maxNrofRS_IndexesToReport int64
		if tmp_int_maxNrofRS_IndexesToReport, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofIndexesToReport}, false); err != nil {
			return utils.WrapError("Decode maxNrofRS_IndexesToReport", err)
		}
		ie.maxNrofRS_IndexesToReport = &tmp_int_maxNrofRS_IndexesToReport
	}
	var tmp_bool_includeBeamMeasurements bool
	if tmp_bool_includeBeamMeasurements, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean includeBeamMeasurements", err)
	}
	ie.includeBeamMeasurements = tmp_bool_includeBeamMeasurements
	var tmp_bool_useAllowedCellList bool
	if tmp_bool_useAllowedCellList, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean useAllowedCellList", err)
	}
	ie.useAllowedCellList = tmp_bool_useAllowedCellList

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

			measRSSI_ReportConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			includeCommonLocationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			includeBT_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			includeWLAN_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			includeSensor_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_DelayValueConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			reportAddNeighMeas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode measRSSI_ReportConfig_r16 optional
			if measRSSI_ReportConfig_r16Present {
				ie.measRSSI_ReportConfig_r16 = new(MeasRSSI_ReportConfig_r16)
				if err = ie.measRSSI_ReportConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode measRSSI_ReportConfig_r16", err)
				}
			}
			// decode includeCommonLocationInfo_r16 optional
			if includeCommonLocationInfo_r16Present {
				ie.includeCommonLocationInfo_r16 = new(PeriodicalReportConfig_includeCommonLocationInfo_r16)
				if err = ie.includeCommonLocationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode includeCommonLocationInfo_r16", err)
				}
			}
			// decode includeBT_Meas_r16 optional
			if includeBT_Meas_r16Present {
				tmp_includeBT_Meas_r16 := utils.SetupRelease[*BT_NameList_r16]{}
				if err = tmp_includeBT_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode includeBT_Meas_r16", err)
				}
				ie.includeBT_Meas_r16 = tmp_includeBT_Meas_r16.Setup
			}
			// decode includeWLAN_Meas_r16 optional
			if includeWLAN_Meas_r16Present {
				tmp_includeWLAN_Meas_r16 := utils.SetupRelease[*WLAN_NameList_r16]{}
				if err = tmp_includeWLAN_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode includeWLAN_Meas_r16", err)
				}
				ie.includeWLAN_Meas_r16 = tmp_includeWLAN_Meas_r16.Setup
			}
			// decode includeSensor_Meas_r16 optional
			if includeSensor_Meas_r16Present {
				tmp_includeSensor_Meas_r16 := utils.SetupRelease[*Sensor_NameList_r16]{}
				if err = tmp_includeSensor_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode includeSensor_Meas_r16", err)
				}
				ie.includeSensor_Meas_r16 = tmp_includeSensor_Meas_r16.Setup
			}
			// decode ul_DelayValueConfig_r16 optional
			if ul_DelayValueConfig_r16Present {
				tmp_ul_DelayValueConfig_r16 := utils.SetupRelease[*UL_DelayValueConfig_r16]{}
				if err = tmp_ul_DelayValueConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_DelayValueConfig_r16", err)
				}
				ie.ul_DelayValueConfig_r16 = tmp_ul_DelayValueConfig_r16.Setup
			}
			// decode reportAddNeighMeas_r16 optional
			if reportAddNeighMeas_r16Present {
				ie.reportAddNeighMeas_r16 = new(PeriodicalReportConfig_reportAddNeighMeas_r16)
				if err = ie.reportAddNeighMeas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportAddNeighMeas_r16", err)
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

			ul_ExcessDelayConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			coarseLocationRequest_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			reportQuantityRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ul_ExcessDelayConfig_r17 optional
			if ul_ExcessDelayConfig_r17Present {
				tmp_ul_ExcessDelayConfig_r17 := utils.SetupRelease[*UL_ExcessDelayConfig_r17]{}
				if err = tmp_ul_ExcessDelayConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_ExcessDelayConfig_r17", err)
				}
				ie.ul_ExcessDelayConfig_r17 = tmp_ul_ExcessDelayConfig_r17.Setup
			}
			// decode coarseLocationRequest_r17 optional
			if coarseLocationRequest_r17Present {
				ie.coarseLocationRequest_r17 = new(PeriodicalReportConfig_coarseLocationRequest_r17)
				if err = ie.coarseLocationRequest_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode coarseLocationRequest_r17", err)
				}
			}
			// decode reportQuantityRelay_r17 optional
			if reportQuantityRelay_r17Present {
				ie.reportQuantityRelay_r17 = new(SL_MeasReportQuantity_r16)
				if err = ie.reportQuantityRelay_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportQuantityRelay_r17", err)
				}
			}
		}
	}
	return nil
}
