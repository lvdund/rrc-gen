package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureCombination_r17 struct {
	redCap_r17           *FeatureCombination_r17_redCap_r17           `optional`
	smallData_r17        *FeatureCombination_r17_smallData_r17        `optional`
	nsag_r17             *NSAG_List_r17                               `optional`
	msg3_Repetitions_r17 *FeatureCombination_r17_msg3_Repetitions_r17 `optional`
	spare4               *FeatureCombination_r17_spare4               `optional`
	spare3               *FeatureCombination_r17_spare3               `optional`
	spare2               *FeatureCombination_r17_spare2               `optional`
	spare1               *FeatureCombination_r17_spare1               `optional`
}

func (ie *FeatureCombination_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.redCap_r17 != nil, ie.smallData_r17 != nil, ie.nsag_r17 != nil, ie.msg3_Repetitions_r17 != nil, ie.spare4 != nil, ie.spare3 != nil, ie.spare2 != nil, ie.spare1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.redCap_r17 != nil {
		if err = ie.redCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode redCap_r17", err)
		}
	}
	if ie.smallData_r17 != nil {
		if err = ie.smallData_r17.Encode(w); err != nil {
			return utils.WrapError("Encode smallData_r17", err)
		}
	}
	if ie.nsag_r17 != nil {
		if err = ie.nsag_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nsag_r17", err)
		}
	}
	if ie.msg3_Repetitions_r17 != nil {
		if err = ie.msg3_Repetitions_r17.Encode(w); err != nil {
			return utils.WrapError("Encode msg3_Repetitions_r17", err)
		}
	}
	if ie.spare4 != nil {
		if err = ie.spare4.Encode(w); err != nil {
			return utils.WrapError("Encode spare4", err)
		}
	}
	if ie.spare3 != nil {
		if err = ie.spare3.Encode(w); err != nil {
			return utils.WrapError("Encode spare3", err)
		}
	}
	if ie.spare2 != nil {
		if err = ie.spare2.Encode(w); err != nil {
			return utils.WrapError("Encode spare2", err)
		}
	}
	if ie.spare1 != nil {
		if err = ie.spare1.Encode(w); err != nil {
			return utils.WrapError("Encode spare1", err)
		}
	}
	return nil
}

func (ie *FeatureCombination_r17) Decode(r *uper.UperReader) error {
	var err error
	var redCap_r17Present bool
	if redCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var smallData_r17Present bool
	if smallData_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nsag_r17Present bool
	if nsag_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msg3_Repetitions_r17Present bool
	if msg3_Repetitions_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var spare4Present bool
	if spare4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var spare3Present bool
	if spare3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var spare2Present bool
	if spare2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var spare1Present bool
	if spare1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if redCap_r17Present {
		ie.redCap_r17 = new(FeatureCombination_r17_redCap_r17)
		if err = ie.redCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode redCap_r17", err)
		}
	}
	if smallData_r17Present {
		ie.smallData_r17 = new(FeatureCombination_r17_smallData_r17)
		if err = ie.smallData_r17.Decode(r); err != nil {
			return utils.WrapError("Decode smallData_r17", err)
		}
	}
	if nsag_r17Present {
		ie.nsag_r17 = new(NSAG_List_r17)
		if err = ie.nsag_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nsag_r17", err)
		}
	}
	if msg3_Repetitions_r17Present {
		ie.msg3_Repetitions_r17 = new(FeatureCombination_r17_msg3_Repetitions_r17)
		if err = ie.msg3_Repetitions_r17.Decode(r); err != nil {
			return utils.WrapError("Decode msg3_Repetitions_r17", err)
		}
	}
	if spare4Present {
		ie.spare4 = new(FeatureCombination_r17_spare4)
		if err = ie.spare4.Decode(r); err != nil {
			return utils.WrapError("Decode spare4", err)
		}
	}
	if spare3Present {
		ie.spare3 = new(FeatureCombination_r17_spare3)
		if err = ie.spare3.Decode(r); err != nil {
			return utils.WrapError("Decode spare3", err)
		}
	}
	if spare2Present {
		ie.spare2 = new(FeatureCombination_r17_spare2)
		if err = ie.spare2.Decode(r); err != nil {
			return utils.WrapError("Decode spare2", err)
		}
	}
	if spare1Present {
		ie.spare1 = new(FeatureCombination_r17_spare1)
		if err = ie.spare1.Decode(r); err != nil {
			return utils.WrapError("Decode spare1", err)
		}
	}
	return nil
}
