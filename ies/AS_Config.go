package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AS_Config struct {
	rrcReconfiguration     []byte                          `madatory`
	sourceRB_SN_Config     *[]byte                         `optional,ext-1`
	sourceSCG_NR_Config    *[]byte                         `optional,ext-1`
	sourceSCG_EUTRA_Config *[]byte                         `optional,ext-1`
	sourceSCG_Configured   *AS_Config_sourceSCG_Configured `optional,ext-2`
	sdt_Config_r17         *SDT_Config_r17                 `optional,ext-3`
}

func (ie *AS_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sourceRB_SN_Config != nil || ie.sourceSCG_NR_Config != nil || ie.sourceSCG_EUTRA_Config != nil || ie.sourceSCG_Configured != nil || ie.sdt_Config_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteOctetString(ie.rrcReconfiguration, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString rrcReconfiguration", err)
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.sourceRB_SN_Config != nil || ie.sourceSCG_NR_Config != nil || ie.sourceSCG_EUTRA_Config != nil, ie.sourceSCG_Configured != nil, ie.sdt_Config_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap AS_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sourceRB_SN_Config != nil, ie.sourceSCG_NR_Config != nil, ie.sourceSCG_EUTRA_Config != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sourceRB_SN_Config optional
			if ie.sourceRB_SN_Config != nil {
				if err = extWriter.WriteOctetString(*ie.sourceRB_SN_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode sourceRB_SN_Config", err)
				}
			}
			// encode sourceSCG_NR_Config optional
			if ie.sourceSCG_NR_Config != nil {
				if err = extWriter.WriteOctetString(*ie.sourceSCG_NR_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode sourceSCG_NR_Config", err)
				}
			}
			// encode sourceSCG_EUTRA_Config optional
			if ie.sourceSCG_EUTRA_Config != nil {
				if err = extWriter.WriteOctetString(*ie.sourceSCG_EUTRA_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Encode sourceSCG_EUTRA_Config", err)
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
			optionals_ext_2 := []bool{ie.sourceSCG_Configured != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sourceSCG_Configured optional
			if ie.sourceSCG_Configured != nil {
				if err = ie.sourceSCG_Configured.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sourceSCG_Configured", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.sdt_Config_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sdt_Config_r17 optional
			if ie.sdt_Config_r17 != nil {
				if err = ie.sdt_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sdt_Config_r17", err)
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

func (ie *AS_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_os_rrcReconfiguration []byte
	if tmp_os_rrcReconfiguration, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString rrcReconfiguration", err)
	}
	ie.rrcReconfiguration = tmp_os_rrcReconfiguration

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			sourceRB_SN_ConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sourceSCG_NR_ConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sourceSCG_EUTRA_ConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sourceRB_SN_Config optional
			if sourceRB_SN_ConfigPresent {
				var tmp_os_sourceRB_SN_Config []byte
				if tmp_os_sourceRB_SN_Config, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode sourceRB_SN_Config", err)
				}
				ie.sourceRB_SN_Config = &tmp_os_sourceRB_SN_Config
			}
			// decode sourceSCG_NR_Config optional
			if sourceSCG_NR_ConfigPresent {
				var tmp_os_sourceSCG_NR_Config []byte
				if tmp_os_sourceSCG_NR_Config, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode sourceSCG_NR_Config", err)
				}
				ie.sourceSCG_NR_Config = &tmp_os_sourceSCG_NR_Config
			}
			// decode sourceSCG_EUTRA_Config optional
			if sourceSCG_EUTRA_ConfigPresent {
				var tmp_os_sourceSCG_EUTRA_Config []byte
				if tmp_os_sourceSCG_EUTRA_Config, err = extReader.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
					return utils.WrapError("Decode sourceSCG_EUTRA_Config", err)
				}
				ie.sourceSCG_EUTRA_Config = &tmp_os_sourceSCG_EUTRA_Config
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sourceSCG_ConfiguredPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sourceSCG_Configured optional
			if sourceSCG_ConfiguredPresent {
				ie.sourceSCG_Configured = new(AS_Config_sourceSCG_Configured)
				if err = ie.sourceSCG_Configured.Decode(extReader); err != nil {
					return utils.WrapError("Decode sourceSCG_Configured", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sdt_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sdt_Config_r17 optional
			if sdt_Config_r17Present {
				ie.sdt_Config_r17 = new(SDT_Config_r17)
				if err = ie.sdt_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sdt_Config_r17", err)
				}
			}
		}
	}
	return nil
}
