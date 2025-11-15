package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSCG_Failure struct {
	measResultPerMOList MeasResultList2NR `madatory`
	locationInfo_r16    *LocationInfo_r16 `optional,ext-1`
}

func (ie *MeasResultSCG_Failure) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.locationInfo_r16 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.measResultPerMOList.Encode(w); err != nil {
		return utils.WrapError("Encode measResultPerMOList", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.locationInfo_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasResultSCG_Failure", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.locationInfo_r16 != nil}
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

func (ie *MeasResultSCG_Failure) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.measResultPerMOList.Decode(r); err != nil {
		return utils.WrapError("Decode measResultPerMOList", err)
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

			locationInfo_r16Present, err := extReader.ReadBool()
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
		}
	}
	return nil
}
