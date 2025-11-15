package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resources struct {
	maxNumberAperiodicSRS_PerBWP              SRS_Resources_maxNumberAperiodicSRS_PerBWP      `madatory`
	maxNumberAperiodicSRS_PerBWP_PerSlot      int64                                           `lb:1,ub:6,madatory`
	maxNumberPeriodicSRS_PerBWP               SRS_Resources_maxNumberPeriodicSRS_PerBWP       `madatory`
	maxNumberPeriodicSRS_PerBWP_PerSlot       int64                                           `lb:1,ub:6,madatory`
	maxNumberSemiPersistentSRS_PerBWP         SRS_Resources_maxNumberSemiPersistentSRS_PerBWP `madatory`
	maxNumberSemiPersistentSRS_PerBWP_PerSlot int64                                           `lb:1,ub:6,madatory`
	maxNumberSRS_Ports_PerResource            SRS_Resources_maxNumberSRS_Ports_PerResource    `madatory`
}

func (ie *SRS_Resources) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberAperiodicSRS_PerBWP.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberAperiodicSRS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.maxNumberAperiodicSRS_PerBWP_PerSlot, &uper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberAperiodicSRS_PerBWP_PerSlot", err)
	}
	if err = ie.maxNumberPeriodicSRS_PerBWP.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPeriodicSRS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.maxNumberPeriodicSRS_PerBWP_PerSlot, &uper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberPeriodicSRS_PerBWP_PerSlot", err)
	}
	if err = ie.maxNumberSemiPersistentSRS_PerBWP.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSemiPersistentSRS_PerBWP", err)
	}
	if err = w.WriteInteger(ie.maxNumberSemiPersistentSRS_PerBWP_PerSlot, &uper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSemiPersistentSRS_PerBWP_PerSlot", err)
	}
	if err = ie.maxNumberSRS_Ports_PerResource.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberSRS_Ports_PerResource", err)
	}
	return nil
}

func (ie *SRS_Resources) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberAperiodicSRS_PerBWP.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberAperiodicSRS_PerBWP", err)
	}
	var tmp_int_maxNumberAperiodicSRS_PerBWP_PerSlot int64
	if tmp_int_maxNumberAperiodicSRS_PerBWP_PerSlot, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberAperiodicSRS_PerBWP_PerSlot", err)
	}
	ie.maxNumberAperiodicSRS_PerBWP_PerSlot = tmp_int_maxNumberAperiodicSRS_PerBWP_PerSlot
	if err = ie.maxNumberPeriodicSRS_PerBWP.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPeriodicSRS_PerBWP", err)
	}
	var tmp_int_maxNumberPeriodicSRS_PerBWP_PerSlot int64
	if tmp_int_maxNumberPeriodicSRS_PerBWP_PerSlot, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberPeriodicSRS_PerBWP_PerSlot", err)
	}
	ie.maxNumberPeriodicSRS_PerBWP_PerSlot = tmp_int_maxNumberPeriodicSRS_PerBWP_PerSlot
	if err = ie.maxNumberSemiPersistentSRS_PerBWP.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSemiPersistentSRS_PerBWP", err)
	}
	var tmp_int_maxNumberSemiPersistentSRS_PerBWP_PerSlot int64
	if tmp_int_maxNumberSemiPersistentSRS_PerBWP_PerSlot, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 6}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSemiPersistentSRS_PerBWP_PerSlot", err)
	}
	ie.maxNumberSemiPersistentSRS_PerBWP_PerSlot = tmp_int_maxNumberSemiPersistentSRS_PerBWP_PerSlot
	if err = ie.maxNumberSRS_Ports_PerResource.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberSRS_Ports_PerResource", err)
	}
	return nil
}
