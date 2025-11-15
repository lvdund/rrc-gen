package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type2_PortSelection struct {
	supportedCSI_RS_ResourceList []SupportedCSI_RS_Resource                                  `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	parameterLx                  int64                                                       `lb:2,ub:4,madatory`
	amplitudeScalingType         CodebookParameters_type2_PortSelection_amplitudeScalingType `madatory`
}

func (ie *CodebookParameters_type2_PortSelection) Encode(w *uper.UperWriter) error {
	var err error
	tmp_supportedCSI_RS_ResourceList := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
	for _, i := range ie.supportedCSI_RS_ResourceList {
		tmp_supportedCSI_RS_ResourceList.Value = append(tmp_supportedCSI_RS_ResourceList.Value, &i)
	}
	if err = tmp_supportedCSI_RS_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode supportedCSI_RS_ResourceList", err)
	}
	if err = w.WriteInteger(ie.parameterLx, &uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger parameterLx", err)
	}
	if err = ie.amplitudeScalingType.Encode(w); err != nil {
		return utils.WrapError("Encode amplitudeScalingType", err)
	}
	return nil
}

func (ie *CodebookParameters_type2_PortSelection) Decode(r *uper.UperReader) error {
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
	var tmp_int_parameterLx int64
	if tmp_int_parameterLx, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger parameterLx", err)
	}
	ie.parameterLx = tmp_int_parameterLx
	if err = ie.amplitudeScalingType.Decode(r); err != nil {
		return utils.WrapError("Decode amplitudeScalingType", err)
	}
	return nil
}
