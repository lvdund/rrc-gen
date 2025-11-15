package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCRelease_CriticalExtensions_Choice_nothing uint64 = iota
	RRCRelease_CriticalExtensions_Choice_rrcRelease
	RRCRelease_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCRelease_CriticalExtensions struct {
	Choice                   uint64
	rrcRelease               *RRCRelease_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCRelease_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCRelease_CriticalExtensions_Choice_rrcRelease:
		if err = ie.rrcRelease.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcRelease", err)
		}
	case RRCRelease_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCRelease_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCRelease_CriticalExtensions_Choice_rrcRelease:
		ie.rrcRelease = new(RRCRelease_IEs)
		if err = ie.rrcRelease.Decode(r); err != nil {
			return utils.WrapError("Decode rrcRelease", err)
		}
	case RRCRelease_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
