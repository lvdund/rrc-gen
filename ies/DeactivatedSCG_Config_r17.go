package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DeactivatedSCG_Config_r17 struct {
	bfd_and_RLM_r17 bool `madatory`
}

func (ie *DeactivatedSCG_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.bfd_and_RLM_r17); err != nil {
		return utils.WrapError("WriteBoolean bfd_and_RLM_r17", err)
	}
	return nil
}

func (ie *DeactivatedSCG_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_bfd_and_RLM_r17 bool
	if tmp_bool_bfd_and_RLM_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean bfd_and_RLM_r17", err)
	}
	ie.bfd_and_RLM_r17 = tmp_bool_bfd_and_RLM_r17
	return nil
}
