package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasConfig_s_MeasureConfig_Choice_nothing uint64 = iota
	MeasConfig_s_MeasureConfig_Choice_ssb_RSRP
	MeasConfig_s_MeasureConfig_Choice_csi_RSRP
)

type MeasConfig_s_MeasureConfig struct {
	Choice   uint64
	ssb_RSRP *RSRP_Range
	csi_RSRP *RSRP_Range
}

func (ie *MeasConfig_s_MeasureConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasConfig_s_MeasureConfig_Choice_ssb_RSRP:
		if err = ie.ssb_RSRP.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_RSRP", err)
		}
	case MeasConfig_s_MeasureConfig_Choice_csi_RSRP:
		if err = ie.csi_RSRP.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RSRP", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasConfig_s_MeasureConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasConfig_s_MeasureConfig_Choice_ssb_RSRP:
		ie.ssb_RSRP = new(RSRP_Range)
		if err = ie.ssb_RSRP.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RSRP", err)
		}
	case MeasConfig_s_MeasureConfig_Choice_csi_RSRP:
		ie.csi_RSRP = new(RSRP_Range)
		if err = ie.csi_RSRP.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RSRP", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
