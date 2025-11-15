package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SubgroupConfig_r17 struct {
	subgroupsNumPerPO_r17   int64  `lb:1,ub:maxNrofPagingSubgroups_r17,madatory`
	subgroupsNumForUEID_r17 *int64 `lb:1,ub:maxNrofPagingSubgroups_r17,optional`
}

func (ie *SubgroupConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.subgroupsNumForUEID_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.subgroupsNumPerPO_r17, &uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
		return utils.WrapError("WriteInteger subgroupsNumPerPO_r17", err)
	}
	if ie.subgroupsNumForUEID_r17 != nil {
		if err = w.WriteInteger(*ie.subgroupsNumForUEID_r17, &uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
			return utils.WrapError("Encode subgroupsNumForUEID_r17", err)
		}
	}
	return nil
}

func (ie *SubgroupConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var subgroupsNumForUEID_r17Present bool
	if subgroupsNumForUEID_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_subgroupsNumPerPO_r17 int64
	if tmp_int_subgroupsNumPerPO_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
		return utils.WrapError("ReadInteger subgroupsNumPerPO_r17", err)
	}
	ie.subgroupsNumPerPO_r17 = tmp_int_subgroupsNumPerPO_r17
	if subgroupsNumForUEID_r17Present {
		var tmp_int_subgroupsNumForUEID_r17 int64
		if tmp_int_subgroupsNumForUEID_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofPagingSubgroups_r17}, false); err != nil {
			return utils.WrapError("Decode subgroupsNumForUEID_r17", err)
		}
		ie.subgroupsNumForUEID_r17 = &tmp_int_subgroupsNumForUEID_r17
	}
	return nil
}
