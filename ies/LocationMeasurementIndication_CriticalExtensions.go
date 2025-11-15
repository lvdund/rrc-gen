package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LocationMeasurementIndication_CriticalExtensions_Choice_nothing uint64 = iota
	LocationMeasurementIndication_CriticalExtensions_Choice_locationMeasurementIndication
	LocationMeasurementIndication_CriticalExtensions_Choice_criticalExtensionsFuture
)

type LocationMeasurementIndication_CriticalExtensions struct {
	Choice                        uint64
	locationMeasurementIndication *LocationMeasurementIndication_IEs
	criticalExtensionsFuture      interface{} `madatory`
}

func (ie *LocationMeasurementIndication_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementIndication_CriticalExtensions_Choice_locationMeasurementIndication:
		if err = ie.locationMeasurementIndication.Encode(w); err != nil {
			err = utils.WrapError("Encode locationMeasurementIndication", err)
		}
	case LocationMeasurementIndication_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *LocationMeasurementIndication_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationMeasurementIndication_CriticalExtensions_Choice_locationMeasurementIndication:
		ie.locationMeasurementIndication = new(LocationMeasurementIndication_IEs)
		if err = ie.locationMeasurementIndication.Decode(r); err != nil {
			return utils.WrapError("Decode locationMeasurementIndication", err)
		}
	case LocationMeasurementIndication_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
