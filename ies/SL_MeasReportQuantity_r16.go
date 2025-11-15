package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_MeasReportQuantity_r16_Choice_nothing uint64 = iota
	SL_MeasReportQuantity_r16_Choice_sl_RSRP_r16
)

type SL_MeasReportQuantity_r16 struct {
	Choice      uint64
	sl_RSRP_r16 bool `madatory`
}

func (ie *SL_MeasReportQuantity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_MeasReportQuantity_r16_Choice_sl_RSRP_r16:
		if err = w.WriteBoolean(ie.sl_RSRP_r16); err != nil {
			err = utils.WrapError("Encode sl_RSRP_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_MeasReportQuantity_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_MeasReportQuantity_r16_Choice_sl_RSRP_r16:
		var tmp_bool_sl_RSRP_r16 bool
		if tmp_bool_sl_RSRP_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode sl_RSRP_r16", err)
		}
		ie.sl_RSRP_r16 = tmp_bool_sl_RSRP_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
