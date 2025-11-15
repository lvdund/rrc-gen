package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_moreThanOneRLC_primaryPath struct {
	cellGroup      *CellGroupId            `optional`
	logicalChannel *LogicalChannelIdentity `optional`
}

func (ie *PDCP_Config_moreThanOneRLC_primaryPath) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellGroup != nil, ie.logicalChannel != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cellGroup != nil {
		if err = ie.cellGroup.Encode(w); err != nil {
			return utils.WrapError("Encode cellGroup", err)
		}
	}
	if ie.logicalChannel != nil {
		if err = ie.logicalChannel.Encode(w); err != nil {
			return utils.WrapError("Encode logicalChannel", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_moreThanOneRLC_primaryPath) Decode(r *uper.UperReader) error {
	var err error
	var cellGroupPresent bool
	if cellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var logicalChannelPresent bool
	if logicalChannelPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cellGroupPresent {
		ie.cellGroup = new(CellGroupId)
		if err = ie.cellGroup.Decode(r); err != nil {
			return utils.WrapError("Decode cellGroup", err)
		}
	}
	if logicalChannelPresent {
		ie.logicalChannel = new(LogicalChannelIdentity)
		if err = ie.logicalChannel.Decode(r); err != nil {
			return utils.WrapError("Decode logicalChannel", err)
		}
	}
	return nil
}
