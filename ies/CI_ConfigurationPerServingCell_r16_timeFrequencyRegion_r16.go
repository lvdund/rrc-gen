package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16 struct {
	timeDurationForCI_r16    *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeDurationForCI_r16   `optional`
	timeGranularityForCI_r16 CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeGranularityForCI_r16 `madatory`
	frequencyRegionForCI_r16 int64                                                                               `lb:0,ub:37949,madatory`
	deltaOffset_r16          int64                                                                               `lb:0,ub:2,madatory`
}

func (ie *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.timeDurationForCI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.timeDurationForCI_r16 != nil {
		if err = ie.timeDurationForCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode timeDurationForCI_r16", err)
		}
	}
	if err = ie.timeGranularityForCI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode timeGranularityForCI_r16", err)
	}
	if err = w.WriteInteger(ie.frequencyRegionForCI_r16, &uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("WriteInteger frequencyRegionForCI_r16", err)
	}
	if err = w.WriteInteger(ie.deltaOffset_r16, &uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger deltaOffset_r16", err)
	}
	return nil
}

func (ie *CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16) Decode(r *uper.UperReader) error {
	var err error
	var timeDurationForCI_r16Present bool
	if timeDurationForCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if timeDurationForCI_r16Present {
		ie.timeDurationForCI_r16 = new(CI_ConfigurationPerServingCell_r16_timeFrequencyRegion_r16_timeDurationForCI_r16)
		if err = ie.timeDurationForCI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode timeDurationForCI_r16", err)
		}
	}
	if err = ie.timeGranularityForCI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode timeGranularityForCI_r16", err)
	}
	var tmp_int_frequencyRegionForCI_r16 int64
	if tmp_int_frequencyRegionForCI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
		return utils.WrapError("ReadInteger frequencyRegionForCI_r16", err)
	}
	ie.frequencyRegionForCI_r16 = tmp_int_frequencyRegionForCI_r16
	var tmp_int_deltaOffset_r16 int64
	if tmp_int_deltaOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger deltaOffset_r16", err)
	}
	ie.deltaOffset_r16 = tmp_int_deltaOffset_r16
	return nil
}
