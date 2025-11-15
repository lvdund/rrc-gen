package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureInfoRLC_Bearer struct {
	cellGroupId            CellGroupId                       `madatory`
	logicalChannelIdentity LogicalChannelIdentity            `madatory`
	failureType            FailureInfoRLC_Bearer_failureType `madatory`
}

func (ie *FailureInfoRLC_Bearer) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cellGroupId.Encode(w); err != nil {
		return utils.WrapError("Encode cellGroupId", err)
	}
	if err = ie.logicalChannelIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode logicalChannelIdentity", err)
	}
	if err = ie.failureType.Encode(w); err != nil {
		return utils.WrapError("Encode failureType", err)
	}
	return nil
}

func (ie *FailureInfoRLC_Bearer) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cellGroupId.Decode(r); err != nil {
		return utils.WrapError("Decode cellGroupId", err)
	}
	if err = ie.logicalChannelIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode logicalChannelIdentity", err)
	}
	if err = ie.failureType.Decode(r); err != nil {
		return utils.WrapError("Decode failureType", err)
	}
	return nil
}
