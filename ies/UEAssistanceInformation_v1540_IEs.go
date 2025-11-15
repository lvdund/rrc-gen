package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1540_IEs struct {
	overheatingAssistance *OverheatingAssistance             `optional`
	nonCriticalExtension  *UEAssistanceInformation_v1610_IEs `optional`
}

func (ie *UEAssistanceInformation_v1540_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.overheatingAssistance != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.overheatingAssistance != nil {
		if err = ie.overheatingAssistance.Encode(w); err != nil {
			return utils.WrapError("Encode overheatingAssistance", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1540_IEs) Decode(r *uper.UperReader) error {
	var err error
	var overheatingAssistancePresent bool
	if overheatingAssistancePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if overheatingAssistancePresent {
		ie.overheatingAssistance = new(OverheatingAssistance)
		if err = ie.overheatingAssistance.Decode(r); err != nil {
			return utils.WrapError("Decode overheatingAssistance", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UEAssistanceInformation_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
