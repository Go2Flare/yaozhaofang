package handler

import (
	"context"
	"go_code/yaozhaofang/service/getArea/model"
	"go_code/yaozhaofang/service/getArea/utils"

	getArea "go_code/yaozhaofang/service/getArea/proto/getArea"
)

type GetArea struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetArea) MicroGetArea(ctx context.Context, req *getArea.Request, rsp *getArea.Response) error {
	//redis获取MySQL中中获取数据
	areas,err := model.GetArea()
	if err != nil {
		//获取数据失败返回前端
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return err
	}
	//得到数据
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)

	//将地区信息回传
	for _,v := range areas{
		var areaInfo getArea.AreaInfo
		areaInfo.Aid = int32(v.Id)
		areaInfo.Aname = v.Name

		rsp.Data = append(rsp.Data,&areaInfo)
	}

	return nil
}
