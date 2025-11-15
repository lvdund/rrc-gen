package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UERadioAccessCapabilityInformation_CriticalExtensions_Choice_nothing uint64 = iota
	UERadioAccessCapabilityInformation_CriticalExtensions_Choice_c1
	UERadioAccessCapabilityInformation_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UERadioAccessCapabilityInformation_CriticalExtensions struct {
	Choice                   uint64
	c1                       *UERadioAccessCapabilityInformation_CriticalExtensions_C1
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *UERadioAccessCapabilityInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UERadioAccessCapabilityInformation_CriticalExtensions_Choice_c1:
		if err = ie.c1.Encode(w); err != nil {
			err = utils.WrapError("Encode c1", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UERadioAccessCapabilityInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UERadioAccessCapabilityInformation_CriticalExtensions_Choice_c1:
		ie.c1 = new(UERadioAccessCapabilityInformation_CriticalExtensions_C1)
		if err = ie.c1.Decode(r); err != nil {
			return utils.WrapError("Decode c1", err)
		}
	case UERadioAccessCapabilityInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
