package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_oneEighth uper.Enumerated = 0
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_oneFourth uper.Enumerated = 1
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_oneHalf   uper.Enumerated = 2
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_one       uper.Enumerated = 3
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_two       uper.Enumerated = 4
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_four      uper.Enumerated = 5
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_eight     uper.Enumerated = 6
	SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion_Enum_sixteen   uper.Enumerated = 7
)

type SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion", err)
	}
	return nil
}

func (ie *SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SI_RequestConfig_rach_OccasionsSI_ssb_perRACH_Occasion", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
