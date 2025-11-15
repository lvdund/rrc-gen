package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_BasedPerfMeas_Parameters_r16 struct {
	barometerMeasReport_r16            *UE_BasedPerfMeas_Parameters_r16_barometerMeasReport_r16            `optional`
	immMeasBT_r16                      *UE_BasedPerfMeas_Parameters_r16_immMeasBT_r16                      `optional`
	immMeasWLAN_r16                    *UE_BasedPerfMeas_Parameters_r16_immMeasWLAN_r16                    `optional`
	loggedMeasBT_r16                   *UE_BasedPerfMeas_Parameters_r16_loggedMeasBT_r16                   `optional`
	loggedMeasurements_r16             *UE_BasedPerfMeas_Parameters_r16_loggedMeasurements_r16             `optional`
	loggedMeasWLAN_r16                 *UE_BasedPerfMeas_Parameters_r16_loggedMeasWLAN_r16                 `optional`
	orientationMeasReport_r16          *UE_BasedPerfMeas_Parameters_r16_orientationMeasReport_r16          `optional`
	speedMeasReport_r16                *UE_BasedPerfMeas_Parameters_r16_speedMeasReport_r16                `optional`
	gnss_Location_r16                  *UE_BasedPerfMeas_Parameters_r16_gnss_Location_r16                  `optional`
	ulPDCP_Delay_r16                   *UE_BasedPerfMeas_Parameters_r16_ulPDCP_Delay_r16                   `optional`
	sigBasedLogMDT_OverrideProtect_r17 *UE_BasedPerfMeas_Parameters_r16_sigBasedLogMDT_OverrideProtect_r17 `optional,ext-1`
	multipleCEF_Report_r17             *UE_BasedPerfMeas_Parameters_r16_multipleCEF_Report_r17             `optional,ext-1`
	excessPacketDelay_r17              *UE_BasedPerfMeas_Parameters_r16_excessPacketDelay_r17              `optional,ext-1`
	earlyMeasLog_r17                   *UE_BasedPerfMeas_Parameters_r16_earlyMeasLog_r17                   `optional,ext-1`
}

