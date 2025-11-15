package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16 struct {
	type1_SinglePanel_r16   *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_SinglePanel_r16   `optional`
	type1_MultiPanel_r16    *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_MultiPanel_r16    `optional`
	type2_r16               *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_r16               `optional`
	type2_PortSelection_r16 *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_PortSelection_r16 `optional`
}

func (ie *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.type1_SinglePanel_r16 != nil, ie.type1_MultiPanel_r16 != nil, ie.type2_r16 != nil, ie.type2_PortSelection_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.type1_SinglePanel_r16 != nil {
		if err = ie.type1_SinglePanel_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1_SinglePanel_r16", err)
		}
	}
	if ie.type1_MultiPanel_r16 != nil {
		if err = ie.type1_MultiPanel_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type1_MultiPanel_r16", err)
		}
	}
	if ie.type2_r16 != nil {
		if err = ie.type2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type2_r16", err)
		}
	}
	if ie.type2_PortSelection_r16 != nil {
		if err = ie.type2_PortSelection_r16.Encode(w); err != nil {
			return utils.WrapError("Encode type2_PortSelection_r16", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16) Decode(r *uper.UperReader) error {
	var err error
	var type1_SinglePanel_r16Present bool
	if type1_SinglePanel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type1_MultiPanel_r16Present bool
	if type1_MultiPanel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type2_r16Present bool
	if type2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type2_PortSelection_r16Present bool
	if type2_PortSelection_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if type1_SinglePanel_r16Present {
		ie.type1_SinglePanel_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_SinglePanel_r16)
		if err = ie.type1_SinglePanel_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type1_SinglePanel_r16", err)
		}
	}
	if type1_MultiPanel_r16Present {
		ie.type1_MultiPanel_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type1_MultiPanel_r16)
		if err = ie.type1_MultiPanel_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type1_MultiPanel_r16", err)
		}
	}
	if type2_r16Present {
		ie.type2_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_r16)
		if err = ie.type2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type2_r16", err)
		}
	}
	if type2_PortSelection_r16Present {
		ie.type2_PortSelection_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16_type2_PortSelection_r16)
		if err = ie.type2_PortSelection_r16.Decode(r); err != nil {
			return utils.WrapError("Decode type2_PortSelection_r16", err)
		}
	}
	return nil
}
