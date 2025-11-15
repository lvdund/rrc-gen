package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1540_IEs_reportCGI_RequestNR struct {
	requestedCellInfo *CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo `optional`
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.requestedCellInfo != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.requestedCellInfo != nil {
		if err = ie.requestedCellInfo.Encode(w); err != nil {
			return utils.WrapError("Encode requestedCellInfo", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR) Decode(r *uper.UperReader) error {
	var err error
	var requestedCellInfoPresent bool
	if requestedCellInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if requestedCellInfoPresent {
		ie.requestedCellInfo = new(CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo)
		if err = ie.requestedCellInfo.Decode(r); err != nil {
			return utils.WrapError("Decode requestedCellInfo", err)
		}
	}
	return nil
}
