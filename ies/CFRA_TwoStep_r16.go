package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_TwoStep_r16 struct {
	occasionsTwoStepRA_r16 *CFRA_TwoStep_r16_occasionsTwoStepRA_r16 `optional`
	msgA_CFRA_PUSCH_r16    MsgA_PUSCH_Resource_r16                  `madatory`
	msgA_TransMax_r16      *CFRA_TwoStep_r16_msgA_TransMax_r16      `optional`
	resourcesTwoStep_r16   CFRA_TwoStep_r16_resourcesTwoStep_r16    `lb:1,ub:maxRA_SSB_Resources,madatory`
}

func (ie *CFRA_TwoStep_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.occasionsTwoStepRA_r16 != nil, ie.msgA_TransMax_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.occasionsTwoStepRA_r16 != nil {
		if err = ie.occasionsTwoStepRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode occasionsTwoStepRA_r16", err)
		}
	}
	if err = ie.msgA_CFRA_PUSCH_r16.Encode(w); err != nil {
		return utils.WrapError("Encode msgA_CFRA_PUSCH_r16", err)
	}
	if ie.msgA_TransMax_r16 != nil {
		if err = ie.msgA_TransMax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_TransMax_r16", err)
		}
	}
	if err = ie.resourcesTwoStep_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resourcesTwoStep_r16", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16) Decode(r *uper.UperReader) error {
	var err error
	var occasionsTwoStepRA_r16Present bool
	if occasionsTwoStepRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_TransMax_r16Present bool
	if msgA_TransMax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if occasionsTwoStepRA_r16Present {
		ie.occasionsTwoStepRA_r16 = new(CFRA_TwoStep_r16_occasionsTwoStepRA_r16)
		if err = ie.occasionsTwoStepRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode occasionsTwoStepRA_r16", err)
		}
	}
	if err = ie.msgA_CFRA_PUSCH_r16.Decode(r); err != nil {
		return utils.WrapError("Decode msgA_CFRA_PUSCH_r16", err)
	}
	if msgA_TransMax_r16Present {
		ie.msgA_TransMax_r16 = new(CFRA_TwoStep_r16_msgA_TransMax_r16)
		if err = ie.msgA_TransMax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_TransMax_r16", err)
		}
	}
	if err = ie.resourcesTwoStep_r16.Decode(r); err != nil {
		return utils.WrapError("Decode resourcesTwoStep_r16", err)
	}
	return nil
}
