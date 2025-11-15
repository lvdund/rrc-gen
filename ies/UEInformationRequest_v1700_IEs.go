package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationRequest_v1700_IEs struct {
	successHO_ReportReq_r17   *UEInformationRequest_v1700_IEs_successHO_ReportReq_r17   `optional`
	coarseLocationRequest_r17 *UEInformationRequest_v1700_IEs_coarseLocationRequest_r17 `optional`
	nonCriticalExtension      interface{}                                               `optional`
}

func (ie *UEInformationRequest_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.successHO_ReportReq_r17 != nil, ie.coarseLocationRequest_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.successHO_ReportReq_r17 != nil {
		if err = ie.successHO_ReportReq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode successHO_ReportReq_r17", err)
		}
	}
	if ie.coarseLocationRequest_r17 != nil {
		if err = ie.coarseLocationRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode coarseLocationRequest_r17", err)
		}
	}
	return nil
}

func (ie *UEInformationRequest_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var successHO_ReportReq_r17Present bool
	if successHO_ReportReq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var coarseLocationRequest_r17Present bool
	if coarseLocationRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if successHO_ReportReq_r17Present {
		ie.successHO_ReportReq_r17 = new(UEInformationRequest_v1700_IEs_successHO_ReportReq_r17)
		if err = ie.successHO_ReportReq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode successHO_ReportReq_r17", err)
		}
	}
	if coarseLocationRequest_r17Present {
		ie.coarseLocationRequest_r17 = new(UEInformationRequest_v1700_IEs_coarseLocationRequest_r17)
		if err = ie.coarseLocationRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode coarseLocationRequest_r17", err)
		}
	}
	return nil
}
