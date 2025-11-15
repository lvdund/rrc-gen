package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_nothing uint64 = iota
	PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_newSearchSpace
	PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_existingSearchSpace
)

type PDCCH_ConfigCommon_sdt_SearchSpace_r17 struct {
	Choice              uint64
	newSearchSpace      *SearchSpace
	existingSearchSpace *SearchSpaceId
}

func (ie *PDCCH_ConfigCommon_sdt_SearchSpace_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_newSearchSpace:
		if err = ie.newSearchSpace.Encode(w); err != nil {
			err = utils.WrapError("Encode newSearchSpace", err)
		}
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_existingSearchSpace:
		if err = ie.existingSearchSpace.Encode(w); err != nil {
			err = utils.WrapError("Encode existingSearchSpace", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDCCH_ConfigCommon_sdt_SearchSpace_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_newSearchSpace:
		ie.newSearchSpace = new(SearchSpace)
		if err = ie.newSearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode newSearchSpace", err)
		}
	case PDCCH_ConfigCommon_sdt_SearchSpace_r17_Choice_existingSearchSpace:
		ie.existingSearchSpace = new(SearchSpaceId)
		if err = ie.existingSearchSpace.Decode(r); err != nil {
			return utils.WrapError("Decode existingSearchSpace", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
