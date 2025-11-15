package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1560_IEs struct {
	scg_Response         *RRCReconfigurationComplete_v1560_IEs_scg_Response `optional`
	nonCriticalExtension *RRCReconfigurationComplete_v1610_IEs              `optional`
}

func (ie *RRCReconfigurationComplete_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scg_Response != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scg_Response != nil {
		if err = ie.scg_Response.Encode(w); err != nil {
			return utils.WrapError("Encode scg_Response", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var scg_ResponsePresent bool
	if scg_ResponsePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scg_ResponsePresent {
		ie.scg_Response = new(RRCReconfigurationComplete_v1560_IEs_scg_Response)
		if err = ie.scg_Response.Decode(r); err != nil {
			return utils.WrapError("Decode scg_Response", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfigurationComplete_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
