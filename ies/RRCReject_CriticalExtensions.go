package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReject_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReject_CriticalExtensions_Choice_rrcReject
	RRCReject_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCReject_CriticalExtensions struct {
	Choice                   uint64
	rrcReject                *RRCReject_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCReject_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReject_CriticalExtensions_Choice_rrcReject:
		if err = ie.rrcReject.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReject", err)
		}
	case RRCReject_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReject_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReject_CriticalExtensions_Choice_rrcReject:
		ie.rrcReject = new(RRCReject_IEs)
		if err = ie.rrcReject.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReject", err)
		}
	case RRCReject_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
