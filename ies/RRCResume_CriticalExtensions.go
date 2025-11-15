package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCResume_CriticalExtensions_Choice_nothing uint64 = iota
	RRCResume_CriticalExtensions_Choice_rrcResume
	RRCResume_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCResume_CriticalExtensions struct {
	Choice                   uint64
	rrcResume                *RRCResume_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCResume_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_CriticalExtensions_Choice_rrcResume:
		if err = ie.rrcResume.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcResume", err)
		}
	case RRCResume_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCResume_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCResume_CriticalExtensions_Choice_rrcResume:
		ie.rrcResume = new(RRCResume_IEs)
		if err = ie.rrcResume.Decode(r); err != nil {
			return utils.WrapError("Decode rrcResume", err)
		}
	case RRCResume_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
