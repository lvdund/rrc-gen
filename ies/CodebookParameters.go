package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters struct {
	type1               *CodebookParameters_type1               `lb:1,ub:maxNrofCSI_RS_Resources,optional`
	type2               *CodebookParameters_type2               `lb:1,ub:maxNrofCSI_RS_Resources,optional`
	type2_PortSelection *CodebookParameters_type2_PortSelection `lb:1,ub:maxNrofCSI_RS_Resources,optional`
}

func (ie *CodebookParameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.type1 != nil, ie.type2 != nil, ie.type2_PortSelection != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.type1 != nil {
		if err = ie.type1.Encode(w); err != nil {
			return utils.WrapError("Encode type1", err)
		}
	}
	if ie.type2 != nil {
		if err = ie.type2.Encode(w); err != nil {
			return utils.WrapError("Encode type2", err)
		}
	}
	if ie.type2_PortSelection != nil {
		if err = ie.type2_PortSelection.Encode(w); err != nil {
			return utils.WrapError("Encode type2_PortSelection", err)
		}
	}
	return nil
}

func (ie *CodebookParameters) Decode(r *uper.UperReader) error {
	var err error
	var type1Present bool
	if type1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type2Present bool
	if type2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var type2_PortSelectionPresent bool
	if type2_PortSelectionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if type1Present {
		ie.type1 = new(CodebookParameters_type1)
		if err = ie.type1.Decode(r); err != nil {
			return utils.WrapError("Decode type1", err)
		}
	}
	if type2Present {
		ie.type2 = new(CodebookParameters_type2)
		if err = ie.type2.Decode(r); err != nil {
			return utils.WrapError("Decode type2", err)
		}
	}
	if type2_PortSelectionPresent {
		ie.type2_PortSelection = new(CodebookParameters_type2_PortSelection)
		if err = ie.type2_PortSelection.Decode(r); err != nil {
			return utils.WrapError("Decode type2_PortSelection", err)
		}
	}
	return nil
}
