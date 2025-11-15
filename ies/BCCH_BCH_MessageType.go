package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BCCH_BCH_MessageType_Choice_nothing uint64 = iota
	BCCH_BCH_MessageType_Choice_mib
	BCCH_BCH_MessageType_Choice_messageClassExtension
)

type BCCH_BCH_MessageType struct {
	Choice                uint64
	mib                   *MIB
	messageClassExtension interface{} `madatory`
}

func (ie *BCCH_BCH_MessageType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_BCH_MessageType_Choice_mib:
		if err = ie.mib.Encode(w); err != nil {
			err = utils.WrapError("Encode mib", err)
		}
	case BCCH_BCH_MessageType_Choice_messageClassExtension:
		// interface{} field of choice - nothing to encode
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BCCH_BCH_MessageType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_BCH_MessageType_Choice_mib:
		ie.mib = new(MIB)
		if err = ie.mib.Decode(r); err != nil {
			return utils.WrapError("Decode mib", err)
		}
	case BCCH_BCH_MessageType_Choice_messageClassExtension:
		// interface{} field of choice - nothing to decode
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
