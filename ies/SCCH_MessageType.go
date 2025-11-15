package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCCH_MessageType_Choice_nothing uint64 = iota
	SCCH_MessageType_Choice_c1
	SCCH_MessageType_Choice_messageClassExtension
)

type SCCH_MessageType struct {
	Choice                uint64
	c1                    *SCCH_MessageType_C1
	messageClassExtension *SCCH_MessageType_MessageClassExtension
}

func (ie *SCCH_MessageType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_Choice_c1:
		if err = ie.c1.Encode(w); err != nil {
			err = utils.WrapError("Encode c1", err)
		}
	case SCCH_MessageType_Choice_messageClassExtension:
		if err = ie.messageClassExtension.Encode(w); err != nil {
			err = utils.WrapError("Encode messageClassExtension", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SCCH_MessageType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SCCH_MessageType_Choice_c1:
		ie.c1 = new(SCCH_MessageType_C1)
		if err = ie.c1.Decode(r); err != nil {
			return utils.WrapError("Decode c1", err)
		}
	case SCCH_MessageType_Choice_messageClassExtension:
		ie.messageClassExtension = new(SCCH_MessageType_MessageClassExtension)
		if err = ie.messageClassExtension.Decode(r); err != nil {
			return utils.WrapError("Decode messageClassExtension", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
