package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Pdsch_ProcessingType_r16 struct {
	scs_15kHz_r16 *PDCCH_MonitoringOccasions_r16 `optional`
	scs_30kHz_r16 *PDCCH_MonitoringOccasions_r16 `optional`
}

func (ie *Pdsch_ProcessingType_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_15kHz_r16 != nil, ie.scs_30kHz_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_15kHz_r16 != nil {
		if err = ie.scs_15kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scs_15kHz_r16", err)
		}
	}
	if ie.scs_30kHz_r16 != nil {
		if err = ie.scs_30kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scs_30kHz_r16", err)
		}
	}
	return nil
}

func (ie *Pdsch_ProcessingType_r16) Decode(r *uper.UperReader) error {
	var err error
	var scs_15kHz_r16Present bool
	if scs_15kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_30kHz_r16Present bool
	if scs_30kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_15kHz_r16Present {
		ie.scs_15kHz_r16 = new(PDCCH_MonitoringOccasions_r16)
		if err = ie.scs_15kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scs_15kHz_r16", err)
		}
	}
	if scs_30kHz_r16Present {
		ie.scs_30kHz_r16 = new(PDCCH_MonitoringOccasions_r16)
		if err = ie.scs_30kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scs_30kHz_r16", err)
		}
	}
	return nil
}
