package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ZP_CSI_RS_ResourceSet struct {
	zp_CSI_RS_ResourceSetId  ZP_CSI_RS_ResourceSetId `madatory`
	zp_CSI_RS_ResourceIdList []ZP_CSI_RS_ResourceId  `lb:1,ub:maxNrofZP_CSI_RS_ResourcesPerSet,madatory`
}

func (ie *ZP_CSI_RS_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.zp_CSI_RS_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode zp_CSI_RS_ResourceSetId", err)
	}
	tmp_zp_CSI_RS_ResourceIdList := utils.NewSequence[*ZP_CSI_RS_ResourceId]([]*ZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourcesPerSet}, false)
	for _, i := range ie.zp_CSI_RS_ResourceIdList {
		tmp_zp_CSI_RS_ResourceIdList.Value = append(tmp_zp_CSI_RS_ResourceIdList.Value, &i)
	}
	if err = tmp_zp_CSI_RS_ResourceIdList.Encode(w); err != nil {
		return utils.WrapError("Encode zp_CSI_RS_ResourceIdList", err)
	}
	return nil
}

func (ie *ZP_CSI_RS_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.zp_CSI_RS_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode zp_CSI_RS_ResourceSetId", err)
	}
	tmp_zp_CSI_RS_ResourceIdList := utils.NewSequence[*ZP_CSI_RS_ResourceId]([]*ZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourcesPerSet}, false)
	fn_zp_CSI_RS_ResourceIdList := func() *ZP_CSI_RS_ResourceId {
		return new(ZP_CSI_RS_ResourceId)
	}
	if err = tmp_zp_CSI_RS_ResourceIdList.Decode(r, fn_zp_CSI_RS_ResourceIdList); err != nil {
		return utils.WrapError("Decode zp_CSI_RS_ResourceIdList", err)
	}
	ie.zp_CSI_RS_ResourceIdList = []ZP_CSI_RS_ResourceId{}
	for _, i := range tmp_zp_CSI_RS_ResourceIdList.Value {
		ie.zp_CSI_RS_ResourceIdList = append(ie.zp_CSI_RS_ResourceIdList, *i)
	}
	return nil
}
