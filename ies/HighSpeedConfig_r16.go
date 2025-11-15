package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedConfig_r16 struct {
	highSpeedMeasFlag_r16  *HighSpeedConfig_r16_highSpeedMeasFlag_r16  `optional`
	highSpeedDemodFlag_r16 *HighSpeedConfig_r16_highSpeedDemodFlag_r16 `optional`
}

func (ie *HighSpeedConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.highSpeedMeasFlag_r16 != nil, ie.highSpeedDemodFlag_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.highSpeedMeasFlag_r16 != nil {
		if err = ie.highSpeedMeasFlag_r16.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedMeasFlag_r16", err)
		}
	}
	if ie.highSpeedDemodFlag_r16 != nil {
		if err = ie.highSpeedDemodFlag_r16.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedDemodFlag_r16", err)
		}
	}
	return nil
}

func (ie *HighSpeedConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var highSpeedMeasFlag_r16Present bool
	if highSpeedMeasFlag_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedDemodFlag_r16Present bool
	if highSpeedDemodFlag_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if highSpeedMeasFlag_r16Present {
		ie.highSpeedMeasFlag_r16 = new(HighSpeedConfig_r16_highSpeedMeasFlag_r16)
		if err = ie.highSpeedMeasFlag_r16.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedMeasFlag_r16", err)
		}
	}
	if highSpeedDemodFlag_r16Present {
		ie.highSpeedDemodFlag_r16 = new(HighSpeedConfig_r16_highSpeedDemodFlag_r16)
		if err = ie.highSpeedDemodFlag_r16.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedDemodFlag_r16", err)
		}
	}
	return nil
}
