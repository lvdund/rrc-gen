package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_nrofCandidates struct {
	aggregationLevel1  SearchSpace_nrofCandidates_aggregationLevel1  `madatory`
	aggregationLevel2  SearchSpace_nrofCandidates_aggregationLevel2  `madatory`
	aggregationLevel4  SearchSpace_nrofCandidates_aggregationLevel4  `madatory`
	aggregationLevel8  SearchSpace_nrofCandidates_aggregationLevel8  `madatory`
	aggregationLevel16 SearchSpace_nrofCandidates_aggregationLevel16 `madatory`
}

func (ie *SearchSpace_nrofCandidates) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.aggregationLevel1.Encode(w); err != nil {
		return utils.WrapError("Encode aggregationLevel1", err)
	}
	if err = ie.aggregationLevel2.Encode(w); err != nil {
		return utils.WrapError("Encode aggregationLevel2", err)
	}
	if err = ie.aggregationLevel4.Encode(w); err != nil {
		return utils.WrapError("Encode aggregationLevel4", err)
	}
	if err = ie.aggregationLevel8.Encode(w); err != nil {
		return utils.WrapError("Encode aggregationLevel8", err)
	}
	if err = ie.aggregationLevel16.Encode(w); err != nil {
		return utils.WrapError("Encode aggregationLevel16", err)
	}
	return nil
}

func (ie *SearchSpace_nrofCandidates) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.aggregationLevel1.Decode(r); err != nil {
		return utils.WrapError("Decode aggregationLevel1", err)
	}
	if err = ie.aggregationLevel2.Decode(r); err != nil {
		return utils.WrapError("Decode aggregationLevel2", err)
	}
	if err = ie.aggregationLevel4.Decode(r); err != nil {
		return utils.WrapError("Decode aggregationLevel4", err)
	}
	if err = ie.aggregationLevel8.Decode(r); err != nil {
		return utils.WrapError("Decode aggregationLevel8", err)
	}
	if err = ie.aggregationLevel16.Decode(r); err != nil {
		return utils.WrapError("Decode aggregationLevel16", err)
	}
	return nil
}
