package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1540_IEs struct {
	waitTime             *RejectWaitTime       `optional`
	nonCriticalExtension *RRCRelease_v1610_IEs `optional`
}

func (ie *RRCRelease_v1540_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.waitTime != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.waitTime != nil {
		if err = ie.waitTime.Encode(w); err != nil {
			return utils.WrapError("Encode waitTime", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1540_IEs) Decode(r *uper.UperReader) error {
	var err error
	var waitTimePresent bool
	if waitTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if waitTimePresent {
		ie.waitTime = new(RejectWaitTime)
		if err = ie.waitTime.Decode(r); err != nil {
			return utils.WrapError("Decode waitTime", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCRelease_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
