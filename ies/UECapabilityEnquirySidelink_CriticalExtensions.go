package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UECapabilityEnquirySidelink_CriticalExtensions_Choice_nothing uint64 = iota
	UECapabilityEnquirySidelink_CriticalExtensions_Choice_ueCapabilityEnquirySidelink_r16
	UECapabilityEnquirySidelink_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UECapabilityEnquirySidelink_CriticalExtensions struct {
	Choice                          uint64
	ueCapabilityEnquirySidelink_r16 *UECapabilityEnquirySidelink_r16_IEs
	criticalExtensionsFuture        interface{} `madatory`
}

func (ie *UECapabilityEnquirySidelink_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_ueCapabilityEnquirySidelink_r16:
		if err = ie.ueCapabilityEnquirySidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ueCapabilityEnquirySidelink_r16", err)
		}
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UECapabilityEnquirySidelink_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_ueCapabilityEnquirySidelink_r16:
		ie.ueCapabilityEnquirySidelink_r16 = new(UECapabilityEnquirySidelink_r16_IEs)
		if err = ie.ueCapabilityEnquirySidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ueCapabilityEnquirySidelink_r16", err)
		}
	case UECapabilityEnquirySidelink_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
