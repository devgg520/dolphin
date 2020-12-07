// Code generated by dol build. Only Generate by tools if not existed.
// source: sys_notification.go

package app

import (
	"errors"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/gin/binding"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/time"
)

// SysNotificationAdd api implementation
// @Summary 添加站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param notification body model.SysNotification false "站内消息信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/notification/add [post]
func SysNotificationAdd(ctx *Context) {
	var payload model.SysNotification
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.ID = null.StringFromUUID()
	payload.CreateTime = null.TimeFrom(time.Now().Value())
	payload.CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.DelFlag = null.IntFrom(0)
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationBatchAdd api implementation
// @Summary 添加站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body []model.SysNotification false "站内消息信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/notification/batch_add [post]
func SysNotificationBatchAdd(ctx *Context) {
	var payload []model.SysNotification
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	for i := range payload {
		payload[i].ID = null.StringFromUUID()
		payload[i].CreateTime = null.TimeFrom(time.Now().Value())
		payload[i].CreateBy = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		payload[i].DelFlag = null.IntFrom(0)
	}
	ret, err := ctx.DB.Insert(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationDel api implementation
// @Summary 删除站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body model.SysNotification false "站内消息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/notification/del [delete]
func SysNotificationDel(ctx *Context) {
	var payload model.SysNotification
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ret, err := ctx.DB.In("id", payload.ID.String).Update(&model.SysNotification{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationBatchDel api implementation
// @Summary 删除站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body []model.SysNotification false "站内消息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/notification/batch_del [delete]
func SysNotificationBatchDel(ctx *Context) {
	var payload []model.SysNotification
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	var ids = funk.Map(payload, func(form model.SysNotification) string { return form.ID.String }).([]string)
	ret, err := ctx.DB.In("id", ids).Update(&model.SysNotification{
		UpdateTime: null.TimeFrom(time.Now().Value()),
		UpdateBy:   null.StringFrom(ctx.GetToken().GetUserID()),
		DelFlag:    null.IntFrom(1),
	})
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationUpdate api implementation
// @Summary 更新站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param notification body model.SysRole false "站内消息信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/notification/update [put]
func SysNotificationUpdate(ctx *Context) {
	var payload model.SysRole
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	payload.UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
	payload.UpdateTime = null.TimeFrom(time.Now().Value())
	ret, err := ctx.DB.ID(payload.ID.String).Update(&payload)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationBatchUpdate api implementation
// @Summary 更新站内消息
// @Tags 站内消息
// @Accept application/json
// @Param Authorization header string false "认证令牌"
// @Param sys_notification body []model.SysNotification false "站内消息信息"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router/api/sys/notification/batch_update [put]
func SysNotificationBatchUpdate(ctx *Context) {
	var payload []model.SysNotification
	var err error
	var ret []int64
	var r int64
	if err := ctx.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	s := ctx.DB.NewSession()
	for i := range payload {
		payload[i].UpdateTime = null.TimeFrom(time.Now().Value())
		payload[i].UpdateBy = null.StringFrom(ctx.GetToken().GetUserID())
		r, err = s.ID(payload[i].ID.String).Update(&payload[i])
		ret = append(ret, r)
	}
	if err != nil {
		s.Rollback()
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	err = s.Commit()
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationPage api implementation
// @Summary 站内消息分页查询
// @Tags 站内消息
// @Param Authorization header string false "认证令牌"
// @Param page  query  int false "页码"
// @Param size  query  int false "单页数"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/notification/page [get]
func SysNotificationPage(ctx *Context) {
	q := ctx.TypeQuery()
	q.SetInt("page", 1)
	q.SetInt("size", 15)
	q.SetRule("sys_notification_page")
	q.SetTags()
	ret, err := ctx.PageSearch(ctx.DB, "sys_notification", "page", "sys_notification", q.Value())
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ctx.Success(ret)
}

// SysNotificationGet api implementation
// @Summary 获取站内消息信息
// @Tags 站内消息
// @Param Authorization header string false "认证令牌"
// @Param id  query  string false "站内消息id"
// @Failure 403 {object} model.Fail
// @Success 200 {object} model.Success
// @Failure 500 {object} model.Fail
// @Router /api/sys/notification/get [get]
func SysNotificationGet(ctx *Context) {
	var entity model.SysNotification
	err := ctx.ShouldBindQuery(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	ext, err := ctx.DB.Get(&entity)
	if err != nil {
		logrus.Error(err)
		ctx.Fail(err)
		return
	}
	if !ext {
		ctx.Fail(errors.New("not found"))
		return
	}
	ctx.Success(entity)
}
