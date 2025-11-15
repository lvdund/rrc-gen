package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16 struct {
	aggregationLevel1_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel1_r16  `optional`
	aggregationLevel2_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16  `optional`
	aggregationLevel4_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel4_r16  `optional`
	aggregationLevel8_r16  *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel8_r16  `optional`
	aggregationLevel16_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel16_r16 `optional`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.aggregationLevel1_r16 != nil, ie.aggregationLevel2_r16 != nil, ie.aggregationLevel4_r16 != nil, ie.aggregationLevel8_r16 != nil, ie.aggregationLevel16_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.aggregationLevel1_r16 != nil {
		if err = ie.aggregationLevel1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel1_r16", err)
		}
	}
	if ie.aggregationLevel2_r16 != nil {
		if err = ie.aggregationLevel2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel2_r16", err)
		}
	}
	if ie.aggregationLevel4_r16 != nil {
		if err = ie.aggregationLevel4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel4_r16", err)
		}
	}
	if ie.aggregationLevel8_r16 != nil {
		if err = ie.aggregationLevel8_r16.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel8_r16", err)
		}
	}
	if ie.aggregationLevel16_r16 != nil {
		if err = ie.aggregationLevel16_r16.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel16_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16) Decode(r *uper.UperReader) error {
	var err error
	var aggregationLevel1_r16Present bool
	if aggregationLevel1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel2_r16Present bool
	if aggregationLevel2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel4_r16Present bool
	if aggregationLevel4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel8_r16Present bool
	if aggregationLevel8_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel16_r16Present bool
	if aggregationLevel16_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if aggregationLevel1_r16Present {
		ie.aggregationLevel1_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel1_r16)
		if err = ie.aggregationLevel1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel1_r16", err)
		}
	}
	if aggregationLevel2_r16Present {
		ie.aggregationLevel2_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16)
		if err = ie.aggregationLevel2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel2_r16", err)
		}
	}
	if aggregationLevel4_r16Present {
		ie.aggregationLevel4_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel4_r16)
		if err = ie.aggregationLevel4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel4_r16", err)
		}
	}
	if aggregationLevel8_r16Present {
		ie.aggregationLevel8_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel8_r16)
		if err = ie.aggregationLevel8_r16.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel8_r16", err)
		}
	}
	if aggregationLevel16_r16Present {
		ie.aggregationLevel16_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel16_r16)
		if err = ie.aggregationLevel16_r16.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel16_r16", err)
		}
	}
	return nil
}
