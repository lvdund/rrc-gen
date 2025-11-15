package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReconfigurationWithSync_rach_ConfigDedicated_Choice_nothing uint64 = iota
	ReconfigurationWithSync_rach_ConfigDedicated_Choice_uplink
	ReconfigurationWithSync_rach_ConfigDedicated_Choice_supplementaryUplink
)

type ReconfigurationWithSync_rach_ConfigDedicated struct {
	Choice              uint64
	uplink              *RACH_ConfigDedicated
	supplementaryUplink *RACH_ConfigDedicated
}

func (ie *ReconfigurationWithSync_rach_ConfigDedicated) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_uplink:
		if err = ie.uplink.Encode(w); err != nil {
			err = utils.WrapError("Encode uplink", err)
		}
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_supplementaryUplink:
		if err = ie.supplementaryUplink.Encode(w); err != nil {
			err = utils.WrapError("Encode supplementaryUplink", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReconfigurationWithSync_rach_ConfigDedicated) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_uplink:
		ie.uplink = new(RACH_ConfigDedicated)
		if err = ie.uplink.Decode(r); err != nil {
			return utils.WrapError("Decode uplink", err)
		}
	case ReconfigurationWithSync_rach_ConfigDedicated_Choice_supplementaryUplink:
		ie.supplementaryUplink = new(RACH_ConfigDedicated)
		if err = ie.supplementaryUplink.Decode(r); err != nil {
			return utils.WrapError("Decode supplementaryUplink", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
