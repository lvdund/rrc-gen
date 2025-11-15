package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasurementReport_CriticalExtensions_Choice_nothing uint64 = iota
	MeasurementReport_CriticalExtensions_Choice_measurementReport
	MeasurementReport_CriticalExtensions_Choice_criticalExtensionsFuture
)

type MeasurementReport_CriticalExtensions struct {
	Choice                   uint64
	measurementReport        *MeasurementReport_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *MeasurementReport_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReport_CriticalExtensions_Choice_measurementReport:
		if err = ie.measurementReport.Encode(w); err != nil {
			err = utils.WrapError("Encode measurementReport", err)
		}
	case MeasurementReport_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasurementReport_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementReport_CriticalExtensions_Choice_measurementReport:
		ie.measurementReport = new(MeasurementReport_IEs)
		if err = ie.measurementReport.Decode(r); err != nil {
			return utils.WrapError("Decode measurementReport", err)
		}
	case MeasurementReport_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
