package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CandidateCell_r17 struct {
	physCellId_r17           PhysCellId `madatory`
	condExecutionCondSCG_r17 *[]byte    `optional`
}

func (ie *CandidateCell_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.condExecutionCondSCG_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.physCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r17", err)
	}
	if ie.condExecutionCondSCG_r17 != nil {
		if err = w.WriteOctetString(*ie.condExecutionCondSCG_r17, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode condExecutionCondSCG_r17", err)
		}
	}
	return nil
}

func (ie *CandidateCell_r17) Decode(r *uper.UperReader) error {
	var err error
	var condExecutionCondSCG_r17Present bool
	if condExecutionCondSCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.physCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r17", err)
	}
	if condExecutionCondSCG_r17Present {
		var tmp_os_condExecutionCondSCG_r17 []byte
		if tmp_os_condExecutionCondSCG_r17, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode condExecutionCondSCG_r17", err)
		}
		ie.condExecutionCondSCG_r17 = &tmp_os_condExecutionCondSCG_r17
	}
	return nil
}
