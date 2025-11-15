package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DLInformationTransfer_v1610_IEs struct {
	referenceTimeInfo_r16 *ReferenceTimeInfo_r16           `optional`
	nonCriticalExtension  *DLInformationTransfer_v1700_IEs `optional`
}

func (ie *DLInformationTransfer_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.referenceTimeInfo_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.referenceTimeInfo_r16 != nil {
		if err = ie.referenceTimeInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode referenceTimeInfo_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *DLInformationTransfer_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var referenceTimeInfo_r16Present bool
	if referenceTimeInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if referenceTimeInfo_r16Present {
		ie.referenceTimeInfo_r16 = new(ReferenceTimeInfo_r16)
		if err = ie.referenceTimeInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode referenceTimeInfo_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(DLInformationTransfer_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
