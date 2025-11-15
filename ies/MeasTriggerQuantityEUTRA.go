package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasTriggerQuantityEUTRA_Choice_nothing uint64 = iota
	MeasTriggerQuantityEUTRA_Choice_rsrp
	MeasTriggerQuantityEUTRA_Choice_rsrq
	MeasTriggerQuantityEUTRA_Choice_sinr
)

type MeasTriggerQuantityEUTRA struct {
	Choice uint64
	rsrp   *RSRP_RangeEUTRA
	rsrq   *RSRQ_RangeEUTRA
	sinr   *SINR_RangeEUTRA
}

func (ie *MeasTriggerQuantityEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityEUTRA_Choice_rsrp:
		if err = ie.rsrp.Encode(w); err != nil {
			err = utils.WrapError("Encode rsrp", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_rsrq:
		if err = ie.rsrq.Encode(w); err != nil {
			err = utils.WrapError("Encode rsrq", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_sinr:
		if err = ie.sinr.Encode(w); err != nil {
			err = utils.WrapError("Encode sinr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasTriggerQuantityEUTRA) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityEUTRA_Choice_rsrp:
		ie.rsrp = new(RSRP_RangeEUTRA)
		if err = ie.rsrp.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_rsrq:
		ie.rsrq = new(RSRQ_RangeEUTRA)
		if err = ie.rsrq.Decode(r); err != nil {
			return utils.WrapError("Decode rsrq", err)
		}
	case MeasTriggerQuantityEUTRA_Choice_sinr:
		ie.sinr = new(SINR_RangeEUTRA)
		if err = ie.sinr.Decode(r); err != nil {
			return utils.WrapError("Decode sinr", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
