package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Orbital_r17 struct {
	semiMajorAxis_r17 int64 `lb:0,ub:8589934591,madatory`
	eccentricity_r17  int64 `lb:0,ub:1048575,madatory`
	periapsis_r17     int64 `lb:0,ub:268435455,madatory`
	longitude_r17     int64 `lb:0,ub:268435455,madatory`
	inclination_r17   int64 `lb:-67108864,ub:67108863,madatory`
	meanAnomaly_r17   int64 `lb:0,ub:268435455,madatory`
}

func (ie *Orbital_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.semiMajorAxis_r17, &uper.Constraint{Lb: 0, Ub: 8589934591}, false); err != nil {
		return utils.WrapError("WriteInteger semiMajorAxis_r17", err)
	}
	if err = w.WriteInteger(ie.eccentricity_r17, &uper.Constraint{Lb: 0, Ub: 1048575}, false); err != nil {
		return utils.WrapError("WriteInteger eccentricity_r17", err)
	}
	if err = w.WriteInteger(ie.periapsis_r17, &uper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("WriteInteger periapsis_r17", err)
	}
	if err = w.WriteInteger(ie.longitude_r17, &uper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("WriteInteger longitude_r17", err)
	}
	if err = w.WriteInteger(ie.inclination_r17, &uper.Constraint{Lb: -67108864, Ub: 67108863}, false); err != nil {
		return utils.WrapError("WriteInteger inclination_r17", err)
	}
	if err = w.WriteInteger(ie.meanAnomaly_r17, &uper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("WriteInteger meanAnomaly_r17", err)
	}
	return nil
}

func (ie *Orbital_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_semiMajorAxis_r17 int64
	if tmp_int_semiMajorAxis_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 8589934591}, false); err != nil {
		return utils.WrapError("ReadInteger semiMajorAxis_r17", err)
	}
	ie.semiMajorAxis_r17 = tmp_int_semiMajorAxis_r17
	var tmp_int_eccentricity_r17 int64
	if tmp_int_eccentricity_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1048575}, false); err != nil {
		return utils.WrapError("ReadInteger eccentricity_r17", err)
	}
	ie.eccentricity_r17 = tmp_int_eccentricity_r17
	var tmp_int_periapsis_r17 int64
	if tmp_int_periapsis_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("ReadInteger periapsis_r17", err)
	}
	ie.periapsis_r17 = tmp_int_periapsis_r17
	var tmp_int_longitude_r17 int64
	if tmp_int_longitude_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("ReadInteger longitude_r17", err)
	}
	ie.longitude_r17 = tmp_int_longitude_r17
	var tmp_int_inclination_r17 int64
	if tmp_int_inclination_r17, err = r.ReadInteger(&uper.Constraint{Lb: -67108864, Ub: 67108863}, false); err != nil {
		return utils.WrapError("ReadInteger inclination_r17", err)
	}
	ie.inclination_r17 = tmp_int_inclination_r17
	var tmp_int_meanAnomaly_r17 int64
	if tmp_int_meanAnomaly_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 268435455}, false); err != nil {
		return utils.WrapError("ReadInteger meanAnomaly_r17", err)
	}
	ie.meanAnomaly_r17 = tmp_int_meanAnomaly_r17
	return nil
}
