package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_TPC_CommandConfig struct {
	startingBitOfFormat2_3    *int64 `lb:1,ub:31,optional`
	fieldTypeFormat2_3        *int64 `lb:0,ub:1,optional`
	startingBitOfFormat2_3SUL *int64 `lb:1,ub:31,optional,ext-1`
}

func (ie *SRS_TPC_CommandConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.startingBitOfFormat2_3SUL != nil
	preambleBits := []bool{hasExtensions, ie.startingBitOfFormat2_3 != nil, ie.fieldTypeFormat2_3 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.startingBitOfFormat2_3 != nil {
		if err = w.WriteInteger(*ie.startingBitOfFormat2_3, &uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			return utils.WrapError("Encode startingBitOfFormat2_3", err)
		}
	}
	if ie.fieldTypeFormat2_3 != nil {
		if err = w.WriteInteger(*ie.fieldTypeFormat2_3, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode fieldTypeFormat2_3", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.startingBitOfFormat2_3SUL != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SRS_TPC_CommandConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.startingBitOfFormat2_3SUL != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode startingBitOfFormat2_3SUL optional
			if ie.startingBitOfFormat2_3SUL != nil {
				if err = extWriter.WriteInteger(*ie.startingBitOfFormat2_3SUL, &uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode startingBitOfFormat2_3SUL", err)
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

func (ie *SRS_TPC_CommandConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var startingBitOfFormat2_3Present bool
	if startingBitOfFormat2_3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var fieldTypeFormat2_3Present bool
	if fieldTypeFormat2_3Present, err = r.ReadBool(); err != nil {
		return err
	}
	if startingBitOfFormat2_3Present {
		var tmp_int_startingBitOfFormat2_3 int64
		if tmp_int_startingBitOfFormat2_3, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode startingBitOfFormat2_3", err)
		}
		ie.startingBitOfFormat2_3 = &tmp_int_startingBitOfFormat2_3
	}
	if fieldTypeFormat2_3Present {
		var tmp_int_fieldTypeFormat2_3 int64
		if tmp_int_fieldTypeFormat2_3, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode fieldTypeFormat2_3", err)
		}
		ie.fieldTypeFormat2_3 = &tmp_int_fieldTypeFormat2_3
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

			startingBitOfFormat2_3SULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode startingBitOfFormat2_3SUL optional
			if startingBitOfFormat2_3SULPresent {
				var tmp_int_startingBitOfFormat2_3SUL int64
				if tmp_int_startingBitOfFormat2_3SUL, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode startingBitOfFormat2_3SUL", err)
				}
				ie.startingBitOfFormat2_3SUL = &tmp_int_startingBitOfFormat2_3SUL
			}
		}
	}
	return nil
}
