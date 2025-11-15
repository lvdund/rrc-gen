package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfigurationSidelink_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReconfigurationSidelink_CriticalExtensions_Choice_rrcReconfigurationSidelink_r16
	RRCReconfigurationSidelink_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCReconfigurationSidelink_CriticalExtensions struct {
	Choice                         uint64
	rrcReconfigurationSidelink_r16 *RRCReconfigurationSidelink_r16_IEs
	criticalExtensionsFuture       interface{} `madatory`
}

func (ie *RRCReconfigurationSidelink_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationSidelink_CriticalExtensions_Choice_rrcReconfigurationSidelink_r16:
		if err = ie.rrcReconfigurationSidelink_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfigurationSidelink_r16", err)
		}
	case RRCReconfigurationSidelink_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfigurationSidelink_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationSidelink_CriticalExtensions_Choice_rrcReconfigurationSidelink_r16:
		ie.rrcReconfigurationSidelink_r16 = new(RRCReconfigurationSidelink_r16_IEs)
		if err = ie.rrcReconfigurationSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfigurationSidelink_r16", err)
		}
	case RRCReconfigurationSidelink_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
