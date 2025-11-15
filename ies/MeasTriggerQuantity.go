package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasTriggerQuantity_Choice_nothing uint64 = iota
	MeasTriggerQuantity_Choice_rsrp
	MeasTriggerQuantity_Choice_rsrq
	MeasTriggerQuantity_Choice_sinr
)

type MeasTriggerQuantity struct {
	Choice uint64
	rsrp   *RSRP_Range
	rsrq   *RSRQ_Range
	sinr   *SINR_Range
}

func (ie *MeasTriggerQuantity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantity_Choice_rsrp:
		if err = ie.rsrp.Encode(w); err != nil {
			err = utils.WrapError("Encode rsrp", err)
		}
	case MeasTriggerQuantity_Choice_rsrq:
		if err = ie.rsrq.Encode(w); err != nil {
			err = utils.WrapError("Encode rsrq", err)
		}
	case MeasTriggerQuantity_Choice_sinr:
		if err = ie.sinr.Encode(w); err != nil {
			err = utils.WrapError("Encode sinr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasTriggerQuantity) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantity_Choice_rsrp:
		ie.rsrp = new(RSRP_Range)
		if err = ie.rsrp.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp", err)
		}
	case MeasTriggerQuantity_Choice_rsrq:
		ie.rsrq = new(RSRQ_Range)
		if err = ie.rsrq.Decode(r); err != nil {
			return utils.WrapError("Decode rsrq", err)
		}
	case MeasTriggerQuantity_Choice_sinr:
		ie.sinr = new(SINR_Range)
		if err = ie.sinr.Decode(r); err != nil {
			return utils.WrapError("Decode sinr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
