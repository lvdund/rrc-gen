package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17 struct {
	aggregationLevel4_r17  *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel4_r17  `optional`
	aggregationLevel8_r17  *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17  `optional`
	aggregationLevel16_r17 *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17 `optional`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.aggregationLevel4_r17 != nil, ie.aggregationLevel8_r17 != nil, ie.aggregationLevel16_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.aggregationLevel4_r17 != nil {
		if err = ie.aggregationLevel4_r17.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel4_r17", err)
		}
	}
	if ie.aggregationLevel8_r17 != nil {
		if err = ie.aggregationLevel8_r17.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel8_r17", err)
		}
	}
	if ie.aggregationLevel16_r17 != nil {
		if err = ie.aggregationLevel16_r17.Encode(w); err != nil {
			return utils.WrapError("Encode aggregationLevel16_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17) Decode(r *uper.UperReader) error {
	var err error
	var aggregationLevel4_r17Present bool
	if aggregationLevel4_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel8_r17Present bool
	if aggregationLevel8_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aggregationLevel16_r17Present bool
	if aggregationLevel16_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if aggregationLevel4_r17Present {
		ie.aggregationLevel4_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel4_r17)
		if err = ie.aggregationLevel4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel4_r17", err)
		}
	}
	if aggregationLevel8_r17Present {
		ie.aggregationLevel8_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel8_r17)
		if err = ie.aggregationLevel8_r17.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel8_r17", err)
		}
	}
	if aggregationLevel16_r17Present {
		ie.aggregationLevel16_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17)
		if err = ie.aggregationLevel16_r17.Decode(r); err != nil {
			return utils.WrapError("Decode aggregationLevel16_r17", err)
		}
	}
	return nil
}
