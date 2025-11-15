package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_RequestConfig struct {
	rach_OccasionsSI    *SI_RequestConfig_rach_OccasionsSI `optional`
	si_RequestPeriod    *SI_RequestConfig_si_RequestPeriod `optional`
	si_RequestResources []SI_RequestResources              `lb:1,ub:maxSI_Message,madatory`
}

func (ie *SI_RequestConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rach_OccasionsSI != nil, ie.si_RequestPeriod != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rach_OccasionsSI != nil {
		if err = ie.rach_OccasionsSI.Encode(w); err != nil {
			return utils.WrapError("Encode rach_OccasionsSI", err)
		}
	}
	if ie.si_RequestPeriod != nil {
		if err = ie.si_RequestPeriod.Encode(w); err != nil {
			return utils.WrapError("Encode si_RequestPeriod", err)
		}
	}
	tmp_si_RequestResources := utils.NewSequence[*SI_RequestResources]([]*SI_RequestResources{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	for _, i := range ie.si_RequestResources {
		tmp_si_RequestResources.Value = append(tmp_si_RequestResources.Value, &i)
	}
	if err = tmp_si_RequestResources.Encode(w); err != nil {
		return utils.WrapError("Encode si_RequestResources", err)
	}
	return nil
}

func (ie *SI_RequestConfig) Decode(r *uper.UperReader) error {
	var err error
	var rach_OccasionsSIPresent bool
	if rach_OccasionsSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var si_RequestPeriodPresent bool
	if si_RequestPeriodPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if rach_OccasionsSIPresent {
		ie.rach_OccasionsSI = new(SI_RequestConfig_rach_OccasionsSI)
		if err = ie.rach_OccasionsSI.Decode(r); err != nil {
			return utils.WrapError("Decode rach_OccasionsSI", err)
		}
	}
	if si_RequestPeriodPresent {
		ie.si_RequestPeriod = new(SI_RequestConfig_si_RequestPeriod)
		if err = ie.si_RequestPeriod.Decode(r); err != nil {
			return utils.WrapError("Decode si_RequestPeriod", err)
		}
	}
	tmp_si_RequestResources := utils.NewSequence[*SI_RequestResources]([]*SI_RequestResources{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	fn_si_RequestResources := func() *SI_RequestResources {
		return new(SI_RequestResources)
	}
	if err = tmp_si_RequestResources.Decode(r, fn_si_RequestResources); err != nil {
		return utils.WrapError("Decode si_RequestResources", err)
	}
	ie.si_RequestResources = []SI_RequestResources{}
	for _, i := range tmp_si_RequestResources.Value {
		ie.si_RequestResources = append(ie.si_RequestResources, *i)
	}
	return nil
}
