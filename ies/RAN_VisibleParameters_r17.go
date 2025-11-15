package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RAN_VisibleParameters_r17 struct {
	ran_VisiblePeriodicity_r17            *RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17 `optional`
	numberOfBufferLevelEntries_r17        *int64                                                `lb:1,ub:8,optional`
	reportPlayoutDelayForMediaStartup_r17 *bool                                                 `optional`
}

func (ie *RAN_VisibleParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ran_VisiblePeriodicity_r17 != nil, ie.numberOfBufferLevelEntries_r17 != nil, ie.reportPlayoutDelayForMediaStartup_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ran_VisiblePeriodicity_r17 != nil {
		if err = ie.ran_VisiblePeriodicity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ran_VisiblePeriodicity_r17", err)
		}
	}
	if ie.numberOfBufferLevelEntries_r17 != nil {
		if err = w.WriteInteger(*ie.numberOfBufferLevelEntries_r17, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode numberOfBufferLevelEntries_r17", err)
		}
	}
	if ie.reportPlayoutDelayForMediaStartup_r17 != nil {
		if err = w.WriteBoolean(*ie.reportPlayoutDelayForMediaStartup_r17); err != nil {
			return utils.WrapError("Encode reportPlayoutDelayForMediaStartup_r17", err)
		}
	}
	return nil
}

func (ie *RAN_VisibleParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var ran_VisiblePeriodicity_r17Present bool
	if ran_VisiblePeriodicity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var numberOfBufferLevelEntries_r17Present bool
	if numberOfBufferLevelEntries_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reportPlayoutDelayForMediaStartup_r17Present bool
	if reportPlayoutDelayForMediaStartup_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ran_VisiblePeriodicity_r17Present {
		ie.ran_VisiblePeriodicity_r17 = new(RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17)
		if err = ie.ran_VisiblePeriodicity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ran_VisiblePeriodicity_r17", err)
		}
	}
	if numberOfBufferLevelEntries_r17Present {
		var tmp_int_numberOfBufferLevelEntries_r17 int64
		if tmp_int_numberOfBufferLevelEntries_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode numberOfBufferLevelEntries_r17", err)
		}
		ie.numberOfBufferLevelEntries_r17 = &tmp_int_numberOfBufferLevelEntries_r17
	}
	if reportPlayoutDelayForMediaStartup_r17Present {
		var tmp_bool_reportPlayoutDelayForMediaStartup_r17 bool
		if tmp_bool_reportPlayoutDelayForMediaStartup_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode reportPlayoutDelayForMediaStartup_r17", err)
		}
		ie.reportPlayoutDelayForMediaStartup_r17 = &tmp_bool_reportPlayoutDelayForMediaStartup_r17
	}
	return nil
}
