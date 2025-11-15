package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UL_GapFR2_Preference_r17 struct {
	ul_GapFR2_PatternPreference_r17 *int64 `lb:0,ub:3,optional`
}

func (ie *UL_GapFR2_Preference_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_GapFR2_PatternPreference_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_GapFR2_PatternPreference_r17 != nil {
		if err = w.WriteInteger(*ie.ul_GapFR2_PatternPreference_r17, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode ul_GapFR2_PatternPreference_r17", err)
		}
	}
	return nil
}

func (ie *UL_GapFR2_Preference_r17) Decode(r *uper.UperReader) error {
	var err error
	var ul_GapFR2_PatternPreference_r17Present bool
	if ul_GapFR2_PatternPreference_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_GapFR2_PatternPreference_r17Present {
		var tmp_int_ul_GapFR2_PatternPreference_r17 int64
		if tmp_int_ul_GapFR2_PatternPreference_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode ul_GapFR2_PatternPreference_r17", err)
		}
		ie.ul_GapFR2_PatternPreference_r17 = &tmp_int_ul_GapFR2_PatternPreference_r17
	}
	return nil
}
