package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityModeFailure_CriticalExtensions_Choice_nothing uint64 = iota
	SecurityModeFailure_CriticalExtensions_Choice_securityModeFailure
	SecurityModeFailure_CriticalExtensions_Choice_criticalExtensionsFuture
)

type SecurityModeFailure_CriticalExtensions struct {
	Choice                   uint64
	securityModeFailure      *SecurityModeFailure_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *SecurityModeFailure_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeFailure_CriticalExtensions_Choice_securityModeFailure:
		if err = ie.securityModeFailure.Encode(w); err != nil {
			err = utils.WrapError("Encode securityModeFailure", err)
		}
	case SecurityModeFailure_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SecurityModeFailure_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeFailure_CriticalExtensions_Choice_securityModeFailure:
		ie.securityModeFailure = new(SecurityModeFailure_IEs)
		if err = ie.securityModeFailure.Decode(r); err != nil {
			return utils.WrapError("Decode securityModeFailure", err)
		}
	case SecurityModeFailure_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
