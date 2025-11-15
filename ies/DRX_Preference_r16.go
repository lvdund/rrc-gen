package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Preference_r16 struct {
	preferredDRX_InactivityTimer_r16 *DRX_Preference_r16_preferredDRX_InactivityTimer_r16 `optional`
	preferredDRX_LongCycle_r16       *DRX_Preference_r16_preferredDRX_LongCycle_r16       `optional`
	preferredDRX_ShortCycle_r16      *DRX_Preference_r16_preferredDRX_ShortCycle_r16      `optional`
	preferredDRX_ShortCycleTimer_r16 *int64                                               `lb:1,ub:16,optional`
}

func (ie *DRX_Preference_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.preferredDRX_InactivityTimer_r16 != nil, ie.preferredDRX_LongCycle_r16 != nil, ie.preferredDRX_ShortCycle_r16 != nil, ie.preferredDRX_ShortCycleTimer_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.preferredDRX_InactivityTimer_r16 != nil {
		if err = ie.preferredDRX_InactivityTimer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode preferredDRX_InactivityTimer_r16", err)
		}
	}
	if ie.preferredDRX_LongCycle_r16 != nil {
		if err = ie.preferredDRX_LongCycle_r16.Encode(w); err != nil {
			return utils.WrapError("Encode preferredDRX_LongCycle_r16", err)
		}
	}
	if ie.preferredDRX_ShortCycle_r16 != nil {
		if err = ie.preferredDRX_ShortCycle_r16.Encode(w); err != nil {
			return utils.WrapError("Encode preferredDRX_ShortCycle_r16", err)
		}
	}
	if ie.preferredDRX_ShortCycleTimer_r16 != nil {
		if err = w.WriteInteger(*ie.preferredDRX_ShortCycleTimer_r16, &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode preferredDRX_ShortCycleTimer_r16", err)
		}
	}
	return nil
}

func (ie *DRX_Preference_r16) Decode(r *uper.UperReader) error {
	var err error
	var preferredDRX_InactivityTimer_r16Present bool
	if preferredDRX_InactivityTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var preferredDRX_LongCycle_r16Present bool
	if preferredDRX_LongCycle_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var preferredDRX_ShortCycle_r16Present bool
	if preferredDRX_ShortCycle_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var preferredDRX_ShortCycleTimer_r16Present bool
	if preferredDRX_ShortCycleTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if preferredDRX_InactivityTimer_r16Present {
		ie.preferredDRX_InactivityTimer_r16 = new(DRX_Preference_r16_preferredDRX_InactivityTimer_r16)
		if err = ie.preferredDRX_InactivityTimer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode preferredDRX_InactivityTimer_r16", err)
		}
	}
	if preferredDRX_LongCycle_r16Present {
		ie.preferredDRX_LongCycle_r16 = new(DRX_Preference_r16_preferredDRX_LongCycle_r16)
		if err = ie.preferredDRX_LongCycle_r16.Decode(r); err != nil {
			return utils.WrapError("Decode preferredDRX_LongCycle_r16", err)
		}
	}
	if preferredDRX_ShortCycle_r16Present {
		ie.preferredDRX_ShortCycle_r16 = new(DRX_Preference_r16_preferredDRX_ShortCycle_r16)
		if err = ie.preferredDRX_ShortCycle_r16.Decode(r); err != nil {
			return utils.WrapError("Decode preferredDRX_ShortCycle_r16", err)
		}
	}
	if preferredDRX_ShortCycleTimer_r16Present {
		var tmp_int_preferredDRX_ShortCycleTimer_r16 int64
		if tmp_int_preferredDRX_ShortCycleTimer_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode preferredDRX_ShortCycleTimer_r16", err)
		}
		ie.preferredDRX_ShortCycleTimer_r16 = &tmp_int_preferredDRX_ShortCycleTimer_r16
	}
	return nil
}
