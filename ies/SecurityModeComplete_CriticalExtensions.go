package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityModeComplete_CriticalExtensions_Choice_nothing uint64 = iota
	SecurityModeComplete_CriticalExtensions_Choice_securityModeComplete
	SecurityModeComplete_CriticalExtensions_Choice_criticalExtensionsFuture
)

type SecurityModeComplete_CriticalExtensions struct {
	Choice                   uint64
	securityModeComplete     *SecurityModeComplete_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *SecurityModeComplete_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeComplete_CriticalExtensions_Choice_securityModeComplete:
		if err = ie.securityModeComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode securityModeComplete", err)
		}
	case SecurityModeComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SecurityModeComplete_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeComplete_CriticalExtensions_Choice_securityModeComplete:
		ie.securityModeComplete = new(SecurityModeComplete_IEs)
		if err = ie.securityModeComplete.Decode(r); err != nil {
			return utils.WrapError("Decode securityModeComplete", err)
		}
	case SecurityModeComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
