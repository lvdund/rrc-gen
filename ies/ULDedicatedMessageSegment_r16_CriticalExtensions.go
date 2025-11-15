package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ULDedicatedMessageSegment_r16_CriticalExtensions_Choice_nothing uint64 = iota
	ULDedicatedMessageSegment_r16_CriticalExtensions_Choice_ulDedicatedMessageSegment_r16
	ULDedicatedMessageSegment_r16_CriticalExtensions_Choice_criticalExtensionsFuture
)

type ULDedicatedMessageSegment_r16_CriticalExtensions struct {
	Choice                        uint64
	ulDedicatedMessageSegment_r16 *ULDedicatedMessageSegment_r16_IEs
	criticalExtensionsFuture      interface{} `madatory`
}

func (ie *ULDedicatedMessageSegment_r16_CriticalExtensions) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULDedicatedMessageSegment_r16_CriticalExtensions_Choice_ulDedicatedMessageSegment_r16:
		if err = ie.ulDedicatedMessageSegment_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode ulDedicatedMessageSegment_r16", err)
		}
	case ULDedicatedMessageSegment_r16_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ULDedicatedMessageSegment_r16_CriticalExtensions) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ULDedicatedMessageSegment_r16_CriticalExtensions_Choice_ulDedicatedMessageSegment_r16:
		ie.ulDedicatedMessageSegment_r16 = new(ULDedicatedMessageSegment_r16_IEs)
		if err = ie.ulDedicatedMessageSegment_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ulDedicatedMessageSegment_r16", err)
		}
	case ULDedicatedMessageSegment_r16_CriticalExtensions_Choice_criticalExtensionsFuture:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
