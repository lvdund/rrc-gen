package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_CodeBlockGroupTransmission struct {
	maxCodeBlockGroupsPerTransportBlock PDSCH_CodeBlockGroupTransmission_maxCodeBlockGroupsPerTransportBlock `madatory`
	codeBlockGroupFlushIndicator        bool                                                                 `madatory`
}

func (ie *PDSCH_CodeBlockGroupTransmission) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxCodeBlockGroupsPerTransportBlock.Encode(w); err != nil {
		return utils.WrapError("Encode maxCodeBlockGroupsPerTransportBlock", err)
	}
	if err = w.WriteBoolean(ie.codeBlockGroupFlushIndicator); err != nil {
		return utils.WrapError("WriteBoolean codeBlockGroupFlushIndicator", err)
	}
	return nil
}

func (ie *PDSCH_CodeBlockGroupTransmission) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxCodeBlockGroupsPerTransportBlock.Decode(r); err != nil {
		return utils.WrapError("Decode maxCodeBlockGroupsPerTransportBlock", err)
	}
	var tmp_bool_codeBlockGroupFlushIndicator bool
	if tmp_bool_codeBlockGroupFlushIndicator, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean codeBlockGroupFlushIndicator", err)
	}
	ie.codeBlockGroupFlushIndicator = tmp_bool_codeBlockGroupFlushIndicator
	return nil
}
