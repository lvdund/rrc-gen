package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1690 struct {
	ul_RRC_Segmentation_r16 *UE_NR_Capability_v1690_ul_RRC_Segmentation_r16 `optional`
	nonCriticalExtension    *UE_NR_Capability_v1700                         `optional`
}

func (ie *UE_NR_Capability_v1690) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ul_RRC_Segmentation_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_RRC_Segmentation_r16 != nil {
		if err = ie.ul_RRC_Segmentation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_RRC_Segmentation_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1690) Decode(r *uper.UperReader) error {
	var err error
	var ul_RRC_Segmentation_r16Present bool
	if ul_RRC_Segmentation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_RRC_Segmentation_r16Present {
		ie.ul_RRC_Segmentation_r16 = new(UE_NR_Capability_v1690_ul_RRC_Segmentation_r16)
		if err = ie.ul_RRC_Segmentation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_RRC_Segmentation_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1700)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
