package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceGroup_r16 struct {
	pucch_ResourceGroupId_r16 PUCCH_ResourceGroupId_r16 `madatory`
	resourcePerGroupList_r16  []PUCCH_ResourceId        `lb:1,ub:maxNrofPUCCH_ResourcesPerGroup_r16,madatory`
}

func (ie *PUCCH_ResourceGroup_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.pucch_ResourceGroupId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_ResourceGroupId_r16", err)
	}
	tmp_resourcePerGroupList_r16 := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerGroup_r16}, false)
	for _, i := range ie.resourcePerGroupList_r16 {
		tmp_resourcePerGroupList_r16.Value = append(tmp_resourcePerGroupList_r16.Value, &i)
	}
	if err = tmp_resourcePerGroupList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resourcePerGroupList_r16", err)
	}
	return nil
}

func (ie *PUCCH_ResourceGroup_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.pucch_ResourceGroupId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_ResourceGroupId_r16", err)
	}
	tmp_resourcePerGroupList_r16 := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerGroup_r16}, false)
	fn_resourcePerGroupList_r16 := func() *PUCCH_ResourceId {
		return new(PUCCH_ResourceId)
	}
	if err = tmp_resourcePerGroupList_r16.Decode(r, fn_resourcePerGroupList_r16); err != nil {
		return utils.WrapError("Decode resourcePerGroupList_r16", err)
	}
	ie.resourcePerGroupList_r16 = []PUCCH_ResourceId{}
	for _, i := range tmp_resourcePerGroupList_r16.Value {
		ie.resourcePerGroupList_r16 = append(ie.resourcePerGroupList_r16, *i)
	}
	return nil
}
