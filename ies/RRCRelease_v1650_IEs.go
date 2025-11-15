package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCRelease_v1650_IEs struct {
	mpsPriorityIndication_r16 *RRCRelease_v1650_IEs_mpsPriorityIndication_r16 `optional`
	nonCriticalExtension      *RRCRelease_v1710_IEs                           `optional`
}

func (ie *RRCRelease_v1650_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mpsPriorityIndication_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mpsPriorityIndication_r16 != nil {
		if err = ie.mpsPriorityIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mpsPriorityIndication_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCRelease_v1650_IEs) Decode(r *uper.UperReader) error {
	var err error
	var mpsPriorityIndication_r16Present bool
	if mpsPriorityIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mpsPriorityIndication_r16Present {
		ie.mpsPriorityIndication_r16 = new(RRCRelease_v1650_IEs_mpsPriorityIndication_r16)
		if err = ie.mpsPriorityIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mpsPriorityIndication_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCRelease_v1710_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
