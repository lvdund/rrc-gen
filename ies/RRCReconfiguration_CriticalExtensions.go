package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfiguration_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReconfiguration_CriticalExtensions_Choice_rrcReconfiguration
	RRCReconfiguration_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCReconfiguration_CriticalExtensions struct {
	Choice                   uint64
	rrcReconfiguration       *RRCReconfiguration_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCReconfiguration_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfiguration_CriticalExtensions_Choice_rrcReconfiguration:
		if err = ie.rrcReconfiguration.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReconfiguration", err)
		}
	case RRCReconfiguration_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfiguration_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfiguration_CriticalExtensions_Choice_rrcReconfiguration:
		ie.rrcReconfiguration = new(RRCReconfiguration_IEs)
		if err = ie.rrcReconfiguration.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReconfiguration", err)
		}
	case RRCReconfiguration_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
