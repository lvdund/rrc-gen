package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type1_multiPanel struct {
	supportedCSI_RS_ResourceList   []SupportedCSI_RS_Resource                     `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	modes                          CodebookParameters_type1_multiPanel_modes      `madatory`
	nrofPanels                     CodebookParameters_type1_multiPanel_nrofPanels `madatory`
	maxNumberCSI_RS_PerResourceSet int64                                          `lb:1,ub:8,madatory`
}

func (ie *CodebookParameters_type1_multiPanel) Encode(w *uper.UperWriter) error {
	var err error
	tmp_supportedCSI_RS_ResourceList := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
	for _, i := range ie.supportedCSI_RS_ResourceList {
		tmp_supportedCSI_RS_ResourceList.Value = append(tmp_supportedCSI_RS_ResourceList.Value, &i)
	}
	if err = tmp_supportedCSI_RS_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode supportedCSI_RS_ResourceList", err)
	}
	if err = ie.modes.Encode(w); err != nil {
		return utils.WrapError("Encode modes", err)
	}
	if err = ie.nrofPanels.Encode(w); err != nil {
		return utils.WrapError("Encode nrofPanels", err)
	}
	if err = w.WriteInteger(ie.maxNumberCSI_RS_PerResourceSet, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberCSI_RS_PerResourceSet", err)
	}
	return nil
}

func (ie *CodebookParameters_type1_multiPanel) Decode(r *uper.UperReader) error {
	var err error
	tmp_supportedCSI_RS_ResourceList := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
	fn_supportedCSI_RS_ResourceList := func() *SupportedCSI_RS_Resource {
		return new(SupportedCSI_RS_Resource)
	}
	if err = tmp_supportedCSI_RS_ResourceList.Decode(r, fn_supportedCSI_RS_ResourceList); err != nil {
		return utils.WrapError("Decode supportedCSI_RS_ResourceList", err)
	}
	ie.supportedCSI_RS_ResourceList = []SupportedCSI_RS_Resource{}
	for _, i := range tmp_supportedCSI_RS_ResourceList.Value {
		ie.supportedCSI_RS_ResourceList = append(ie.supportedCSI_RS_ResourceList, *i)
	}
	if err = ie.modes.Decode(r); err != nil {
		return utils.WrapError("Decode modes", err)
	}
	if err = ie.nrofPanels.Decode(r); err != nil {
		return utils.WrapError("Decode nrofPanels", err)
	}
	var tmp_int_maxNumberCSI_RS_PerResourceSet int64
	if tmp_int_maxNumberCSI_RS_PerResourceSet, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberCSI_RS_PerResourceSet", err)
	}
	ie.maxNumberCSI_RS_PerResourceSet = tmp_int_maxNumberCSI_RS_PerResourceSet
	return nil
}
