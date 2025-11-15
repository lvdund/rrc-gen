package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedParameters_r16 struct {
	measurementEnhancement_r16  *HighSpeedParameters_r16_measurementEnhancement_r16  `optional`
	demodulationEnhancement_r16 *HighSpeedParameters_r16_demodulationEnhancement_r16 `optional`
}

func (ie *HighSpeedParameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measurementEnhancement_r16 != nil, ie.demodulationEnhancement_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measurementEnhancement_r16 != nil {
		if err = ie.measurementEnhancement_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measurementEnhancement_r16", err)
		}
	}
	if ie.demodulationEnhancement_r16 != nil {
		if err = ie.demodulationEnhancement_r16.Encode(w); err != nil {
			return utils.WrapError("Encode demodulationEnhancement_r16", err)
		}
	}
	return nil
}

func (ie *HighSpeedParameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var measurementEnhancement_r16Present bool
	if measurementEnhancement_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var demodulationEnhancement_r16Present bool
	if demodulationEnhancement_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measurementEnhancement_r16Present {
		ie.measurementEnhancement_r16 = new(HighSpeedParameters_r16_measurementEnhancement_r16)
		if err = ie.measurementEnhancement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measurementEnhancement_r16", err)
		}
	}
	if demodulationEnhancement_r16Present {
		ie.demodulationEnhancement_r16 = new(HighSpeedParameters_r16_demodulationEnhancement_r16)
		if err = ie.demodulationEnhancement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode demodulationEnhancement_r16", err)
		}
	}
	return nil
}
