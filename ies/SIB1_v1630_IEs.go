package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1630_IEs struct {
	uac_BarringInfo_v1630 *SIB1_v1630_IEs_uac_BarringInfo_v1630 `lb:2,ub:maxPLMN,optional`
	nonCriticalExtension  *SIB1_v1700_IEs                       `optional`
}

func (ie *SIB1_v1630_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uac_BarringInfo_v1630 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uac_BarringInfo_v1630 != nil {
		if err = ie.uac_BarringInfo_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode uac_BarringInfo_v1630", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB1_v1630_IEs) Decode(r *uper.UperReader) error {
	var err error
	var uac_BarringInfo_v1630Present bool
	if uac_BarringInfo_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if uac_BarringInfo_v1630Present {
		ie.uac_BarringInfo_v1630 = new(SIB1_v1630_IEs_uac_BarringInfo_v1630)
		if err = ie.uac_BarringInfo_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode uac_BarringInfo_v1630", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(SIB1_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
