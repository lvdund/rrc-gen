package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookConfig_r17_codebookType_type2 struct {
	typeII_PortSelection_r17 *CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17 `lb:4,ub:4,optional`
}

func (ie *CodebookConfig_r17_codebookType_type2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.typeII_PortSelection_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.typeII_PortSelection_r17 != nil {
		if err = ie.typeII_PortSelection_r17.Encode(w); err != nil {
			return utils.WrapError("Encode typeII_PortSelection_r17", err)
		}
	}
	return nil
}

func (ie *CodebookConfig_r17_codebookType_type2) Decode(r *uper.UperReader) error {
	var err error
	var typeII_PortSelection_r17Present bool
	if typeII_PortSelection_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if typeII_PortSelection_r17Present {
		ie.typeII_PortSelection_r17 = new(CodebookConfig_r17_codebookType_type2_typeII_PortSelection_r17)
		if err = ie.typeII_PortSelection_r17.Decode(r); err != nil {
			return utils.WrapError("Decode typeII_PortSelection_r17", err)
		}
	}
	return nil
}
