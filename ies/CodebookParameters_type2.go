package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type2 struct {
	supportedCSI_RS_ResourceList []SupportedCSI_RS_Resource                           `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	parameterLx                  int64                                                `lb:2,ub:4,madatory`
	amplitudeScalingType         CodebookParameters_type2_amplitudeScalingType        `madatory`
	amplitudeSubsetRestriction   *CodebookParameters_type2_amplitudeSubsetRestriction `optional`
}

func (ie *CodebookParameters_type2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.amplitudeSubsetRestriction != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
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
	if ie.amplitudeSubsetRestriction != nil {
		if err = ie.amplitudeSubsetRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode amplitudeSubsetRestriction", err)
		}
	}
	return nil
}

func (ie *CodebookParameters_type2) Decode(r *uper.UperReader) error {
	var err error
	var amplitudeSubsetRestrictionPresent bool
	if amplitudeSubsetRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
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
	if amplitudeSubsetRestrictionPresent {
		ie.amplitudeSubsetRestriction = new(CodebookParameters_type2_amplitudeSubsetRestriction)
		if err = ie.amplitudeSubsetRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode amplitudeSubsetRestriction", err)
		}
	}
	return nil
}
