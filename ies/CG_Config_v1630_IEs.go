package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1630_IEs struct {
	selectedToffset_r16  *T_Offset_r16        `optional`
	nonCriticalExtension *CG_Config_v1640_IEs `optional`
}

func (ie *CG_Config_v1630_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.selectedToffset_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.selectedToffset_r16 != nil {
		if err = ie.selectedToffset_r16.Encode(w); err != nil {
			return utils.WrapError("Encode selectedToffset_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1630_IEs) Decode(r *uper.UperReader) error {
	var err error
	var selectedToffset_r16Present bool
	if selectedToffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if selectedToffset_r16Present {
		ie.selectedToffset_r16 = new(T_Offset_r16)
		if err = ie.selectedToffset_r16.Decode(r); err != nil {
			return utils.WrapError("Decode selectedToffset_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1640_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
