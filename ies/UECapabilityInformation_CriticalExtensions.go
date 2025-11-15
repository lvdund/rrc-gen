package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UECapabilityInformation_CriticalExtensions_Choice_nothing uint64 = iota
	UECapabilityInformation_CriticalExtensions_Choice_ueCapabilityInformation
	UECapabilityInformation_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UECapabilityInformation_CriticalExtensions struct {
	Choice                   uint64
	ueCapabilityInformation  *UECapabilityInformation_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *UECapabilityInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityInformation_CriticalExtensions_Choice_ueCapabilityInformation:
		if err = ie.ueCapabilityInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode ueCapabilityInformation", err)
		}
	case UECapabilityInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UECapabilityInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityInformation_CriticalExtensions_Choice_ueCapabilityInformation:
		ie.ueCapabilityInformation = new(UECapabilityInformation_IEs)
		if err = ie.ueCapabilityInformation.Decode(r); err != nil {
			return utils.WrapError("Decode ueCapabilityInformation", err)
		}
	case UECapabilityInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
