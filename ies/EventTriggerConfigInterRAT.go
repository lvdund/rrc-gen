package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigInterRAT struct {
	eventId                       EventTriggerConfigInterRAT_eventId                        `madatory`
	rsType                        NR_RS_Type                                                `madatory,ext`
	reportInterval                ReportInterval                                            `madatory,ext`
	reportAmount                  EventTriggerConfigInterRAT_reportAmount                   `madatory,ext`
	reportQuantity                MeasReportQuantity                                        `madatory,ext`
	maxReportCells                int64                                                     `lb:1,ub:maxCellReport,madatory,ext`
	reportQuantityUTRA_FDD_r16    *MeasReportQuantityUTRA_FDD_r16                           `optional,ext-3`
	includeCommonLocationInfo_r16 *EventTriggerConfigInterRAT_includeCommonLocationInfo_r16 `optional,ext-4`
	includeBT_Meas_r16            *BT_NameList_r16                                          `optional,ext-4,setuprelease`
	includeWLAN_Meas_r16          *WLAN_NameList_r16                                        `optional,ext-4,setuprelease`
	includeSensor_Meas_r16        *Sensor_NameList_r16                                      `optional,ext-4,setuprelease`
	reportQuantityRelay_r17       *SL_MeasReportQuantity_r16                                `optional,ext-5`
}

func (ie *EventTriggerConfigInterRAT) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.reportQuantityUTRA_FDD_r16 != nil || ie.includeCommonLocationInfo_r16 != nil || ie.includeBT_Meas_r16 != nil || ie.includeWLAN_Meas_r16 != nil || ie.includeSensor_Meas_r16 != nil || ie.reportQuantityRelay_r17 != nil
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
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.reportQuantityUTRA_FDD_r16 != nil, ie.includeCommonLocationInfo_r16 != nil || ie.includeBT_Meas_r16 != nil || ie.includeWLAN_Meas_r16 != nil || ie.includeSensor_Meas_r16 != nil, ie.reportQuantityRelay_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap EventTriggerConfigInterRAT", err)
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.reportQuantityUTRA_FDD_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode reportQuantityUTRA_FDD_r16 optional
			if ie.reportQuantityUTRA_FDD_r16 != nil {
				if err = ie.reportQuantityUTRA_FDD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportQuantityUTRA_FDD_r16", err)
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
			optionals_ext_4 := []bool{ie.includeCommonLocationInfo_r16 != nil, ie.includeBT_Meas_r16 != nil, ie.includeWLAN_Meas_r16 != nil, ie.includeSensor_Meas_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.reportQuantityRelay_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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

func (ie *EventTriggerConfigInterRAT) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.eventId.Decode(r); err != nil {
		return utils.WrapError("Decode eventId", err)
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			reportQuantityUTRA_FDD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode reportQuantityUTRA_FDD_r16 optional
			if reportQuantityUTRA_FDD_r16Present {
				ie.reportQuantityUTRA_FDD_r16 = new(MeasReportQuantityUTRA_FDD_r16)
				if err = ie.reportQuantityUTRA_FDD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportQuantityUTRA_FDD_r16", err)
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
			// decode includeCommonLocationInfo_r16 optional
			if includeCommonLocationInfo_r16Present {
				ie.includeCommonLocationInfo_r16 = new(EventTriggerConfigInterRAT_includeCommonLocationInfo_r16)
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
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			reportQuantityRelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
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
