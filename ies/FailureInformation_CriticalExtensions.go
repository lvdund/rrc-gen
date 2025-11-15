package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FailureInformation_CriticalExtensions_Choice_nothing uint64 = iota
	FailureInformation_CriticalExtensions_Choice_failureInformation
	FailureInformation_CriticalExtensions_Choice_criticalExtensionsFuture
)

type FailureInformation_CriticalExtensions struct {
	Choice                   uint64
	failureInformation       *FailureInformation_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *FailureInformation_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FailureInformation_CriticalExtensions_Choice_failureInformation:
		if err = ie.failureInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode failureInformation", err)
		}
	case FailureInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *FailureInformation_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FailureInformation_CriticalExtensions_Choice_failureInformation:
		ie.failureInformation = new(FailureInformation_IEs)
		if err = ie.failureInformation.Decode(r); err != nil {
			return utils.WrapError("Decode failureInformation", err)
		}
	case FailureInformation_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
