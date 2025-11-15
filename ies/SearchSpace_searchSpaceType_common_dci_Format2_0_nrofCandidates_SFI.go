package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI struct {
	aggregationLevel1  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel1  `optional`
	aggregationLevel2  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel2  `optional`
	aggregationLevel4  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4  `optional`
	aggregationLevel8  *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8  `optional`
	aggregationLevel16 *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel16 `optional`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.aggregationLevel1 != nil, ie.aggregationLevel2 != nil, ie.aggregationLevel4 != nil, ie.aggregationLevel8 != nil, ie.aggregationLevel16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.aggregationLevel1 != nil {
		if err = ie.aggregationLevel1.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel1", err)
		}
	}
	if ie.aggregationLevel2 != nil {
		if err = ie.aggregationLevel2.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel2", err)
		}
	}
	if ie.aggregationLevel4 != nil {
		if err = ie.aggregationLevel4.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel4", err)
		}
	}
	if ie.aggregationLevel8 != nil {
		if err = ie.aggregationLevel8.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel8", err)
		}
	}
	if ie.aggregationLevel16 != nil {
		if err = ie.aggregationLevel16.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel16", err)
		}
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI) Decode(r *uper.UperReader) error {
	var err error
	var aggregationLevel1Present bool
	if aggregationLevel1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel2Present bool
	if aggregationLevel2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel4Present bool
	if aggregationLevel4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel8Present bool
	if aggregationLevel8Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel16Present bool
	if aggregationLevel16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if aggregationLevel1Present {
		ie.aggregationLevel1 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel1)
		if err = ie.aggregationLevel1.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel1", err)
		}
	}
	if aggregationLevel2Present {
		ie.aggregationLevel2 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel2)
		if err = ie.aggregationLevel2.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel2", err)
		}
	}
	if aggregationLevel4Present {
		ie.aggregationLevel4 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel4)
		if err = ie.aggregationLevel4.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel4", err)
		}
	}
	if aggregationLevel8Present {
		ie.aggregationLevel8 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8)
		if err = ie.aggregationLevel8.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel8", err)
		}
	}
	if aggregationLevel16Present {
		ie.aggregationLevel16 = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel16)
		if err = ie.aggregationLevel16.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel16", err)
		}
	}
	return nil
}
