package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_MonitoringOccasions_r16 struct {
	period7span3_r16 *PDCCH_MonitoringOccasions_r16_period7span3_r16 `optional`
	period4span3_r16 *PDCCH_MonitoringOccasions_r16_period4span3_r16 `optional`
	period2span2_r16 *PDCCH_MonitoringOccasions_r16_period2span2_r16 `optional`
}

func (ie *PDCCH_MonitoringOccasions_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.period7span3_r16 != nil, ie.period4span3_r16 != nil, ie.period2span2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.period7span3_r16 != nil {
		if err = ie.period7span3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode period7span3_r16", err)
		}
	}
	if ie.period4span3_r16 != nil {
		if err = ie.period4span3_r16.Encode(w); err != nil {
			return utils.WrapError("Encode period4span3_r16", err)
		}
	}
	if ie.period2span2_r16 != nil {
		if err = ie.period2span2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode period2span2_r16", err)
		}
	}
	return nil
}

func (ie *PDCCH_MonitoringOccasions_r16) Decode(r *uper.UperReader) error {
	var err error
	var period7span3_r16Present bool
	if period7span3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var period4span3_r16Present bool
	if period4span3_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var period2span2_r16Present bool
	if period2span2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if period7span3_r16Present {
		ie.period7span3_r16 = new(PDCCH_MonitoringOccasions_r16_period7span3_r16)
		if err = ie.period7span3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode period7span3_r16", err)
		}
	}
	if period4span3_r16Present {
		ie.period4span3_r16 = new(PDCCH_MonitoringOccasions_r16_period4span3_r16)
		if err = ie.period4span3_r16.Decode(r); err != nil {
			return utils.WrapError("Decode period4span3_r16", err)
		}
	}
	if period2span2_r16Present {
		ie.period2span2_r16 = new(PDCCH_MonitoringOccasions_r16_period2span2_r16)
		if err = ie.period2span2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode period2span2_r16", err)
		}
	}
	return nil
}
