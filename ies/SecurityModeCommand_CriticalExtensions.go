package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SecurityModeCommand_CriticalExtensions_Choice_nothing uint64 = iota
	SecurityModeCommand_CriticalExtensions_Choice_securityModeCommand
	SecurityModeCommand_CriticalExtensions_Choice_criticalExtensionsFuture
)

type SecurityModeCommand_CriticalExtensions struct {
	Choice                   uint64
	securityModeCommand      *SecurityModeCommand_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *SecurityModeCommand_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeCommand_CriticalExtensions_Choice_securityModeCommand:
		if err = ie.securityModeCommand.Encode(w); err != nil {
			err = utils.WrapError("Encode securityModeCommand", err)
		}
	case SecurityModeCommand_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SecurityModeCommand_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SecurityModeCommand_CriticalExtensions_Choice_securityModeCommand:
		ie.securityModeCommand = new(SecurityModeCommand_IEs)
		if err = ie.securityModeCommand.Decode(r); err != nil {
			return utils.WrapError("Decode securityModeCommand", err)
		}
	case SecurityModeCommand_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
