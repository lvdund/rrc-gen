package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureReportMCG_r16 struct {
	failureType_r16                *FailureReportMCG_r16_failureType_r16 `optional`
	measResultFreqList_r16         *MeasResultList2NR                    `optional`
	measResultFreqListEUTRA_r16    *MeasResultList2EUTRA                 `optional`
	measResultSCG_r16              *[]byte                               `optional`
	measResultSCG_EUTRA_r16        *[]byte                               `optional`
	measResultFreqListUTRA_FDD_r16 *MeasResultList2UTRA                  `optional`
}

func (ie *FailureReportMCG_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.failureType_r16 != nil, ie.measResultFreqList_r16 != nil, ie.measResultFreqListEUTRA_r16 != nil, ie.measResultSCG_r16 != nil, ie.measResultSCG_EUTRA_r16 != nil, ie.measResultFreqListUTRA_FDD_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.failureType_r16 != nil {
		if err = ie.failureType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode failureType_r16", err)
		}
	}
	if ie.measResultFreqList_r16 != nil {
		if err = ie.measResultFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultFreqList_r16", err)
		}
	}
	if ie.measResultFreqListEUTRA_r16 != nil {
		if err = ie.measResultFreqListEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultFreqListEUTRA_r16", err)
		}
	}
	if ie.measResultSCG_r16 != nil {
		if err = w.WriteOctetString(*ie.measResultSCG_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode measResultSCG_r16", err)
		}
	}
	if ie.measResultSCG_EUTRA_r16 != nil {
		if err = w.WriteOctetString(*ie.measResultSCG_EUTRA_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode measResultSCG_EUTRA_r16", err)
		}
	}
	if ie.measResultFreqListUTRA_FDD_r16 != nil {
		if err = ie.measResultFreqListUTRA_FDD_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultFreqListUTRA_FDD_r16", err)
		}
	}
	return nil
}

func (ie *FailureReportMCG_r16) Decode(r *uper.UperReader) error {
	var err error
	var failureType_r16Present bool
	if failureType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultFreqList_r16Present bool
	if measResultFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultFreqListEUTRA_r16Present bool
	if measResultFreqListEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultSCG_r16Present bool
	if measResultSCG_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultSCG_EUTRA_r16Present bool
	if measResultSCG_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultFreqListUTRA_FDD_r16Present bool
	if measResultFreqListUTRA_FDD_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if failureType_r16Present {
		ie.failureType_r16 = new(FailureReportMCG_r16_failureType_r16)
		if err = ie.failureType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode failureType_r16", err)
		}
	}
	if measResultFreqList_r16Present {
		ie.measResultFreqList_r16 = new(MeasResultList2NR)
		if err = ie.measResultFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultFreqList_r16", err)
		}
	}
	if measResultFreqListEUTRA_r16Present {
		ie.measResultFreqListEUTRA_r16 = new(MeasResultList2EUTRA)
		if err = ie.measResultFreqListEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultFreqListEUTRA_r16", err)
		}
	}
	if measResultSCG_r16Present {
		var tmp_os_measResultSCG_r16 []byte
		if tmp_os_measResultSCG_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode measResultSCG_r16", err)
		}
		ie.measResultSCG_r16 = &tmp_os_measResultSCG_r16
	}
	if measResultSCG_EUTRA_r16Present {
		var tmp_os_measResultSCG_EUTRA_r16 []byte
		if tmp_os_measResultSCG_EUTRA_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode measResultSCG_EUTRA_r16", err)
		}
		ie.measResultSCG_EUTRA_r16 = &tmp_os_measResultSCG_EUTRA_r16
	}
	if measResultFreqListUTRA_FDD_r16Present {
		ie.measResultFreqListUTRA_FDD_r16 = new(MeasResultList2UTRA)
		if err = ie.measResultFreqListUTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultFreqListUTRA_FDD_r16", err)
		}
	}
	return nil
}
