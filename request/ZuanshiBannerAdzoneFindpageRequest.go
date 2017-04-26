package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerAdzoneFindpageRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdzoneFindpageRequest() (req *ZuanshiBannerAdzoneFindpageRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adzone.findpage", Params: make(map[string]interface{}, 9)}
	req = &ZuanshiBannerAdzoneFindpageRequest{
		Request: &request,
	}
	return
}

// 广告位id列表
func (req *ZuanshiBannerAdzoneFindpageRequest) SetAdzoneIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["adzone_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetAdzoneIdList() []uint64 {
	var ids []uint64
	if idStr, found := req.Request.Params["adzone_id_list"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, aId)
			}
		}
	}
	return ids
}

// 广告位尺寸列表
func (req *ZuanshiBannerAdzoneFindpageRequest) SetAdzoneSizeList(sizes []string) {
	req.Request.Params["adzone_size_list"] = strings.Join(sizes, ",")
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetAdzoneSizeList() []string {
	var sizes []string
	if sizeStr, found := req.Request.Params["adzone_size_list"]; found {
		sizes = strings.Split(sizeStr.(string), ",")
	}
	return sizes
}

// 分页大小, 默认值：15 最小值：1 最大长度：200
func (req *ZuanshiBannerAdzoneFindpageRequest) SetPageSize(pageSize uint) {
	if pageSize == 0 {
		pageSize = 15
	}
	req.Request.Params["page_size"] = pageSize
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetPageSize() uint {
	if pageSize, found := req.Request.Params["page_size"]; found {
		return pageSize.(uint)
	}
	return 15
}

// 当前页码, 默认值：1
func (req *ZuanshiBannerAdzoneFindpageRequest) SetPageNum(pageNum uint) {
	if pageNum == 0 {
		pageNum = 1
	}
	req.Request.Params["page_num"] = pageNum
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetPageNum() uint {
	if pageNum, found := req.Request.Params["page_num"]; found {
		return pageNum.(uint)
	}
	return 1
}

// 允许的创意类型
func (req *ZuanshiBannerAdzoneFindpageRequest) SetAllowAdFormatList(formatList []uint) {
	var formatsStr []string
	for _, f := range formatList {
		formatsStr = append(formatsStr, strconv.FormatUint(uint64(f), 10))
	}
	req.Request.Params["allow_ad_format_list"] = strings.Join(formatsStr, ",")
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetAllowAdFormatList() []uint {
	var formatList []uint
	if flStr, found := req.Request.Params["allow_ad_format_list"]; found {
		fs := strings.Split(flStr.(string), ",")
		for _, f := range fs {
			if aF, err := strconv.ParseUint(f, 10, 64); err == nil && aF > 0 {
				formatList = append(formatList, uint(aF))
			}
		}
	}
	return formatList
}

// 媒体类型列表
func (req *ZuanshiBannerAdzoneFindpageRequest) SetMediaTypeList(mediaTypes []uint) {
	var mediaTypesStr []string
	for _, mt := range mediaTypes {
		mediaTypesStr = append(mediaTypesStr, strconv.FormatUint(uint64(mt), 10))
	}
	req.Request.Params["media_type_list"] = strings.Join(mediaTypesStr, ",")
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetMediaTypeList() []uint {
	var mediaTypes []uint
	if mediaTypeStr, found := req.Request.Params["media_type_list"]; found {
		mts := strings.Split(mediaTypeStr.(string), ",")
		for _, mt := range mts {
			if aMT, err := strconv.ParseUint(mt, 10, 64); err == nil && aMT > 0 {
				mediaTypes = append(mediaTypes, uint(aMT))
			}
		}
	}
	return mediaTypes
}

// 广告位类型，1表示广告位，2表示广告位包
func (req *ZuanshiBannerAdzoneFindpageRequest) SetAdzoneType(t uint8) {
	req.Request.Params["adzone_type"] = t
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetAdzoneType() uint8 {
	if t, found := req.Request.Params["adzone_type"]; found {
		return t.(uint8)
	}
	return 0
}

// 广告位名称
func (req *ZuanshiBannerAdzoneFindpageRequest) SetAdzoneName(name string) {
	req.Request.Params["adzone_name"] = name
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetAdzoneName() string {
	if name, found := req.Request.Params["adzone_name"]; found {
		return name.(string)
	}
	return ""
}

// 结算类型，1表示CPM 2表示CPC
func (req *ZuanshiBannerAdzoneFindpageRequest) SetSettleTypeList(settleTypes []uint8) {
	var list []string
	for _, t := range settleTypes {
		list = append(list, strconv.FormatUint(uint64(t), 10))
	}
	req.Request.Params["settle_type_list"] = strings.Join(list, ",")
}

func (req *ZuanshiBannerAdzoneFindpageRequest) GetSettleTypeList() []uint8 {
	var settleTypes []uint8
	if listStr, found := req.Request.Params["settle_type_list"]; found {
		types := strings.Split(listStr.(string), ",")
		for _, t := range types {
			if aType, err := strconv.ParseUint(t, 10, 64); err == nil && aType > 0 {
				settleTypes = append(settleTypes, uint8(aType))
			}
		}
	}
	return settleTypes
}
