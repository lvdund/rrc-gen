package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HandoverPreparationInformation_IEs struct {
	ue_CapabilityRAT_List UE_CapabilityRAT_ContainerList `madatory`
	sourceConfig          *AS_Config                     `optional`
	rrm_Config            *RRM_Config                    `optional`
	as_Context            *AS_Context                    `optional`
	nonCriticalExtension  interface{}                    `optional`
}

func (ie *HandoverPreparationInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sourceConfig != nil, ie.rrm_Config != nil, ie.as_Context != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sourceConfig != nil {
		if err = ie.sourceConfig.Encode(w); err != nil {
			return utils.WrapError("Encode sourceConfig", err)
		}
	}
	if ie.rrm_Config != nil {
		if err = ie.rrm_Config.Encode(w); err != nil {
			return utils.WrapError("Encode rrm_Config", err)
		}
	}
	if ie.as_Context != nil {
		if err = ie.as_Context.Encode(w); err != nil {
			return utils.WrapError("Encode as_Context", err)
		}
	}
	return nil
}

func (ie *HandoverPreparationInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var sourceConfigPresent bool
	if sourceConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rrm_ConfigPresent bool
	if rrm_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var as_ContextPresent bool
	if as_ContextPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sourceConfigPresent {
		ie.sourceConfig = new(AS_Config)
		if err = ie.sourceConfig.Decode(r); err != nil {
			return utils.WrapError("Decode sourceConfig", err)
		}
	}
	if rrm_ConfigPresent {
		ie.rrm_Config = new(RRM_Config)
		if err = ie.rrm_Config.Decode(r); err != nil {
			return utils.WrapError("Decode rrm_Config", err)
		}
	}
	if as_ContextPresent {
		ie.as_Context = new(AS_Context)
		if err = ie.as_Context.Decode(r); err != nil {
			return utils.WrapError("Decode as_Context", err)
		}
	}
	return nil
}
