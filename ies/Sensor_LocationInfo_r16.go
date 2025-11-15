package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Sensor_LocationInfo_r16 struct {
	sensor_MeasurementInformation_r16 *[]byte `optional`
	sensor_MotionInformation_r16      *[]byte `optional`
}

func (ie *Sensor_LocationInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sensor_MeasurementInformation_r16 != nil, ie.sensor_MotionInformation_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sensor_MeasurementInformation_r16 != nil {
		if err = w.WriteOctetString(*ie.sensor_MeasurementInformation_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sensor_MeasurementInformation_r16", err)
		}
	}
	if ie.sensor_MotionInformation_r16 != nil {
		if err = w.WriteOctetString(*ie.sensor_MotionInformation_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sensor_MotionInformation_r16", err)
		}
	}
	return nil
}

func (ie *Sensor_LocationInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var sensor_MeasurementInformation_r16Present bool
	if sensor_MeasurementInformation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sensor_MotionInformation_r16Present bool
	if sensor_MotionInformation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sensor_MeasurementInformation_r16Present {
		var tmp_os_sensor_MeasurementInformation_r16 []byte
		if tmp_os_sensor_MeasurementInformation_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sensor_MeasurementInformation_r16", err)
		}
		ie.sensor_MeasurementInformation_r16 = &tmp_os_sensor_MeasurementInformation_r16
	}
	if sensor_MotionInformation_r16Present {
		var tmp_os_sensor_MotionInformation_r16 []byte
		if tmp_os_sensor_MotionInformation_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sensor_MotionInformation_r16", err)
		}
		ie.sensor_MotionInformation_r16 = &tmp_os_sensor_MotionInformation_r16
	}
	return nil
}
