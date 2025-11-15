package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Sensor_NameList_r16 struct {
	measUncomBarPre_r16 *Sensor_NameList_r16_measUncomBarPre_r16 `optional`
	measUeSpeed         *Sensor_NameList_r16_measUeSpeed         `optional`
	measUeOrientation   *Sensor_NameList_r16_measUeOrientation   `optional`
}

func (ie *Sensor_NameList_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measUncomBarPre_r16 != nil, ie.measUeSpeed != nil, ie.measUeOrientation != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measUncomBarPre_r16 != nil {
		if err = ie.measUncomBarPre_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measUncomBarPre_r16", err)
		}
	}
	if ie.measUeSpeed != nil {
		if err = ie.measUeSpeed.Encode(w); err != nil {
			return utils.WrapError("Encode measUeSpeed", err)
		}
	}
	if ie.measUeOrientation != nil {
		if err = ie.measUeOrientation.Encode(w); err != nil {
			return utils.WrapError("Encode measUeOrientation", err)
		}
	}
	return nil
}

func (ie *Sensor_NameList_r16) Decode(r *uper.UperReader) error {
	var err error
	var measUncomBarPre_r16Present bool
	if measUncomBarPre_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measUeSpeedPresent bool
	if measUeSpeedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measUeOrientationPresent bool
	if measUeOrientationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if measUncomBarPre_r16Present {
		ie.measUncomBarPre_r16 = new(Sensor_NameList_r16_measUncomBarPre_r16)
		if err = ie.measUncomBarPre_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measUncomBarPre_r16", err)
		}
	}
	if measUeSpeedPresent {
		ie.measUeSpeed = new(Sensor_NameList_r16_measUeSpeed)
		if err = ie.measUeSpeed.Decode(r); err != nil {
			return utils.WrapError("Decode measUeSpeed", err)
		}
	}
	if measUeOrientationPresent {
		ie.measUeOrientation = new(Sensor_NameList_r16_measUeOrientation)
		if err = ie.measUeOrientation.Decode(r); err != nil {
			return utils.WrapError("Decode measUeOrientation", err)
		}
	}
	return nil
}
