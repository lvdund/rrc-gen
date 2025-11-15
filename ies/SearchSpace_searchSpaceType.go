package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_Choice_nothing uint64 = iota
	SearchSpace_searchSpaceType_Choice_common
	SearchSpace_searchSpaceType_Choice_ue_Specific
)

type SearchSpace_searchSpaceType struct {
	Choice      uint64
	common      *SearchSpace_searchSpaceType_common
	ue_Specific *SearchSpace_searchSpaceType_ue_Specific
}

func (ie *SearchSpace_searchSpaceType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpace_searchSpaceType_Choice_common:
		if err = ie.common.Encode(w); err != nil {
			err = utils.WrapError("Encode common", err)
		}
	case SearchSpace_searchSpaceType_Choice_ue_Specific:
		if err = ie.ue_Specific.Encode(w); err != nil {
			err = utils.WrapError("Encode ue_Specific", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SearchSpace_searchSpaceType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpace_searchSpaceType_Choice_common:
		ie.common = new(SearchSpace_searchSpaceType_common)
		if err = ie.common.Decode(r); err != nil {
			return utils.WrapError("Decode common", err)
		}
	case SearchSpace_searchSpaceType_Choice_ue_Specific:
		ie.ue_Specific = new(SearchSpace_searchSpaceType_ue_Specific)
		if err = ie.ue_Specific.Decode(r); err != nil {
			return utils.WrapError("Decode ue_Specific", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
