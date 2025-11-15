package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReestablishment_CriticalExtensions_Choice_nothing uint64 = iota
	RRCReestablishment_CriticalExtensions_Choice_rrcReestablishment
	RRCReestablishment_CriticalExtensions_Choice_criticalExtensionsFuture
)

type RRCReestablishment_CriticalExtensions struct {
	Choice                   uint64
	rrcReestablishment       *RRCReestablishment_IEs
	criticalExtensionsFuture interface{} `madatory`
}

func (ie *RRCReestablishment_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReestablishment_CriticalExtensions_Choice_rrcReestablishment:
		if err = ie.rrcReestablishment.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReestablishment", err)
		}
	case RRCReestablishment_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReestablishment_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReestablishment_CriticalExtensions_Choice_rrcReestablishment:
		ie.rrcReestablishment = new(RRCReestablishment_IEs)
		if err = ie.rrcReestablishment.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReestablishment", err)
		}
	case RRCReestablishment_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
