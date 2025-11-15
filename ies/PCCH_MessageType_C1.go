package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCCH_MessageType_C1_Choice_nothing uint64 = iota
	PCCH_MessageType_C1_Choice_paging
	PCCH_MessageType_C1_Choice_spare1
)

type PCCH_MessageType_C1 struct {
	Choice uint64
	paging *Paging
	spare1 uper.NULL `madatory`
}

func (ie *PCCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_MessageType_C1_Choice_paging:
		if err = ie.paging.Encode(w); err != nil {
			err = utils.WrapError("Encode paging", err)
		}
	case PCCH_MessageType_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PCCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_MessageType_C1_Choice_paging:
		ie.paging = new(Paging)
		if err = ie.paging.Decode(r); err != nil {
			return utils.WrapError("Decode paging", err)
		}
	case PCCH_MessageType_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
