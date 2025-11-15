package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasurementTimingConfiguration_IEs struct {
	measTiming           *MeasTimingList                           `optional`
	nonCriticalExtension *MeasurementTimingConfiguration_v1550_IEs `optional`
}

func (ie *MeasurementTimingConfiguration_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measTiming != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measTiming != nil {
		if err = ie.measTiming.Encode(w); err != nil {
			return utils.WrapError("Encode measTiming", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MeasurementTimingConfiguration_IEs) Decode(r *uper.UperReader) error {
	var err error
	var measTimingPresent bool
	if measTimingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measTimingPresent {
		ie.measTiming = new(MeasTimingList)
		if err = ie.measTiming.Decode(r); err != nil {
			return utils.WrapError("Decode measTiming", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(MeasurementTimingConfiguration_v1550_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