func (ie *UE_BasedPerfMeas_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sigBasedLogMDT_OverrideProtect_r17 != nil || ie.multipleCEF_Report_r17 != nil || ie.excessPacketDelay_r17 != nil || ie.earlyMeasLog_r17 != nil
	preambleBits := []bool{hasExtensions, ie.barometerMeasReport_r16 != nil, ie.immMeasBT_r16 != nil, ie.immMeasWLAN_r16 != nil, ie.loggedMeasBT_r16 != nil, ie.loggedMeasurements_r16 != nil, ie.loggedMeasWLAN_r16 != nil, ie.orientationMeasReport_r16 != nil, ie.speedMeasReport_r16 != nil, ie.gnss_Location_r16 != nil, ie.ulPDCP_Delay_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.barometerMeasReport_r16 != nil {
		if err = ie.barometerMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode barometerMeasReport_r16", err)
		}
	}
	if ie.immMeasBT_r16 != nil {
		if err = ie.immMeasBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode immMeasBT_r16", err)
		}
	}
	if ie.immMeasWLAN_r16 != nil {
		if err = ie.immMeasWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode immMeasWLAN_r16", err)
		}
	}
	if ie.loggedMeasBT_r16 != nil {
		if err = ie.loggedMeasBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode loggedMeasBT_r16", err)
		}
	}
	if ie.loggedMeasurements_r16 != nil {
		if err = ie.loggedMeasurements_r16.Encode(w); err != nil {
			return utils.WrapError("Encode loggedMeasurements_r16", err)
		}
	}
	if ie.loggedMeasWLAN_r16 != nil {
		if err = ie.loggedMeasWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode loggedMeasWLAN_r16", err)
		}
	}
	if ie.orientationMeasReport_r16 != nil {
		if err = ie.orientationMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode orientationMeasReport_r16", err)
		}
	}
	if ie.speedMeasReport_r16 != nil {
		if err = ie.speedMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode speedMeasReport_r16", err)
		}
	}
	if ie.gnss_Location_r16 != nil {
		if err = ie.gnss_Location_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gnss_Location_r16", err)
		}
	}
	if ie.ulPDCP_Delay_r16 != nil {
		if err = ie.ulPDCP_Delay_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ulPDCP_Delay_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sigBasedLogMDT_OverrideProtect_r17 != nil || ie.multipleCEF_Report_r17 != nil || ie.excessPacketDelay_r17 != nil || ie.earlyMeasLog_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UE_BasedPerfMeas_Parameters_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sigBasedLogMDT_OverrideProtect_r17 != nil, ie.multipleCEF_Report_r17 != nil, ie.excessPacketDelay_r17 != nil, ie.earlyMeasLog_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sigBasedLogMDT_OverrideProtect_r17 optional
			if ie.sigBasedLogMDT_OverrideProtect_r17 != nil {
				if err = ie.sigBasedLogMDT_OverrideProtect_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sigBasedLogMDT_OverrideProtect_r17", err)
				}
			}
			// encode multipleCEF_Report_r17 optional
			if ie.multipleCEF_Report_r17 != nil {
				if err = ie.multipleCEF_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multipleCEF_Report_r17", err)
				}
			}
			// encode excessPacketDelay_r17 optional
			if ie.excessPacketDelay_r17 != nil {
				if err = ie.excessPacketDelay_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode excessPacketDelay_r17", err)
				}
			}
			// encode earlyMeasLog_r17 optional
			if ie.earlyMeasLog_r17 != nil {
				if err = ie.earlyMeasLog_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode earlyMeasLog_r17", err)
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

func (ie *UE_BasedPerfMeas_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var barometerMeasReport_r16Present bool
	if barometerMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var immMeasBT_r16Present bool
	if immMeasBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var immMeasWLAN_r16Present bool
	if immMeasWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var loggedMeasBT_r16Present bool
	if loggedMeasBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var loggedMeasurements_r16Present bool
	if loggedMeasurements_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var loggedMeasWLAN_r16Present bool
	if loggedMeasWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var orientationMeasReport_r16Present bool
	if orientationMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var speedMeasReport_r16Present bool
	if speedMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var gnss_Location_r16Present bool
	if gnss_Location_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ulPDCP_Delay_r16Present bool
	if ulPDCP_Delay_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if barometerMeasReport_r16Present {
		ie.barometerMeasReport_r16 = new(UE_BasedPerfMeas_Parameters_r16_barometerMeasReport_r16)
		if err = ie.barometerMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode barometerMeasReport_r16", err)
		}
	}
	if immMeasBT_r16Present {
		ie.immMeasBT_r16 = new(UE_BasedPerfMeas_Parameters_r16_immMeasBT_r16)
		if err = ie.immMeasBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode immMeasBT_r16", err)
		}
	}
	if immMeasWLAN_r16Present {
		ie.immMeasWLAN_r16 = new(UE_BasedPerfMeas_Parameters_r16_immMeasWLAN_r16)
		if err = ie.immMeasWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode immMeasWLAN_r16", err)
		}
	}
	if loggedMeasBT_r16Present {
		ie.loggedMeasBT_r16 = new(UE_BasedPerfMeas_Parameters_r16_loggedMeasBT_r16)
		if err = ie.loggedMeasBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode loggedMeasBT_r16", err)
		}
	}
	if loggedMeasurements_r16Present {
		ie.loggedMeasurements_r16 = new(UE_BasedPerfMeas_Parameters_r16_loggedMeasurements_r16)
		if err = ie.loggedMeasurements_r16.Decode(r); err != nil {
			return utils.WrapError("Decode loggedMeasurements_r16", err)
		}
	}
	if loggedMeasWLAN_r16Present {
		ie.loggedMeasWLAN_r16 = new(UE_BasedPerfMeas_Parameters_r16_loggedMeasWLAN_r16)
		if err = ie.loggedMeasWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode loggedMeasWLAN_r16", err)
		}
	}
	if orientationMeasReport_r16Present {
		ie.orientationMeasReport_r16 = new(UE_BasedPerfMeas_Parameters_r16_orientationMeasReport_r16)
		if err = ie.orientationMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode orientationMeasReport_r16", err)
		}
	}
	if speedMeasReport_r16Present {
		ie.speedMeasReport_r16 = new(UE_BasedPerfMeas_Parameters_r16_speedMeasReport_r16)
		if err = ie.speedMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode speedMeasReport_r16", err)
		}
	}
	if gnss_Location_r16Present {
		ie.gnss_Location_r16 = new(UE_BasedPerfMeas_Parameters_r16_gnss_Location_r16)
		if err = ie.gnss_Location_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gnss_Location_r16", err)
		}
	}
	if ulPDCP_Delay_r16Present {
		ie.ulPDCP_Delay_r16 = new(UE_BasedPerfMeas_Parameters_r16_ulPDCP_Delay_r16)
		if err = ie.ulPDCP_Delay_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ulPDCP_Delay_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			sigBasedLogMDT_OverrideProtect_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multipleCEF_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			excessPacketDelay_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			earlyMeasLog_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sigBasedLogMDT_OverrideProtect_r17 optional
			if sigBasedLogMDT_OverrideProtect_r17Present {
				ie.sigBasedLogMDT_OverrideProtect_r17 = new(UE_BasedPerfMeas_Parameters_r16_sigBasedLogMDT_OverrideProtect_r17)
				if err = ie.sigBasedLogMDT_OverrideProtect_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sigBasedLogMDT_OverrideProtect_r17", err)
				}
			}
			// decode multipleCEF_Report_r17 optional
			if multipleCEF_Report_r17Present {
				ie.multipleCEF_Report_r17 = new(UE_BasedPerfMeas_Parameters_r16_multipleCEF_Report_r17)
				if err = ie.multipleCEF_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode multipleCEF_Report_r17", err)
				}
			}
			// decode excessPacketDelay_r17 optional
			if excessPacketDelay_r17Present {
				ie.excessPacketDelay_r17 = new(UE_BasedPerfMeas_Parameters_r16_excessPacketDelay_r17)
				if err = ie.excessPacketDelay_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode excessPacketDelay_r17", err)
				}
			}
			// decode earlyMeasLog_r17 optional
			if earlyMeasLog_r17Present {
				ie.earlyMeasLog_r17 = new(UE_BasedPerfMeas_Parameters_r16_earlyMeasLog_r17)
				if err = ie.earlyMeasLog_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode earlyMeasLog_r17", err)
				}
			}
		}
	}
	return nil
}
