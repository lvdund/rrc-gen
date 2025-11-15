package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_RS_Resource_Mobility_slotConfig_Choice_nothing uint64 = iota
	CSI_RS_Resource_Mobility_slotConfig_Choice_ms4
	CSI_RS_Resource_Mobility_slotConfig_Choice_ms5
	CSI_RS_Resource_Mobility_slotConfig_Choice_ms10
	CSI_RS_Resource_Mobility_slotConfig_Choice_ms20
	CSI_RS_Resource_Mobility_slotConfig_Choice_ms40
)

type CSI_RS_Resource_Mobility_slotConfig struct {
	Choice uint64
	ms4    int64 `lb:0,ub:31,madatory`
	ms5    int64 `lb:0,ub:39,madatory`
	ms10   int64 `lb:0,ub:79,madatory`
	ms20   int64 `lb:0,ub:159,madatory`
	ms40   int64 `lb:0,ub:319,madatory`
}

func (ie *CSI_RS_Resource_Mobility_slotConfig) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms4:
		if err = w.WriteInteger(int64(ie.ms4), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode ms4", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms5:
		if err = w.WriteInteger(int64(ie.ms5), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode ms5", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms10:
		if err = w.WriteInteger(int64(ie.ms10), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode ms10", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms20:
		if err = w.WriteInteger(int64(ie.ms20), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode ms20", err)
		}
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms40:
		if err = w.WriteInteger(int64(ie.ms40), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode ms40", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_RS_Resource_Mobility_slotConfig) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms4:
		var tmp_int_ms4 int64
		if tmp_int_ms4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode ms4", err)
		}
		ie.ms4 = tmp_int_ms4
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms5:
		var tmp_int_ms5 int64
		if tmp_int_ms5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode ms5", err)
		}
		ie.ms5 = tmp_int_ms5
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms10:
		var tmp_int_ms10 int64
		if tmp_int_ms10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode ms10", err)
		}
		ie.ms10 = tmp_int_ms10
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms20:
		var tmp_int_ms20 int64
		if tmp_int_ms20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode ms20", err)
		}
		ie.ms20 = tmp_int_ms20
	case CSI_RS_Resource_Mobility_slotConfig_Choice_ms40:
		var tmp_int_ms40 int64
		if tmp_int_ms40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode ms40", err)
		}
		ie.ms40 = tmp_int_ms40
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
