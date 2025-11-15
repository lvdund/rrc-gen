package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RadioBearerConfig struct {
	srb_ToAddModList      *SRB_ToAddModList                     `optional`
	srb3_ToRelease        *RadioBearerConfig_srb3_ToRelease     `optional`
	drb_ToAddModList      *DRB_ToAddModList                     `optional`
	drb_ToReleaseList     *DRB_ToReleaseList                    `optional`
	securityConfig        *SecurityConfig                       `optional`
	mrb_ToAddModList_r17  *MRB_ToAddModList_r17                 `optional,ext-1`
	mrb_ToReleaseList_r17 *MRB_ToReleaseList_r17                `optional,ext-1`
	srb4_ToAddMod_r17     *SRB_ToAddMod                         `optional,ext-1`
	srb4_ToRelease_r17    *RadioBearerConfig_srb4_ToRelease_r17 `optional,ext-1`
}

func (ie *RadioBearerConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.mrb_ToAddModList_r17 != nil || ie.mrb_ToReleaseList_r17 != nil || ie.srb4_ToAddMod_r17 != nil || ie.srb4_ToRelease_r17 != nil
	preambleBits := []bool{hasExtensions, ie.srb_ToAddModList != nil, ie.srb3_ToRelease != nil, ie.drb_ToAddModList != nil, ie.drb_ToReleaseList != nil, ie.securityConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srb_ToAddModList != nil {
		if err = ie.srb_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode srb_ToAddModList", err)
		}
	}
	if ie.srb3_ToRelease != nil {
		if err = ie.srb3_ToRelease.Encode(w); err != nil {
			return utils.WrapError("Encode srb3_ToRelease", err)
		}
	}
	if ie.drb_ToAddModList != nil {
		if err = ie.drb_ToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode drb_ToAddModList", err)
		}
	}
	if ie.drb_ToReleaseList != nil {
		if err = ie.drb_ToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode drb_ToReleaseList", err)
		}
	}
	if ie.securityConfig != nil {
		if err = ie.securityConfig.Encode(w); err != nil {
			return utils.WrapError("Encode securityConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.mrb_ToAddModList_r17 != nil || ie.mrb_ToReleaseList_r17 != nil || ie.srb4_ToAddMod_r17 != nil || ie.srb4_ToRelease_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RadioBearerConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.mrb_ToAddModList_r17 != nil, ie.mrb_ToReleaseList_r17 != nil, ie.srb4_ToAddMod_r17 != nil, ie.srb4_ToRelease_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode mrb_ToAddModList_r17 optional
			if ie.mrb_ToAddModList_r17 != nil {
				if err = ie.mrb_ToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mrb_ToAddModList_r17", err)
				}
			}
			// encode mrb_ToReleaseList_r17 optional
			if ie.mrb_ToReleaseList_r17 != nil {
				if err = ie.mrb_ToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mrb_ToReleaseList_r17", err)
				}
			}
			// encode srb4_ToAddMod_r17 optional
			if ie.srb4_ToAddMod_r17 != nil {
				if err = ie.srb4_ToAddMod_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srb4_ToAddMod_r17", err)
				}
			}
			// encode srb4_ToRelease_r17 optional
			if ie.srb4_ToRelease_r17 != nil {
				if err = ie.srb4_ToRelease_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srb4_ToRelease_r17", err)
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

func (ie *RadioBearerConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var srb_ToAddModListPresent bool
	if srb_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srb3_ToReleasePresent bool
	if srb3_ToReleasePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var drb_ToAddModListPresent bool
	if drb_ToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var drb_ToReleaseListPresent bool
	if drb_ToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var securityConfigPresent bool
	if securityConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if srb_ToAddModListPresent {
		ie.srb_ToAddModList = new(SRB_ToAddModList)
		if err = ie.srb_ToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode srb_ToAddModList", err)
		}
	}
	if srb3_ToReleasePresent {
		ie.srb3_ToRelease = new(RadioBearerConfig_srb3_ToRelease)
		if err = ie.srb3_ToRelease.Decode(r); err != nil {
			return utils.WrapError("Decode srb3_ToRelease", err)
		}
	}
	if drb_ToAddModListPresent {
		ie.drb_ToAddModList = new(DRB_ToAddModList)
		if err = ie.drb_ToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode drb_ToAddModList", err)
		}
	}
	if drb_ToReleaseListPresent {
		ie.drb_ToReleaseList = new(DRB_ToReleaseList)
		if err = ie.drb_ToReleaseList.Decode(r); err != nil {
			return utils.WrapError("Decode drb_ToReleaseList", err)
		}
	}
	if securityConfigPresent {
		ie.securityConfig = new(SecurityConfig)
		if err = ie.securityConfig.Decode(r); err != nil {
			return utils.WrapError("Decode securityConfig", err)
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

			mrb_ToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mrb_ToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srb4_ToAddMod_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srb4_ToRelease_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode mrb_ToAddModList_r17 optional
			if mrb_ToAddModList_r17Present {
				ie.mrb_ToAddModList_r17 = new(MRB_ToAddModList_r17)
				if err = ie.mrb_ToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mrb_ToAddModList_r17", err)
				}
			}
			// decode mrb_ToReleaseList_r17 optional
			if mrb_ToReleaseList_r17Present {
				ie.mrb_ToReleaseList_r17 = new(MRB_ToReleaseList_r17)
				if err = ie.mrb_ToReleaseList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mrb_ToReleaseList_r17", err)
				}
			}
			// decode srb4_ToAddMod_r17 optional
			if srb4_ToAddMod_r17Present {
				ie.srb4_ToAddMod_r17 = new(SRB_ToAddMod)
				if err = ie.srb4_ToAddMod_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srb4_ToAddMod_r17", err)
				}
			}
			// decode srb4_ToRelease_r17 optional
			if srb4_ToRelease_r17Present {
				ie.srb4_ToRelease_r17 = new(RadioBearerConfig_srb4_ToRelease_r17)
				if err = ie.srb4_ToRelease_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srb4_ToRelease_r17", err)
				}
			}
		}
	}
	return nil
}
