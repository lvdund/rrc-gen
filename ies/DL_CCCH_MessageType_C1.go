package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_CCCH_MessageType_C1_Choice_nothing uint64 = iota
	DL_CCCH_MessageType_C1_Choice_rrcReject
	DL_CCCH_MessageType_C1_Choice_rrcSetup
	DL_CCCH_MessageType_C1_Choice_spare2
	DL_CCCH_MessageType_C1_Choice_spare1
)

type DL_CCCH_MessageType_C1 struct {
	Choice    uint64
	rrcReject *RRCReject
	rrcSetup  *RRCSetup
	spare2    uper.NULL `madatory`
	spare1    uper.NULL `madatory`
}

func (ie *DL_CCCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_CCCH_MessageType_C1_Choice_rrcReject:
		if err = ie.rrcReject.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcReject", err)
		}
	case DL_CCCH_MessageType_C1_Choice_rrcSetup:
		if err = ie.rrcSetup.Encode(w); err != nil {
			err = utils.WrapError("Encode rrcSetup", err)
		}
	case DL_CCCH_MessageType_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case DL_CCCH_MessageType_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_CCCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_CCCH_MessageType_C1_Choice_rrcReject:
		ie.rrcReject = new(RRCReject)
		if err = ie.rrcReject.Decode(r); err != nil {
			return utils.WrapError("Decode rrcReject", err)
		}
	case DL_CCCH_MessageType_C1_Choice_rrcSetup:
		ie.rrcSetup = new(RRCSetup)
		if err = ie.rrcSetup.Decode(r); err != nil {
			return utils.WrapError("Decode rrcSetup", err)
		}
	case DL_CCCH_MessageType_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case DL_CCCH_MessageType_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
