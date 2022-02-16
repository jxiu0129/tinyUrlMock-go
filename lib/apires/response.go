package apires

import (
	"time"
)

type (
	Base struct {
		Code    int    `json:"code" description:"API Return Code"`
		Message string `json:"message" description:"API Return Message"`
	}

	Errs struct {
		Base
		Errs map[string][]interface{} `json:"errs" description:"Multiple errors"`
	}

	Data struct {
		Base
		Data interface{} `json:"data" description:"API output data"`
	}

	List struct {
		Base
		List  interface{} `json:"data" description:"data list"`
		Count uint64      `json:"total_count" description:"total count of data"`
	}

	ExportList struct {
		Base
		Count      uint64      `json:"total_count" description:"total count of data"`
		ExportID   interface{} `json:"export_id" description:"export id"`
		ExportPath string      `json:"export_path" description:"path for download"`
	}

	TimezoneList struct {
		Base
		Count    uint64      `json:"total_count" description:"total count of data"`
		Timezone string      `json:"timezone"`
		List     interface{} `json:"data" description:"data list"`
	}

	PageByIDList struct {
		Base
		DataList interface{} `json:"data" description:"data list"`
	}

	LoginData struct {
		UID      uint64 `json:"uid" description:"Logged in UID"`
		Username string `json:"username" description:"Logged in username (AuthID)"`
	}

	LoginRedirect struct {
		RedirectUrl string `json:"redirect_url"`
	}

	DocStatusData struct {
		ID        uint64    `json:"id"`
		Status    int       `json:"status"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	DocStatus struct {
		Base
		Data DocStatusData `json:"data"`
	}

	DocCoverImageData struct {
		ID         uint64    `json:"id"`
		CoverImage string    `json:"cover_image"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	DocCoverImage struct {
		Base
		Data DocCoverImageData `json:"data"`
	}

	DocImgData struct {
		ID        uint64    `json:"id"`
		Img       string    `json:"img"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	DocImg struct {
		Base
		Data DocImgData `json:"data"`
	}

	MultiDataRes struct {
		SuccessID []interface{}            `json:"success_id"`
		ErrID     []interface{}            `json:"err_id"`
		ErrData   map[string][]interface{} `json:"err_data"`
	}

	MultiData struct {
		Base
		Data MultiDataRes `json:"data"`
	}

	DocMultiDataRes struct {
		SuccessID []uint64                 `json:"success_id"`
		ErrID     []uint64                 `json:"err_id"`
		ErrData   map[string][]interface{} `json:"err_data"`
	}

	DocMultiData struct {
		Base
		Data DocMultiDataRes `json:"data"`
	}
)
