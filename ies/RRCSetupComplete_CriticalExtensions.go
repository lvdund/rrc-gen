package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetupComplete_CriticalExtensions_Choice_nothing uint64 = iota
	RRCSetupComplete_CriticalExtensions_Choice_rrcSetupComplete
	RRCSetupComplete_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCSetupComplete_CriticalExtensions struct {
	Choice                   uint64
	rrcSetupComplete         *RRCSetupComplete_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCSetupComplete_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_CriticalExtensions_Choice_rrcSetupComplete:
		if err = ie.rrcSetupComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcSetupComplete", err)
		}
	case RRCSetupComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSetupComplete_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetupComplete_CriticalExtensions_Choice_rrcSetupComplete:
		ie.rrcSetupComplete = new(RRCSetupComplete_IEs)
		if err = ie.rrcSetupComplete.Decode(r); err != nil {
			return utils.WrapError("Decode rrcSetupComplete", err)
		}
	case RRCSetupComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
