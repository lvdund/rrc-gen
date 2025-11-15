package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_resources_ssb struct {
	ssb_ResourceList         []CFRA_SSB_Resource `lb:1,ub:maxRA_SSB_Resources,madatory`
	ra_ssb_OccasionMaskIndex int64               `lb:0,ub:15,madatory`
}

func (ie *CFRA_resources_ssb) Encode(w *uper.UperWriter) error {
	var err error
	tmp_ssb_ResourceList := utils.NewSequence[*CFRA_SSB_Resource]([]*CFRA_SSB_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_SSB_Resources}, false)
	for _, i := range ie.ssb_ResourceList {
		tmp_ssb_ResourceList.Value = append(tmp_ssb_ResourceList.Value, &i)
	}
	if err = tmp_ssb_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_ResourceList", err)
	}
	if err = w.WriteInteger(ie.ra_ssb_OccasionMaskIndex, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger ra_ssb_OccasionMaskIndex", err)
	}
	return nil
}

func (ie *CFRA_resources_ssb) Decode(r *uper.UperReader) error {
	var err error
	tmp_ssb_ResourceList := utils.NewSequence[*CFRA_SSB_Resource]([]*CFRA_SSB_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_SSB_Resources}, false)
	fn_ssb_ResourceList := func() *CFRA_SSB_Resource {
		return new(CFRA_SSB_Resource)
	}
	if err = tmp_ssb_ResourceList.Decode(r, fn_ssb_ResourceList); err != nil {
		return utils.WrapError("Decode ssb_ResourceList", err)
	}
	ie.ssb_ResourceList = []CFRA_SSB_Resource{}
	for _, i := range tmp_ssb_ResourceList.Value {
		ie.ssb_ResourceList = append(ie.ssb_ResourceList, *i)
	}
	var tmp_int_ra_ssb_OccasionMaskIndex int64
	if tmp_int_ra_ssb_OccasionMaskIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger ra_ssb_OccasionMaskIndex", err)
	}
	ie.ra_ssb_OccasionMaskIndex = tmp_int_ra_ssb_OccasionMaskIndex
	return nil
}
