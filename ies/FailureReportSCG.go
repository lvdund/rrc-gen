package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureReportSCG struct {
	failureType           FailureReportSCG_failureType        `madatory`
	measResultFreqList    *MeasResultFreqList                 `optional`
	measResultSCG_Failure *[]byte                             `optional`
	locationInfo_r16      *LocationInfo_r16                   `optional,ext-1`
	failureType_v1610     *FailureReportSCG_failureType_v1610 `optional,ext-1`
	previousPSCellId_r17  *PSCellId_r17                       `optional,ext-2`
	failedPSCellId_r17    *PSCellId_r17                       `optional,ext-2`
	timeSCGFailure_r17    *int64                              `lb:0,ub:1023,optional,ext-2`
	perRAInfoList_r17     *PerRAInfoList_r16                  `optional,ext-2`
}

func (ie *FailureReportSCG) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.locationInfo_r16 != nil || ie.failureType_v1610 != nil || ie.previousPSCellId_r17 != nil || ie.failedPSCellId_r17 != nil || ie.timeSCGFailure_r17 != nil || ie.perRAInfoList_r17 != nil
	preambleBits := []bool{hasExtensions, ie.measResultFreqList != nil, ie.measResultSCG_Failure != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.failureType.Encode(w); err != nil {
		return utils.WrapError("Encode failureType", err)
	}
	if ie.measResultFreqList != nil {
		if err = ie.measResultFreqList.Encode(w); err != nil {
			return utils.WrapError("Encode measResultFreqList", err)
		}
	}
	if ie.measResultSCG_Failure != nil {
		if err = w.WriteOctetString(*ie.measResultSCG_Failure, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode measResultSCG_Failure", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.locationInfo_r16 != nil || ie.failureType_v1610 != nil, ie.previousPSCellId_r17 != nil || ie.failedPSCellId_r17 != nil || ie.timeSCGFailure_r17 != nil || ie.perRAInfoList_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap FailureReportSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.locationInfo_r16 != nil, ie.failureType_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode locationInfo_r16 optional
			if ie.locationInfo_r16 != nil {
				if err = ie.locationInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode locationInfo_r16", err)
				}
			}
			// encode failureType_v1610 optional
			if ie.failureType_v1610 != nil {
				if err = ie.failureType_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode failureType_v1610", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.previousPSCellId_r17 != nil, ie.failedPSCellId_r17 != nil, ie.timeSCGFailure_r17 != nil, ie.perRAInfoList_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode previousPSCellId_r17 optional
			if ie.previousPSCellId_r17 != nil {
				if err = ie.previousPSCellId_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode previousPSCellId_r17", err)
				}
			}
			// encode failedPSCellId_r17 optional
			if ie.failedPSCellId_r17 != nil {
				if err = ie.failedPSCellId_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode failedPSCellId_r17", err)
				}
			}
			// encode timeSCGFailure_r17 optional
			if ie.timeSCGFailure_r17 != nil {
				if err = extWriter.WriteInteger(*ie.timeSCGFailure_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Encode timeSCGFailure_r17", err)
				}
			}
			// encode perRAInfoList_r17 optional
			if ie.perRAInfoList_r17 != nil {
				if err = ie.perRAInfoList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode perRAInfoList_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *FailureReportSCG) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultFreqListPresent bool
	if measResultFreqListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultSCG_FailurePresent bool
	if measResultSCG_FailurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.failureType.Decode(r); err != nil {
		return utils.WrapError("Decode failureType", err)
	}
	if measResultFreqListPresent {
		ie.measResultFreqList = new(MeasResultFreqList)
		if err = ie.measResultFreqList.Decode(r); err != nil {
			return utils.WrapError("Decode measResultFreqList", err)
		}
	}
	if measResultSCG_FailurePresent {
		var tmp_os_measResultSCG_Failure []byte
		if tmp_os_measResultSCG_Failure, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode measResultSCG_Failure", err)
		}
		ie.measResultSCG_Failure = &tmp_os_measResultSCG_Failure
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			locationInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			failureType_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode locationInfo_r16 optional
			if locationInfo_r16Present {
				ie.locationInfo_r16 = new(LocationInfo_r16)
				if err = ie.locationInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode locationInfo_r16", err)
				}
			}
			// decode failureType_v1610 optional
			if failureType_v1610Present {
				ie.failureType_v1610 = new(FailureReportSCG_failureType_v1610)
				if err = ie.failureType_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode failureType_v1610", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			previousPSCellId_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			failedPSCellId_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			timeSCGFailure_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			perRAInfoList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode previousPSCellId_r17 optional
			if previousPSCellId_r17Present {
				ie.previousPSCellId_r17 = new(PSCellId_r17)
				if err = ie.previousPSCellId_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode previousPSCellId_r17", err)
				}
			}
			// decode failedPSCellId_r17 optional
			if failedPSCellId_r17Present {
				ie.failedPSCellId_r17 = new(PSCellId_r17)
				if err = ie.failedPSCellId_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode failedPSCellId_r17", err)
				}
			}
			// decode timeSCGFailure_r17 optional
			if timeSCGFailure_r17Present {
				var tmp_int_timeSCGFailure_r17 int64
				if tmp_int_timeSCGFailure_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Decode timeSCGFailure_r17", err)
				}
				ie.timeSCGFailure_r17 = &tmp_int_timeSCGFailure_r17
			}
			// decode perRAInfoList_r17 optional
			if perRAInfoList_r17Present {
				ie.perRAInfoList_r17 = new(PerRAInfoList_r16)
				if err = ie.perRAInfoList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode perRAInfoList_r17", err)
				}
			}
		}
	}
	return nil
}
