package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRM_Config struct {
	ue_InactiveTime               *RRM_Config_ue_InactiveTime      `optional`
	candidateCellInfoList         *MeasResultList2NR               `optional`
	candidateCellInfoListSN_EUTRA *MeasResultServFreqListEUTRA_SCG `optional,ext-1`
}

func (ie *RRM_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.candidateCellInfoListSN_EUTRA != nil
	preambleBits := []bool{hasExtensions, ie.ue_InactiveTime != nil, ie.candidateCellInfoList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ue_InactiveTime != nil {
		if err = ie.ue_InactiveTime.Encode(w); err != nil {
			return utils.WrapError("Encode ue_InactiveTime", err)
		}
	}
	if ie.candidateCellInfoList != nil {
		if err = ie.candidateCellInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode candidateCellInfoList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.candidateCellInfoListSN_EUTRA != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RRM_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.candidateCellInfoListSN_EUTRA != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode candidateCellInfoListSN_EUTRA optional
			if ie.candidateCellInfoListSN_EUTRA != nil {
				if err = ie.candidateCellInfoListSN_EUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode candidateCellInfoListSN_EUTRA", err)
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

func (ie *RRM_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_InactiveTimePresent bool
	if ue_InactiveTimePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateCellInfoListPresent bool
	if candidateCellInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ue_InactiveTimePresent {
		ie.ue_InactiveTime = new(RRM_Config_ue_InactiveTime)
		if err = ie.ue_InactiveTime.Decode(r); err != nil {
			return utils.WrapError("Decode ue_InactiveTime", err)
		}
	}
	if candidateCellInfoListPresent {
		ie.candidateCellInfoList = new(MeasResultList2NR)
		if err = ie.candidateCellInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode candidateCellInfoList", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			candidateCellInfoListSN_EUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode candidateCellInfoListSN_EUTRA optional
			if candidateCellInfoListSN_EUTRAPresent {
				ie.candidateCellInfoListSN_EUTRA = new(MeasResultServFreqListEUTRA_SCG)
				if err = ie.candidateCellInfoListSN_EUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode candidateCellInfoListSN_EUTRA", err)
				}
			}
		}
	}
	return nil
}
