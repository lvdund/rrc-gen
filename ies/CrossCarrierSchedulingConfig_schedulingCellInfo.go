package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_nothing uint64 = iota
	CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_own
	CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_other
)

type CrossCarrierSchedulingConfig_schedulingCellInfo struct {
	Choice uint64
	own    *CrossCarrierSchedulingConfig_schedulingCellInfo_own
	other  *CrossCarrierSchedulingConfig_schedulingCellInfo_other
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_own:
		if err = ie.own.Encode(w); err != nil {
			err = utils.WrapError("Encode own", err)
		}
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_other:
		if err = ie.other.Encode(w); err != nil {
			err = utils.WrapError("Encode other", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CrossCarrierSchedulingConfig_schedulingCellInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_own:
		ie.own = new(CrossCarrierSchedulingConfig_schedulingCellInfo_own)
		if err = ie.own.Decode(r); err != nil {
			return utils.WrapError("Decode own", err)
		}
	case CrossCarrierSchedulingConfig_schedulingCellInfo_Choice_other:
		ie.other = new(CrossCarrierSchedulingConfig_schedulingCellInfo_other)
		if err = ie.other.Decode(r); err != nil {
			return utils.WrapError("Decode other", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
