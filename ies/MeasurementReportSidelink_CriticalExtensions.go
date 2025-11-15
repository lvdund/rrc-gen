package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasurementReportSidelink_CriticalExtensions_Choice_nothing uint64 = iota
	MeasurementReportSidelink_CriticalExtensions_Choice_measurementReportSidelink_r16
	MeasurementReportSidelink_CriticalExtensions_Choice_criticalExtensionsFuture
)

type MeasurementReportSidelink_CriticalExtensions struct {
	Choice                        uint64
	measurementReportSidelink_r16 *MeasurementReportSidelink_r16_IEs
	criticalExtensionsFuture      interface{} `madatory`
}

func (ie *MeasurementReportSidelink_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReportSidelink_CriticalExtensions_Choice_measurementReportSidelink_r16:
		if err = ie.measurementReportSidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode measurementReportSidelink_r16", err)
		}
	case MeasurementReportSidelink_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasurementReportSidelink_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReportSidelink_CriticalExtensions_Choice_measurementReportSidelink_r16:
		ie.measurementReportSidelink_r16 = new(MeasurementReportSidelink_r16_IEs)
		if err = ie.measurementReportSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measurementReportSidelink_r16", err)
		}
	case MeasurementReportSidelink_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
