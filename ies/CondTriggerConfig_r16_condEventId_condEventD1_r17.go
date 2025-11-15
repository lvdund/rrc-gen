package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventD1_r17 struct {
	distanceThreshFromReference1_r17 int64                  `lb:0,ub:65525,madatory`
	distanceThreshFromReference2_r17 int64                  `lb:0,ub:65525,madatory`
	referenceLocation1_r17           ReferenceLocation_r17  `madatory`
	referenceLocation2_r17           ReferenceLocation_r17  `madatory`
	hysteresisLocation_r17           HysteresisLocation_r17 `madatory`
	timeToTrigger_r17                TimeToTrigger          `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventD1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.distanceThreshFromReference1_r17, &uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("WriteInteger distanceThreshFromReference1_r17", err)
	}
	if err = w.WriteInteger(ie.distanceThreshFromReference2_r17, &uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("WriteInteger distanceThreshFromReference2_r17", err)
	}
	if err = ie.referenceLocation1_r17.Encode(w); err != nil {
		return utils.WrapError("Encode referenceLocation1_r17", err)
	}
	if err = ie.referenceLocation2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode referenceLocation2_r17", err)
	}
	if err = ie.hysteresisLocation_r17.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresisLocation_r17", err)
	}
	if err = ie.timeToTrigger_r17.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger_r17", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventD1_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_distanceThreshFromReference1_r17 int64
	if tmp_int_distanceThreshFromReference1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("ReadInteger distanceThreshFromReference1_r17", err)
	}
	ie.distanceThreshFromReference1_r17 = tmp_int_distanceThreshFromReference1_r17
	var tmp_int_distanceThreshFromReference2_r17 int64
	if tmp_int_distanceThreshFromReference2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65525}, false); err != nil {
		return utils.WrapError("ReadInteger distanceThreshFromReference2_r17", err)
	}
	ie.distanceThreshFromReference2_r17 = tmp_int_distanceThreshFromReference2_r17
	if err = ie.referenceLocation1_r17.Decode(r); err != nil {
		return utils.WrapError("Decode referenceLocation1_r17", err)
	}
	if err = ie.referenceLocation2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode referenceLocation2_r17", err)
	}
	if err = ie.hysteresisLocation_r17.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresisLocation_r17", err)
	}
	if err = ie.timeToTrigger_r17.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger_r17", err)
	}
	return nil
}
