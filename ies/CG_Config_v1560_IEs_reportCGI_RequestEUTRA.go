package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1560_IEs_reportCGI_RequestEUTRA struct {
	requestedCellInfoEUTRA *CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA `optional`
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.requestedCellInfoEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.requestedCellInfoEUTRA != nil {
		if err = ie.requestedCellInfoEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode requestedCellInfoEUTRA", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var requestedCellInfoEUTRAPresent bool
	if requestedCellInfoEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if requestedCellInfoEUTRAPresent {
		ie.requestedCellInfoEUTRA = new(CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA)
		if err = ie.requestedCellInfoEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode requestedCellInfoEUTRA", err)
		}
	}
	return nil
}
