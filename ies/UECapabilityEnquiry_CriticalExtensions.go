package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UECapabilityEnquiry_CriticalExtensions_Choice_nothing uint64 = iota
	UECapabilityEnquiry_CriticalExtensions_Choice_ueCapabilityEnquiry
	UECapabilityEnquiry_CriticalExtensions_Choice_criticalExtensionsFuture
)

type UECapabilityEnquiry_CriticalExtensions struct {
	Choice                   uint64
	ueCapabilityEnquiry      *UECapabilityEnquiry_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *UECapabilityEnquiry_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityEnquiry_CriticalExtensions_Choice_ueCapabilityEnquiry:
		if err = ie.ueCapabilityEnquiry.Encode(w); err != nil {
			err = utils.WrapError("Encode ueCapabilityEnquiry", err)
		}
	case UECapabilityEnquiry_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UECapabilityEnquiry_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UECapabilityEnquiry_CriticalExtensions_Choice_ueCapabilityEnquiry:
		ie.ueCapabilityEnquiry = new(UECapabilityEnquiry_IEs)
		if err = ie.ueCapabilityEnquiry.Decode(r); err != nil {
			return utils.WrapError("Decode ueCapabilityEnquiry", err)
		}
	case UECapabilityEnquiry_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
