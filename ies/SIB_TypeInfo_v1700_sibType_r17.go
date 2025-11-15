package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB_TypeInfo_v1700_sibType_r17_Choice_nothing uint64 = iota
	SIB_TypeInfo_v1700_sibType_r17_Choice_type1_r17
	SIB_TypeInfo_v1700_sibType_r17_Choice_type2_r17
)

type SIB_TypeInfo_v1700_sibType_r17 struct {
	Choice    uint64
	type1_r17 *SIB_TypeInfo_v1700_sibType_r17_type1_r17
	type2_r17 *SIB_TypeInfo_v1700_sibType_r17_type2_r17
}

func (ie *SIB_TypeInfo_v1700_sibType_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SIB_TypeInfo_v1700_sibType_r17_Choice_type1_r17:
		if err = ie.type1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode type1_r17", err)
		}
	case SIB_TypeInfo_v1700_sibType_r17_Choice_type2_r17:
		if err = ie.type2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode type2_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SIB_TypeInfo_v1700_sibType_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SIB_TypeInfo_v1700_sibType_r17_Choice_type1_r17:
		ie.type1_r17 = new(SIB_TypeInfo_v1700_sibType_r17_type1_r17)
		if err = ie.type1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode type1_r17", err)
		}
	case SIB_TypeInfo_v1700_sibType_r17_Choice_type2_r17:
		ie.type2_r17 = new(SIB_TypeInfo_v1700_sibType_r17_type2_r17)
		if err = ie.type2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode type2_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
