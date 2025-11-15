package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig struct {
	eventId                       EventTriggerConfig_eventId                        `lb:1,ub:65525,madatory`
	rsType                        NR_RS_Type                                        `madatory,ext`
	reportInterval                ReportInterval                                    `madatory,ext`
	reportAmount                  EventTriggerConfig_reportAmount                   `madatory,ext`
	reportQuantityCell            MeasReportQuantity                                `madatory,ext`
	maxReportCells                int64                                             `lb:1,ub:maxCellReport,madatory,ext`
	reportQuantityRS_Indexes      *MeasReportQuantity                               `optional,ext`
	maxNrofRS_IndexesToReport     *int64                                            `lb:1,ub:maxNrofIndexesToReport,optional,ext`
	includeBeamMeasurements       bool                                              `madatory,ext`
	reportAddNeighMeas            *EventTriggerConfig_reportAddNeighMeas            `optional,ext`
	measRSSI_ReportConfig_r16     *MeasRSSI_ReportConfig_r16                        `optional,ext-2`
	useT312_r16                   *bool                                             `optional,ext-2`
	includeCommonLocationInfo_r16 *EventTriggerConfig_includeCommonLocationInfo_r16 `optional,ext-2`
	includeBT_Meas_r16            *BT_NameList_r16                                  `optional,ext-2,setuprelease`
	includeWLAN_Meas_r16          *WLAN_NameList_r16                                `optional,ext-2,setuprelease`
	includeSensor_Meas_r16        *Sensor_NameList_r16                              `optional,ext-2,setuprelease`
	coarseLocationRequest_r17     *EventTriggerConfig_coarseLocationRequest_r17     `optional,ext-3`
	reportQuantityRelay_r17       *SL_MeasReportQuantity_r16                        `optional,ext-3`
}

func (ie *EventTriggerConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.measRSSI_ReportConfig_r16 != nil || ie.useT312_r16 != nil || ie.includeCommonLocationInfo_r16 != nil || ie.includeBT_Meas_r16 != nil || ie.includeWLAN_Meas_r16 != nil || ie.includeSensor_Meas_r16 != nil || ie.coarseLocationRequest_r17 != nil || ie.reportQuantityRelay_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.eventId.Encode(w); err != nil {
		return utils.WrapError("Encode eventId", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.measRSSI_ReportConfig_r16 != nil || ie.useT312_r16 != nil || ie.includeCommonLocationInfo_r16 != nil || ie.includeBT_Meas_r16 != nil || ie.includeWLAN_Meas_r16 != nil || ie.includeSensor_Meas_r16 != nil, ie.coarseLocationRequest_r17 != nil || ie.reportQuantityRelay_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap EventTriggerConfig", err)
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.measRSSI_ReportConfig_r16 != nil, ie.useT312_r16 != nil, ie.includeCommonLocationInfo_r16 != nil, ie.includeBT_Meas_r16 != nil, ie.includeWLAN_Meas_r16 != nil, ie.includeSensor_Meas_r16 != nil}
			for _, bit := range optionals_ext_2 {
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
			// encode useT312_r16 optional
			if ie.useT312_r16 != nil {
				if err = extWriter.WriteBoolean(*ie.useT312_r16); err != nil {
					return utils.WrapError("Encode useT312_r16", err)
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
			optionals_ext_3 := []bool{ie.coarseLocationRequest_r17 != nil, ie.reportQuantityRelay_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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

func (ie *EventTriggerConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.eventId.Decode(r); err != nil {
		return utils.WrapError("Decode eventId", err)
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			measRSSI_ReportConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			useT312_r16Present, err := extReader.ReadBool()
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
			// decode measRSSI_ReportConfig_r16 optional
			if measRSSI_ReportConfig_r16Present {
				ie.measRSSI_ReportConfig_r16 = new(MeasRSSI_ReportConfig_r16)
				if err = ie.measRSSI_ReportConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode measRSSI_ReportConfig_r16", err)
				}
			}
			// decode useT312_r16 optional
			if useT312_r16Present {
				var tmp_bool_useT312_r16 bool
				if tmp_bool_useT312_r16, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode useT312_r16", err)
				}
				ie.useT312_r16 = &tmp_bool_useT312_r16
			}
			// decode includeCommonLocationInfo_r16 optional
			if includeCommonLocationInfo_r16Present {
				ie.includeCommonLocationInfo_r16 = new(EventTriggerConfig_includeCommonLocationInfo_r16)
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
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			coarseLocationRequest_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			reportQuantityRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode coarseLocationRequest_r17 optional
			if coarseLocationRequest_r17Present {
				ie.coarseLocationRequest_r17 = new(EventTriggerConfig_coarseLocationRequest_r17)
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
