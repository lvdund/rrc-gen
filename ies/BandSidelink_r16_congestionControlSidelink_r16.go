package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelink_r16_congestionControlSidelink_r16 struct {
	cbr_ReportSidelink_r16       *BandSidelink_r16_congestionControlSidelink_r16_cbr_ReportSidelink_r16      `optional`
	cbr_CR_TimeLimitSidelink_r16 BandSidelink_r16_congestionControlSidelink_r16_cbr_CR_TimeLimitSidelink_r16 `madatory`
}

func (ie *BandSidelink_r16_congestionControlSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cbr_ReportSidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cbr_ReportSidelink_r16 != nil {
		if err = ie.cbr_ReportSidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cbr_ReportSidelink_r16", err)
		}
	}
	if err = ie.cbr_CR_TimeLimitSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cbr_CR_TimeLimitSidelink_r16", err)
	}
	return nil
}

func (ie *BandSidelink_r16_congestionControlSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var cbr_ReportSidelink_r16Present bool
	if cbr_ReportSidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if cbr_ReportSidelink_r16Present {
		ie.cbr_ReportSidelink_r16 = new(BandSidelink_r16_congestionControlSidelink_r16_cbr_ReportSidelink_r16)
		if err = ie.cbr_ReportSidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cbr_ReportSidelink_r16", err)
		}
	}
	if err = ie.cbr_CR_TimeLimitSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cbr_CR_TimeLimitSidelink_r16", err)
	}
	return nil
}
