package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCSetup_CriticalExtensions_Choice_nothing uint64 = iota
	RRCSetup_CriticalExtensions_Choice_rrcSetup
	RRCSetup_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCSetup_CriticalExtensions struct {
	Choice                   uint64
	rrcSetup                 *RRCSetup_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCSetup_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetup_CriticalExtensions_Choice_rrcSetup:
		if err = ie.rrcSetup.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcSetup", err)
		}
	case RRCSetup_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCSetup_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCSetup_CriticalExtensions_Choice_rrcSetup:
		ie.rrcSetup = new(RRCSetup_IEs)
		if err = ie.rrcSetup.Decode(r); err != nil {
			return utils.WrapError("Decode rrcSetup", err)
		}
	case RRCSetup_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
