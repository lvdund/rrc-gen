package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationCompleteSidelink_v1710_IEs struct {
	dummy                RRCReconfigurationCompleteSidelink_v1710_IEs_dummy `madatory`
	nonCriticalExtension *RRCReconfigurationCompleteSidelink_v1720_IEs      `optional`
}

func (ie *RRCReconfigurationCompleteSidelink_v1710_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.dummy.Encode(w); err != nil {
		return utils.WrapError("Encode dummy", err)
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationCompleteSidelink_v1710_IEs) Decode(r *uper.UperReader) error {
	var err error
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.dummy.Decode(r); err != nil {
		return utils.WrapError("Decode dummy", err)
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfigurationCompleteSidelink_v1720_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
