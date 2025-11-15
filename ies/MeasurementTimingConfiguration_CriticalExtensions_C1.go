package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_nothing uint64 = iota
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_measTimingConf
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare3
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare2
	MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare1
)

type MeasurementTimingConfiguration_CriticalExtensions_C1 struct {
	Choice         uint64
	measTimingConf *MeasurementTimingConfiguration_IEs
	spare3         uper.NULL `madatory`
	spare2         uper.NULL `madatory`
	spare1         uper.NULL `madatory`
}

func (ie *MeasurementTimingConfiguration_CriticalExtensions_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_measTimingConf:
		if err = ie.measTimingConf.Encode(w); err != nil {
			err = utils.WrapError("Encode measTimingConf", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare3:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare3", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare2", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode spare1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasurementTimingConfiguration_CriticalExtensions_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_measTimingConf:
		ie.measTimingConf = new(MeasurementTimingConfiguration_IEs)
		if err = ie.measTimingConf.Decode(r); err != nil {
			return utils.WrapError("Decode measTimingConf", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare3:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	case MeasurementTimingConfiguration_CriticalExtensions_C1_Choice_spare1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
