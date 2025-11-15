package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarConditionalReconfig struct {
	condReconfigList *CondReconfigToAddModList_r16 `optional`
}

func (ie *VarConditionalReconfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.condReconfigList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.condReconfigList != nil {
		if err = ie.condReconfigList.Encode(w); err != nil {
			return utils.WrapError("Encode condReconfigList", err)
		}
	}
	return nil
}

func (ie *VarConditionalReconfig) Decode(r *uper.UperReader) error {
	var err error
	var condReconfigListPresent bool
	if condReconfigListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if condReconfigListPresent {
		ie.condReconfigList = new(CondReconfigToAddModList_r16)
		if err = ie.condReconfigList.Decode(r); err != nil {
			return utils.WrapError("Decode condReconfigList", err)
		}
	}
	return nil
}
