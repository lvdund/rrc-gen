package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfigurationComplete_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReconfigurationComplete_CriticalExtensions_Choice_rrcReconfigurationComplete
	RRCReconfigurationComplete_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCReconfigurationComplete_CriticalExtensions struct {
	Choice                     uint64
	rrcReconfigurationComplete *RRCReconfigurationComplete_IEs
	criticalExtensionsFuture   interface{} `madatory`
}

func (ie *RRCReconfigurationComplete_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationComplete_CriticalExtensions_Choice_rrcReconfigurationComplete:
		if err = ie.rrcReconfigurationComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfigurationComplete", err)
		}
	case RRCReconfigurationComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfigurationComplete_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationComplete_CriticalExtensions_Choice_rrcReconfigurationComplete:
		ie.rrcReconfigurationComplete = new(RRCReconfigurationComplete_IEs)
		if err = ie.rrcReconfigurationComplete.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfigurationComplete", err)
		}
	case RRCReconfigurationComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
