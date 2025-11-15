package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_ue_Specific struct {
	dci_Formats        SearchSpace_searchSpaceType_ue_Specific_dci_Formats         `madatory`
	dci_Formats_MT_r16 *SearchSpace_searchSpaceType_ue_Specific_dci_Formats_MT_r16 `optional`
	dci_FormatsSL_r16  *SearchSpace_searchSpaceType_ue_Specific_dci_FormatsSL_r16  `optional`
	dci_FormatsExt_r16 *SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16 `optional`
}

func (ie *SearchSpace_searchSpaceType_ue_Specific) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dci_Formats_MT_r16 != nil, ie.dci_FormatsSL_r16 != nil, ie.dci_FormatsExt_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.dci_Formats.Encode(w); err != nil {
		return utils.WrapError("Encode dci_Formats", err)
	}
	if ie.dci_Formats_MT_r16 != nil {
		if err = ie.dci_Formats_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dci_Formats_MT_r16", err)
		}
	}
	if ie.dci_FormatsSL_r16 != nil {
		if err = ie.dci_FormatsSL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dci_FormatsSL_r16", err)
		}
	}
	if ie.dci_FormatsExt_r16 != nil {
		if err = ie.dci_FormatsExt_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dci_FormatsExt_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_ue_Specific) Decode(r *uper.UperReader) error {
	var err error
	var dci_Formats_MT_r16Present bool
	if dci_Formats_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dci_FormatsSL_r16Present bool
	if dci_FormatsSL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dci_FormatsExt_r16Present bool
	if dci_FormatsExt_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.dci_Formats.Decode(r); err != nil {
		return utils.WrapError("Decode dci_Formats", err)
	}
	if dci_Formats_MT_r16Present {
		ie.dci_Formats_MT_r16 = new(SearchSpace_searchSpaceType_ue_Specific_dci_Formats_MT_r16)
		if err = ie.dci_Formats_MT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dci_Formats_MT_r16", err)
		}
	}
	if dci_FormatsSL_r16Present {
		ie.dci_FormatsSL_r16 = new(SearchSpace_searchSpaceType_ue_Specific_dci_FormatsSL_r16)
		if err = ie.dci_FormatsSL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dci_FormatsSL_r16", err)
		}
	}
	if dci_FormatsExt_r16Present {
		ie.dci_FormatsExt_r16 = new(SearchSpace_searchSpaceType_ue_Specific_dci_FormatsExt_r16)
		if err = ie.dci_FormatsExt_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dci_FormatsExt_r16", err)
		}
	}
	return nil
}
