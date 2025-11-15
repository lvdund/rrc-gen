package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_CCCH1_MessageType_C1_Choice_nothing uint64 = iota
	UL_CCCH1_MessageType_C1_Choice_rrcResumeRequest1
	UL_CCCH1_MessageType_C1_Choice_spare3
	UL_CCCH1_MessageType_C1_Choice_spare2
	UL_CCCH1_MessageType_C1_Choice_spare1
)

type UL_CCCH1_MessageType_C1 struct {
	Choice            uint64
	rrcResumeRequest1 *RRCResumeRequest1
	spare3            uper.NULL `madatory`
	spare2            uper.NULL `madatory`
	spare1            uper.NULL `madatory`
}

func (ie *UL_CCCH1_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH1_MessageType_C1_Choice_rrcResumeRequest1:
		if err = ie.rrcResumeRequest1.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcResumeRequest1", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UL_CCCH1_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UL_CCCH1_MessageType_C1_Choice_rrcResumeRequest1:
		ie.rrcResumeRequest1 = new(RRCResumeRequest1)
		if err = ie.rrcResumeRequest1.Decode(r); err != nil {
			return utils.WrapError("Decode rrcResumeRequest1", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case UL_CCCH1_MessageType_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
