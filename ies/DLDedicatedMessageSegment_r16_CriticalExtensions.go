package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DLDedicatedMessageSegment_r16_CriticalExtensions_Choice_nothing uint64 = iota
	DLDedicatedMessageSegment_r16_CriticalExtensions_Choice_dlDedicatedMessageSegment_r16
	DLDedicatedMessageSegment_r16_CriticalExtensions_Choice_criticalExtensionsFuture
)

type DLDedicatedMessageSegment_r16_CriticalExtensions struct {
	Choice                        uint64
	dlDedicatedMessageSegment_r16 *DLDedicatedMessageSegment_r16_IEs
	criticalExtensionsFuture      interface{} `madatory`
}

func (ie *DLDedicatedMessageSegment_r16_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DLDedicatedMessageSegment_r16_CriticalExtensions_Choice_dlDedicatedMessageSegment_r16:
		if err = ie.dlDedicatedMessageSegment_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode dlDedicatedMessageSegment_r16", err)
		}
	case DLDedicatedMessageSegment_r16_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DLDedicatedMessageSegment_r16_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DLDedicatedMessageSegment_r16_CriticalExtensions_Choice_dlDedicatedMessageSegment_r16:
		ie.dlDedicatedMessageSegment_r16 = new(DLDedicatedMessageSegment_r16_IEs)
		if err = ie.dlDedicatedMessageSegment_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dlDedicatedMessageSegment_r16", err)
		}
	case DLDedicatedMessageSegment_r16_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
