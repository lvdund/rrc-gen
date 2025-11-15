package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasTiming_frequencyAndTiming struct {
	carrierFreq                        ARFCN_ValueNR        `madatory`
	ssbSubcarrierSpacing               SubcarrierSpacing    `madatory`
	ssb_MeasurementTimingConfiguration SSB_MTC              `madatory`
	ss_RSSI_Measurement                *SS_RSSI_Measurement `optional`
}

func (ie *MeasTiming_frequencyAndTiming) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ss_RSSI_Measurement != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq", err)
	}
	if err = ie.ssbSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode ssbSubcarrierSpacing", err)
	}
	if err = ie.ssb_MeasurementTimingConfiguration.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_MeasurementTimingConfiguration", err)
	}
	if ie.ss_RSSI_Measurement != nil {
		if err = ie.ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode ss_RSSI_Measurement", err)
		}
	}
	return nil
}

func (ie *MeasTiming_frequencyAndTiming) Decode(r *uper.UperReader) error {
	var err error
	var ss_RSSI_MeasurementPresent bool
	if ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq", err)
	}
	if err = ie.ssbSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode ssbSubcarrierSpacing", err)
	}
	if err = ie.ssb_MeasurementTimingConfiguration.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_MeasurementTimingConfiguration", err)
	}
	if ss_RSSI_MeasurementPresent {
		ie.ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode ss_RSSI_Measurement", err)
		}
	}
	return nil
}
