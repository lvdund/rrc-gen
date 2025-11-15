package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_MessageType_r17_C1_Choice_nothing uint64 = iota
	MCCH_MessageType_r17_C1_Choice_mbsBroadcastConfiguration_r17
	MCCH_MessageType_r17_C1_Choice_spare1
)

type MCCH_MessageType_r17_C1 struct {
	Choice                        uint64
	mbsBroadcastConfiguration_r17 *MBSBroadcastConfiguration_r17
	spare1                        uper.NULL `madatory`
}

func (ie *MCCH_MessageType_r17_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_MessageType_r17_C1_Choice_mbsBroadcastConfiguration_r17:
		if err = ie.mbsBroadcastConfiguration_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode mbsBroadcastConfiguration_r17", err)
		}
	case MCCH_MessageType_r17_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MCCH_MessageType_r17_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_MessageType_r17_C1_Choice_mbsBroadcastConfiguration_r17:
		ie.mbsBroadcastConfiguration_r17 = new(MBSBroadcastConfiguration_r17)
		if err = ie.mbsBroadcastConfiguration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mbsBroadcastConfiguration_r17", err)
		}
	case MCCH_MessageType_r17_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
