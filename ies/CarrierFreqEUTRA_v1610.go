package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CarrierFreqEUTRA_v1610 struct {
	highSpeedEUTRACarrier_r16 *CarrierFreqEUTRA_v1610_highSpeedEUTRACarrier_r16 `optional`
}

func (ie *CarrierFreqEUTRA_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.highSpeedEUTRACarrier_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.highSpeedEUTRACarrier_r16 != nil {
		if err = ie.highSpeedEUTRACarrier_r16.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedEUTRACarrier_r16", err)
		}
	}
	return nil
}

func (ie *CarrierFreqEUTRA_v1610) Decode(r *uper.UperReader) error {
	var err error
	var highSpeedEUTRACarrier_r16Present bool
	if highSpeedEUTRACarrier_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if highSpeedEUTRACarrier_r16Present {
		ie.highSpeedEUTRACarrier_r16 = new(CarrierFreqEUTRA_v1610_highSpeedEUTRACarrier_r16)
		if err = ie.highSpeedEUTRACarrier_r16.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedEUTRACarrier_r16", err)
		}
	}
	return nil
}
