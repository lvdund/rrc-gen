package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_DCCH_MessageType_MessageClassExtension_Choice_nothing uint64 = iota
	UL_DCCH_MessageType_MessageClassExtension_Choice_c2
	UL_DCCH_MessageType_MessageClassExtension_Choice_messageClassExtensionFuture_r16
)

type UL_DCCH_MessageType_MessageClassExtension struct {
	Choice                          uint64
	c2                              *UL_DCCH_MessageType_MessageClassExtension_C2
	messageClassExtensionFuture_r16 interface{} `madatory`
}

func (ie *UL_DCCH_MessageType_MessageClassExtension) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_MessageClassExtension_Choice_c2:
		if err = ie.c2.Encode(w); err != nil {
			err = utils.WrapError("Encode c2", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_Choice_messageClassExtensionFuture_r16:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_DCCH_MessageType_MessageClassExtension) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_DCCH_MessageType_MessageClassExtension_Choice_c2:
		ie.c2 = new(UL_DCCH_MessageType_MessageClassExtension_C2)
		if err = ie.c2.Decode(r); err != nil {
			return utils.WrapError("Decode c2", err)
		}
	case UL_DCCH_MessageType_MessageClassExtension_Choice_messageClassExtensionFuture_r16:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
