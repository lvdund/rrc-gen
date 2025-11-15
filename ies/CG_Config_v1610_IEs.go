package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1610_IEs struct {
	drx_InfoSCG2         *DRX_Info2           `optional`
	nonCriticalExtension *CG_Config_v1620_IEs `optional`
}

func (ie *CG_Config_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.drx_InfoSCG2 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.drx_InfoSCG2 != nil {
		if err = ie.drx_InfoSCG2.Encode(w); err != nil {
			return utils.WrapError("Encode drx_InfoSCG2", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var drx_InfoSCG2Present bool
	if drx_InfoSCG2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if drx_InfoSCG2Present {
		ie.drx_InfoSCG2 = new(DRX_Info2)
		if err = ie.drx_InfoSCG2.Decode(r); err != nil {
			return utils.WrapError("Decode drx_InfoSCG2", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1620_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
