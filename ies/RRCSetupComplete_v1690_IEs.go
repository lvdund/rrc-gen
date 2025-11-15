package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupComplete_v1690_IEs struct {
	ul_RRC_Segmentation_r16 *RRCSetupComplete_v1690_IEs_ul_RRC_Segmentation_r16 `optional`
	nonCriticalExtension    *RRCSetupComplete_v1700_IEs                         `optional`
}

func (ie *RRCSetupComplete_v1690_IEs) Encode(w *uper.UperWriter) error {
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

func (ie *RRCSetupComplete_v1690_IEs) Decode(r *uper.UperReader) error {
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
		ie.ul_RRC_Segmentation_r16 = new(RRCSetupComplete_v1690_IEs_ul_RRC_Segmentation_r16)
		if err = ie.ul_RRC_Segmentation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_RRC_Segmentation_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCSetupComplete_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
