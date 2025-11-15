package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReestablishmentComplete_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReestablishmentComplete_CriticalExtensions_Choice_rrcReestablishmentComplete
	RRCReestablishmentComplete_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCReestablishmentComplete_CriticalExtensions struct {
	Choice                     uint64
	rrcReestablishmentComplete *RRCReestablishmentComplete_IEs
	criticalExtensionsFuture   interface{} `madatory`
}

func (ie *RRCReestablishmentComplete_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReestablishmentComplete_CriticalExtensions_Choice_rrcReestablishmentComplete:
		if err = ie.rrcReestablishmentComplete.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReestablishmentComplete", err)
		}
	case RRCReestablishmentComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReestablishmentComplete_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReestablishmentComplete_CriticalExtensions_Choice_rrcReestablishmentComplete:
		ie.rrcReestablishmentComplete = new(RRCReestablishmentComplete_IEs)
		if err = ie.rrcReestablishmentComplete.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReestablishmentComplete", err)
		}
	case RRCReestablishmentComplete_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
