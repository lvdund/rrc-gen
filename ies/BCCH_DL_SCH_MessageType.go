package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BCCH_DL_SCH_MessageType_Choice_nothing uint64 = iota
	BCCH_DL_SCH_MessageType_Choice_c1
	BCCH_DL_SCH_MessageType_Choice_messageClassExtension
)

type BCCH_DL_SCH_MessageType struct {
	Choice                uint64
	c1                    *BCCH_DL_SCH_MessageType_C1
	messageClassExtension interface{} `madatory`
}

func (ie *BCCH_DL_SCH_MessageType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_DL_SCH_MessageType_Choice_c1:
		if err = ie.c1.Encode(w); err != nil {
			err = utils.WrapError("Encode c1", err)
		}
	case BCCH_DL_SCH_MessageType_Choice_messageClassExtension:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BCCH_DL_SCH_MessageType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_DL_SCH_MessageType_Choice_c1:
		ie.c1 = new(BCCH_DL_SCH_MessageType_C1)
		if err = ie.c1.Decode(r); err != nil {
			return utils.WrapError("Decode c1", err)
		}
	case BCCH_DL_SCH_MessageType_Choice_messageClassExtension:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
