package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MCCH_Config_r17 struct {
	mcch_RepetitionPeriodAndOffset_r17 MCCH_RepetitionPeriodAndOffset_r17          `madatory`
	mcch_WindowStartSlot_r17           int64                                       `lb:0,ub:79,madatory`
	mcch_WindowDuration_r17            *MCCH_Config_r17_mcch_WindowDuration_r17    `optional`
	mcch_ModificationPeriod_r17        MCCH_Config_r17_mcch_ModificationPeriod_r17 `madatory`
}

func (ie *MCCH_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mcch_WindowDuration_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.mcch_RepetitionPeriodAndOffset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mcch_RepetitionPeriodAndOffset_r17", err)
	}
	if err = w.WriteInteger(ie.mcch_WindowStartSlot_r17, &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
		return utils.WrapError("WriteInteger mcch_WindowStartSlot_r17", err)
	}
	if ie.mcch_WindowDuration_r17 != nil {
		if err = ie.mcch_WindowDuration_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mcch_WindowDuration_r17", err)
		}
	}
	if err = ie.mcch_ModificationPeriod_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mcch_ModificationPeriod_r17", err)
	}
	return nil
}

func (ie *MCCH_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var mcch_WindowDuration_r17Present bool
	if mcch_WindowDuration_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.mcch_RepetitionPeriodAndOffset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mcch_RepetitionPeriodAndOffset_r17", err)
	}
	var tmp_int_mcch_WindowStartSlot_r17 int64
	if tmp_int_mcch_WindowStartSlot_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
		return utils.WrapError("ReadInteger mcch_WindowStartSlot_r17", err)
	}
	ie.mcch_WindowStartSlot_r17 = tmp_int_mcch_WindowStartSlot_r17
	if mcch_WindowDuration_r17Present {
		ie.mcch_WindowDuration_r17 = new(MCCH_Config_r17_mcch_WindowDuration_r17)
		if err = ie.mcch_WindowDuration_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mcch_WindowDuration_r17", err)
		}
	}
	if err = ie.mcch_ModificationPeriod_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mcch_ModificationPeriod_r17", err)
	}
	return nil
}
