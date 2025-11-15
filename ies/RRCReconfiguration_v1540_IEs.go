package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1540_IEs struct {
	otherConfig_v1540    *OtherConfig_v1540            `optional`
	nonCriticalExtension *RRCReconfiguration_v1560_IEs `optional`
}

func (ie *RRCReconfiguration_v1540_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.otherConfig_v1540 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.otherConfig_v1540 != nil {
		if err = ie.otherConfig_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode otherConfig_v1540", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1540_IEs) Decode(r *uper.UperReader) error {
	var err error
	var otherConfig_v1540Present bool
	if otherConfig_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if otherConfig_v1540Present {
		ie.otherConfig_v1540 = new(OtherConfig_v1540)
		if err = ie.otherConfig_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode otherConfig_v1540", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfiguration_v1560_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
