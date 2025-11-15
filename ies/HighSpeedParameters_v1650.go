package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	HighSpeedParameters_v1650_Choice_nothing uint64 = iota
	HighSpeedParameters_v1650_Choice_intraNR_MeasurementEnhancement_r16
	HighSpeedParameters_v1650_Choice_interRAT_MeasurementEnhancement_r16
)

type HighSpeedParameters_v1650 struct {
	Choice                              uint64
	intraNR_MeasurementEnhancement_r16  *HighSpeedParameters_v1650_intraNR_MeasurementEnhancement_r16
	interRAT_MeasurementEnhancement_r16 *HighSpeedParameters_v1650_interRAT_MeasurementEnhancement_r16
}

func (ie *HighSpeedParameters_v1650) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case HighSpeedParameters_v1650_Choice_intraNR_MeasurementEnhancement_r16:
		if err = ie.intraNR_MeasurementEnhancement_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode intraNR_MeasurementEnhancement_r16", err)
		}
	case HighSpeedParameters_v1650_Choice_interRAT_MeasurementEnhancement_r16:
		if err = ie.interRAT_MeasurementEnhancement_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode interRAT_MeasurementEnhancement_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *HighSpeedParameters_v1650) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case HighSpeedParameters_v1650_Choice_intraNR_MeasurementEnhancement_r16:
		ie.intraNR_MeasurementEnhancement_r16 = new(HighSpeedParameters_v1650_intraNR_MeasurementEnhancement_r16)
		if err = ie.intraNR_MeasurementEnhancement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraNR_MeasurementEnhancement_r16", err)
		}
	case HighSpeedParameters_v1650_Choice_interRAT_MeasurementEnhancement_r16:
		ie.interRAT_MeasurementEnhancement_r16 = new(HighSpeedParameters_v1650_interRAT_MeasurementEnhancement_r16)
		if err = ie.interRAT_MeasurementEnhancement_r16.Decode(r); err != nil {
			return utils.WrapError("Decode interRAT_MeasurementEnhancement_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
