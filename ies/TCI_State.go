package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TCI_State struct {
	tci_StateId                TCI_StateId                 `madatory`
	qcl_Type1                  QCL_Info                    `madatory`
	qcl_Type2                  *QCL_Info                   `optional`
	additionalPCI_r17          *AdditionalPCIIndex_r17     `optional,ext-1`
	pathlossReferenceRS_Id_r17 *PathlossReferenceRS_Id_r17 `optional,ext-1`
	ul_powerControl_r17        *Uplink_powerControlId_r17  `optional,ext-1`
}

func (ie *TCI_State) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.additionalPCI_r17 != nil || ie.pathlossReferenceRS_Id_r17 != nil || ie.ul_powerControl_r17 != nil
	preambleBits := []bool{hasExtensions, ie.qcl_Type2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.tci_StateId.Encode(w); err != nil {
		return utils.WrapError("Encode tci_StateId", err)
	}
	if err = ie.qcl_Type1.Encode(w); err != nil {
		return utils.WrapError("Encode qcl_Type1", err)
	}
	if ie.qcl_Type2 != nil {
		if err = ie.qcl_Type2.Encode(w); err != nil {
			return utils.WrapError("Encode qcl_Type2", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.additionalPCI_r17 != nil || ie.pathlossReferenceRS_Id_r17 != nil || ie.ul_powerControl_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap TCI_State", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.additionalPCI_r17 != nil, ie.pathlossReferenceRS_Id_r17 != nil, ie.ul_powerControl_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode additionalPCI_r17 optional
			if ie.additionalPCI_r17 != nil {
				if err = ie.additionalPCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode additionalPCI_r17", err)
				}
			}
			// encode pathlossReferenceRS_Id_r17 optional
			if ie.pathlossReferenceRS_Id_r17 != nil {
				if err = ie.pathlossReferenceRS_Id_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pathlossReferenceRS_Id_r17", err)
				}
			}
			// encode ul_powerControl_r17 optional
			if ie.ul_powerControl_r17 != nil {
				if err = ie.ul_powerControl_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_powerControl_r17", err)
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

func (ie *TCI_State) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var qcl_Type2Present bool
	if qcl_Type2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.tci_StateId.Decode(r); err != nil {
		return utils.WrapError("Decode tci_StateId", err)
	}
	if err = ie.qcl_Type1.Decode(r); err != nil {
		return utils.WrapError("Decode qcl_Type1", err)
	}
	if qcl_Type2Present {
		ie.qcl_Type2 = new(QCL_Info)
		if err = ie.qcl_Type2.Decode(r); err != nil {
			return utils.WrapError("Decode qcl_Type2", err)
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

			additionalPCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pathlossReferenceRS_Id_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_powerControl_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode additionalPCI_r17 optional
			if additionalPCI_r17Present {
				ie.additionalPCI_r17 = new(AdditionalPCIIndex_r17)
				if err = ie.additionalPCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode additionalPCI_r17", err)
				}
			}
			// decode pathlossReferenceRS_Id_r17 optional
			if pathlossReferenceRS_Id_r17Present {
				ie.pathlossReferenceRS_Id_r17 = new(PathlossReferenceRS_Id_r17)
				if err = ie.pathlossReferenceRS_Id_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pathlossReferenceRS_Id_r17", err)
				}
			}
			// decode ul_powerControl_r17 optional
			if ul_powerControl_r17Present {
				ie.ul_powerControl_r17 = new(Uplink_powerControlId_r17)
				if err = ie.ul_powerControl_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_powerControl_r17", err)
				}
			}
		}
	}
	return nil
}
