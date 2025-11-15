package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_IM_ResourceSet struct {
	csi_IM_ResourceSetId CSI_IM_ResourceSetId `madatory`
	csi_IM_Resources     []CSI_IM_ResourceId  `lb:1,ub:maxNrofCSI_IM_ResourcesPerSet,madatory`
}

func (ie *CSI_IM_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.csi_IM_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode csi_IM_ResourceSetId", err)
	}
	tmp_csi_IM_Resources := utils.NewSequence[*CSI_IM_ResourceId]([]*CSI_IM_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourcesPerSet}, false)
	for _, i := range ie.csi_IM_Resources {
		tmp_csi_IM_Resources.Value = append(tmp_csi_IM_Resources.Value, &i)
	}
	if err = tmp_csi_IM_Resources.Encode(w); err != nil {
		return utils.WrapError("Encode csi_IM_Resources", err)
	}
	return nil
}

func (ie *CSI_IM_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.csi_IM_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode csi_IM_ResourceSetId", err)
	}
	tmp_csi_IM_Resources := utils.NewSequence[*CSI_IM_ResourceId]([]*CSI_IM_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourcesPerSet}, false)
	fn_csi_IM_Resources := func() *CSI_IM_ResourceId {
		return new(CSI_IM_ResourceId)
	}
	if err = tmp_csi_IM_Resources.Decode(r, fn_csi_IM_Resources); err != nil {
		return utils.WrapError("Decode csi_IM_Resources", err)
	}
	ie.csi_IM_Resources = []CSI_IM_ResourceId{}
	for _, i := range tmp_csi_IM_Resources.Value {
		ie.csi_IM_Resources = append(ie.csi_IM_Resources, *i)
	}
	return nil
}
