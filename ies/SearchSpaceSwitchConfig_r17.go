package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceSwitchConfig_r17 struct {
	searchSpaceSwitchTimer_r17 *SCS_SpecificDuration_r17 `optional`
	searchSpaceSwitchDelay_r17 *int64                    `lb:10,ub:52,optional`
}

func (ie *SearchSpaceSwitchConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.searchSpaceSwitchTimer_r17 != nil, ie.searchSpaceSwitchDelay_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.searchSpaceSwitchTimer_r17 != nil {
		if err = ie.searchSpaceSwitchTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceSwitchTimer_r17", err)
		}
	}
	if ie.searchSpaceSwitchDelay_r17 != nil {
		if err = w.WriteInteger(*ie.searchSpaceSwitchDelay_r17, &uper.Constraint{Lb: 10, Ub: 52}, false); err != nil {
			return utils.WrapError("Encode searchSpaceSwitchDelay_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceSwitchConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var searchSpaceSwitchTimer_r17Present bool
	if searchSpaceSwitchTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceSwitchDelay_r17Present bool
	if searchSpaceSwitchDelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if searchSpaceSwitchTimer_r17Present {
		ie.searchSpaceSwitchTimer_r17 = new(SCS_SpecificDuration_r17)
		if err = ie.searchSpaceSwitchTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceSwitchTimer_r17", err)
		}
	}
	if searchSpaceSwitchDelay_r17Present {
		var tmp_int_searchSpaceSwitchDelay_r17 int64
		if tmp_int_searchSpaceSwitchDelay_r17, err = r.ReadInteger(&uper.Constraint{Lb: 10, Ub: 52}, false); err != nil {
			return utils.WrapError("Decode searchSpaceSwitchDelay_r17", err)
		}
		ie.searchSpaceSwitchDelay_r17 = &tmp_int_searchSpaceSwitchDelay_r17
	}
	return nil
}
