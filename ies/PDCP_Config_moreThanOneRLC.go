package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_Config_moreThanOneRLC struct {
	primaryPath           *PDCP_Config_moreThanOneRLC_primaryPath `optional`
	ul_DataSplitThreshold *UL_DataSplitThreshold                  `optional`
	pdcp_Duplication      *bool                                   `optional`
}

func (ie *PDCP_Config_moreThanOneRLC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.primaryPath != nil, ie.ul_DataSplitThreshold != nil, ie.pdcp_Duplication != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.primaryPath != nil {
		if err = ie.primaryPath.Encode(w); err != nil {
			return utils.WrapError("Encode primaryPath", err)
		}
	}
	if ie.ul_DataSplitThreshold != nil {
		if err = ie.ul_DataSplitThreshold.Encode(w); err != nil {
			return utils.WrapError("Encode ul_DataSplitThreshold", err)
		}
	}
	if ie.pdcp_Duplication != nil {
		if err = w.WriteBoolean(*ie.pdcp_Duplication); err != nil {
			return utils.WrapError("Encode pdcp_Duplication", err)
		}
	}
	return nil
}

func (ie *PDCP_Config_moreThanOneRLC) Decode(r *uper.UperReader) error {
	var err error
	var primaryPathPresent bool
	if primaryPathPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_DataSplitThresholdPresent bool
	if ul_DataSplitThresholdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_DuplicationPresent bool
	if pdcp_DuplicationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if primaryPathPresent {
		ie.primaryPath = new(PDCP_Config_moreThanOneRLC_primaryPath)
		if err = ie.primaryPath.Decode(r); err != nil {
			return utils.WrapError("Decode primaryPath", err)
		}
	}
	if ul_DataSplitThresholdPresent {
		ie.ul_DataSplitThreshold = new(UL_DataSplitThreshold)
		if err = ie.ul_DataSplitThreshold.Decode(r); err != nil {
			return utils.WrapError("Decode ul_DataSplitThreshold", err)
		}
	}
	if pdcp_DuplicationPresent {
		var tmp_bool_pdcp_Duplication bool
		if tmp_bool_pdcp_Duplication, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode pdcp_Duplication", err)
		}
		ie.pdcp_Duplication = &tmp_bool_pdcp_Duplication
	}
	return nil
}
