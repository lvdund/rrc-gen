package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_resources_csirs struct {
	csirs_ResourceList   []CFRA_CSIRS_Resource `lb:1,ub:maxRA_CSIRS_Resources,madatory`
	rsrp_ThresholdCSI_RS RSRP_Range            `madatory`
}

func (ie *CFRA_resources_csirs) Encode(w *uper.UperWriter) error {
	var err error
	tmp_csirs_ResourceList := utils.NewSequence[*CFRA_CSIRS_Resource]([]*CFRA_CSIRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_CSIRS_Resources}, false)
	for _, i := range ie.csirs_ResourceList {
		tmp_csirs_ResourceList.Value = append(tmp_csirs_ResourceList.Value, &i)
	}
	if err = tmp_csirs_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode csirs_ResourceList", err)
	}
	if err = ie.rsrp_ThresholdCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode rsrp_ThresholdCSI_RS", err)
	}
	return nil
}

func (ie *CFRA_resources_csirs) Decode(r *uper.UperReader) error {
	var err error
	tmp_csirs_ResourceList := utils.NewSequence[*CFRA_CSIRS_Resource]([]*CFRA_CSIRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_CSIRS_Resources}, false)
	fn_csirs_ResourceList := func() *CFRA_CSIRS_Resource {
		return new(CFRA_CSIRS_Resource)
	}
	if err = tmp_csirs_ResourceList.Decode(r, fn_csirs_ResourceList); err != nil {
		return utils.WrapError("Decode csirs_ResourceList", err)
	}
	ie.csirs_ResourceList = []CFRA_CSIRS_Resource{}
	for _, i := range tmp_csirs_ResourceList.Value {
		ie.csirs_ResourceList = append(ie.csirs_ResourceList, *i)
	}
	if err = ie.rsrp_ThresholdCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode rsrp_ThresholdCSI_RS", err)
	}
	return nil
}
