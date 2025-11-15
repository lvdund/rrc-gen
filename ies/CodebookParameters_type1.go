package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type1 struct {
	singlePanel CodebookParameters_type1_singlePanel `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	multiPanel  *CodebookParameters_type1_multiPanel `lb:1,ub:maxNrofCSI_RS_Resources,optional`
}

func (ie *CodebookParameters_type1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.multiPanel != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.singlePanel.Encode(w); err != nil {
		return utils.WrapError("Encode singlePanel", err)
	}
	if ie.multiPanel != nil {
		if err = ie.multiPanel.Encode(w); err != nil {
			return utils.WrapError("Encode multiPanel", err)
		}
	}
	return nil
}

func (ie *CodebookParameters_type1) Decode(r *uper.UperReader) error {
	var err error
	var multiPanelPresent bool
	if multiPanelPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.singlePanel.Decode(r); err != nil {
		return utils.WrapError("Decode singlePanel", err)
	}
	if multiPanelPresent {
		ie.multiPanel = new(CodebookParameters_type1_multiPanel)
		if err = ie.multiPanel.Decode(r); err != nil {
			return utils.WrapError("Decode multiPanel", err)
		}
	}
	return nil
}
